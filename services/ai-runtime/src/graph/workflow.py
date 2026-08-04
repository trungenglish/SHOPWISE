import json
from typing import Any, TypedDict

from langchain_core.runnables import RunnableLambda
from langgraph.graph import END, START, StateGraph
from pydantic import ValidationError

from src.core.prompts import PromptManager
from src.llm.provider import LLMProvider
from src.models.schemas import (
    AgentDraft,
    AgentResponse,
    CatalogProduct,
    hydrate_agent_response,
)
from src.tools.interfaces import ToolProxy

MAX_SELF_CORRECTION_RETRIES = 2


class GraphState(TypedDict, total=False):
    messages: list[dict[str, Any]]
    session_id: str
    retry_count: int
    catalog: list[CatalogProduct]
    allowed_comparison_ids: list[str]
    response: AgentResponse
    error: str | None


def create_workflow(provider: LLMProvider, tool_proxy: ToolProxy, model: str) -> Any:
    workflow = StateGraph(GraphState)

    async def load_catalog(_state: GraphState) -> GraphState:
        return {"catalog": await tool_proxy.catalog_search()}

    async def invoke_llm(state: GraphState) -> GraphState:
        catalog_json = json.dumps(
            [product.model_dump() for product in state["catalog"]],
            ensure_ascii=False,
        )
        messages = [
            {
                "role": "system",
                "content": PromptManager.get_system_prompt(
                    catalog_json,
                    json.dumps(state.get("allowed_comparison_ids", [])),
                ),
            },
            *state["messages"],
        ]
        if state.get("error"):
            messages.append(
                {
                    "role": "system",
                    "content": (
                        "Your previous response was invalid. Return JSON matching "
                        f"the required schema. Validation error: {state['error']}"
                    ),
                }
            )

        raw_response = await provider.chat_completion(
            messages=messages,
            model=model,
            session_id=state["session_id"],
            response_format={
                "type": "json_schema",
                "json_schema": {
                    "name": "AgentDraft",
                    "schema": AgentDraft.model_json_schema(),
                    "strict": True,
                },
            },
        )

        try:
            draft = AgentDraft.model_validate_json(raw_response)
            response = hydrate_agent_response(
                draft,
                state["catalog"],
                allowed_comparison_ids=set(state.get("allowed_comparison_ids", [])),
                language_source=next(
                    (
                        str(message.get("content", ""))
                        for message in reversed(state["messages"])
                        if message.get("role") == "user"
                    ),
                    "",
                ),
            )
            return {"response": response, "error": None}
        except (ValidationError, ValueError) as error:
            return {
                "error": str(error),
                "retry_count": state.get("retry_count", 0) + 1,
            }

    def route_after_llm(state: GraphState) -> str:
        if state.get("error") is None:
            return END
        if state.get("retry_count", 0) <= MAX_SELF_CORRECTION_RETRIES:
            return "invoke_llm"
        return END

    workflow.add_node("load_catalog", RunnableLambda(load_catalog))
    workflow.add_node("invoke_llm", RunnableLambda(invoke_llm))
    workflow.add_edge(START, "load_catalog")
    workflow.add_edge("load_catalog", "invoke_llm")
    workflow.add_conditional_edges("invoke_llm", route_after_llm)
    return workflow.compile()

import json
from collections.abc import AsyncGenerator
from typing import Literal, cast

import httpx
from fastapi import APIRouter, Depends, HTTPException
from pydantic import BaseModel, Field
from sse_starlette.sse import EventSourceResponse

from src.api.streaming import generate_agent_sse
from src.core.config import settings
from src.graph.workflow import create_workflow
from src.llm.openai_provider import OpenAIProvider
from src.models.schemas import AgentResponse
from src.tools.interfaces import ToolProxy

router = APIRouter()


class ConversationMessage(BaseModel):
    role: Literal["user", "assistant"]
    content: str = Field(min_length=1)


class ChatRequest(BaseModel):
    session_id: str = Field(min_length=1)
    messages: list[ConversationMessage] = Field(min_length=1)
    allowed_comparison_ids: list[str] = Field(default_factory=list)


def get_provider() -> OpenAIProvider:
    return OpenAIProvider(
        api_key=settings.openai_api_key.get_secret_value(),
        base_url=settings.openai_base_url,
    )


def get_tool_proxy() -> ToolProxy:
    return ToolProxy(settings.backend_url)


@router.post("/chat", response_model=AgentResponse, response_model_exclude_none=True)
async def chat_endpoint(
    request: ChatRequest,
    provider: OpenAIProvider = Depends(get_provider),
    tool_proxy: ToolProxy = Depends(get_tool_proxy),
) -> AgentResponse:
    workflow = create_workflow(provider, tool_proxy, settings.model)
    try:
        result = await workflow.ainvoke(
            {
                "messages": [message.model_dump() for message in request.messages],
                "session_id": request.session_id,
                "retry_count": 0,
                "allowed_comparison_ids": request.allowed_comparison_ids,
            }
        )
    except httpx.HTTPError as error:
        raise HTTPException(status_code=502, detail="catalog unavailable") from error

    response = result.get("response")
    if response is None:
        raise HTTPException(status_code=502, detail="invalid model response")
    return cast(AgentResponse, response)


@router.post("/chat/stream")
async def chat_stream_endpoint(
    request: ChatRequest,
    provider: OpenAIProvider = Depends(get_provider),
    tool_proxy: ToolProxy = Depends(get_tool_proxy),
) -> EventSourceResponse:
    async def events() -> AsyncGenerator[dict[str, str], None]:
        try:
            response = await chat_endpoint(request, provider, tool_proxy)
            async for event in generate_agent_sse(response):
                yield event
        except HTTPException as error:
            yield {
                "event": "error",
                "data": json.dumps({"message": str(error.detail)}),
            }
            yield {"event": "done", "data": "{}"}
        except Exception:
            yield {
                "event": "error",
                "data": json.dumps({"message": "Agent request failed"}),
            }
            yield {"event": "done", "data": "{}"}

    return EventSourceResponse(events())

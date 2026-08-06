import json
from collections.abc import AsyncGenerator
from typing import Literal, cast

import httpx
from fastapi import APIRouter, Depends, HTTPException
from pydantic import BaseModel, ConfigDict, Field
from sse_starlette.sse import EventSourceResponse

from src.api.streaming import generate_agent_sse
from src.core.config import settings
from src.core.prompts import PromptManager
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


class GreetingRequest(BaseModel):
    locale: str = Field(default="en", max_length=20)
    display_name: str | None = Field(default=None, max_length=100)
    last_active_at: str | None = Field(default=None, max_length=50)
    product_name: str | None = Field(default=None, max_length=255)
    recent_interests: list[str] = Field(default_factory=list, max_length=5)


class GreetingDraft(BaseModel):
    model_config = ConfigDict(extra="forbid")
    message: str = Field(min_length=1, max_length=1000)


def get_provider() -> OpenAIProvider:
    return OpenAIProvider(
        api_key=settings.openai_api_key.get_secret_value(),
        base_url=settings.openai_base_url,
    )


def get_tool_proxy() -> ToolProxy:
    return ToolProxy(settings.backend_url)


def fallback_greeting(request: GreetingRequest) -> str:
    name = f" {request.display_name.strip()}" if request.display_name else ""
    if request.locale.lower().startswith("vi"):
        return f"Chào{name}! 👋 Hôm nay mình có thể giúp gì cho bạn?"
    return f"Hey{name}! 👋 What can I help you with today?"


@router.post("/greeting")
async def greeting_endpoint(
    request: GreetingRequest,
    provider: OpenAIProvider = Depends(get_provider),
) -> dict[str, str]:
    facts = request.model_dump(exclude={"locale"}, exclude_none=True)
    try:
        raw_response = await provider.chat_completion(
            messages=[
                {
                    "role": "system",
                    "content": PromptManager.get_greeting_prompt(
                        json.dumps(facts, ensure_ascii=False), request.locale
                    ),
                }
            ],
            model=settings.model,
            session_id="greeting",
            response_format={
                "type": "json_schema",
                "json_schema": {
                    "name": "GreetingDraft",
                    "schema": GreetingDraft.model_json_schema(),
                    "strict": True,
                },
            },
        )
        return {"message": GreetingDraft.model_validate_json(raw_response).message}
    except Exception:
        return {"message": fallback_greeting(request)}


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

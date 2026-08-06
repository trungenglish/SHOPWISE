import asyncio
from collections.abc import AsyncGenerator, Awaitable, Callable, Iterable
from typing import Any, TypeVar, cast

from openai import APIConnectionError, APIError, AsyncOpenAI, RateLimitError
from openai.types.chat import ChatCompletionMessageParam

from src.core.observability import log_error, log_interaction
from src.llm.provider import LLMProvider

Result = TypeVar("Result")


class OpenAIProvider(LLMProvider):
    def __init__(
        self,
        api_key: str,
        base_url: str | None = None,
        client: AsyncOpenAI | None = None,
    ) -> None:
        self.client = client or AsyncOpenAI(api_key=api_key, base_url=base_url)

    async def _execute_with_retry(
        self,
        operation: Callable[[], Awaitable[Result]],
        session_id: str,
        max_retries: int = 3,
    ) -> Result:
        """Execute a coroutine with exponential backoff for transient errors."""
        base_delay = 1.0
        for attempt in range(max_retries):
            try:
                return await operation()
            except (APIConnectionError, RateLimitError, APIError) as error:
                log_error(session_id, "openai", type(error).__name__)
                if attempt == max_retries - 1:
                    raise
                await asyncio.sleep(base_delay * (2**attempt))
        raise RuntimeError("OpenAI retry loop ended unexpectedly")

    async def chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float | None = None,
        max_tokens: int | None = None,
        session_id: str = "unknown",
        response_format: dict[str, Any] | None = None,
    ) -> str:
        async def call_openai() -> str:
            kwargs: dict[str, Any] = {
                "model": model,
                "messages": cast(Iterable[ChatCompletionMessageParam], messages),
            }
            if temperature is not None:
                kwargs["temperature"] = temperature
            if max_tokens is not None:
                kwargs["max_completion_tokens"] = max_tokens
            if response_format:
                kwargs["response_format"] = response_format

            response = cast(Any, await self.client.chat.completions.create(**kwargs))
            return response.choices[0].message.content or ""

        result = await self._execute_with_retry(call_openai, session_id)
        log_interaction(session_id, "openai", event="chat_completion")
        return result

    async def stream_chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float | None = None,
        max_tokens: int | None = None,
        session_id: str = "unknown",
        response_format: dict[str, Any] | None = None,
    ) -> AsyncGenerator[str, None]:
        async def call_openai() -> Any:
            kwargs: dict[str, Any] = {
                "model": model,
                "messages": cast(Iterable[ChatCompletionMessageParam], messages),
                "stream": True,
            }
            if temperature is not None:
                kwargs["temperature"] = temperature
            if max_tokens is not None:
                kwargs["max_completion_tokens"] = max_tokens
            if response_format:
                kwargs["response_format"] = response_format

            return cast(Any, await self.client.chat.completions.create(**kwargs))

        stream = await self._execute_with_retry(call_openai, session_id)
        log_interaction(session_id, "openai", event="stream_chat_completion_start")

        async for chunk in stream:
            if chunk.choices:
                delta = chunk.choices[0].delta.content
                if delta:
                    yield delta

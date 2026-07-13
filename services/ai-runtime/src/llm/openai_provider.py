import asyncio
from typing import Any, AsyncGenerator, cast, Iterable
from openai import AsyncOpenAI, APIError, APIConnectionError, RateLimitError
from openai.types.chat import ChatCompletionMessageParam
from src.llm.provider import LLMProvider
from src.core.observability import log_interaction, log_error

class OpenAIProvider(LLMProvider):
    def __init__(self, api_key: str, base_url: str | None = None, client: AsyncOpenAI | None = None):
        if client:
            self.client = client
        else:
            self.client = AsyncOpenAI(api_key=api_key, base_url=base_url)
            
    async def _execute_with_retry(self, coro, session_id: str, max_retries: int = 3):
        """Execute a coroutine with exponential backoff for transient errors."""
        base_delay = 1.0
        for attempt in range(max_retries):
            try:
                return await coro()
            except (APIConnectionError, RateLimitError, APIError) as e:
                log_error(session_id, "openai", type(e).__name__, str(e))
                if attempt == max_retries - 1:
                    raise
                # Exponential backoff: 1s, 2s, 4s...
                await asyncio.sleep(base_delay * (2 ** attempt))
                
    async def chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float = 0.7,
        max_tokens: int | None = None,
        session_id: str = "unknown",
        response_format: dict[str, Any] | None = None
    ) -> str:
        async def _call():
            kwargs: dict[str, Any] = {
                "model": model,
                "messages": cast(Iterable[ChatCompletionMessageParam], messages),
                "temperature": temperature,
            }
            if max_tokens:
                kwargs["max_tokens"] = max_tokens
            if response_format:
                kwargs["response_format"] = response_format
                
            response = cast(Any, await self.client.chat.completions.create(**kwargs))
            return response.choices[0].message.content or ""
            
        result = await self._execute_with_retry(_call, session_id)
        # Log interaction without PII
        log_interaction(session_id, "openai", event="chat_completion")
        return result
        
    async def stream_chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float = 0.7,
        max_tokens: int | None = None,
        session_id: str = "unknown",
        response_format: dict[str, Any] | None = None
    ) -> AsyncGenerator[str, None]:
        async def _call():
            kwargs: dict[str, Any] = {
                "model": model,
                "messages": cast(Iterable[ChatCompletionMessageParam], messages),
                "temperature": temperature,
                "stream": True,
            }
            if max_tokens:
                kwargs["max_tokens"] = max_tokens
            if response_format:
                kwargs["response_format"] = response_format
                
            response_stream = cast(Any, await self.client.chat.completions.create(**kwargs))
            return response_stream
            
        stream = await self._execute_with_retry(_call, session_id)
        
        # Log start of stream interaction
        log_interaction(session_id, "openai", event="stream_chat_completion_start")
        
        async for chunk in stream:
            if chunk.choices and len(chunk.choices) > 0:
                delta = chunk.choices[0].delta.content
                if delta:
                    yield delta

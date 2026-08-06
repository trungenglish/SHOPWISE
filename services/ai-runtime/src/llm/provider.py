from abc import ABC, abstractmethod
from collections.abc import AsyncGenerator
from typing import Any


class LLMProvider(ABC):
    """
    Base interface for all LLM providers.
    Ensures the AI Runtime can hot-swap providers without changing business logic.
    """

    @abstractmethod
    async def chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float | None = None,
        max_tokens: int | None = None,
        session_id: str = "unknown",
        response_format: dict[str, Any] | None = None,
    ) -> str:
        """Execute a blocking chat completion."""
        raise NotImplementedError

    @abstractmethod
    def stream_chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float | None = None,
        max_tokens: int | None = None,
        session_id: str = "unknown",
        response_format: dict[str, Any] | None = None,
    ) -> AsyncGenerator[str, None]:
        """Execute a streaming chat completion yielding string chunks."""
        raise NotImplementedError

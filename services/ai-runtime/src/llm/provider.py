from abc import ABC, abstractmethod
from typing import Any, AsyncGenerator

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
        temperature: float = 0.7,
        max_tokens: int | None = None,
    ) -> str:
        """Execute a blocking chat completion."""
        pass
        
    @abstractmethod
    async def stream_chat_completion(
        self,
        messages: list[dict[str, Any]],
        model: str,
        temperature: float = 0.7,
        max_tokens: int | None = None,
    ) -> AsyncGenerator[str, None]:
        """Execute a streaming chat completion yielding string chunks."""
        pass

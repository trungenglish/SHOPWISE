from typing import Literal

from pydantic import BaseModel


class AIErrorState(BaseModel):
    """
    Structured error response returned via Dynamic UI Protocol when the AI Runtime fails.
    """

    error: bool = True
    type: Literal["provider_error", "validation_error", "rate_limit", "timeout", "internal_error"]
    message: str
    retryable: bool = False

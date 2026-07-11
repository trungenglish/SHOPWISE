from pydantic import BaseModel, Field
from typing import Any

class BaseStructuredOutput(BaseModel):
    """
    Base model for all structured LLM responses.
    """
    message: str = Field(..., description="Message for the user")
    tool_calls: list[dict[str, Any]] = Field(default_factory=list, description="Optional tools to invoke")

class RetailAction(BaseModel):
    """
    Example schema for specific retail actions that might be returned.
    """
    action_name: str
    parameters: dict[str, Any]

from pydantic import BaseModel, Field
from typing import Any, List, Optional

class BaseStructuredOutput(BaseModel):
    """
    Base model for all structured LLM responses.
    """
    message: str = Field(..., description="Message for the user")
    tool_calls: list[dict[str, Any]] = Field(default_factory=list, description="Optional tools to invoke")

class ProductSpecs(BaseModel):
    gpu: str
    ram: str
    cooling: str
    cpu: str
    screen: str
    warranty: str

class Product(BaseModel):
    id: str
    name: str
    price: int
    image: str
    matchScore: int
    specs: ProductSpecs
    aiPerf: int
    rendering: int
    thermals: int
    matchExplanation: str

class AuditLog(BaseModel):
    time: str
    message: str
    status: str = Field(description="Can be 'done', 'running', or 'pending'")

class AgentStatus(BaseModel):
    id: str
    name: str
    progress: int
    statusMessage: str
    type: str

class Accessory(BaseModel):
    name: str
    price: int
    category: str
    reason: str
    image: str

class DecisionResponse(BaseModel):
    products: List[Product]
    logs: List[AuditLog]
    agents: List[AgentStatus]
    accessories: List[Accessory]
    reasoning: str
    trustScore: int

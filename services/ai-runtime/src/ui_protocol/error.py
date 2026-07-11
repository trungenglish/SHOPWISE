from typing import Optional, List, Literal
from pydantic import BaseModel

class ErrorStateProps(BaseModel):
    title: str
    message: str
    code: Optional[str] = None
    severity: Optional[Literal["info", "warning", "error", "critical"]] = None

class RetryActionProps(BaseModel):
    label: Optional[str] = None
    actionId: Optional[str] = None

class ContinueCachedDataProps(BaseModel):
    label: Optional[str] = None

class ModifySearchProps(BaseModel):
    label: Optional[str] = None
    originalQuery: Optional[str] = None

class AlternativeRecommendationProps(BaseModel):
    message: Optional[str] = None
    productIds: Optional[List[str]] = None

from datetime import datetime
from typing import Optional, List, Dict, Any, Literal
from pydantic import BaseModel, Field

class ValidationRules(BaseModel):
    required: Optional[bool] = None
    type: Optional[str] = None
    min: Optional[float] = None
    max: Optional[float] = None
    minLength: Optional[int] = None
    maxLength: Optional[int] = None
    pattern: Optional[str] = None

class ActionDefinition(BaseModel):
    trigger: str
    action: str
    payloadSchema: Optional[Dict[str, Any]] = None

class ComponentNode(BaseModel):
    id: str
    type: str
    props: Dict[str, Any]
    state: Optional[Dict[str, Any]] = None
    validation: Optional[ValidationRules] = None
    children: Optional[List['ComponentNode']] = None
    actions: Optional[List[ActionDefinition]] = None

class DynamicUIDocument(BaseModel):
    protocol: Literal["dynamic-ui"] = "dynamic-ui"
    schemaVersion: Literal["1.0"] = "1.0"
    documentId: str
    timestamp: datetime
    operation: Literal["replace", "append", "patch", "remove"]
    targetId: Optional[str] = None
    metadata: Optional[Dict[str, Any]] = None
    root: ComponentNode

class InteractionEvent(BaseModel):
    componentId: str
    action: str
    payload: Dict[str, Any]
    metadata: Optional[Dict[str, Any]] = None

ComponentNode.model_rebuild()

from typing import Optional, List, Literal
from pydantic import BaseModel, Field

class ButtonProps(BaseModel):
    text: str
    variant: Optional[Literal["primary", "secondary", "danger", "outline"]] = None

class QuickReplyProps(BaseModel):
    label: str

class SelectOption(BaseModel):
    label: Optional[str] = None
    value: Optional[str] = None

class SelectProps(BaseModel):
    options: Optional[List[SelectOption]] = None
    placeholder: Optional[str] = None

class RadioGroupProps(BaseModel):
    options: Optional[List[SelectOption]] = None
    name: Optional[str] = None

class CheckboxProps(BaseModel):
    label: Optional[str] = None
    checked: Optional[bool] = None

class TextInputProps(BaseModel):
    placeholder: Optional[str] = None
    defaultValue: Optional[str] = None
    type: Optional[Literal["text", "email", "password", "tel"]] = None

class NumberInputProps(BaseModel):
    placeholder: Optional[str] = None
    defaultValue: Optional[float] = None

class BudgetSliderProps(BaseModel):
    minPriceVND: int = Field(..., ge=0)
    maxPriceVND: int = Field(..., ge=0)
    step: Optional[int] = Field(None, ge=1)

class Store(BaseModel):
    storeId: Optional[str] = None
    name: Optional[str] = None
    address: Optional[str] = None
    distanceKm: Optional[float] = None

class StorePickerProps(BaseModel):
    stores: Optional[List[Store]] = None

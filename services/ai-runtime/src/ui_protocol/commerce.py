from typing import Optional, List, Dict, Any, Literal
from pydantic import BaseModel, Field

# AI Components

class ChatMessageProps(BaseModel):
    text: str
    sender: Literal["user", "ai", "system"]

class ThinkingIndicatorProps(BaseModel):
    statusText: Optional[str] = None

class RecommendationExplanationProps(BaseModel):
    text: Optional[str] = None
    highlights: Optional[List[str]] = None

class DecisionReasoningProps(BaseModel):
    reasoning: Optional[str] = None

class ConfidenceIndicatorProps(BaseModel):
    score: Optional[float] = Field(None, ge=0, le=100)
    label: Optional[str] = None


# Commerce Components

class ProductCardProps(BaseModel):
    productId: str
    name: str
    priceVND: int = Field(..., ge=0)
    imageUrl: Optional[str] = None
    rating: Optional[float] = None

class ProductCarouselProps(BaseModel):
    title: Optional[str] = None

class ProductComparisonProps(BaseModel):
    productIds: Optional[List[str]] = None
    features: Optional[List[str]] = None

class SpecificationTableProps(BaseModel):
    specs: Optional[Dict[str, str]] = None

class PromotionBannerProps(BaseModel):
    text: Optional[str] = None
    discountCode: Optional[str] = None

class WarrantyInformationProps(BaseModel):
    months: Optional[int] = None
    details: Optional[str] = None

class InventoryStatusProps(BaseModel):
    inStock: Optional[bool] = None
    quantity: Optional[int] = None
    storeId: Optional[str] = None

class CheckoutSummaryItem(BaseModel):
    productId: Optional[str] = None
    quantity: Optional[int] = None
    priceVND: Optional[int] = Field(None, ge=0)

class CheckoutSummaryProps(BaseModel):
    items: Optional[List[CheckoutSummaryItem]] = None
    totalPriceVND: Optional[int] = Field(None, ge=0)


# Commerce Actions

class AddToComparisonPayload(BaseModel):
    productId: str

class RemoveProductPayload(BaseModel):
    productId: str

class SaveDecisionPayload(BaseModel):
    decisionId: str
    productId: str

class ResumeConversationPayload(BaseModel):
    conversationId: str

class CheckoutReadinessPayload(BaseModel):
    cartId: Optional[str] = None
    ready: bool

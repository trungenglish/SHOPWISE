from typing import Annotated, Any, Literal

from pydantic import AliasChoices, BaseModel, ConfigDict, Field, RootModel, model_validator


class CatalogProduct(BaseModel):
    id: str = Field(validation_alias=AliasChoices("ID", "id"))
    name: str = Field(validation_alias=AliasChoices("Name", "name"))
    price: int = Field(ge=0, validation_alias=AliasChoices("Price", "price"))
    specifications: dict[str, Any] = Field(
        default_factory=dict,
        validation_alias=AliasChoices("Specifications", "specifications"),
    )
    metadata: dict[str, Any] = Field(
        default_factory=dict,
        validation_alias=AliasChoices("Metadata", "metadata"),
    )


class StrictDraftModel(BaseModel):
    model_config = ConfigDict(extra="forbid")


class RecommendationSelection(StrictDraftModel):
    product_id: str
    match_score: int = Field(ge=0, le=100)
    explanation: str


class AgentDraft(StrictDraftModel):
    type: Literal["question", "recommendation", "comparison", "checkout_ready"]
    message: str
    reasoning: str | None
    selections: list[RecommendationSelection] | None

    @model_validator(mode="after")
    def validate_shape(self) -> "AgentDraft":
        if self.type == "question":
            if self.reasoning is not None or self.selections is not None:
                raise ValueError("question must use null reasoning and selections")
            return self
        if not self.reasoning or not self.selections:
            raise ValueError(f"{self.type} requires reasoning and selections")
        if self.type == "comparison" and len(self.selections) < 2:
            raise ValueError("comparison requires at least two selections")
        if self.type == "checkout_ready" and len(self.selections) != 1:
            raise ValueError("checkout_ready requires exactly one selection")
        return self


class RecommendedProduct(BaseModel):
    id: str
    name: str
    price: int
    image: str
    specifications: dict[str, Any]
    match_score: int
    match_explanation: str


class RecommendationDecision(BaseModel):
    products: list[RecommendedProduct]
    reasoning: str


class QuestionResponse(BaseModel):
    type: Literal["question"]
    message: str


class RecommendationResponse(BaseModel):
    type: Literal["recommendation"]
    message: str
    decision: RecommendationDecision


class ComparisonResponse(BaseModel):
    type: Literal["comparison"]
    message: str
    decision: RecommendationDecision


class CheckoutReadyResponse(BaseModel):
    type: Literal["checkout_ready"]
    message: str
    decision: RecommendationDecision


Response = Annotated[
    QuestionResponse | RecommendationResponse | ComparisonResponse | CheckoutReadyResponse,
    Field(discriminator="type"),
]


class AgentResponse(RootModel[Response]):
    pass


def hydrate_agent_response(
    draft: AgentDraft,
    catalog: list[CatalogProduct],
    allowed_comparison_ids: set[str] | None = None,
) -> AgentResponse:
    if draft.type == "question":
        return AgentResponse(root=QuestionResponse(type="question", message=draft.message))

    products_by_id = {product.id: product for product in catalog}
    hydrated_products: list[RecommendedProduct] = []
    for selection in draft.selections or []:
        is_session_only_action = draft.type in {"comparison", "checkout_ready"}
        if is_session_only_action and selection.product_id not in (allowed_comparison_ids or set()):
            raise ValueError(f"product is not available in this session: {selection.product_id}")
        product = products_by_id.get(selection.product_id)
        if product is None:
            raise ValueError(f"unknown catalog product: {selection.product_id}")
        hydrated_products.append(
            RecommendedProduct(
                id=product.id,
                name=product.name,
                price=product.price,
                image=str(product.metadata.get("image_url", "")),
                specifications=product.specifications,
                match_score=selection.match_score,
                match_explanation=selection.explanation,
            )
        )

    decision = RecommendationDecision(
        products=hydrated_products,
        reasoning=draft.reasoning or "",
    )
    if draft.type == "comparison":
        return AgentResponse(
            root=ComparisonResponse(
                type="comparison",
                message=draft.message,
                decision=decision,
            )
        )

    if draft.type == "checkout_ready":
        return AgentResponse(
            root=CheckoutReadyResponse(
                type="checkout_ready",
                message=draft.message,
                decision=decision,
            )
        )

    return AgentResponse(
        root=RecommendationResponse(
            type="recommendation",
            message=draft.message,
            decision=decision,
        )
    )

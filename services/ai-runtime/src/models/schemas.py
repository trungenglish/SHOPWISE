from datetime import datetime
from typing import Annotated, Any, Literal
from uuid import uuid4

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


class OfferLine(BaseModel):
    retailer_offer_id: str
    name: str
    original_price: int = Field(ge=0)
    price: int = Field(ge=0)
    required: bool
    default_selected: bool


class ScheduledCampaign(BaseModel):
    name: str
    starts_at: datetime
    ends_at: datetime
    discount_percent: int = Field(gt=0, le=100)
    sale_price: int = Field(ge=0)
    savings: int = Field(ge=0)


class OfferComparison(BaseModel):
    id: str
    product_id: str
    currency: Literal["VND"] = "VND"
    lines: list[OfferLine]
    default_total: int = Field(ge=0)
    base_total: int = Field(ge=0)
    scheduled_campaign: ScheduledCampaign


class StrictDraftModel(BaseModel):
    model_config = ConfigDict(extra="forbid")


class RecommendationSelection(StrictDraftModel):
    product_id: str
    match_score: int = Field(ge=0, le=100)
    explanation: str


class QuestionDraft(StrictDraftModel):
    mode: Literal["single", "multiple"]
    options: list[str] = Field(min_length=2, max_length=4)
    free_text_allowed: bool
    input_label: str
    input_placeholder: str
    submit_label: str


class AgentDraft(StrictDraftModel):
    type: Literal[
        "question", "recommendation", "comparison", "offer_comparison", "checkout_ready"
    ]
    message: str
    reasoning: str | None
    selections: list[RecommendationSelection] | None
    question: QuestionDraft | None = None

    @model_validator(mode="after")
    def validate_shape(self) -> "AgentDraft":
        if self.type == "question":
            if self.reasoning is not None or self.selections is not None:
                raise ValueError("question must use null reasoning and selections")
            if self.question is None:
                raise ValueError("question controls are required")
            return self
        if self.question is not None:
            raise ValueError(f"{self.type} must use null question controls")
        if not self.reasoning or not self.selections:
            raise ValueError(f"{self.type} requires reasoning and selections")
        if self.type == "comparison" and len(self.selections) < 2:
            raise ValueError("comparison requires at least two selections")
        if self.type == "offer_comparison" and len(self.selections) != 1:
            raise ValueError("offer_comparison requires exactly one selection")
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


class ClarificationOption(BaseModel):
    id: str
    label: str


class ClarificationQuestion(BaseModel):
    id: str
    mode: Literal["single", "multiple"] = "single"
    options: list[ClarificationOption] = Field(min_length=2, max_length=4)
    free_text_allowed: bool = True
    input_label: str = "Your answer"
    input_placeholder: str = "Add details in your own words"
    submit_label: str = "Submit"


class DynamicUIComponent(BaseModel):
    id: str
    type: str
    props: dict[str, Any] = Field(default_factory=dict)
    children: list["DynamicUIComponent"] = Field(default_factory=list)
    interaction: dict[str, Any] | None = None


class UIOperation(BaseModel):
    id: str
    sequence: int = Field(ge=0)
    operation: Literal["append", "replace"]
    target: Literal["surface", "component"]
    target_id: str | None = None
    component: DynamicUIComponent


class DynamicUIResponse(BaseModel):
    schema_version: Literal["1.1"] = "1.1"
    turn_id: str = Field(default_factory=lambda: str(uuid4()))
    revision: int = 1
    conversation_state: Literal[
        "collecting_requirements",
        "recommending",
        "comparing",
        "offering",
        "checkout_ready",
        "error",
    ]
    ui_state: dict[str, Any] = Field(default_factory=lambda: {"status": "ready"})
    ui_operations: list[UIOperation] = Field(default_factory=list)


class QuestionResponse(DynamicUIResponse):
    type: Literal["question"]
    message: str
    conversation_state: Literal["collecting_requirements"] = "collecting_requirements"
    question: ClarificationQuestion | None = None


class RecommendationResponse(DynamicUIResponse):
    type: Literal["recommendation"]
    message: str
    decision: RecommendationDecision
    conversation_state: Literal["recommending"] = "recommending"


class ComparisonResponse(DynamicUIResponse):
    type: Literal["comparison"]
    message: str
    decision: RecommendationDecision
    conversation_state: Literal["comparing"] = "comparing"


class OfferComparisonResponse(DynamicUIResponse):
    type: Literal["offer_comparison"]
    message: str
    decision: RecommendationDecision
    offer: OfferComparison
    conversation_state: Literal["offering"] = "offering"


class CheckoutReadyResponse(DynamicUIResponse):
    type: Literal["checkout_ready"]
    message: str
    decision: RecommendationDecision
    conversation_state: Literal["checkout_ready"] = "checkout_ready"


Response = Annotated[
    QuestionResponse
    | RecommendationResponse
    | ComparisonResponse
    | OfferComparisonResponse
    | CheckoutReadyResponse,
    Field(discriminator="type"),
]


class AgentResponse(RootModel[Response]):
    pass


def _question_response(message: str, draft: QuestionDraft) -> QuestionResponse:
    question_id = str(uuid4())
    question = ClarificationQuestion(
        id=question_id,
        mode=draft.mode,
        options=[
            ClarificationOption(id=str(uuid4()), label=label) for label in draft.options
        ],
        free_text_allowed=draft.free_text_allowed,
        input_label=draft.input_label,
        input_placeholder=draft.input_placeholder,
        submit_label=draft.submit_label,
    )
    options = question.options
    option_components = [
        DynamicUIComponent(
            id=str(uuid4()),
            type="radio_group",
            props={"options": [option.model_dump() for option in options], "name": question_id},
        ),
    ]
    if question.free_text_allowed:
        option_components.append(
            DynamicUIComponent(
                id=str(uuid4()),
                type="text_input",
                props={"label": question.input_label, "name": "free_text"},
            )
        )
    option_components.append(
        DynamicUIComponent(
            id=str(uuid4()),
            type="button",
            props={"label": question.submit_label},
            interaction={
                "event": "submit",
                "action": "question.answer",
                "payload": {"question_id": question_id},
                "binding": question_id,
                "state": {"disabled": False},
            },
        )
    )
    card = DynamicUIComponent(
        id=question_id,
        type="card",
        props={"title": message, "active": True},
        children=option_components,
    )
    return QuestionResponse(
        type="question",
        message=message,
        question=question,
        ui_operations=[
            UIOperation(
                id=str(uuid4()),
                sequence=0,
                operation="append",
                target="surface",
                component=card,
            )
        ],
    )


def _decision_ui(
    response_type: Literal[
        "recommendation", "comparison", "offer_comparison", "checkout_ready"
    ],
    message: str,
    decision: RecommendationDecision,
) -> list[UIOperation]:
    component_type = {
        "recommendation": "product_carousel",
        "comparison": "product_comparison",
        "offer_comparison": "promotion_banner",
        "checkout_ready": "checkout_summary",
    }[response_type]
    component = DynamicUIComponent(
        id=str(uuid4()),
        type=component_type,
        props={
            "message": message,
            "reasoning": decision.reasoning,
            "products": [product.model_dump() for product in decision.products],
        },
    )
    return [
        UIOperation(
            id=str(uuid4()),
            sequence=0,
            operation="replace",
            target="surface",
            component=component,
        )
    ]


def hydrate_agent_response(
    draft: AgentDraft,
    catalog: list[CatalogProduct],
    allowed_comparison_ids: set[str] | None = None,
    language_source: str = "",
    offers: dict[str, OfferComparison] | None = None,
) -> AgentResponse:
    if draft.type == "question":
        if draft.question is None:
            raise ValueError("question controls are required")
        return AgentResponse(root=_question_response(draft.message, draft.question))

    products_by_id = {product.id: product for product in catalog}
    hydrated_products: list[RecommendedProduct] = []
    for selection in draft.selections or []:
        is_session_only_action = draft.type in {
            "comparison",
            "offer_comparison",
            "checkout_ready",
        }
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
                ui_operations=_decision_ui("comparison", draft.message, decision),
            )
        )

    if draft.type == "offer_comparison":
        product_id = hydrated_products[0].id
        offer = (offers or {}).get(product_id)
        if offer is None:
            raise ValueError(f"offer is not available for product: {product_id}")
        return AgentResponse(
            root=OfferComparisonResponse(
                type="offer_comparison",
                message=draft.message,
                decision=decision,
                offer=offer,
                ui_operations=_decision_ui("offer_comparison", draft.message, decision),
            )
        )

    if draft.type == "checkout_ready":
        return AgentResponse(
            root=CheckoutReadyResponse(
                type="checkout_ready",
                message=draft.message,
                decision=decision,
                ui_operations=_decision_ui("checkout_ready", draft.message, decision),
            )
        )

    return AgentResponse(
        root=RecommendationResponse(
            type="recommendation",
            message=draft.message,
            decision=decision,
            ui_operations=_decision_ui("recommendation", draft.message, decision),
        )
    )

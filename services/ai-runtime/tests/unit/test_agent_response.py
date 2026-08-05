import pytest

from src.models.schemas import (
    AgentDraft,
    CatalogProduct,
    OfferComparison,
    hydrate_agent_response,
)


def test_agent_draft_schema_is_strict_for_openai():
    schema = AgentDraft.model_json_schema()
    object_schemas: list[dict[str, object]] = []

    def collect(value: object) -> None:
        if isinstance(value, dict):
            if value.get("type") == "object":
                object_schemas.append(value)
            for child in value.values():
                collect(child)
        elif isinstance(value, list):
            for child in value:
                collect(child)

    collect(schema)

    assert object_schemas
    assert all(item.get("additionalProperties") is False for item in object_schemas)
    assert "oneOf" not in str(schema)
    assert schema.get("type") == "object"


def test_question_draft_becomes_question_response():
    draft = AgentDraft.model_validate(
        {
            "type": "question",
            "message": "What's your gaming style?",
            "reasoning": None,
            "selections": None,
            "question": {
                "mode": "single",
                "options": ["AAA story games", "Competitive esports", "Casual games"],
                "free_text_allowed": True,
                "input_label": "Your games",
                "input_placeholder": "e.g. Wuthering Waves",
                "submit_label": "Continue",
            },
        }
    )

    response = hydrate_agent_response(draft, [])

    assert response.root.type == "question"
    assert response.root.message == "What's your gaming style?"
    assert response.root.schema_version == "1.1"
    assert response.root.conversation_state == "collecting_requirements"
    assert response.root.question is not None
    assert len(response.root.question.options) == 3
    assert response.root.question.options[0].label == "AAA story games"
    assert response.root.question.input_placeholder == "e.g. Wuthering Waves"
    assert response.root.ui_operations[0].component.type == "card"


def test_question_requires_structured_controls():
    with pytest.raises(ValueError, match="question controls"):
        AgentDraft.model_validate(
            {
                "type": "question",
                "message": "What is your budget?",
                "reasoning": None,
                "selections": None,
            }
        )


def test_question_controls_follow_the_user_language():
    draft = AgentDraft.model_validate(
        {
            "type": "question",
            "message": "Ngân sách của bạn là bao nhiêu?",
            "reasoning": None,
            "selections": None,
            "question": {
                "mode": "single",
                "options": ["Dưới 25 triệu VND", "25–40 triệu VND"],
                "free_text_allowed": True,
                "input_label": "Câu trả lời của bạn",
                "input_placeholder": "Nhập ngân sách",
                "submit_label": "Gửi",
            },
        }
    )

    response = hydrate_agent_response(draft, [], language_source="Tôi muốn tìm laptop chơi game")

    assert response.root.question is not None
    assert response.root.question.submit_label == "Gửi"
    assert response.root.question.options[0].label == "Dưới 25 triệu VND"


def test_recommendation_uses_only_authoritative_catalog_data():
    catalog = [
        CatalogProduct.model_validate(
            {
                "ID": "product-1",
                "Name": "Catalog Laptop",
                "Price": 37_475_000,
                "Specifications": {"cpu": "Core Ultra 9", "gpu": "RTX 5070"},
                "Metadata": {"image_url": "https://example.com/laptop.png"},
            }
        )
    ]
    draft = AgentDraft.model_validate(
        {
            "type": "recommendation",
            "message": "This matches your needs.",
            "reasoning": "Strong GPU within budget.",
            "selections": [
                {
                    "product_id": "product-1",
                    "match_score": 91,
                    "explanation": "Best fit for gaming.",
                }
            ],
        }
    )

    response = hydrate_agent_response(draft, catalog)
    product = response.root.decision.products[0]

    assert product.name == "Catalog Laptop"
    assert product.price == 37_475_000
    assert product.match_explanation == "Best fit for gaming."


def test_recommendation_rejects_unknown_product_id():
    draft = AgentDraft.model_validate(
        {
            "type": "recommendation",
            "message": "Try this.",
            "reasoning": "Looks suitable.",
            "selections": [
                {
                    "product_id": "invented-id",
                    "match_score": 90,
                    "explanation": "Invented product.",
                }
            ],
        }
    )

    with pytest.raises(ValueError, match="invented-id"):
        hydrate_agent_response(draft, [])


def test_comparison_rejects_catalog_product_not_in_session():
    catalog = [
        CatalogProduct.model_validate({"ID": product_id, "Name": product_id, "Price": 25_000_000})
        for product_id in ("product-1", "product-2")
    ]
    draft = AgentDraft.model_validate(
        {
            "type": "comparison",
            "message": "Here is the comparison.",
            "reasoning": "Product 2 is faster.",
            "selections": [
                {
                    "product_id": product_id,
                    "match_score": 90,
                    "explanation": "Session comparison.",
                }
                for product_id in ("product-1", "product-2")
            ],
        }
    )

    with pytest.raises(ValueError, match="product-2"):
        hydrate_agent_response(draft, catalog, allowed_comparison_ids={"product-1"})


def test_offer_comparison_uses_only_authoritative_offer_data():
    catalog = [
        CatalogProduct.model_validate(
            {"ID": "product-1", "Name": "ROG Strix G16", "Price": 42_000_000}
        )
    ]
    offer = OfferComparison.model_validate(
        {
            "id": "rog-bundle",
            "product_id": "product-1",
            "currency": "VND",
            "lines": [
                {
                    "retailer_offer_id": "mouse-offer",
                    "name": "Gaming Mouse",
                    "original_price": 1_500_000,
                    "price": 0,
                    "required": True,
                    "default_selected": True,
                }
            ],
            "default_total": 43_000_000,
            "base_total": 42_000_000,
            "scheduled_campaign": {
                "name": "Back to School",
                "starts_at": "2026-09-01T00:00:00Z",
                "ends_at": "2026-10-01T00:00:00Z",
                "discount_percent": 5,
                "sale_price": 39_900_000,
                "savings": 2_100_000,
            },
        }
    )
    draft = AgentDraft.model_validate(
        {
            "type": "offer_comparison",
            "message": "Today's complete setup beats waiting.",
            "reasoning": "The bundle adds useful gear now.",
            "selections": [
                {
                    "product_id": "product-1",
                    "match_score": 95,
                    "explanation": "The selected laptop has an active bundle.",
                }
            ],
            "question": None,
        }
    )

    response = hydrate_agent_response(
        draft,
        catalog,
        allowed_comparison_ids={"product-1"},
        offers={"product-1": offer},
    )

    assert response.root.type == "offer_comparison"
    assert response.root.offer.default_total == 43_000_000
    assert response.root.offer.scheduled_campaign.savings == 2_100_000


def test_offer_comparison_rejects_product_without_authoritative_offer():
    draft = AgentDraft.model_validate(
        {
            "type": "offer_comparison",
            "message": "Buy today.",
            "reasoning": "There is a deal.",
            "selections": [
                {"product_id": "product-1", "match_score": 90, "explanation": "Deal"}
            ],
            "question": None,
        }
    )

    with pytest.raises(ValueError, match="offer is not available"):
        hydrate_agent_response(
            draft,
            [CatalogProduct.model_validate({"ID": "product-1", "Name": "ROG", "Price": 1})],
            allowed_comparison_ids={"product-1"},
            offers={},
        )

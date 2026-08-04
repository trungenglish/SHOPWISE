import pytest

from src.models.schemas import (
    AgentDraft,
    CatalogProduct,
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
            "message": "What is your budget?",
            "reasoning": None,
            "selections": None,
        }
    )

    response = hydrate_agent_response(draft, [])

    assert response.root.type == "question"
    assert response.root.message == "What is your budget?"


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

import json

import pytest

from src.api.streaming import generate_agent_sse
from src.models.schemas import (
    AgentDraft,
    CatalogProduct,
    hydrate_agent_response,
)


@pytest.mark.asyncio
async def test_recommendation_sse_has_token_recommendation_and_done_events():
    response = hydrate_agent_response(
        AgentDraft.model_validate(
            {
                "type": "recommendation",
                "message": "Catalog match found.",
                "reasoning": "Strong GPU.",
                "selections": [
                    {
                        "product_id": "product-1",
                        "match_score": 92,
                        "explanation": "Fits the budget.",
                    }
                ],
            }
        ),
        [
            CatalogProduct(
                id="product-1",
                name="Catalog Laptop",
                price=37_475_000,
                specifications={"gpu": "RTX 5070"},
            )
        ],
    )

    events = [event async for event in generate_agent_sse(response)]

    assert [event["event"] for event in events] == [
        "token",
        "recommendation",
        "ui_operation",
        "envelope",
        "done",
    ]
    assert json.loads(events[0]["data"]) == {"text": "Catalog match found."}
    assert json.loads(events[1]["data"])["products"][0]["id"] == "product-1"
    assert "target_id" not in json.loads(events[2]["data"])
    assert json.loads(events[3]["data"])["schema_version"] == "1.1"
    assert "target_id" not in json.loads(events[3]["data"])["ui_operations"][0]

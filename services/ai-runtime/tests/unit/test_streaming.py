import json

import pytest

from src.api.streaming import generate_agent_sse
from src.models.schemas import (
    AgentResponse,
    RecommendationDecision,
    RecommendationResponse,
    RecommendedProduct,
)


@pytest.mark.asyncio
async def test_recommendation_sse_has_token_recommendation_and_done_events():
    response = AgentResponse(
        root=RecommendationResponse(
            type="recommendation",
            message="Catalog match found.",
            decision=RecommendationDecision(
                products=[
                    RecommendedProduct(
                        id="product-1",
                        name="Catalog Laptop",
                        price=37_475_000,
                        image="",
                        specifications={"gpu": "RTX 5070"},
                        match_score=92,
                        match_explanation="Fits the budget.",
                    )
                ],
                reasoning="Strong GPU.",
            ),
        )
    )

    events = [event async for event in generate_agent_sse(response)]

    assert [event["event"] for event in events] == [
        "token",
        "recommendation",
        "done",
    ]
    assert json.loads(events[0]["data"]) == {"text": "Catalog match found."}
    assert json.loads(events[1]["data"])["products"][0]["id"] == "product-1"

import pytest

from src.graph.workflow import create_workflow
from src.models.schemas import CatalogProduct, OfferComparison


class FakeProvider:
    def __init__(self, response: str) -> None:
        self.response = response
        self.model = ""
        self.messages: list[dict[str, str]] = []

    async def chat_completion(self, **kwargs) -> str:
        self.model = kwargs["model"]
        self.messages = kwargs["messages"]
        return self.response


class FakeToolProxy:
    async def catalog_search(self) -> list[CatalogProduct]:
        return [
            CatalogProduct.model_validate(
                {
                    "ID": "product-1",
                    "Name": "Catalog Laptop",
                    "Price": 37_475_000,
                    "Specifications": {"gpu": "RTX 5070"},
                    "Metadata": {},
                }
            )
        ]

    async def offer_comparison(self, product_id: str) -> OfferComparison:
        return OfferComparison.model_validate(
            {
                "id": "rog-bundle",
                "product_id": product_id,
                "currency": "VND",
                "lines": [],
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


@pytest.mark.asyncio
async def test_workflow_uses_configured_model_and_complete_history():
    provider = FakeProvider(
        '{"type":"question","message":"What games do you play?","reasoning":null,'
        '"selections":null,"question":{"mode":"single","options":'
        '["AAA games","Esports","Casual games"],"free_text_allowed":true,'
        '"input_label":"Your games","input_placeholder":"Name a game",'
        '"submit_label":"Continue"}}'
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")
    history = [
        {"role": "user", "content": "I need a laptop."},
        {"role": "assistant", "content": "What will you use it for?"},
        {"role": "user", "content": "Gaming."},
    ]

    result = await workflow.ainvoke(
        {"messages": history, "session_id": "session-1", "retry_count": 0}
    )

    assert result["response"].root.type == "question"
    assert result["response"].root.question.options[0].label == "AAA games"
    assert provider.model == "gpt-5.4-mini"
    assert provider.messages[-3:] == history


@pytest.mark.asyncio
async def test_workflow_hydrates_recommendation_from_catalog():
    provider = FakeProvider(
        """{
          "type":"recommendation",
          "message":"This is the best fit.",
          "reasoning":"Strong gaming GPU.",
          "selections":[{
            "product_id":"product-1",
            "match_score":92,
            "explanation":"Fits the budget and workload."
          }]
        }"""
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [{"role": "user", "content": "Gaming under 40m VND"}],
            "session_id": "session-1",
            "retry_count": 0,
        }
    )

    product = result["response"].root.decision.products[0]
    assert product.id == "product-1"
    assert product.price == 37_475_000


@pytest.mark.asyncio
async def test_workflow_hydrates_offer_from_backend_tool():
    provider = FakeProvider(
        '{"type":"offer_comparison","message":"Buy now or wait.",'
        '"reasoning":"The bundle is available today.","selections":'
        '[{"product_id":"product-1","match_score":95,"explanation":"Active offer"}],'
        '"question":null}'
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [{"role": "user", "content": "Should I wait for the sale?"}],
            "session_id": "session-1",
            "retry_count": 0,
            "allowed_comparison_ids": ["product-1"],
        }
    )

    assert result["response"].root.type == "offer_comparison"
    assert result["response"].root.offer.default_total == 43_000_000
    assert "43" in provider.messages[0]["content"]


@pytest.mark.asyncio
async def test_system_prompt_asks_gaming_workload_before_budget():
    """A gaming request must instruct the LLM to collect workload first, not budget."""
    provider = FakeProvider(
        '{"type":"question","message":"What kind of games do you play?",'
        '"reasoning":null,"selections":null,"question":{"mode":"single",'
        '"options":["AAA story games","Competitive esports","Casual / indie"],'
        '"free_text_allowed":true,"input_label":"Your games",'
        '"input_placeholder":"e.g. Wuthering Waves","submit_label":"Continue"}}'
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [{"role": "user", "content": "I want a gaming laptop."}],
            "session_id": "session-gaming-1",
            "retry_count": 0,
        }
    )

    assert result["response"].root.type == "question"
    system_prompt = provider.messages[0]["content"]
    # Prompt must list workload, ambition, and budget dimensions
    assert "workload" in system_prompt
    assert "ambition" in system_prompt
    assert "budget" in system_prompt
    # workload must appear before budget in the priority ordering
    assert system_prompt.index("workload") < system_prompt.index("budget")


@pytest.mark.asyncio
async def test_system_prompt_instructs_one_missing_dimension_per_turn():
    """System prompt must instruct the LLM to ask exactly one missing dimension."""
    provider = FakeProvider(
        '{"type":"question","message":"How competitive are you?",'
        '"reasoning":null,"selections":null,"question":{"mode":"single",'
        '"options":["Casual fun","Semi-pro","Full competitive"],'
        '"free_text_allowed":false,"input_label":"Your style",'
        '"input_placeholder":"Pick one","submit_label":"Continue"}}'
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [
                {"role": "user", "content": "I play AAA games."},
                {"role": "assistant", "content": "What gaming style do you have?"},
            ],
            "session_id": "session-dim-1",
            "retry_count": 0,
        }
    )

    assert result["response"].root.type == "question"
    system_prompt = provider.messages[0]["content"]
    # Prompt must say to ask exactly one missing dimension at a time
    assert "one missing dimension" in system_prompt or "exactly one" in system_prompt


@pytest.mark.asyncio
async def test_system_prompt_treats_flexible_as_complete_dimension():
    """'No preference' / 'flexible' must count as a satisfied dimension."""
    provider = FakeProvider(
        '{"type":"question","message":"What is your budget?",'
        '"reasoning":null,"selections":null,"question":{"mode":"single",'
        '"options":["Under 25M VND","25-40M VND","Over 40M VND"],'
        '"free_text_allowed":true,"input_label":"Budget",'
        '"input_placeholder":"Enter your budget","submit_label":"Continue"}}'
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [
                {"role": "user", "content": "Gaming laptop for competitive esports."},
                {"role": "assistant", "content": "How serious are you?"},
                {"role": "user", "content": "I have no preference, just casual fun."},
            ],
            "session_id": "session-flex-1",
            "retry_count": 0,
        }
    )

    assert result["response"].root.type == "question"
    system_prompt = provider.messages[0]["content"]
    # Prompt must explicitly say flexible or no preference is answered
    assert "flexible" in system_prompt or "no preference" in system_prompt


@pytest.mark.asyncio
async def test_workflow_all_dimensions_in_initial_message_produces_recommendation():
    """All three dimensions in first message → LLM recommends directly without re-asking."""
    provider = FakeProvider(
        '{"type":"recommendation","message":"Perfect gaming setup found.",'
        '"reasoning":"Matches all three criteria.","selections":[{'
        '"product_id":"product-1","match_score":94,'
        '"explanation":"Ideal for AAA competitive gaming under 40M VND."}],'
        '"question":null}'
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [
                {
                    "role": "user",
                    "content": (
                        "I want a gaming laptop for AAA games, competitive play, "
                        "budget under 40 million VND."
                    ),
                }
            ],
            "session_id": "session-alldims-1",
            "retry_count": 0,
        }
    )

    assert result["response"].root.type == "recommendation"
    assert result["response"].root.decision.products[0].id == "product-1"

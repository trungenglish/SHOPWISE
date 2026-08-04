import pytest

from src.graph.workflow import create_workflow
from src.models.schemas import CatalogProduct


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


@pytest.mark.asyncio
async def test_workflow_uses_configured_model_and_complete_history():
    provider = FakeProvider(
        '{"type":"question","message":"What is your budget?","reasoning":null,"selections":null}'
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

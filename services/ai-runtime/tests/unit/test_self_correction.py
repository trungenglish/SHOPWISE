import pytest
from test_graph import FakeToolProxy

from src.graph.workflow import create_workflow


class SequencedProvider:
    def __init__(self, responses: list[str]) -> None:
        self.responses = iter(responses)
        self.call_count = 0

    async def chat_completion(self, **_kwargs) -> str:
        self.call_count += 1
        return next(self.responses)


@pytest.mark.asyncio
async def test_self_correction_retries_invalid_json_twice_at_most():
    provider = SequencedProvider(
        [
            "invalid json",
            "still invalid",
            '{"type":"question","message":"What is your budget?",'
            '"reasoning":null,"selections":null}',
        ]
    )
    workflow = create_workflow(provider, FakeToolProxy(), "gpt-5.4-mini")

    result = await workflow.ainvoke(
        {
            "messages": [{"role": "user", "content": "Laptop"}],
            "session_id": "session-1",
            "retry_count": 0,
        }
    )

    assert result["response"].root.type == "question"
    assert result["retry_count"] == 2
    assert provider.call_count == 3

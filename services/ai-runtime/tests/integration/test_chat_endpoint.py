from fastapi.testclient import TestClient

from src.api.chat import get_provider, get_tool_proxy
from src.main import app
from src.models.schemas import CatalogProduct


class FakeProvider:
    async def chat_completion(self, **_kwargs) -> str:
        return (
            '{"type":"question","message":"What is your budget?",'
            '"reasoning":null,"selections":null}'
        )


class FakeToolProxy:
    async def catalog_search(self) -> list[CatalogProduct]:
        return []


def test_chat_endpoint_accepts_history_and_returns_agent_envelope():
    app.dependency_overrides[get_provider] = lambda: FakeProvider()
    app.dependency_overrides[get_tool_proxy] = lambda: FakeToolProxy()
    client = TestClient(app)

    response = client.post(
        "/api/v1/chat",
        json={
            "session_id": "session-1",
            "messages": [
                {"role": "user", "content": "I need a laptop"},
                {"role": "assistant", "content": "What is it for?"},
                {"role": "user", "content": "Gaming"},
            ],
        },
    )

    app.dependency_overrides.clear()
    assert response.status_code == 200
    body = response.json()
    assert body["type"] == "question"
    assert body["message"] == "What is your budget?"
    assert body["schema_version"] == "1.0"
    assert body["question"]["mode"] == "single"
    assert len(body["question"]["options"]) == 3


def test_chat_stream_endpoint_uses_named_sse_events():
    app.dependency_overrides[get_provider] = lambda: FakeProvider()
    app.dependency_overrides[get_tool_proxy] = lambda: FakeToolProxy()
    client = TestClient(app)

    response = client.post(
        "/api/v1/chat/stream",
        json={
            "session_id": "session-1",
            "messages": [{"role": "user", "content": "I need a laptop"}],
        },
    )

    app.dependency_overrides.clear()
    assert response.status_code == 200
    assert "event: token" in response.text
    assert "event: ui_operation" in response.text
    assert "event: envelope" in response.text
    assert "event: done" in response.text

import pytest
from fastapi.testclient import TestClient
from src.main import app

client = TestClient(app)

def test_chat_endpoint_no_auth():
    # Since API key is missing in test environment, it should return 500 LLM API key not configured
    response = client.post("/api/v1/chat", json={"session_id": "123", "message": "Hi", "model": "gpt-4o-mini"})
    assert response.status_code == 500
    assert "LLM API key not configured" in response.text

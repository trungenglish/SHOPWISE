import pytest
from unittest.mock import AsyncMock
from src.graph.workflow import create_workflow

@pytest.mark.asyncio
async def test_self_correction_loop():
    mock_provider = AsyncMock()
    # First returns invalid JSON, then valid
    mock_provider.chat_completion.side_effect = [
        "invalid json text",
        '{"message": "fixed", "tool_calls": []}'
    ]
    
    workflow = create_workflow(mock_provider)
    state = {
        "messages": [{"role": "user", "content": "Hi"}],
        "session_id": "123",
        "retry_count": 0,
        "final_response": None,
        "error": None
    }
    
    result = await workflow.ainvoke(state)
    
    assert result["error"] is None
    assert "fixed" in result["final_response"]
    assert result["retry_count"] == 1
    assert mock_provider.chat_completion.call_count == 2

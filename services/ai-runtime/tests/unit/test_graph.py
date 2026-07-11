import pytest
from unittest.mock import AsyncMock
from src.graph.workflow import create_workflow

@pytest.mark.asyncio
async def test_langgraph_workflow_valid_json():
    mock_provider = AsyncMock()
    mock_provider.chat_completion.return_value = '{"message": "success", "tool_calls": []}'
    
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
    assert "success" in result["final_response"]

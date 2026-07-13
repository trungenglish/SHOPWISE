import pytest
from src.api.streaming import generate_sse
import json

@pytest.mark.asyncio
async def test_generate_sse():
    async def mock_generator():
        yield "chunk1"
        yield "chunk2"
        
    sse_gen = generate_sse(mock_generator())
    
    events = [event async for event in sse_gen]
    
    assert len(events) == 3
    assert events[0]["event"] == "message"
    assert "chunk1" in events[0]["data"]

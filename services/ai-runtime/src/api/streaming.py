import json
from typing import AsyncGenerator

async def generate_sse(generator: AsyncGenerator[str, None]) -> AsyncGenerator[dict, None]:
    """
    Wraps an async generator of strings into an SSE-compatible dictionary format
    for EventSourceResponse.
    """
    async for chunk in generator:
        yield {
            "event": "message",
            "data": json.dumps({"text": chunk})
        }
    
    # End of stream indicator
    yield {
        "event": "done",
        "data": "[DONE]"
    }

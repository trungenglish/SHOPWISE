import json
from collections.abc import AsyncGenerator

from src.models.schemas import AgentResponse


async def generate_agent_sse(
    response: AgentResponse,
) -> AsyncGenerator[dict[str, str], None]:
    # ponytail: one validated message token; stream model deltas when partial JSON is safe.
    yield {
        "event": "token",
        "data": json.dumps({"text": response.root.message}),
    }

    if response.root.type in {"recommendation", "comparison", "checkout_ready"}:
        yield {
            "event": response.root.type,
            "data": response.root.decision.model_dump_json(),
        }

    for operation in response.root.ui_operations:
        yield {
            "event": "ui_operation",
            "data": operation.model_dump_json(exclude_none=True),
        }

    yield {
        "event": "envelope",
        "data": response.root.model_dump_json(exclude_none=True),
    }

    yield {
        "event": "done",
        "data": "{}",
    }

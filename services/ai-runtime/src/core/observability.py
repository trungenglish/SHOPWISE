import logging
import sys
from typing import Any

logger = logging.getLogger("ai_runtime")
logger.setLevel(logging.INFO)
handler = logging.StreamHandler(sys.stdout)
formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
handler.setFormatter(formatter)
logger.addHandler(handler)


def log_interaction(
    session_id: str,
    provider: str,
    tokens_used: int | None = None,
    latency_ms: int | None = None,
    event: str = "llm_invocation",
) -> None:
    """
    Logs LLM interactions securely.
    Must NOT contain any user or AI message content.
    """
    log_data: dict[str, Any] = {
        "event": event,
        "session_id": session_id,
        "provider": provider,
    }
    if tokens_used is not None:
        log_data["tokens_used"] = tokens_used
    if latency_ms is not None:
        log_data["latency_ms"] = latency_ms

    logger.info(f"LLM Interaction: {log_data}")


def log_error(session_id: str, provider: str, error_type: str) -> None:
    """
    Logs errors related to LLM invocation without leaking PII.
    """
    logger.error(
        "LLM Error - Session: %s | Provider: %s | Type: %s",
        session_id,
        provider,
        error_type,
    )

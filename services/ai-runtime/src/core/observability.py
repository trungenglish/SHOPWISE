import logging
import sys
from typing import Any

# Configure structured logging
# We ensure we NEVER log message content (PII sanitation)

class AILogger(logging.Logger):
    def trace_span(self, name: str):
        """
        Create a tracing span for OpenTelemetry. 
        In MVP, this is a mock context manager. 
        """
        class MockSpan:
            def __enter__(self):
                pass
            def __exit__(self, exc_type, exc_val, exc_tb):
                pass
        return MockSpan()

# Singleton logger instance
logger = AILogger("ai_runtime")
logger.setLevel(logging.INFO)
handler = logging.StreamHandler(sys.stdout)
formatter = logging.Formatter(
    "%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
handler.setFormatter(formatter)
logger.addHandler(handler)

def log_interaction(session_id: str, provider: str, tokens_used: int | None = None, latency_ms: int | None = None, event: str = "llm_invocation") -> None:
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

def log_error(session_id: str, provider: str, error_type: str, error_message: str) -> None:
    """
    Logs errors related to LLM invocation without leaking PII.
    """
    logger.error(f"LLM Error - Session: {session_id} | Provider: {provider} | Type: {error_type} | Message: {error_message}")

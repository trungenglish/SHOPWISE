# AI Runtime

Python FastAPI service providing LLM integration for the SHOPWISE platform.

## Features
- **Stateless design:** All conversation state comes from the client
- **Pydantic Validation:** Strict request/response schemas
- **Provider Abstraction:** Hot-swappable LLM providers
- **LangGraph Orchestration:** Tools, reasoning, and self-correction
- **Streaming:** Server-Sent Events (SSE) stream support
- **Privacy:** PII-safe logging out of the box

## Quickstart

### Prerequisites
- Python 3.11+
- `uv` package manager

### Run Server
```bash
cd services/ai-runtime
uv run uvicorn src.main:app --reload --port 8000
```

### Run Tests
```bash
uv run pytest
```

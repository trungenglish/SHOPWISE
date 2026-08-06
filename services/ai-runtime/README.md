# AI Runtime

Python FastAPI service providing LLM integration for the SHOPWISE platform.

## Shopping conversation contract

Chat envelopes use `schema_version: "1.1"`. The runtime asks one structured
question at a time until workload, gaming ambition, and budget are known;
question interaction requests remain at `"1.0"`. Product, offer, campaign,
and checkout identifiers are hydrated from backend tools, so generated copy
cannot supply its own prices or offer IDs. `POST /api/v1/greeting` is
ephemeral and consumes only caller-supplied, limited facts.

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

# Quickstart: AI Runtime LLM Integration

This guide explains how to validate the LLM Integration locally.

## Prerequisites
- Docker & Docker Compose
- `uv` (Python package manager)
- A valid OpenAI API Key

## Setup

1. **Configure Environment Variables**:
   In `services/ai-runtime`, copy the example environment file:
   ```bash
   cp .env.example .env
   ```
   Add your OpenAI API key to `.env`:
   ```env
   OPENAI_API_KEY=sk-...
   MODEL=gpt-5.4-mini
   ```

2. **Start the AI Runtime**:
   ```bash
   cd services/ai-runtime
   uv run uvicorn src.main:app --reload --host 0.0.0.0 --port 8000
   ```

## Validation Scenarios

### Scenario 1: Basic Chat Stream
Test that the AI Runtime can connect to OpenAI and stream a basic response.

**Command**:
```bash
curl -N -X POST http://localhost:8000/api/v1/chat/stream \
-H "Content-Type: application/json" \
-d '{
  "session_id": "test-1",
  "messages": [{"role": "user", "content": "I need a laptop."}]
}'
```

**Expected Outcome**:
You should see validated SSE events ending with `done`. The API key and message content must not be printed in server logs.
```text
event: token
data: {"text": "What is your budget?"}

event: done
data: {}
```

### Scenario 2: Error Handling (Invalid API Key)
Test that the AI Runtime catches authentication errors and returns a structured AI Error State instead of crashing.

1. Temporarily change `OPENAI_API_KEY` in `.env` to `sk-invalid`.
2. Restart the server.
3. Run the same curl command from Scenario 1.

**Expected Outcome**:
```text
event: error
data: {"message": "Agent request failed"}

event: done
data: {}
```

### Scenario 3: Structured Recommendation / Self-Correction
Test that the LLM returns a validated recommendation hydrated from the catalog.

**Command**:
Ask for a laptop with a concrete use case and VND budget.

**Expected Outcome**:
The AI Runtime should emit `token`, `recommendation`, and `done`. Every product ID
and price must match the Go catalog.

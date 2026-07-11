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
   LLM_PROVIDER=openai
   LLM_MODEL=gpt-4o-mini
   ```

2. **Start the AI Runtime**:
   ```bash
   cd services/ai-runtime
   uv run uvicorn main:app --reload --host 0.0.0.0 --port 8000
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
  "messages": [{"role": "user", "content": "Hello, respond with the word SUCCESS."}],
  "tools": []
}'
```

**Expected Outcome**:
You should see a stream of SSE events ending with `[DONE]`. The API key should not be printed in the server logs.
```text
event: token
data: {"content": "SUCCESS"}

event: done
data: "[DONE]"
```

### Scenario 2: Error Handling (Invalid API Key)
Test that the AI Runtime catches authentication errors and returns a structured AI Error State instead of crashing.

1. Temporarily change `OPENAI_API_KEY` in `.env` to `sk-invalid`.
2. Restart the server.
3. Run the same curl command from Scenario 1.

**Expected Outcome**:
```text
event: error
data: {"error_code": "INVALID_CREDENTIALS", "message": "The provided API key is invalid.", "retryable": false}

event: done
data: "[DONE]"
```

### Scenario 3: Structured Output / Self-Correction
*(Requires tool orchestration to be fully implemented)*
Test that the LLM successfully forces structured output for a tool.

**Command**:
Send a chat request including a `compare_products` tool schema, and ask "Compare iPhone and Samsung".

**Expected Outcome**:
The AI Runtime should emit a `tool_call` event with strict JSON arguments matching the schema, even if it had to self-correct internally.

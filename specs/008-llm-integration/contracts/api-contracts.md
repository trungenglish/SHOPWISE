# API Contracts: AI Runtime LLM Integration

## 1. Chat Completion Endpoint (Streaming)

**POST** `/api/v1/chat/stream`

This endpoint accepts a conversation history and streams back the LLM's response using Server-Sent Events (SSE). It handles tool calling orchestration internally.

### Request Body (`ChatRequest`)
```json
{
  "session_id": "req-12345",
  "messages": [
    {
      "role": "user",
      "content": "Compare the iPhone 15 and Galaxy S24"
    }
  ],
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "compare_products",
        "description": "Compare two or more products",
        "parameters": {
          "type": "object",
          "properties": {
            "product_ids": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        }
      }
    }
  ],
  "temperature": 0.7
}
```

### Response (Server-Sent Events)

**Content-Type:** `text/event-stream`

The stream yields events of different types: `token`, `tool_call`, `json_patch`, `error`, `done`.

#### Event: Token
```text
event: token
data: {"content": "The "}

event: token
data: {"content": "iPhone "}
```

#### Event: Tool Call (When the AI decides to use a tool)
```text
event: tool_call
data: {"tool_name": "compare_products", "arguments": {"product_ids": ["PROD-1", "PROD-2"]}}
```

#### Event: JSON Patch (For Dynamic UI Protocol updates)
```text
event: json_patch
data: {"op": "add", "path": "/ui_elements/-", "value": {"type": "loading", "message": "Comparing products..."}}
```

#### Event: Error
```text
event: error
data: {"error_code": "PROVIDER_TIMEOUT", "message": "OpenAI API timed out.", "retryable": true}
```

#### Event: Done
```text
event: done
data: "[DONE]"
```

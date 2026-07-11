# Observability & Telemetry

## Observations

The LLM does not receive raw database or API responses. It receives an `Observation`.

```json
{
 "summary": "3 products found matching 'laptop'",
 "data": {
    "items": [...]
 },
 "confidence": 0.98
}
```

This makes the LLM's reasoning process significantly cheaper (fewer tokens) and faster.

## Telemetry

Each Tool execution tracks standard metrics:

```text
duration
  ↓
input_size
  ↓
output_size
  ↓
errors
  ↓
retry
  ↓
cost
```

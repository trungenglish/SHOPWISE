# Quickstart: Dynamic UI Protocol Validation

This guide provides scenarios to validate the schemas and data contracts across different layers.

## Prerequisites

- Node.js (v18+) for validating the frontend schemas
- Go (1.21+) for validating backend structs
- Python (3.11+) for validating AI Runtime Pydantic models
- `ajv-cli` (installed via `npm i -g ajv-cli`) for JSON Schema validation.

## Scenario 1: Validate a complete Product Recommendation Payload

**Goal**: Ensure a generated `DynamicUIDocument` containing a ProductCard validates against the JSON Schema.

1. Create a test payload `test-payload.json`:
```json
{
  "protocol": "dynamic-ui",
  "schemaVersion": "1.0",
  "documentId": "doc-123",
  "timestamp": "2026-07-12T10:00:00Z",
  "operation": "replace",
  "root": {
    "type": "ProductCard",
    "id": "prod-456",
    "props": {
      "title": "Lenovo ThinkPad X1",
      "price": 35000000,
      "currency": "VND"
    },
    "actions": [
      {
        "trigger": "onClick",
        "action": "VIEW_PRODUCT_DETAILS"
      }
    ]
  }
}
```

2. Save the downstream schema from `contracts/ui-contracts.md` as `schema.json`.

3. Run the validation:
```bash
ajv validate -s schema.json -d test-payload.json
```

**Expected Outcome**: `test-payload.json is valid`.

## Scenario 2: Validate an Interaction Event

**Goal**: Ensure the frontend correctly formats an `InteractionEvent` before sending to the backend.

1. Create a test event `test-event.json`:
```json
{
  "componentId": "btn-add-to-cart",
  "action": "ADD_TO_CART",
  "payload": {
    "productId": "prod-456",
    "quantity": 1
  },
  "metadata": {
    "timestamp": "2026-07-12T10:05:00Z",
    "clientContext": {
      "platform": "web"
    }
  }
}
```

2. Save the upstream schema from `contracts/ui-contracts.md` as `event-schema.json`.

3. Run the validation:
```bash
ajv validate -s event-schema.json -d test-event.json
```

**Expected Outcome**: `test-event.json is valid`.

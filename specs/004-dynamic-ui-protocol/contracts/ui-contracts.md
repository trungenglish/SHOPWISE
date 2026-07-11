# UI Protocol Contract

This contract defines the JSON structures exchanged between the backend/AI and the frontend.

## Downstream: `DynamicUIDocument` (Backend -> Frontend)

This is the payload received by the frontend over the chosen transport layer.

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "DynamicUIDocument",
  "type": "object",
  "required": ["protocol", "schemaVersion", "documentId", "operation", "root"],
  "properties": {
    "protocol": {
      "type": "string",
      "const": "dynamic-ui"
    },
    "schemaVersion": {
      "type": "string",
      "pattern": "^[0-9]+\\.[0-9]+$"
    },
    "documentId": {
      "type": "string"
    },
    "timestamp": {
      "type": "string",
      "format": "date-time"
    },
    "operation": {
      "type": "string",
      "enum": ["replace", "append", "patch", "remove"]
    },
    "targetId": {
      "type": "string",
      "description": "Required if operation is not 'replace' (at root)"
    },
    "metadata": {
      "type": "object",
      "additionalProperties": true
    },
    "root": {
      "$ref": "#/definitions/ComponentNode"
    }
  },
  "definitions": {
    "ComponentNode": {
      "type": "object",
      "required": ["type"],
      "properties": {
        "id": { "type": "string" },
        "type": { "type": "string" },
        "props": { "type": "object" },
        "state": { "type": "object" },
        "validation": { "type": "object" },
        "children": {
          "type": "array",
          "items": { "$ref": "#/definitions/ComponentNode" }
        },
        "actions": {
          "type": "array",
          "items": {
            "type": "object",
            "required": ["trigger", "action"],
            "properties": {
              "trigger": { "type": "string" },
              "action": { "type": "string" },
              "payloadSchema": { "type": "object" }
            }
          }
        }
      }
    }
  }
}
```

## Upstream: `InteractionEvent` (Frontend -> Backend)

This is the payload emitted by the frontend when a user interacts with a component.

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "InteractionEvent",
  "type": "object",
  "required": ["componentId", "action", "payload"],
  "properties": {
    "componentId": {
      "type": "string"
    },
    "action": {
      "type": "string"
    },
    "payload": {
      "type": "object",
      "additionalProperties": true
    },
    "metadata": {
      "type": "object",
      "properties": {
        "timestamp": { "type": "string", "format": "date-time" },
        "clientContext": { "type": "object" }
      }
    }
  }
}
```

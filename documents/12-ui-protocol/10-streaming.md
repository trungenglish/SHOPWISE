# Streaming & Incremental Patching

SAUP handles progressive UI rendering over time. Since LLMs generate output sequentially, the UI Composer streams the SAUP AST.

## Initial State

```text
Root
↓
Workspace
↓
Recommendation
↓
loading...
```

## Incremental Patch (JSON Patch RFC6902)

Instead of re-sending the whole tree, SAUP sends JSON Patches.

```json
{
 "op": "replace",
 "path": "/workspace/children/0/children",
 "value": [
   { "type": "product.card" },
   { "type": "product.card" },
   { "type": "product.card" }
 ]
}
```

## Timeline Example
*   **0ms**: Layout container appears with loading skeletons.
*   **500ms**: `Recommendation` populates with 3 products.
*   **800ms**: `Reasoning` node appears explaining the recommendations.
*   **1200ms**: `Promotion` banner streams in.

This ensures the user sees an interactive UI as fast as possible without waiting for the full reasoning generation.

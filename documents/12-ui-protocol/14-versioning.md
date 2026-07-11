# Versioning

SAUP includes versioning in the `UIRoot` to manage backwards compatibility.

```text
1.0
 ↓
1.1
 ↓
2.0
```

## Unknown Nodes

If the UI Composer specifies a Node Type that the older Client Renderer does not recognize (e.g., `feature.new_panel`), the Client Renderer falls back to a `FallbackComponent` that either:

- Displays a placeholder
- Prompts the user to update their app
- Ignores the node gracefully depending on context

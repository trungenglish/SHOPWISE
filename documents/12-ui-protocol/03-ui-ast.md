# UI AST (Abstract Syntax Tree)

SAUP replaces the legacy `Surface[]` flat list with a deeply nested **UI Tree**, inspired by React Fiber, HTML DOM, and Flutter Widget Trees.

## The Root Object

Every SAUP payload starts with a `UIRoot`:

```typescript
interface UIRoot {
    version: string;
    sessionId: string;
    conversationId: string;
    tree: UINode;
    metadata: Record<string, unknown>;
}
```

## The UINode

Everything in SAUP is a Node. There is no limit to the nesting depth.

```typescript
interface UINode {
    id: string;
    type: string;
    props: Record<string, any>;
    state: UIState;
    actions: UIAction[];
    children: UINode[];
}
```

## Tree Example

Conceptually, the AST represents a nested structure:

```text
Workspace
├── Recommendation
│    ├── Product Card
│    ├── Product Card
│    └── Product Card
├── Reasoning
└── Checkout
```

As JSON payload:

```json
{
  "type": "workspace",
  "children": [
      {
        "type": "recommendation.panel",
        "children": [
            { "type": "product.card" },
            { "type": "product.card" }
        ]
      }
  ]
}
```

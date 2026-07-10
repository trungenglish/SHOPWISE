# Actions

Actions represent the operations a user can perform on a UINode.

## Definition

```typescript
interface UIAction {
    id: string;
    label: string;
    type: string;
    payload: any;
}
```

## Example

An action to compare products attached to a `product.card`:

```json
{
 "type": "tool",
 "label": "Compare",
 "payload": {
      "tool": "compare_products",
      "args": { "product_id": "prod_123" }
 }
}
```

When triggered, the Frontend executes the action via the Conversation API.

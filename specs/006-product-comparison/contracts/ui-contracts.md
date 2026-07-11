# UI Contracts: Dynamic UI Protocol

## `ComparisonWorkspace` Component Schema
```json
{
  "type": "ComparisonWorkspace",
  "props": {
    "workspace_id": "uuid",
    "category_name": "Laptops",
    "products": [
      {
        "id": "prod_1",
        "name": "ThinkPad X1 Carbon",
        "price_vnd": 45000000,
        "images": ["url1"],
        "specs": {"CPU": "Core i7", "RAM": "16GB"}
      }
    ],
    "highlights": {
      "RAM": {
        "prod_1": "EQUAL"
      }
    },
    "ai_explanation": "The ThinkPad is better for business users...",
    "state": "ACTIVE" // or PENDING_REPLACEMENT
  },
  "events": {
    "onRemoveProduct": {"type": "REMOVE_PRODUCT", "payload": {"product_id": "string"}},
    "onReplaceProduct": {"type": "REPLACE_PRODUCT", "payload": {"old_product_id": "string", "new_product_id": "string"}},
    "onProceedToCheckout": {"type": "CHECKOUT_READY", "payload": {"product_id": "string"}},
    "onAskFollowUp": {"type": "ASK_AI", "payload": {"query": "string"}}
  }
}
```

## `ErrorState` Component Schema
```json
{
  "type": "ErrorState",
  "props": {
    "title": "Category Mismatch",
    "message": "You can only compare products within the same category. The selected product is a Mouse, but you are currently comparing Laptops.",
    "action": "Dismiss"
  }
}
```

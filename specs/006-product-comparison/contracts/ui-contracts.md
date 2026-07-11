# UI Contracts: Dynamic UI Protocol

**Traceability**: This contract directly implements requirements from [spec.md](file:///d:/Github/SHOPWISE/specs/006-product-comparison/spec.md) (FR-002, FR-003, FR-007, FR-009).

## `ComparisonWorkspace` Component Schema
```json
{
  "type": "ComparisonWorkspace",
  "props": {
    "workspace_id": "uuid",
    "category_name": "Laptops",
    "state": "ACTIVE", // States: INIT, ACTIVE, PENDING_REPLACEMENT
    "loading_state": {
      "ai_explanation": false, // When true, UI renders a multi-line text skeleton spinner
      "workspace": false
    },
    "products": [
      {
        "id": "prod_1",
        "name": "ThinkPad X1 Carbon",
        "price_vnd": 45000000,
        "availability": "IN_STOCK", // OR "OUT_OF_STOCK"
        "images": ["url1"],
        "specs": {"CPU": "Core i7", "RAM": "16GB"}
      }
    ],
    "highlights": {
      "RAM": {
        "prod_1": "EQUAL" // Valid values: BETTER, WORSE, EQUAL, MISSING
      }
    },
    "ai_explanation": "The ThinkPad is better for business users..."
  },
  "events": {
    "onRemoveProduct": {"type": "REMOVE_PRODUCT", "payload": {"product_id": "string"}},
    "onReplaceProduct": {"type": "REPLACE_PRODUCT", "payload": {"old_product_id": "string", "new_product_id": "string"}},
    "onProceedToCheckout": {"type": "CHECKOUT_READY", "payload": {"product_id": "string"}},
    "onAskFollowUp": {"type": "ASK_AI", "payload": {"query": "string", "context": "ComparisonWorkspace"}}
  }
}
```

## `ErrorState` Component Schema
```json
{
  "type": "ErrorState",
  "props": {
    "title": "Category Mismatch",
    "message": "You can only compare products within the same category.",
    "severity": "warning",
    "action": "Dismiss"
  }
}
```

---

## Visual & Behavior Requirements

### 1. Component States
- **INIT**: When 0 products are selected. UI displays empty placeholder columns with an "Add products to compare" prompt.
- **ACTIVE**: Standard view rendering up to 4 products.
- **PENDING_REPLACEMENT**: When adding a 5th product, the UI overlays a modal or highlighted border across the 4 existing products, prompting the user to select one to evict. The new product replaces the pending candidate if another 5th product is added during this state.

### 2. Highlight Mapping (Meaningful Differences)
Differences are visually emphasized to avoid looking like a standard table:
- **BETTER**: Green text/background tint with a checkmark icon.
- **WORSE**: Red text/background tint with an 'X' icon.
- **EQUAL**: Neutral text color with a dash icon.
- **MISSING**: Grey text displaying "Not specified" or "-".

### 3. Structural Constraints
- **AI Explanation**: Supports Markdown (bold, italics, bullet points). Maximum 500 characters.
- **Responsive Breakpoints**: On mobile (width < 768px), switch from a 4-column grid to a horizontal scrolling carousel or an accordion view.
- **Product Card Consistency**: Must reuse the standard application product card components for consistency.
- **Out of Stock**: If a product is `OUT_OF_STOCK`, its card is visually greyed out and the "Proceed to Checkout" button is disabled.

### 4. Edge Cases & Error Handling
- **Image Failure**: Display a standardized application placeholder image if product images fail to load.
- **Missing Specs**: Display "Not specified" or "-" rather than throwing an error.
- **AI Explanation Failure**: Display fallback text: "Explanation currently unavailable."
- **Network Error (Checkout)**: If `onProceedToCheckout` fails, display a toast notification and keep the button enabled for manual retry.
- **Error Dialog Dismissal**: The `ErrorState` dialog is dismissed via the "Dismiss" action button or clicking the modal backdrop. Consistently matches global error handling.

### 5. Performance & Accessibility
- **Performance**: The workspace must render under 200ms when transitioning states. The AI explanation streams progressively.
- **Accessibility**: Wrapper must use `role="region"` with `aria-label="Product Comparison"`. Users must be able to navigate between product columns using the `Tab` key.
- **Measurability**: The `PENDING_REPLACEMENT` resolution is objectively verified when the schema returns an `ACTIVE` state with the new `products` array mapping correctly.

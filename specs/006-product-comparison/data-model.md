# Data Model: Product Comparison Workspace

## Entities

### `ComparisonWorkspace` (Persisted in Decision Memory)
Represents a user's active comparison session.
- `workspace_id`: string (UUID)
- `user_id`: string
- `category_id`: string (Enforces same-category comparison)
- `product_ids`: list of strings (Max 4)
- `created_at`: timestamp
- `updated_at`: timestamp

### `ComparisonProductDetail` (Ephemeral / Fetched from Retail Backend)
Represents the runtime data for a product in the comparison view.
- `product_id`: string
- `name`: string
- `brand`: string
- `category`: string
- `price_vnd`: integer
- `promotion`: object (optional)
- `availability`: enum (IN_STOCK, OUT_OF_STOCK, PREORDER)
- `warranty`: string
- `images`: list of strings (URLs)
- `specifications`: dictionary of string -> string (e.g., {"CPU": "M3", "RAM": "16GB"})

### `ComparisonDifferenceHighlight` (Computed by Backend/AI)
Represents the visual differences for a specific specification across the compared products.
- `specification_key`: string (e.g., "RAM")
- `highlights`: dictionary mapping `product_id` -> `HighlightType`
- `HighlightType`: enum (BETTER, WORSE, EQUAL, MISSING, CATEGORY_ADVANTAGE)

## State Transitions

- `INIT`: Empty workspace created.
- `ADDING_PRODUCT`: Validating category and adding to `product_ids`.
- `MAX_CAPACITY`: 4 products reached. Trying to add a 5th triggers the pending replacement state.
- `PENDING_REPLACEMENT`: A 5th product is selected, awaiting user decision on which existing product to replace.
- `COMPLETED`: A product is selected to move to Checkout Readiness.

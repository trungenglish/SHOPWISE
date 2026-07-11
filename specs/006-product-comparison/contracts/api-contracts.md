# API Contracts: Product Comparison Workspace

## Tool Protocol (AI Runtime -> Go Backend)

### `fetch_comparison_data`
Retrieves comparison data for a list of products.
- **Request**:
  - `product_ids`: list of strings
- **Response**:
  - `products`: list of `ComparisonProductDetail` objects (defined in data-model.md)
  - `highlights`: dictionary of computed visual differences

### `update_workspace_state`
Persists the updated product list to the Decision Memory.
- **Request**:
  - `workspace_id`: string
  - `product_ids`: list of strings
- **Response**:
  - `status`: string (SUCCESS/ERROR)

## Backend Routing Events (Frontend -> Go Backend)
Handled via semantic events in the UI schema.

### `ADD_TO_COMPARISON`
- **Payload**: `product_id`
- **Backend Logic**:
  1. Fetch workspace from Decision Memory.
  2. Validate category matches. If not, return `ErrorState` schema.
  3. Validate count < 4. If 4, return `ComparisonWorkspace` schema with state `PENDING_REPLACEMENT`.
  4. If valid, append `product_id`, update Decision Memory, return updated `ComparisonWorkspace` schema.

### `REPLACE_PRODUCT`
- **Payload**: `old_product_id`, `new_product_id`
- **Backend Logic**: Swap IDs in Decision Memory and return updated UI.

### `REMOVE_PRODUCT`
- **Payload**: `product_id`
- **Backend Logic**: Remove `product_id` from Decision Memory and return updated `ComparisonWorkspace` schema. If 0 products remain, return initial empty workspace state.

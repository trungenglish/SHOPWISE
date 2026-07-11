# Validation Quickstart: Product Comparison Workspace

This guide validates the end-to-end functionality of the Product Comparison Workspace using backend API calls and AI Runtime invocations.

## Prerequisites
- Local Go Backend running (`make run-backend`)
- Local Python AI Runtime running (`make run-ai`)

## Validation Scenario 1: Adding Products to Workspace
1. Select 2 products in the same category (e.g., Laptops).
2. Trigger the `ADD_TO_COMPARISON` event payload (defined in `contracts/api-contracts.md`).
3. Verify the Backend returns a `ComparisonWorkspace` UI schema.
4. Verify Decision Memory has persisted the new state in the backend database.

## Validation Scenario 2: Enforcing Max 4 Products
1. Add 4 products to the workspace.
2. Trigger an `ADD_TO_COMPARISON` event for a 5th product.
3. Verify the Backend returns the UI schema with state `PENDING_REPLACEMENT` and no product was automatically deleted.

## Validation Scenario 3: Cross-Category Error
1. Create a workspace with a Laptop.
2. Trigger an `ADD_TO_COMPARISON` event for a Mouse.
3. Verify the Backend returns an `ErrorState` UI schema rejecting the cross-category addition.

## Validation Scenario 4: AI Explanation Generation
1. With 2 products in the workspace, submit a semantic request to the AI Runtime asking "Compare these options".
2. Verify the AI Runtime calls the `fetch_comparison_data` tool to gather the specs.
3. Verify the returned explanation strictly uses the fetched data without hallucinating specs.

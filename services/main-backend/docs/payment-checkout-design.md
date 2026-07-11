# Payment Checkout Design

## Scope

Add an authenticated `POST /api/v1/checkout` endpoint whose sole responsibility is validating and persisting an order and its line items. Payment-provider authorization, inventory reservation, promotions, and catalog-price reconciliation are outside this change.

## Architecture

The feature follows the existing feature-local flow used by `users` and `identity`:

`handler -> usecase service -> repository interface -> PostgreSQL repository`

The handler obtains the customer UUID from the existing JWT middleware. The request cannot choose another customer, order status, order ID, creation time, or total. The service validates line items, rejects duplicate products, calculates an overflow-safe `int64` total in minor currency units, and creates an order with `PENDING` status. The repository inserts the order and all items in one GORM transaction.

## HTTP Contract

Request fields:

- `items`: required nonempty array.
- `product_id`: UUID string.
- `quantity`: positive integer.
- `unit_price`: positive `int64` integer in minor units.

Response fields:

- `order_id`: server-generated UUID string.
- `customer_id`: authenticated user UUID string.
- `created_at`: server-generated UTC timestamp.
- `items`: persisted line items.
- `total_amount`: server-calculated `int64` integer in minor units.
- `status`: `PENDING`.

## Persistence

The `orders` table stores order ID, customer ID, creation time, total amount, and status. The `order_items` table stores order ID, product ID, quantity, and unit price. A composite primary key on `(order_id, product_id)` enforces one line per product; the service returns a validation error before persistence for duplicates. Deleting an order cascades to its items.

The obsolete platform-level order model is removed so the feature-owned migration is the only schema definition.

## Error Handling

Malformed JSON and invalid order data use the existing `VALIDATION_FAILED` problem response. Missing or invalid JWTs use `UNAUTHORIZED`. Repository failures are wrapped as `INTERNAL_ERROR` without leaking database details.

## Verification

Usecase tests cover defaults, totals, validation, duplicate products, overflow, and repository errors. Handler tests cover authentication, successful response mapping, malformed requests, and service errors. A PostgreSQL integration test covers persisted order/item data and transaction rollback.

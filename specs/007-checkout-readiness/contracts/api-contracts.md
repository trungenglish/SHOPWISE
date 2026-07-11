# API & Event Contracts: Checkout Readiness

## Overview

This document specifies the Semantic Events emitted by the React Frontend, the REST endpoints exposed by the Go Backend, and the Tools registered for the AI Runtime.

## Semantic Events (Frontend -> Backend/AI)

These events are transported via WebSocket or REST to the Go Backend, which then updates state and routes them to the AI Runtime if conversational reasoning is needed.

### 1. `PromotionExpired`
**Endpoint:** `POST /api/events/promotion-expired`
Emitted by the frontend when the live countdown reaches zero.
```json
{
  "type": "PROMOTION_EXPIRED",
  "payload": {
    "promotionId": "promo_flash"
  }
}
```

### 2. `StoreSelected`
**Endpoint:** `POST /api/events/store-selected`
Emitted when the user selects a store for pickup.
```json
{
  "type": "STORE_SELECTED",
  "payload": {
    "storeId": "store_district1"
  }
}
```

### 3. `ContinueToCheckout`
**Endpoint:** `POST /api/events/continue-checkout`
Emitted when the user is in the `READY` state and clicks the final button.
```json
{
  "type": "CONTINUE_RETAIL_CHECKOUT",
  "payload": {
    "sessionId": "..."
  }
}
```
**Response handling:**
- If state is `READY`: `200 OK` (Proceeds to Retail Checkout).
- If state is NOT `READY` (e.g., Error, Missing Information): Returns a `400 Bad Request` containing the specific missing checklist items to flash in the UI.

## Go Backend Endpoints

### `POST /api/checkout-readiness/customer-info`
Secure endpoint for the `CustomerInformationForm` component to submit PII.
- **Request Body**: `{ "name": "...", "phoneNumber": "...", "deliveryPreference": "HOME_DELIVERY" }`
- **Response**: `200 OK` (Updates Decision Memory and triggers an internal event to the AI).

### `GET /api/stores/nearby?lat=...&lng=...`
Returns a list of eligible stores for pickup near the given coordinates.

## AI Runtime Tools (Tool Protocol)

### `ValidateCheckoutReadiness`
The AI calls this tool to ask the Go Backend to run the retail checks (inventory, price, promotions).
- **Arguments**: `sessionId`
- **Returns**: `CheckoutReadinessState` object (see Data Model).

### `GetCheckoutChecklist`
Retrieves the current progress of the user's checklist without triggering a full external API validation.
- **Arguments**: `sessionId`
- **Returns**: Missing fields or `READY` status.

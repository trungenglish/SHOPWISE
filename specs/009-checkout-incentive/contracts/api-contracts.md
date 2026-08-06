# API Contracts: Checkout Incentive

## Endpoints (Go Gin API)

### 1. Start Promotion
`POST /api/v1/checkout-incentive/start`
Triggered by the frontend when a product is saved or checkout begins.
**Request**:
```json
{
  "trigger_type": "saved_product", // or "started_checkout"
  "session_id": "string",
  "device_id": "string" // optional, for anonymous users
}
```
**Response** (200 OK or 201 Created):
```json
{
  "promotion_id": "uuid",
  "status": "ACTIVE",
  "expires_at": "2026-08-05T01:00:00Z",
  "campaign_code": "ACCESSORY_NEXT_PURCHASE_15",
  "reward_type": "percentage",
  "reward_value": 15,
  "reward_scope": "accessories_next_purchase"
}
```
*Note: If a promotion is already active for this session, it returns the existing one idempotently. If on cooldown, returns 403 Forbidden or status INELIGIBLE.*

### 2. Get Promotion Status
`GET /api/v1/checkout-incentive/status?session_id=...&device_id=...`
**Response** (200 OK):
```json
{
  "promotion_id": "uuid",
  "status": "ACTIVE",
  "expires_at": "2026-08-05T01:00:00Z",
  "voucher_code": "optional_string",
  "campaign_code": "ACCESSORY_NEXT_PURCHASE_15",
  "reward_type": "percentage",
  "reward_value": 15,
  "reward_scope": "accessories_next_purchase"
}
```

### 3. Payment Completed Event (Internal or Webhook)
Handled internally or via webhook from the payment service.
**Payload**:
```json
{
  "event_type": "payment_completed",
  "session_id": "string",
  "user_id": "uuid",
  "timestamp": "2026-08-05T00:55:00Z"
}
```

### 4. Retry Voucher Issuance
`POST /api/v1/checkout-incentive/{promotion_id}/retry`
**Response**: Same as status endpoint.

### 5. Dev Simulation Endpoint (Development Only)
`POST /api/v1/checkout-incentive/dev/simulate-payment`
```json
{
  "session_id": "string"
}
```

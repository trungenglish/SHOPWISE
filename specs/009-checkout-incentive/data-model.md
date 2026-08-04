# Database Model: Checkout Incentive

## Tables

### `checkout_promotions`

Stores the lifecycle state of a checkout incentive promotion for a user or anonymous session.

| Column | Type | Constraints | Description |
|---|---|---|---|
| `promotion_id` | UUID | PRIMARY KEY | Unique identifier for the promotion |
| `session_id` | String | NOT NULL, INDEX | Shopping session identifier |
| `user_id` | UUID | NULLABLE, INDEX | Authenticated user ID (if applicable) |
| `device_id` | String | NULLABLE, INDEX | Device/Browser fingerprint (for anonymous) |
| `trigger_type` | String | NOT NULL | Action that triggered it (`saved_product`, `started_checkout`) |
| `status` | String | NOT NULL | Current state (`ACTIVE`, `EXPIRED`, `COMPLETED_PENDING_VOUCHER`, `VOUCHER_ISSUED`, `ERROR_RECOVERABLE`) |
| `started_at` | Timestamp | NOT NULL | Time the 15-minute countdown started |
| `expires_at` | Timestamp | NOT NULL | `started_at` + 15 minutes |
| `completed_at` | Timestamp | NULLABLE | Time payment was completed (if before expires_at) |
| `voucher_id` | UUID | NULLABLE | ID of the generated voucher (if issued) |
| `voucher_status` | String | NULLABLE | Status of voucher issuance (`PENDING`, `ISSUED`, `FAILED`) |
| `campaign_code` | String | NOT NULL | MVP Default: `ACCESSORY_NEXT_PURCHASE_15` |
| `reward_type` | String | NOT NULL | MVP Default: `percentage` |
| `reward_value` | Integer | NOT NULL | MVP Default: `15` |
| `reward_scope` | String | NOT NULL | MVP Default: `accessories_next_purchase` |
| `cooldown_until` | Timestamp | NULLABLE | Timestamp until which the user cannot receive another promotion (default 30 days) |
| `created_at` | Timestamp | NOT NULL | Record creation time |
| `updated_at` | Timestamp | NOT NULL | Last update time |

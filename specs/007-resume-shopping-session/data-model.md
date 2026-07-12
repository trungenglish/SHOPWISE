# Data Model: Resume Shopping Session

## Entities

### 1. `ResumeToken`
Tracks the generated resume tokens to ensure single-use constraint.

| Field | Type | Description |
|---|---|---|
| `id` | UUID | Primary key (matches JTI) |
| `session_id` | UUID | Foreign key to `DecisionSession` |
| `user_id` | UUID | Nullable. Foreign key linking to authenticated owner |
| `token_hash` | String | Hashed version of the JWT (for quick lookup and revocation) |
| `expires_at` | Timestamp | When the token expires (24h from creation) |
| `consumed_at` | Timestamp | Nullable. Set when the token is successfully used |
| `revoked_at` | Timestamp | Nullable. Set if the token is revoked (e.g. superseded) |
| `issued_context_hash` | String | Nullable. Privacy-preserving hash of network metadata/IP upon issuance |
| `consumed_context_hash` | String | Nullable. Privacy-preserving hash of network metadata/IP upon consumption |
| `created_at` | Timestamp | Creation time |

*Indexes:*
- `idx_resume_token_session` on `session_id`
- `idx_resume_token_hash` on `token_hash`

### 2. `NotificationLog`
Tracks the delivery status of notifications for auditing and debugging.

| Field | Type | Description |
|---|---|---|
| `id` | UUID | Primary key |
| `session_id` | UUID | Foreign key to `DecisionSession` |
| `provider` | String | e.g., 'Zalo' |
| `status` | String | 'Accepted', 'Failed', 'RetryableFailure' |
| `error_details` | JSONB | Nullable. Stores API error responses |
| `created_at` | Timestamp | Creation time |

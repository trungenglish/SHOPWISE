# Data Model: Decision Memory

## PostgreSQL Entities (Go Backend)

### `decision_sessions`
Stores the metadata for a shopping decision session.
- `id` (UUID, PK)
- `user_id` (UUID, FK to users, NULLable for anonymous)
- `anonymous_id` (String, for local storage sync)
- `title` (String) - Auto-generated based on first query.
- `status` (Enum: `active`, `archived`, `deleted`)
- `parent_session_id` (UUID, FK to decision_sessions, NULLable) - For branched sessions.
- `created_at` (Timestamp)
- `updated_at` (Timestamp) - Used for LWW conflict resolution.

### `session_messages`
Stores individual conversation turns within a session.
- `id` (UUID, PK)
- `session_id` (UUID, FK to decision_sessions)
- `role` (Enum: `user`, `assistant`, `system`)
- `content` (Text)
- `reasoning_graph` (JSONB, NULLable) - Stores the AI thought process steps.
- `pinned_products` (JSONB, NULLable) - Array of product IDs pinned at this step.
- `created_at` (Timestamp)

### `user_preferences`
Stores extracted implicit and explicit preferences.
- `id` (UUID, PK)
- `user_id` (UUID, FK to users)
- `category` (String) - e.g., 'budget', 'brand', 'usage'.
- `value` (JSONB) - e.g., `{"max": 1000}` or `["Asus", "Dell"]`.
- `source_session_id` (UUID, FK to decision_sessions, NULLable) - Where this preference was extracted.
- `created_at` (Timestamp)
- `updated_at` (Timestamp)

## Client-Side Entities (LocalStorage)
For unauthenticated users, `decision_sessions` and `session_messages` are serialized as a single JSON array under the key `shopwise_anonymous_sessions`. A rolling window of length 10 is enforced.

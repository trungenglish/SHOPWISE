# API Contracts: Decision Memory

## Base URL
`/api/v1`

## Common Headers
- `Authorization`: Bearer `<token>` (Optional, if missing falls back to anonymous storage strategy on client)
- `X-Client-Timestamp`: ISO8601 string used for Conflict Resolution.

## Endpoints

### 1. Create a Session
`POST /sessions`
Creates a new decision session.

**Request Body:**
```json
{
  "initial_message": "I'm looking for a gaming laptop under $1500"
}
```

**Response (201 Created):**
```json
{
  "id": "uuid",
  "title": "Gaming laptop under $1500",
  "status": "active",
  "created_at": "2026-07-11T12:00:00Z"
}
```

### 2. Get Sessions
`GET /sessions?status=active|archived`
Lists sessions for the current user.

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": "uuid",
      "title": "Gaming laptop under $1500",
      "status": "active",
      "updated_at": "2026-07-11T12:00:00Z"
    }
  ],
  "meta": { "total": 1 }
}
```

### 3. Get Session Details
`GET /sessions/:id`
Retrieves full session history including reasoning.

**Response (200 OK):**
```json
{
  "id": "uuid",
  "title": "Gaming laptop",
  "messages": [
    {
      "id": "uuid",
      "role": "assistant",
      "content": "Here are 3 options...",
      "pinned_products": ["prod-1", "prod-2"],
      "reasoning_graph": {
        "steps": ["Extracted budget < $1500", "Filtered GPUs > RTX 4060"]
      }
    }
  ]
}
```

### 4. Update Session (Auto-Save)
`PUT /sessions/:id`
Updates the state of a session. Uses Last-Write-Wins based on `X-Client-Timestamp`.

**Request Body:**
```json
{
  "pinned_products": ["prod-1", "prod-3"]
}
```

**Response (200 OK):** Success.
**Response (409 Conflict):** If `X-Client-Timestamp` is older than server's `updated_at`.

### 5. Branch Session
`POST /sessions/:id/branch`
Creates a new session from a specific message point.

**Request Body:**
```json
{
  "message_id": "uuid-of-branch-point"
}
```

**Response (201 Created):**
```json
{
  "id": "new-session-uuid",
  "parent_session_id": "old-session-uuid"
}
```

### 6. Manage Preferences
`GET /preferences`
`DELETE /preferences/:id` (For Undo actions)

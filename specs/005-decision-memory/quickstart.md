# Validation Quickstart: Decision Memory

This guide provides steps to validate the Decision Memory APIs once implemented locally.

## Prerequisites
- Local Postgres database running.
- Go backend running (`pnpm --filter server dev`).
- Optional: Web frontend running (`pnpm --filter web dev`).

## Scenario 1: Create and Auto-Save Session
1. **Create Session**
   ```bash
   curl -X POST http://localhost:8080/api/v1/sessions \
     -H "Content-Type: application/json" \
     -d '{"initial_message": "Looking for a laptop"}'
   ```
   *Expected: Returns 201 with session ID.*

2. **Auto-Save (Valid Timestamp)**
   ```bash
   curl -X PUT http://localhost:8080/api/v1/sessions/<SESSION_ID> \
     -H "Content-Type: application/json" \
     -H "X-Client-Timestamp: 2099-01-01T00:00:00Z" \
     -d '{"pinned_products": ["p1"]}'
   ```
   *Expected: Returns 200 OK.*

## Scenario 2: Test Conflict Resolution
1. **Trigger Conflict (Old Timestamp)**
   ```bash
   curl -X PUT http://localhost:8080/api/v1/sessions/<SESSION_ID> \
     -H "Content-Type: application/json" \
     -H "X-Client-Timestamp: 2000-01-01T00:00:00Z" \
     -d '{"pinned_products": ["p1", "p2"]}'
   ```
   *Expected: Returns 409 Conflict. The server's `updated_at` is newer.*

## Scenario 3: Branch Session
1. **Branch from an existing message**
   ```bash
   curl -X POST http://localhost:8080/api/v1/sessions/<SESSION_ID>/branch \
     -H "Content-Type: application/json" \
     -d '{"message_id": "<MESSAGE_ID>"}'
   ```
   *Expected: Returns 201 with a NEW session ID, and `parent_session_id` pointing to the original.*

## Scenario 4: UI Validation
1. Open the frontend at `http://localhost:5173`.
2. Start a chat. Look for the "Saved to Decision Memory" indicator.
3. Open a new tab, navigate to the site. Verify the session is restored.

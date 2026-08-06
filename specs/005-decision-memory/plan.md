# Implementation Plan: Decision Memory

## Technical Context
- **Frontend**: React + Vite + TypeScript
- **Backend**: Go + Gin
- **AI Runtime**: Python + FastAPI + LangGraph
- **State Management**: Persistence exclusively in Go backend. AI Runtime is stateless.

## Technical Architecture
1. **Frontend (`apps/web`)**: 
   - Uses TanStack Query for session data fetching and optimistic updates.
   - Debounced auto-save mutations triggered on conversation/preference changes.
   - Renders declarative UI schemas for Reasoning Replay.
2. **Backend (`apps/server`)**: 
   - Go Gin REST API.
   - PostgreSQL handles persistence.
   - Implements Session Lifecycle (90-day archive, 10-session anonymous limit).
   - Handles Branching (full-copy deep insert).
   - Handles Conflict-resolution (Last-Write-Wins via `updated_at` validation).
3. **AI Runtime (`apps/ai-service`)**:
   - Stateless Python LangGraph engine.
   - Receives full session context + preferences from Go backend per request.

## Workflows

### 1. Auto-Save & Conflict Resolution Flow
- **Client**: Debounces user actions (new message, pin product). Sends `PUT /sessions/:id` with `client_updated_at`.
- **Server**: Compares `client_updated_at` with DB `updated_at`. If client is newer or equal, apply. If older, reject with 409 Conflict.
- **Client**: On 409, refetches session to get latest state (Last-Write-Wins). UI shows subtle "Saved to Decision Memory" indicator on success.

### 2. Session Branching Flow
- **Client**: User clicks "Branch from here" on an old message. Sends `POST /sessions/:id/branch` with `message_id`.
- **Server**: Duplicates the session row and all messages up to `message_id`. Sets `parent_session_id`. Returns new `session_id`.
- **Client**: Navigates to new `session_id`.

### 3. Archive and Restore Flow
- **Server Cron**: Nightly job flags `status = 'archived'` for sessions where `updated_at < NOW() - 90 days`.
- **Client**: Can query `/sessions?status=archived`. Clicking an archived session sends `POST /sessions/:id/restore` to flip status back to `active`.

### 4. Preferences Flow
- **AI Runtime**: Detects implicit preference in conversation. Returns it in response payload.
- **Server**: Saves to `user_preferences`.
- **Client**: Shows non-blocking toast: "Preference saved: Budget under $1000 [Undo]". Clicking Undo sends `DELETE /preferences/:id`.

## Implementation Phases

**Phase 1: Database & Core Models**
- Create migrations for `decision_sessions`, `session_messages`, `user_preferences`.
- Implement GORM models and repository functions (including the full-copy branch function).

**Phase 2: Go REST API**
- Implement session CRUD endpoints.
- Implement conflict resolution middleware/logic.
- Implement preference extraction endpoints.

**Phase 3: Frontend Integration**
- Add `@shopwise/api-types` for contracts.
- Build TanStack Query hooks.
- Implement auto-save debouncing and the subtle "Saved" indicator.

**Phase 4: UI Development**
- Build Session Sidebar (History).
- Build Reasoning Replay UI (showing AI thought steps).
- Preference Manager dialog.

**Phase 5: Testing & Polish**
- E2E tests for branching and conflict resolution.
- Verify anonymous local-storage to authenticated DB sync.

# Key Workflows

This page describes the main workflows that exist today, with notes where the product specs go beyond the current implementation.

## 1. Landing page to dashboard exploration

### Current code path
- Route `/` renders `LandingPage` from `/apps/web/src/routes/index.tsx`
- The main interactive workspace lives at `/dashboard` in `/apps/web/src/routes/dashboard.tsx`

### What the dashboard currently does
The dashboard is a composite workspace for agent-assisted shopping. It manages:
- user intent
- simulated product matching results
- audit logs
- trust score/factors
- graph/spatial workspace data
- accessory recommendations
- retail-account toggles
- price alerts
- modal-driven comparison/reasoning/retail/checkout flows
- order history display

This is implemented mostly as route-local React state today. The route is large and acts as an orchestration layer rather than a thin route wrapper.

### Git context
Recent commits indicate the dashboard has been the primary surface for rapid product iteration:
- `721655c` added trust center, audit trail, and spatial workspace
- `7451524` added React Flow visualization support
- `f140d9c` connected dashboard work to catalog/store and AI runtime infrastructure
- `37d307b` consolidated routing and core UI around audit trails and price monitoring

## 2. Decision memory session lifecycle

### Current API surface
The web app uses `/apps/web/src/api/decision-memory.ts` to talk to backend routes under `/api/v1/sessions`.

Supported operations include:
- create session
- list sessions
- get session
- rename session
- delete session
- branch session
- restore session
- auto-save session metadata
- list/update/delete preferences

### Identity model
`/services/main-backend/internal/decision_memory/handler/handler.go` supports both:
- authenticated users via bearer token
- anonymous users via `X-Anonymous-ID`

If neither is present, the request is rejected.

### Important business rule
Anonymous users are capped at **10 sessions**. `CreateSession` in `/services/main-backend/internal/decision_memory/usecase/service.go` counts anonymous sessions and deletes the oldest when the limit is reached.

### Concurrency model
The decision memory spec chooses **last-write-wins based on timestamp**, and the handler/service shape reflects that direction via `X-Client-Timestamp` and update conflict handling.

### What is implemented vs intended
Implemented:
- session CRUD shape
- branching entrypoint
- restore entrypoint
- anonymous identity handling
- preference endpoints

Still thin or partial:
- richer persistence of pinned products/workspace state during autosave
- deeper reasoning replay integration across the full dashboard
- mature session-search/sort/archive UX described in the spec

## 3. Session resume via secure token

### Current flow
1. The frontend enters on `/resume` with a `token` query parameter in `/apps/web/src/routes/resume.tsx`.
2. It calls `resumeSessionFromToken()` in `/apps/web/src/api/decision-memory.ts`.
3. The backend validates the token in `/services/main-backend/internal/resume_session/usecase/service.go`.
4. On success, the route clears the token from the URL, invalidates session queries, and navigates to `/dashboard` with `sessionId` search params.
5. On failure, the route renders `ExpiredTokenUI` with error-specific behavior.

### Security model in code
`ResumeSession` currently checks:
- JWT validity
- token hash lookup
- token JTI/session match
- revoked state
- consumed state
- expiry window

It also records a context mismatch as a risk signal without blocking, which matches the spec's cross-device resume goal.

### Notification path
`TriggerInactivityNotification()` exists in the resume service and:
- finds the session
- finds a verified notification identity
- generates a new resume token
- sends it through a pluggable provider
- stores a notification log

The current bootstrap wires a Zalo provider placeholder in `/services/main-backend/internal/bootstrap/bootstrap.go`, using dummy constructor values in the inspected code path.

### Spec vs implementation
The spec in `/specs/007-resume-shopping-session/spec.md` describes a more complete inactivity-triggered workflow than the currently visible code. The current implementation covers token issuance/verification and notification service structure, but not a fully visible scheduler-driven inactivity pipeline in the files inspected here.

## 4. Chat streaming through the backend to the AI runtime

### Current flow
1. Client posts a message for a decision session.
2. `/services/main-backend/internal/decision_memory/handler/chat.go` validates the session and stores the user message.
3. The handler proxies a request to the AI runtime streaming endpoint.
4. SSE data is streamed back to the client.
5. The accumulated assistant text is stored as a backend session message after the stream completes.

### Important caveat
The proxy currently targets `http://localhost:8001/chat/stream`, while the AI runtime itself exposes `/api/v1/chat/stream` in `/services/ai-runtime/src/api/chat.py` and its README documents port `8000`. This is a current integration mismatch and should be treated as a known issue, not as intended architecture.

## 5. AI runtime structured response workflow

### Current flow in Python
`/services/ai-runtime/src/api/chat.py`:
- validates the request
- resolves the OpenAI provider from env config
- builds a system prompt
- creates a LangGraph workflow
- invokes the workflow with session/message state

`/services/ai-runtime/src/graph/workflow.py`:
- calls the provider for a structured response
- attempts JSON parsing
- retries the LLM loop up to 2 times on malformed JSON
- returns the final response or an error state marker

### Response shape
The intended structured output is defined by `/services/ai-runtime/src/models/schemas.py`, including:
- products
- audit logs
- agent statuses
- accessories
- reasoning
- trust score

This aligns closely with the dashboard's trust/audit/agent-based product storytelling.

## 6. Checkout and order history

### Current flow
- Checkout creation and order listing are handled by `/services/main-backend/internal/orders/handler/handler.go`.
- The web order history experience lives in `/apps/web/src/features/orders/components/order-history-page.tsx`.
- The dashboard can display order history alongside the broader workspace.

### Operational detail
The backend supports a development auth bypass toggle via `CHECKOUT_AUTH_BYPASS`, wired in bootstrap. This is useful for local development, but docs and tests should clearly distinguish bypass mode from real authenticated behavior.

## 7. Catalog and store lookup

Recent backend work added lightweight handlers for:
- `/api/v1/products` via `/services/main-backend/internal/catalog/handler/handler.go`
- `/api/v1/stores` via `/services/main-backend/internal/stores/handler/handler.go`

These are simple but important because the dashboard now expects product/store surfaces to exist.

## Change guidance by workflow

- **Landing/dashboard**: keep route-level complexity from growing unless the change truly coordinates multiple features.
- **Decision memory**: update both web API helpers and backend handler/usecase logic together.
- **Resume flow**: preserve single-use token guarantees and error-code mapping to the resume UI.
- **AI streaming**: verify endpoint paths and ports across both services before debugging payload issues.
- **Checkout/orders**: test both authenticated mode and any local bypass assumptions.

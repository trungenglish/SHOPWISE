# Architecture Overview

## Top-level shape

SHOPWISE is organized as a product monorepo:

- `apps/web`: primary customer-facing web UI
- `apps/mobile`: mobile app shell, present but not the focus of recent work
- `services/main-backend`: Go API and business logic
- `services/ai-runtime`: Python FastAPI service for LLM-backed structured responses
- `packages/*`: shared libraries such as UI, env, schemas, SDK, protocols, and types
- `specs/*`: specification-first product artifacts for major feature epics

The workspace is managed by `pnpm` and `turbo` (`/package.json`, `/pnpm-workspace.yaml`, `/turbo.json`).

## Runtime architecture

### Web app
The web app boots in `/apps/web/src/main.tsx`:
- creates a TanStack Router instance from generated routes
- wraps the app in React Query, theme, and font providers
- uses a loading component as the default pending UI

The router root in `/apps/web/src/routes/__root.tsx` adds:
- global theme and tooltip providers
- toast notifications
- TanStack devtools in development
- shared not-found and error components

Routing is file-based under `/apps/web/src/routes`:
- `/` → landing page
- `/dashboard` → main agent-assisted shopping workspace
- `/chat` → simple conversation route using `ChatWorkspace`
- `/resume` → secure session resume entrypoint
- `/get-started` → onboarding flow

### Go backend
The Go backend is a modular monolith. `cmd/retail/main.go` delegates almost everything to `internal/bootstrap/bootstrap.go`, which:
- loads config
- connects Postgres and Redis
- creates repositories and use-case services
- wires HTTP handlers
- registers `/api/v1` route groups
- starts a Gin-based HTTP server with graceful shutdown

Current registered API areas in bootstrap:
- `identity`
- `users`
- `checkout`
- `orders`
- `files`
- `products`
- `stores`
- `sessions` (decision memory)
- `session` (resume flow)
- `admin`

This file is the best single place to understand real system composition.

### AI runtime
The AI runtime is intentionally separate from the Go backend. It exposes FastAPI endpoints from `/services/ai-runtime/src/main.py` and `/services/ai-runtime/src/api/chat.py`.

Key architectural choices visible in code and git history:
- recent refactor `8044a29` moved the service into a cleaner `src/` layout
- provider-specific logic is isolated behind `OpenAIProvider`
- prompt construction lives in `src/core/prompts.py`
- the workflow uses LangGraph in `src/graph/workflow.py`
- response shape is governed by Pydantic models in `src/models/schemas.py`

The runtime is described as stateless in `/services/ai-runtime/README.md`, which aligns with the architecture: session persistence belongs in the Go backend, not in Python process memory.

## Integration boundaries

### Web ↔ Go backend
The web app calls backend APIs directly for stateful product features. Examples:
- `/apps/web/src/api/decision-memory.ts` calls `/api/v1/sessions` and `/api/v1/session/resume`
- orders UI calls checkout/order APIs through its feature services
- the dashboard route directly fetches products from a backend endpoint

The Go backend is therefore the main source of truth for:
- sessions and messages
- users and identity
- orders and checkout
- product and store catalog surfaces
- resume token validation

### Go backend ↔ AI runtime
The clearest current integration is `/services/main-backend/internal/decision_memory/handler/chat.go`, which:
- accepts a message for a saved session
- writes the user message to backend storage
- proxies a streaming request to the AI runtime
- relays the SSE response back to the client
- saves the assistant response after streaming finishes

This shows the intended ownership split:
- backend owns persistence and session identity
- AI runtime owns model invocation and structured reasoning output

### Shared packages
The most important shared packages for navigation are:
- `packages/ui`: shared visual primitives and contexts
- `packages/env`: web/mobile environment validation exports
- `packages/api-types`, `packages/schemas`, `packages/types`: cross-package type surfaces

The repo uses these packages to keep app-level code thinner, though most product-specific logic still sits in `apps/web` and `services/main-backend`.

## Major architecture patterns

### 1. Modular monolith for business logic
The backend uses `internal/<domain>` folders with handler/usecase/repository layering. This is visible in domains such as:
- `decision_memory`
- `resume_session`
- `orders`
- `identity`
- `users`

This keeps domain logic in one deployable service while preserving separation between transport, orchestration, and persistence.

### 2. Spec-first product development
The `specs/*` tree is unusually important here. For several features, the specs explain product behavior more completely than the source currently enforces. Engineers should treat them as product intent, not as guaranteed implementation truth.

### 3. Stateful backend + stateless AI service
A recurring pattern across recent changes is that customer/session state is persisted in the Go service, while the Python runtime remains replaceable and focused on model interaction.

### 4. UI-first iteration in the dashboard
Recent commits show rapid iteration in the dashboard route and feature components. The architecture is functional, but the route still carries a lot of orchestration and demo logic that may eventually need extraction into hooks/services.

## Architectural tensions and inconsistencies

These are the main things future changes should watch:

- **Port mismatches** between README/env defaults and hardcoded callers (`18080` vs `8080`, `8000` vs `8001`).
- **Placeholder and mock-heavy frontend areas**, especially `ChatWorkspace` and large parts of `dashboard.tsx`.
- **Spec/code gaps**, especially around robust inactivity detection, bounded retries, and production-grade notification handling.
- **Stale scaffold docs** in the root README that no longer fully describe the active service layout.

## Change guidance

- When changing system wiring, start with `/services/main-backend/internal/bootstrap/bootstrap.go`.
- When changing data ownership, default to the Go backend unless the change is purely model-runtime behavior.
- When changing conversation shape or tool/response schema, update both `/services/ai-runtime/src/models/schemas.py` and the frontend consumers.
- When changing dashboard behavior, check whether the logic belongs in route state, a feature component, or an API/service layer before adding more code to `dashboard.tsx`.

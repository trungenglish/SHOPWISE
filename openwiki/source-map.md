# Source Map

A practical navigation map for the repository.

## Top-level directories

- `/apps` — application entrypoints
- `/services` — backend and AI services
- `/packages` — shared internal libraries
- `/specs` — feature specifications and implementation planning artifacts
- `/infra`, `/deploy`, `/scripts` — supporting infrastructure and tooling
- `/openwiki` — generated repository wiki

## Web app (`/apps/web`)

### Start here
- `/apps/web/src/main.tsx` — app bootstrap
- `/apps/web/src/routes/__root.tsx` — router root, global providers, devtools
- `/apps/web/src/routes/index.tsx` — landing page route
- `/apps/web/src/routes/dashboard.tsx` — main workspace route and current orchestration hotspot
- `/apps/web/src/routes/resume.tsx` — token-based session resumption UI

### Feature areas
- `/apps/web/src/features/dashboard` — dashboard-specific components and types
- `/apps/web/src/features/decision-memory` — reasoning replay, branching, preference UI
- `/apps/web/src/features/orders` — order history and order detail UI
- `/apps/web/src/features/get-started` — onboarding/auth-related setup flows
- `/apps/web/src/features/landing` — marketing/landing page sections
- `/apps/web/src/features/errors` — route-level error displays

### API helpers and hooks
- `/apps/web/src/api/decision-memory.ts` — session CRUD, branching, restore, resume
- `/apps/web/src/api/checkout.ts` — checkout interactions
- `/apps/web/src/hooks` — reusable UI/data hooks

## Go backend (`/services/main-backend`)

### Start here
- `/services/main-backend/cmd/retail/main.go` — service entrypoint
- `/services/main-backend/internal/bootstrap/bootstrap.go` — full dependency and route wiring
- `/services/main-backend/.env.example` — non-secret config shape
- `/services/main-backend/package.json` — developer commands

### Important domains
- `/services/main-backend/internal/decision_memory` — session persistence, preferences, branching, and AI-runtime chat proxying
- `/services/main-backend/internal/resume_session` — single-use token flow and notification plumbing
- `/services/main-backend/internal/orders` — checkout/order APIs
- `/services/main-backend/internal/catalog` — product listing handler
- `/services/main-backend/internal/stores` — store listing handler
- `/services/main-backend/internal/identity` — auth and Google OAuth integration
- `/services/main-backend/internal/users` — user management
- `/services/main-backend/internal/platform` — shared backend infrastructure: config, DB, cache, middleware, logging, router

### Useful backend file entrypoints by task
- **Understand route registration**: `internal/bootstrap/bootstrap.go`
- **Change session API behavior**: `internal/decision_memory/handler/handler.go`
- **Change session business rules**: `internal/decision_memory/usecase/service.go`
- **Change AI streaming proxy**: `internal/decision_memory/handler/chat.go`
- **Change resume token logic**: `internal/resume_session/usecase/service.go`
- **Change checkout/order behavior**: `internal/orders/handler/handler.go` and matching usecase files

## AI runtime (`/services/ai-runtime`)

### Start here
- `/services/ai-runtime/src/main.py` — FastAPI app, middleware, exception handling
- `/services/ai-runtime/src/api/chat.py` — chat and streaming endpoints
- `/services/ai-runtime/src/graph/workflow.py` — LangGraph orchestration
- `/services/ai-runtime/src/models/schemas.py` — structured response schema
- `/services/ai-runtime/README.md` — local run/test basics

### Other areas
- `/services/ai-runtime/src/core` — config, prompts, observability
- `/services/ai-runtime/src/llm` — provider abstraction and OpenAI implementation
- `/services/ai-runtime/tests/unit` — workflow/provider unit tests
- `/services/ai-runtime/tests/integration` — endpoint-level tests

## Shared packages (`/packages`)

Most likely to matter in day-to-day work:
- `/packages/ui` — shared UI primitives and providers
- `/packages/env` — env exports for web/mobile
- `/packages/api-types` — shared API type surfaces
- `/packages/schemas` — schema utilities/contracts
- `/packages/types` — shared type definitions
- `/packages/sdk` and `/packages/protocols` — extension points for richer integration work

## Specs (`/specs`)

High-signal specs to read before editing product flows:
- `/specs/005-decision-memory/spec.md`
- `/specs/007-checkout-readiness/spec.md`
- `/specs/007-resume-shopping-session/spec.md`
- `/specs/008-llm-integration/spec.md`

These specs are especially helpful when the code is incomplete, recently changing, or still mock-heavy.

## Testing locations

- Web tests: `/apps/web/src/**/*.test.tsx`
- Backend tests: domain-specific `_test.go` files under `/services/main-backend/internal/**`
- AI runtime tests: `/services/ai-runtime/tests/unit` and `/services/ai-runtime/tests/integration`

## Best first reads by change type

- **Dashboard UX**: route + `features/dashboard/components`
- **Decision memory**: web API helper + backend handler/usecase + spec 005
- **Resume links**: resume route + backend resume service + spec 007 resume
- **AI response schema**: runtime `models/schemas.py` + web consumers + spec 008
- **Developer setup**: root README + backend `.env.example` + runbook page

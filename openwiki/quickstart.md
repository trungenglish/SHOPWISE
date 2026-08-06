# SHOPWISE OpenWiki Quickstart

SHOPWISE is a monorepo for an AI-assisted shopping and procurement experience. The current implementation centers on a React web app, a Go backend that owns persisted state and business APIs, and a Python AI runtime that handles LLM-backed structured responses.

This wiki is a practical map for engineers and future agents. Start here, then follow the section links for detail.

## What this repository currently contains

- **Web app (`apps/web`)**: TanStack Router + React Query frontend with a landing page, dashboard workspace, resume flow, and order history UI. Entry points include `/apps/web/src/main.tsx`, `/apps/web/src/routes/index.tsx`, `/apps/web/src/routes/dashboard.tsx`, and `/apps/web/src/routes/resume.tsx`.
- **Main backend (`services/main-backend`)**: Go modular monolith exposing REST APIs under `/api/v1`, bootstrapped in `/services/main-backend/internal/bootstrap/bootstrap.go` and started from `/services/main-backend/cmd/retail/main.go`.
- **AI runtime (`services/ai-runtime`)**: FastAPI service with provider abstraction and a LangGraph workflow for structured LLM output in `/services/ai-runtime/src/api/chat.py` and `/services/ai-runtime/src/graph/workflow.py`.
- **Shared packages (`packages/*`)**: UI primitives, environment validation, API types, schemas, and supporting internal libraries.
- **Specs (`specs/*`)**: Product intent for major features such as decision memory, checkout readiness, resume shopping sessions, and LLM integration.

## Read this next

- [Architecture overview](./architecture/overview.md)
- [Key workflows](./workflows/key-workflows.md)
- [Domain concepts](./domain/domain-concepts.md)
- [Operations runbook](./operations/runbook.md)
- [Testing guide](./testing/testing-guide.md)
- [Integration points](./integrations/integration-points.md)
- [Source map](./source-map.md)

## Major technical and product domains

### 1. AI-assisted shopping workspace
The dashboard route is the current center of the product. It composes a sidebar, audit trail, trust center header, spatial workspace, agent hub, comparison modal, reasoning modal, retail modal, checkout modal, and price alert modal inside `/apps/web/src/routes/dashboard.tsx` and `apps/web/src/features/dashboard/components/*`.

Recent git history shows this area is evolving fastest:
- `721655c`: trust center, audit trail, spatial workspace
- `7451524`: React Flow-based spatial workspace
- `f140d9c`: catalog/store wiring and AI runtime infrastructure support
- `37d307b`: dashboard routing and core UI consolidation including price monitoring

### 2. Decision memory and session continuity
Decision memory is both a product idea and a real backend surface area. The current code supports creating, listing, renaming, restoring, branching, and deleting sessions through `/api/v1/sessions`, backed by `/services/main-backend/internal/decision_memory/*` and consumed by `/apps/web/src/api/decision-memory.ts`.

Important implementation details:
- anonymous users are identified via `X-Anonymous-ID`
- anonymous session history is capped at **10 sessions** and oldest sessions are evicted in `/services/main-backend/internal/decision_memory/usecase/service.go`
- update conflict handling is timestamp-based and aligns with the spec's last-write-wins direction

### 3. Resume shopping sessions
Resume links are implemented as single-use JWT-backed tokens in `/services/main-backend/internal/resume_session/usecase/service.go` and surfaced in the web app through `/apps/web/src/routes/resume.tsx`.

The product intent comes from `/specs/007-resume-shopping-session/spec.md`: proactive re-engagement after inactivity, Zalo delivery, and secure cross-device resumption.

### 4. Checkout and orders
Checkout and order listing live in the Go backend under `/services/main-backend/internal/orders/*`, with web UI in `/apps/web/src/features/orders/*` and supporting checkout UI inside the dashboard feature set.

### 5. AI runtime orchestration
The Python service is designed to stay stateless: requests include the conversational input, and the runtime produces structured output defined by `/services/ai-runtime/src/models/schemas.py`. Its workflow currently performs an LLM call, attempts JSON parsing, and retries on malformed JSON in `/services/ai-runtime/src/graph/workflow.py`.

## Important current-state caveats

These are worth knowing before making changes:

- **README drift exists.** The root `/README.md` still contains scaffold-era notes such as `services/agentic`, but the active AI service is now `services/ai-runtime`.
- **Port mismatches exist in code.** The backend defaults to port `18080` in `/README.md` and `/services/main-backend/.env.example`, but the dashboard still fetches products from `http://localhost:8080/api/v1/products` inside `/apps/web/src/routes/dashboard.tsx`.
- **AI runtime proxying is inconsistent.** `/services/ai-runtime/README.md` uses port `8000`, while `/services/main-backend/internal/decision_memory/handler/chat.go` proxies to `http://localhost:8001/chat/stream` and omits the runtime's `/api/v1` prefix.
- **Some UI is still placeholder/demo-heavy.** `ChatWorkspace` is currently a placeholder in `/apps/web/src/components/chat-workspace.tsx`, and much of the dashboard route still uses simulated state and mock orchestration.
- **Specs are ahead of code in some areas.** The specs provide richer intended behavior than the currently wired implementation, especially around inactivity detection, delivery retries, and mature tool-calling flows.

## Where to start for common tasks

- **Change dashboard UX or workspace composition**: start at `/apps/web/src/routes/dashboard.tsx`, then inspect `apps/web/src/features/dashboard/components/*`.
- **Change session persistence or memory behavior**: start at `/services/main-backend/internal/decision_memory/handler/handler.go` and `/services/main-backend/internal/decision_memory/usecase/service.go`.
- **Change resume-link behavior**: start at `/services/main-backend/internal/resume_session/usecase/service.go` and `/apps/web/src/routes/resume.tsx`.
- **Change backend route wiring**: start at `/services/main-backend/internal/bootstrap/bootstrap.go`.
- **Change AI runtime behavior**: start at `/services/ai-runtime/src/api/chat.py`, `/services/ai-runtime/src/graph/workflow.py`, and `/services/ai-runtime/src/models/schemas.py`.
- **Understand product intent before editing a feature**: read the matching spec under `/specs`.

## Suggested reading order for future agents

1. This page
2. [Architecture overview](./architecture/overview.md)
3. [Key workflows](./workflows/key-workflows.md)
4. [Domain concepts](./domain/domain-concepts.md)
5. The relevant spec for the feature you are touching
6. The matching source files from the [Source map](./source-map.md)

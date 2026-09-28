---
type: quickstart
title: Repository Quickstart
description: Practical routing guide for coding agents working in the SHOPWISE monorepo. Use it to find the smallest set of repo-specific pages and source entrypoints for UI work, Go backend changes, AI runtime work, schema/protocol updates, operations, and testing.
tags: [quickstart, onboarding, monorepo, web, backend, ai-runtime, testing]
openwiki_generated: true
verified:
  - by: openwiki/0.6.0
    at: 2026-09-28T02:48:19.098Z
sources:
  - id: openwiki-source-8037e2358a2c4f9b2c722a11
    resource: repo://AGENTS.md
  - id: openwiki-source-99de51df25f29bfc72caf823
    resource: repo://apps/web/package.json
  - id: openwiki-source-970773f08ee82250bf68a237
    resource: repo://apps/web/src/api/decision-memory.test.ts
  - id: openwiki-source-313b77d3166c4ef5abe361e6
    resource: repo://apps/web/src/api/decision-memory.ts
  - id: openwiki-source-cd9de6d1c0464c8699d7bc88
    resource: repo://apps/web/src/features/checkout-incentive/components/CheckoutIncentivePanel.test.tsx
  - id: openwiki-source-8e8b395281c4996e784ae3b5
    resource: repo://apps/web/src/main.tsx
  - id: openwiki-source-5f4beeeff5a41258071dfac9
    resource: repo://apps/web/src/routes/dashboard.tsx
  - id: openwiki-source-8c5e10b4fcf9b1ec4222179e
    resource: repo://apps/web/src/routes/resume.test.tsx
  - id: openwiki-source-6487f656f578243180fa4c8a
    resource: repo://apps/web/src/routes/resume.tsx
  - id: openwiki-source-5b54a58d1b51cd490b0e7162
    resource: repo://package.json
  - id: openwiki-source-40275cb92c3610938f16ade3
    resource: repo://pnpm-workspace.yaml
  - id: openwiki-source-23775c3de52f3ab95a13cb8b
    resource: repo://README.md
  - id: openwiki-source-8387cfa6d142c9c9d41e63ec
    resource: repo://services/ai-runtime/pyproject.toml
  - id: openwiki-source-85cedde5d4fbfc25098c14b5
    resource: repo://services/ai-runtime/src/api/chat.py
  - id: openwiki-source-95c5cf8058175b6f67241cb5
    resource: repo://services/ai-runtime/src/graph/workflow.py
  - id: openwiki-source-33b5df8a895d8620b414409a
    resource: repo://services/ai-runtime/src/main.py
  - id: openwiki-source-6dca1297d7084a2d4bb69450
    resource: repo://services/ai-runtime/tests/unit/test_graph.py
  - id: openwiki-source-640bb3dc56ec28e911b3c25b
    resource: repo://services/ai-runtime/tests/unit/test_self_correction.py
  - id: openwiki-source-955812e23455ffe14dd9caf2
    resource: repo://services/main-backend/cmd/retail/main.go
  - id: openwiki-source-4d5a7af019fe73163caed918
    resource: repo://services/main-backend/internal/bootstrap/bootstrap.go
  - id: openwiki-source-8c46307bb52f1d31e366fb3c
    resource: repo://services/main-backend/internal/decision_memory/handler/chat_test.go
  - id: openwiki-source-4f2c192847e0fbe283a0e627
    resource: repo://services/main-backend/internal/decision_memory/handler/chat.go
  - id: openwiki-source-ef5feebdf4490c56cedcb217
    resource: repo://services/main-backend/internal/decision_memory/handler/greeting.go
  - id: openwiki-source-4b8a5d433e7f00b922044ec6
    resource: repo://services/main-backend/internal/decision_memory/handler/interaction.go
  - id: openwiki-source-aab2acaa3eabb8d48be370be
    resource: repo://services/main-backend/internal/orders/usecase/service_test.go
  - id: openwiki-source-f49b429fbb5c73081f714c22
    resource: repo://services/main-backend/internal/platform/config/config.go
  - id: openwiki-source-9c0e62431401592f7587d658
    resource: repo://services/main-backend/internal/promotions/service_test.go
  - id: openwiki-source-c107c02cfc1ca9c0522cab76
    resource: repo://services/main-backend/internal/resume_session/usecase/service.go
  - id: openwiki-source-b76a40124a4e52afe2d5f0f9
    resource: repo://services/main-backend/package.json
  - id: openwiki-source-440ae1e215cb02721dda855c
    resource: repo://turbo.json
generated: { by: "openwiki/0.6.0", at: "2026-09-28T02:48:19.098Z" }
---

SHOPWISE is a pnpm + Turborepo monorepo with three primary runtimes that matter for most tasks: a React web app in `apps/web`, a Go backend in `services/main-backend`, and a Python FastAPI AI runtime in `services/ai-runtime`.

Start here for routing, then move to the linked wiki pages and the specific source entrypoints below.

## Read this next

- [Architecture overview](./architecture/overview.md) — system shape and service boundaries
- [Key workflows](./workflows/key-workflows.md) — end-to-end request and state flow
- [Domain concepts](./domain/domain-concepts.md) — repository-specific product concepts and invariants
- [Integration points](./integrations/integration-points.md) — external APIs and internal service seams
- [Operations runbook](./operations/runbook.md) — local setup, env, migrations, seeding, and runtime debugging
- [Testing guide](./testing/testing-guide.md) — narrowest effective validation by stack
- [Source map](./source-map.md) — where owned code lives without mirroring the tree

## Fast orientation

- Workspace/package management is driven by `pnpm-workspace.yaml`; it includes `apps/*`, `packages/*`, `services/*`, `conformance/*`, `deploy`, `infra`, and `scripts`.
- Root developer flow is centered on Turborepo: `pnpm dev`, `pnpm build`, `pnpm check-types`, `pnpm deploy`, and `pnpm destroy` fan out through `turbo` tasks.
- The root `package.json` still contains some stale script names like `dev:server`; current backend package naming lives under `services/main-backend`, so prefer package-local scripts and the runbook over the scaffold-era README when commands disagree.

## Task-routing map

### If you are changing the web UI
Start with:
- [Source map](./source-map.md)
- [Testing guide](./testing/testing-guide.md)
- [Domain concepts](./domain/domain-concepts.md) for conversation, checkout, and resume behavior

Primary entrypoints:
- `apps/web/src/main.tsx` bootstraps TanStack Router, React Query, theme, and shared font context.
- `apps/web/src/routes/dashboard.tsx` is the main shopping workspace route and composes the dashboard surface, AI interactions, accessories, checkout, order history, and checkout incentive UI.
- `apps/web/src/api/decision-memory.ts` is the main client seam for sessions, AI turns, greetings, resume flow, and interaction streaming.
- `packages/ui` owns shared UI primitives; `packages/api-types`, `packages/protocols`, and related packages hold shared TypeScript contracts.

Use this path for:
- dashboard UX changes
- AI response rendering
- session/resume screens
- shared design-system work
- frontend wiring for promotions or checkout

### If you are changing the Go backend
Start with:
- [Architecture overview](./architecture/overview.md)
- [Key workflows](./workflows/key-workflows.md)
- [Operations runbook](./operations/runbook.md)

Primary entrypoints:
- `services/main-backend/cmd/retail/main.go` starts the HTTP server and also supports migration and seed flags.
- `services/main-backend/internal/bootstrap/bootstrap.go` is the best top-level map of real backend ownership: config, database, Redis, jobs, route registration, and domain service wiring all meet here.
- `services/main-backend/internal/platform/config/config.go` defines required env, default ports/URLs, and feature toggles such as `AI_RUNTIME_URL`, `CHECKOUT_AUTH_BYPASS`, and `PHONGVU_CONNECTOR_ENABLED`.

Use this path for:
- REST route changes
- database-backed business logic
- checkout/orders/promotions work
- auth/identity integration
- resume session token handling
- cross-service proxying to the AI runtime

### If you are changing the AI runtime
Start with:
- [Architecture overview](./architecture/overview.md)
- [Integration points](./integrations/integration-points.md)
- [Key workflows](./workflows/key-workflows.md)

Primary entrypoints:
- `services/ai-runtime/src/main.py` mounts the FastAPI app, `/api/v1` router, health endpoint, request timing header, and global error handler.
- `services/ai-runtime/src/api/chat.py` defines the `greeting`, `chat`, and `chat/stream` endpoints and creates the provider/tool-proxy backed workflow per request.
- `services/ai-runtime/src/graph/workflow.py` is the core orchestration: load catalog and offers from backend tools, build the system prompt, call the LLM with JSON schema output, hydrate IDs against backend data, and retry invalid outputs up to two times.

Use this path for:
- prompt changes
- output schema behavior
- tool-backed catalog or offer hydration
- streaming semantics
- LLM retry/self-correction logic

### If you are changing protocols or shared schemas
Start with:
- [Domain concepts](./domain/domain-concepts.md)
- [Integration points](./integrations/integration-points.md)
- [Key workflows](./workflows/key-workflows.md)

Touch all of these before shipping:
- `packages/protocols` for shared TS protocol exports
- `apps/web/src/api/decision-memory.ts` for client parsing/validation
- `services/main-backend/internal/decision_memory/handler/*` for request validation, persistence, and proxying
- `services/ai-runtime/src/models/*` and `src/graph/workflow.py` for runtime schema generation/hydration

Safe change rule: treat web, Go backend, and Python runtime as one contract surface. A protocol change is not complete until all three layers agree on accepted input, persisted envelope shape, and streamed/non-streamed output behavior.

### If you are debugging deployment or runtime configuration
Start with:
- [Operations runbook](./operations/runbook.md)
- [Integration points](./integrations/integration-points.md)

Most relevant source:
- root `package.json`, `pnpm-workspace.yaml`, and `turbo.json` for workspace/task behavior
- `services/main-backend/internal/platform/config/config.go` for backend env defaults and required secrets
- `services/ai-runtime/README.md` and `pyproject.toml` for AI runtime local execution and Python toolchain
- root `README.md` only as a secondary reference because it still contains scaffold-era drift

### If you are choosing tests
Start with:
- [Testing guide](./testing/testing-guide.md)

Quick stack map:
- Web: `apps/web` uses Vitest and Testing Library.
- Go backend: `services/main-backend` has package-local Go tests, plus integration-tagged suites.
- AI runtime: `services/ai-runtime/tests` uses pytest with unit and integration coverage.

## Smallest useful starting set by task

- **Dashboard/product experience work:** this page → [Source map](./source-map.md) → `apps/web/src/routes/dashboard.tsx` → [Testing guide](./testing/testing-guide.md)
- **Go API/domain change:** this page → [Architecture overview](./architecture/overview.md) → `services/main-backend/internal/bootstrap/bootstrap.go` → [Testing guide](./testing/testing-guide.md)
- **AI runtime/prompt/output change:** this page → [Key workflows](./workflows/key-workflows.md) → `services/ai-runtime/src/api/chat.py` and `src/graph/workflow.py` → [Testing guide](./testing/testing-guide.md)
- **Resume-link/session continuity work:** this page → [Domain concepts](./domain/domain-concepts.md) → `services/main-backend/internal/resume_session/usecase/service.go` and `apps/web/src/routes/resume.tsx`
- **Checkout/promotion work:** this page → [Key workflows](./workflows/key-workflows.md) → backend `internal/orders` and `internal/promotions` → frontend checkout incentive components/tests
- **Cross-service contract change:** this page → [Integration points](./integrations/integration-points.md) → [Key workflows](./workflows/key-workflows.md) → update web + backend + AI runtime together

## What is actually important in each runtime

### Web runtime
The web app is not just a thin client. It owns route composition, client-side session identity for anonymous users, React Query cache invalidation, SSE consumption, and schema validation of AI envelopes before UI rendering.

When debugging AI UX, inspect the browser-side session API wrapper first, not just the dashboard component.

### Go backend
The Go service is the state owner and boundary enforcer. It loads config, connects Postgres and Redis, wires job infrastructure, registers the REST API under `/api/v1`, enforces session ownership, persists conversation history, and proxies AI calls to the Python runtime.

When debugging a failed end-to-end feature, `bootstrap.go` usually tells you which domain package and integration seam own the behavior.

### AI runtime
The AI runtime is request-scoped orchestration, not the system of record. It takes conversation history and allowed comparison IDs from callers, fetches catalog/offer context through backend tools, asks the LLM for strict JSON output, hydrates generated selections against backend-owned entities, and streams SSE events for chat consumers.

When debugging hallucination or invalid-structure issues, check prompt construction and hydration in the workflow before changing frontend rendering.

## Important safe-change checkpoints

- The backend currently defaults `AI_RUNTIME_URL` to `http://localhost:8000`, and its decision-memory handlers proxy to `/api/v1/greeting`, `/api/v1/chat`, and `/api/v1/chat/stream`; if you rename runtime routes, backend chat, greeting, and interaction replay must all move together.
- Session chat requests persist the user message in Go before forwarding the complete session history to the AI runtime, and assistant responses are persisted back into session history, so history-related bugs usually span both the proxy handler and the runtime workflow.
- Clarification interactions are not free-form UI callbacks: the backend validates schema version `1.0`, active question identity, component identity, option rules, and an interaction lease before replaying the answer through the AI runtime.
- Resume tokens are single-use in the backend: successful resume marks the token consumed, and the web resume route removes the token from the URL before redirecting to the dashboard session.

## Focused test suggestions before you stop

- **Web-only change:** run the nearest Vitest file in `apps/web/src`, especially around `decision-memory`, `resume`, dashboard components, or checkout incentive UI.
- **Backend handler/change in chat or session flow:** run decision-memory Go tests, especially `internal/decision_memory/handler/chat_test.go` and related interaction tests.
- **AI runtime workflow or prompt change:** run `services/ai-runtime/tests/unit/test_graph.py` and `test_self_correction.py`, then any affected integration tests.
- **Checkout/orders/promotions change:** run relevant Go service tests under `internal/orders` or `internal/promotions`, plus affected frontend checkout incentive tests if the UI changes.
- **Cross-service contract change:** validate one test per layer instead of only one layer.

## Known navigation caveat

`AGENTS.md` now explicitly says to use OpenWiki just-in-time and to treat source/tests as authoritative. Follow that guidance here too: use this page to find the right subsystem quickly, then verify behavior in code and tests before making cross-service assumptions.

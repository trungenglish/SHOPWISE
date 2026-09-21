---
type: quickstart
title: SHOPWISE OpenWiki Quickstart
description: Fast path for coding agents to orient in the Shopwise monorepo, choose the right runtime, and avoid current repo-specific setup and integration traps.
tags: [quickstart, monorepo, routing, setup, operations]
openwiki_generated: true
verified:
  - by: openwiki/0.5.2
    at: 2026-09-21T02:34:46.016Z
sources:
  - id: openwiki-source-e86fe7b76c693666bc2cb828
    resource: repo://apps/mobile/package.json
  - id: openwiki-source-816936db8aa0e22abd7917b3
    resource: repo://apps/web/src/components/chat-workspace.tsx
  - id: openwiki-source-b8ceb4071640ca03055981af
    resource: repo://apps/web/src/routes/chat.tsx
  - id: openwiki-source-5f4beeeff5a41258071dfac9
    resource: repo://apps/web/src/routes/dashboard.tsx
  - id: openwiki-source-5b54a58d1b51cd490b0e7162
    resource: repo://package.json
  - id: openwiki-source-40275cb92c3610938f16ade3
    resource: repo://pnpm-workspace.yaml
  - id: openwiki-source-23775c3de52f3ab95a13cb8b
    resource: repo://README.md
  - id: openwiki-source-e240c53f054a5afd835b3c31
    resource: repo://services/ai-runtime/README.md
  - id: openwiki-source-7a0147618a7c896970573895
    resource: repo://services/main-backend/.env.example
  - id: openwiki-source-4f2c192847e0fbe283a0e627
    resource: repo://services/main-backend/internal/decision_memory/handler/chat.go
  - id: openwiki-source-b76a40124a4e52afe2d5f0f9
    resource: repo://services/main-backend/package.json
  - id: openwiki-source-440ae1e215cb02721dda855c
    resource: repo://turbo.json
generated: { by: "openwiki/0.5.2", at: "2026-09-21T02:34:46.016Z" }
---

# SHOPWISE OpenWiki Quickstart

Use this page to answer three questions quickly:

1. **Which runtime owns my change?**
2. **Which wiki page should I read next?**
3. **Which current repo traps should I verify before coding?**

## Monorepo in one minute

Shopwise is a multi-runtime monorepo with four practical buckets:

- **`apps/web`** — primary product UI: React + TanStack Router, including the landing page, dashboard workspace, chat page, and resume flow.
- **`services/main-backend`** — stateful Go API and system of record: sessions, identity, resume tokens, orders, promotions, and route wiring under `/api/v1`.
- **`services/ai-runtime`** — stateless Python FastAPI service for structured LLM responses and SSE output.
- **`apps/mobile`** — Expo shell that shares workspace conventions, but is not the main path for the current shopping/session flows.

Repository tooling is workspace-driven: `pnpm-workspace.yaml` includes apps, packages, services, conformance, deploy, infra, and scripts; root `package.json` delegates common tasks to Turbo; `turbo.json` makes `dev` persistent and uncached.

## Read next by task

| If you need to... | Read this page next | Why |
|---|---|---|
| Understand the runtime split and ownership boundaries | [Architecture Overview](./architecture/overview.md) | Best high-level map of web, backend, AI runtime, mobile, and shared packages |
| Understand product/session concepts before changing behavior | [Domain Concepts](./domain/domain-concepts.md) | Covers session, identity, branching, offers, orders, and trust concepts |
| Trace an end-to-end request or user flow | [Key Workflows](./workflows/key-workflows.md) | Best for session chat, resume, checkout, and startup flow tracing |
| Find the first source files to open | [Source Map](./source-map.md) | Fastest navigation page for real code entrypoints |
| Run services locally or debug env/setup | [Operations Runbook](./operations/runbook.md) | Covers commands, envs, dependencies, and dev footguns |
| Verify boundary contracts and external dependencies | [Integration Points](./integrations/integration-points.md) | Best for web↔backend, backend↔AI, DB/Redis/jobs, and third-party boundaries |
| Decide what tests to run for a change | [Testing Guide](./testing/testing-guide.md) | Maps the strongest tests and minimum verification paths |

## Runtime routing map

### Web UI change

Start in:

- `/apps/web/src/main.tsx`
- `/apps/web/src/routes/__root.tsx`
- `/apps/web/src/routes/dashboard.tsx`
- `/apps/web/src/routes/resume.tsx`
- `/apps/web/src/api/decision-memory.ts`

Use this route when changing dashboard UX, session/resume screens, client-side SSE handling, or shared browser providers.

### Go backend change

Start in:

- `/services/main-backend/internal/bootstrap/bootstrap.go`
- `/services/main-backend/internal/decision_memory/handler/handler.go`
- `/services/main-backend/internal/decision_memory/handler/chat.go`
- `/services/main-backend/internal/resume_session/usecase/service.go`
- `/services/main-backend/internal/orders/**`

Use this route when changing persisted state, auth/anonymous ownership, session rules, resume-token semantics, order/checkout APIs, or route registration.

### AI runtime change

Start in:

- `/services/ai-runtime/src/main.py`
- `/services/ai-runtime/src/api/chat.py`
- `/services/ai-runtime/src/graph/workflow.py`
- `/services/ai-runtime/src/models/schemas.py`

Use this route when changing prompt orchestration, structured output validation, SSE event shaping, or backend tool usage inside the Python runtime.

### Shared contract/package change

Start in:

- `/packages/api-types/src/index.ts`
- `/packages/protocols/src/index.ts`
- `/packages/ui`
- `/packages/env`

Use this route when a change crosses runtimes and needs shared TypeScript contracts, UI protocol semantics, or environment helpers.

### Mobile change

Start in:

- `/apps/mobile/app/_layout.tsx`
- `/apps/mobile/app/(drawer)/(tabs)/index.tsx`
- `/apps/mobile/package.json`

Treat mobile as a separate client shell unless you confirm a feature is already wired into the main web/backend/AI flows.

## Current repo caveats worth checking first

These are repository-evidenced mismatches, not theoretical risks:

- **The root README is partly scaffold-stale.** It still describes the repo as a Better-T-Stack starter, references `services/agentic` in the project structure, and points multiple setup steps at `apps/server`, while the active backend lives in `services/main-backend` and the active AI service lives in `services/ai-runtime`.
- **Use workspace config, not root `package.json`, to understand repo shape.** Root `package.json` lists only `apps/*` and `packages/*` in its `workspaces`, but `pnpm-workspace.yaml` is the broader authoritative workspace file and includes `services/*` plus other top-level directories.
- **The AI runtime local port is currently aligned at `8000`, not `8001`.** `services/ai-runtime/README.md` starts FastAPI on port `8000`, and `services/main-backend/.env.example` sets `AI_RUNTIME_URL=http://localhost:8000`.
- **The backend already proxies AI calls through `/api/v1/chat` and `/api/v1/chat/stream`.** If you are debugging integration, start from `services/main-backend/internal/decision_memory/handler/chat.go`, not older assumptions about missing prefixes.
- **The standalone `/chat` page is still placeholder-level UI.** `apps/web/src/components/chat-workspace.tsx` currently renders only `ChatWorkspace Placeholder`, so the main product flow is the dashboard, not the chat route.
- **The dashboard is real but still state-heavy on the client.** `apps/web/src/routes/dashboard.tsx` mixes real API calls like `fetchGreeting`, `runAgentTurnStream`, and `sendInteractionStream` with substantial local UI state, so verify whether a behavior is source-of-truth backend logic or route-local presentation state before changing it.

## Minimal reading plans

### I need to change the dashboard or shopping workspace

1. This page
2. [Source Map](./source-map.md)
3. [Architecture Overview](./architecture/overview.md)
4. [Key Workflows](./workflows/key-workflows.md)
5. Then open `/apps/web/src/routes/dashboard.tsx`

### I need to change sessions, memory, or chat persistence

1. This page
2. [Domain Concepts](./domain/domain-concepts.md)
3. [Integration Points](./integrations/integration-points.md)
4. [Testing Guide](./testing/testing-guide.md)
5. Then open `/services/main-backend/internal/decision_memory/handler/chat.go` and `/services/main-backend/internal/decision_memory/usecase/service.go`

### I need to change resume-link behavior

1. This page
2. [Key Workflows](./workflows/key-workflows.md)
3. [Domain Concepts](./domain/domain-concepts.md)
4. [Testing Guide](./testing/testing-guide.md)
5. Then open `/apps/web/src/routes/resume.tsx` and `/services/main-backend/internal/resume_session/usecase/service.go`

### I need to change LLM/runtime behavior

1. This page
2. [Architecture Overview](./architecture/overview.md)
3. [Integration Points](./integrations/integration-points.md)
4. [Testing Guide](./testing/testing-guide.md)
5. Then open `/services/ai-runtime/src/api/chat.py`, `/services/ai-runtime/src/graph/workflow.py`, and `/services/ai-runtime/src/models/schemas.py`

### I need to run or debug the repo locally

1. This page
2. [Operations Runbook](./operations/runbook.md)
3. [Source Map](./source-map.md)
4. Then verify the exact package script you need in root `package.json`, `services/main-backend/package.json`, or `apps/mobile/package.json`

## Quick command orientation

At repo root, the main shared commands are:

- `pnpm run dev`
- `pnpm run build`
- `pnpm run check-types`
- `pnpm run dev:web`
- `pnpm run dev:mobile`
- `pnpm run dev:server`
- `pnpm run dev:server:docker`

Backend-specific commands live in `services/main-backend/package.json`, including `dev`, `dev:app`, `dev:worker`, `dev:deps`, `test`, `test:integration`, and `swagger:gen`.

## Default rule of thumb

If you are unsure where to start:

- read [Source Map](./source-map.md) to find files,
- read [Architecture Overview](./architecture/overview.md) to confirm ownership,
- and prefer the dashboard + Go backend + AI runtime path over README-era scaffold assumptions.

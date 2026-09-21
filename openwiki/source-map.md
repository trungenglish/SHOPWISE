---
type: reference
title: Source Map
description: Practical navigation map to the repo’s highest-signal entrypoints, feature hotspots, shared contracts, and tests. Use this page to choose the first files to open for common product, backend, runtime, and contract changes.
tags: [source-map, navigation, entrypoints, tests, contracts]
openwiki_generated: true
verified:
  - by: openwiki/0.5.2
    at: 2026-09-21T02:34:46.016Z
sources:
  - id: openwiki-source-3499d03c25905d042c8b4ce8
    resource: repo://apps/mobile/app/_layout.tsx
  - id: openwiki-source-0660ff7766bf8418ee1b577a
    resource: repo://apps/mobile/app/(drawer)/(tabs)/index.tsx
  - id: openwiki-source-970773f08ee82250bf68a237
    resource: repo://apps/web/src/api/decision-memory.test.ts
  - id: openwiki-source-313b77d3166c4ef5abe361e6
    resource: repo://apps/web/src/api/decision-memory.ts
  - id: openwiki-source-c8f480dc0e39a3fb58a0a718
    resource: repo://apps/web/src/api/dynamic-ui-protocol.test.ts
  - id: openwiki-source-8e8b395281c4996e784ae3b5
    resource: repo://apps/web/src/main.tsx
  - id: openwiki-source-bc5a450920bf02bea94318ed
    resource: repo://apps/web/src/routes/__root.tsx
  - id: openwiki-source-5f4beeeff5a41258071dfac9
    resource: repo://apps/web/src/routes/dashboard.tsx
  - id: openwiki-source-8c5e10b4fcf9b1ec4222179e
    resource: repo://apps/web/src/routes/resume.test.tsx
  - id: openwiki-source-6487f656f578243180fa4c8a
    resource: repo://apps/web/src/routes/resume.tsx
  - id: openwiki-source-2588c8e46ad945f13a352f60
    resource: repo://packages/api-types/src/identity.ts
  - id: openwiki-source-14dd435365d0b940a0f02278
    resource: repo://packages/api-types/src/index.ts
  - id: openwiki-source-2b81fb17fe237694265d4a08
    resource: repo://packages/api-types/src/promotions.ts
  - id: openwiki-source-00ff56e965b7281717d3c1c2
    resource: repo://packages/protocols/src/index.ts
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
  - id: openwiki-source-4a343e84e31e7cb9e9bb9d9c
    resource: repo://services/main-backend/internal/decision_memory/handler/handler.go
  - id: openwiki-source-4b8a5d433e7f00b922044ec6
    resource: repo://services/main-backend/internal/decision_memory/handler/interaction.go
  - id: openwiki-source-2b135185faa7dcba48f4362c
    resource: repo://services/main-backend/internal/platform/testutil/contract/openapi_test.go
  - id: openwiki-source-831f6e54c15356919121aff1
    resource: repo://services/main-backend/internal/resume_session/usecase/service_test.go
  - id: openwiki-source-c107c02cfc1ca9c0522cab76
    resource: repo://services/main-backend/internal/resume_session/usecase/service.go
  - id: openwiki-source-a66ce5dcdffb897240fb5be0
    resource: repo://specs/005-decision-memory/spec.md
  - id: openwiki-source-76e8441941a0de278eb0fd3c
    resource: repo://specs/007-resume-shopping-session/spec.md
  - id: openwiki-source-070d8c04caccda1f5c5c3ca6
    resource: repo://specs/008-llm-integration/spec.md
generated: { by: "openwiki/0.5.2", at: "2026-09-21T02:34:46.016Z" }
---

# Source Map

This page is intentionally navigation-first. It points to the files and directories that usually matter **first** when you need to change behavior, trace a request path, or verify a contract.

## Top-level orientation

- `/apps/web/src` — primary browser client, including route entrypoints, feature UI, and frontend API helpers.
- `/apps/mobile/app` — Expo Router shell for the native app; currently more scaffold than product hotspot.
- `/services/main-backend` — Go API, persistence, auth, decision-memory state, checkout, promotions, and resume-session flows.
- `/services/ai-runtime/src` — Python FastAPI service that turns session history into structured agent responses.
- `/packages` — shared UI, API types, protocol packages, env/config, and generated artifacts.
- `/specs` — feature specs and contracts worth reading before changing active flows.

## Start here for common change types

| Change you need | Open first | Then follow into |
|---|---|---|
| Web app bootstrap, providers, router behavior | `/apps/web/src/main.tsx` | `/apps/web/src/routes/__root.tsx` |
| Dashboard / conversational shopping UX | `/apps/web/src/routes/dashboard.tsx` | `features/dashboard`, `api/decision-memory.ts`, backend session/chat handlers |
| Resume-link UX and token error handling | `/apps/web/src/routes/resume.tsx` | `components/expired-token-ui.tsx`, backend resume-session service/tests |
| Session CRUD, branching, restore, autosave | `/apps/web/src/api/decision-memory.ts` | `/services/main-backend/internal/decision_memory/handler/handler.go`, usecase/repository |
| AI chat request/streaming path | `/services/main-backend/internal/decision_memory/handler/chat.go` | `/services/ai-runtime/src/api/chat.py`, `/services/ai-runtime/src/graph/workflow.py` |
| Dynamic UI interactions | `/services/main-backend/internal/decision_memory/handler/interaction.go` | `@shopwise/protocols`, web protocol tests |
| Resume-token security / inactivity notification | `/services/main-backend/internal/resume_session/usecase/service.go` | handler, provider, tests, spec 007 |
| Checkout / order / retailer offers | `/services/main-backend/internal/bootstrap/bootstrap.go` | `/internal/orders/**`, `/internal/accessories/**`, `/internal/promotions/**` |
| Shared API payload types for TS clients | `/packages/api-types/src/index.ts` | `decision-memory.ts`, `identity.ts`, `promotions.ts` |
| Shared dynamic UI contract | `/packages/protocols/src/index.ts` | `/packages/ui-protocol/src/schema`, frontend tests, spec 004 |
| AI runtime request handling and schema validation | `/services/ai-runtime/src/api/chat.py` | `/src/models/schemas.py`, `/src/graph/workflow.py`, tests |

## Web app (`/apps/web/src`)

### Entrypoints

- `/apps/web/src/main.tsx` — creates the TanStack router and React Query client, then mounts the app under `QueryClientProvider`, dark `ThemeProvider`, and shared `FontProvider`.
- `/apps/web/src/routes/__root.tsx` — global route shell: tooltip provider, toaster, checkout-incentive provider/panel, and devtools gated by `VITE_NODE_ENV`.
- `/apps/web/src/routes/dashboard.tsx` — current conversational workspace hotspot. It fetches the greeting, starts/continues agent turns, handles interaction streams, and coordinates the comparison / checkout / reasoning / accessory modals.
- `/apps/web/src/routes/resume.tsx` — token-driven session resumption page. It consumes the `token` query param, calls the resume API via React Query, clears the token from the URL, invalidates relevant session queries, and redirects to `/dashboard` with `sessionId`.

### Feature hotspots

Prefer these folders over scanning the entire route tree:

- `/apps/web/src/features/dashboard` — shopping workspace UI and agent result rendering.
- `/apps/web/src/features/orders` — order history surfaces used from the dashboard route.
- `/apps/web/src/features/checkout-incentive` — provider/panel for saved-product or checkout-triggered promotions.
- `/apps/web/src/features/errors` — app-level error/not-found components.
- `/apps/web/src/components` — reusable shell components like loaders and token-expiry UI.

### Frontend API files worth opening first

- `/apps/web/src/api/decision-memory.ts`
  - Adds `Authorization` when present.
  - Ensures every request carries `X-Anonymous-ID`, generating and persisting one in `localStorage` if missing.
  - Owns session CRUD, autosave, branch/restore, resume-token consumption, greeting fetch, chat, streaming chat, and agent-envelope validation.
- `/apps/web/src/api/checkout.ts` — checkout API surface.
- `/apps/web/src/api/accessories.ts` — accessory recommendation fetches used by dashboard flows.

### Important web tests

Open these before changing user-visible behavior:

- `/apps/web/src/api/decision-memory.test.ts` — verifies first-turn session creation, anonymous identity headers, SSE parsing, and agent-envelope validation.
- `/apps/web/src/api/dynamic-ui-protocol.test.ts` — verifies deterministic UI operation ordering and nested component replacement.
- `/apps/web/src/routes/resume.test.tsx` — covers successful resume redirect and invalid/expired-token UI behavior.
- `/apps/web/src/components/expired-token-ui.test.tsx` — focused token recovery UI coverage.

## Mobile app (`/apps/mobile`)

This area is currently a shell, not a major product hotspot.

### Open first

- `/apps/mobile/app/_layout.tsx` — Expo Router root stack plus native providers (`GestureHandlerRootView`, `KeyboardProvider`, theme provider, `HeroUINativeProvider`).
- `/apps/mobile/app/(drawer)/(tabs)/index.tsx` — current placeholder home tab.
- `/apps/mobile/app/(drawer)/_layout.tsx` and `/apps/mobile/app/(drawer)/(tabs)/_layout.tsx` — drawer/tab composition when you need navigation wiring.

Use the mobile app mainly when you are adding native navigation structure or trying to understand parity gaps versus the web client.

## Main backend (`/services/main-backend`)

### Process entrypoints

- `/services/main-backend/cmd/retail/main.go` — CLI entrypoint. Supports three modes: normal server run, `-migrate`, and `-seed`.
- `/services/main-backend/internal/bootstrap/bootstrap.go` — highest-signal backend file. It loads config, connects Postgres/Redis/job client, builds services/handlers, and registers almost every route.
- `/services/main-backend/.env.example` — practical config shape.
- `/services/main-backend/package.json` — developer commands.

### Route and wiring hotspot

`/services/main-backend/internal/bootstrap/bootstrap.go` is the best first read when you need to know:

- which modules are live,
- which dependencies a service is constructed with,
- which route group owns an endpoint,
- whether a feature is debug-only,
- and where AI runtime, promotions, accessories, or resume-session services attach.

### Domain hotspots by responsibility

- `/services/main-backend/internal/decision_memory`
  - session CRUD,
  - message persistence,
  - optional auth + anonymous ownership,
  - AI-runtime proxying,
  - dynamic interaction deduplication/leases.
- `/services/main-backend/internal/resume_session`
  - single-use resume tokens,
  - token revocation/consumption,
  - notification provider integration,
  - inactivity-resume flow.
- `/services/main-backend/internal/orders`
  - checkout APIs, order listing, order confirmation, retailer offer refresh integration.
- `/services/main-backend/internal/accessories`
  - accessory catalog and recommendation support.
- `/services/main-backend/internal/promotions`
  - promotion issuance and voucher generation plumbing.
- `/services/main-backend/internal/identity` and `/internal/users`
  - auth, JWT, Google OAuth, verified user identity, and user lifecycle.
- `/services/main-backend/internal/platform`
  - cross-cutting infrastructure: config, DB, cache, jobs, middleware, logging, health, router, and contract test helpers.

### Best first backend files by task

- **Understand all route registration**: `/services/main-backend/internal/bootstrap/bootstrap.go`
- **Change session REST behavior**: `/services/main-backend/internal/decision_memory/handler/handler.go`
- **Change AI request forwarding / SSE proxying**: `/services/main-backend/internal/decision_memory/handler/chat.go`
- **Change interaction submission semantics**: `/services/main-backend/internal/decision_memory/handler/interaction.go`
- **Change decision-memory business rules**: `/services/main-backend/internal/decision_memory/usecase/service.go`
- **Change resume token lifecycle**: `/services/main-backend/internal/resume_session/usecase/service.go`
- **Change checkout/order flows**: `/services/main-backend/internal/orders/handler/handler.go` and `/services/main-backend/internal/orders/usecase/service.go`
- **Change config or environment validation**: `/services/main-backend/internal/platform/config`

### Control-flow shortcuts that matter

- **Web session/chat path**: web `api/decision-memory.ts` → backend `decision_memory/handler/chat.go` → AI runtime `/api/v1/chat` or `/api/v1/chat/stream`.
- **Resume flow**: web `/routes/resume.tsx` → frontend resume helper → backend `resume_session` handler/service → session fetch → redirect back to dashboard with `sessionId`.
- **Interaction flow**: dashboard interaction submit → backend `interaction.go` validates question linkage and deduplicates by `interaction_id` → AI runtime → stored envelope replay or completion.

### High-signal backend tests

- `/services/main-backend/internal/decision_memory/handler/chat_test.go` — verifies history forwarding, ownership checks, extracted comparison IDs, and SSE persistence behavior.
- `/services/main-backend/internal/decision_memory/handler/interaction_test.go` — interaction idempotency/validation coverage.
- `/services/main-backend/internal/resume_session/usecase/service_test.go` — token consumed/revoked/expired handling and cross-device context mismatch behavior.
- `/services/main-backend/internal/orders/usecase/service_test.go` and handler tests — checkout behavior.
- `/services/main-backend/internal/platform/testutil/contract/openapi_test.go` — OpenAPI/spec coverage sanity checks.

## AI runtime (`/services/ai-runtime/src`)

### Entrypoints

- `/services/ai-runtime/src/main.py` — FastAPI app, request timing middleware, global exception mapping, chat router mount, and health endpoint.
- `/services/ai-runtime/src/api/chat.py` — the runtime’s real entrypoint for product behavior. Owns request models, greeting generation, direct chat invocation, and SSE streaming.
- `/services/ai-runtime/src/graph/workflow.py` — LangGraph workflow that loads backend data, invokes the model, validates structured JSON, and retries self-correction when needed.
- `/services/ai-runtime/src/models/schemas.py` — structured response schema and hydration logic for agent envelopes.

### What to open for common runtime questions

- **Why did the agent ask this question?** → `core/prompts.py` and `graph/workflow.py`
- **Why did the response fail validation?** → `models/schemas.py` and `tests/unit/test_self_correction.py`
- **How are offers/catalog injected into prompts?** → `graph/workflow.py` and `tools/interfaces.py`
- **How does streaming behave?** → `api/chat.py` and `api/streaming.py`

### Runtime test map

- `/services/ai-runtime/tests/unit/test_graph.py` — verifies prompt/history handling, catalog hydration, and offer-comparison hydration.
- `/services/ai-runtime/tests/unit/test_self_correction.py` — verifies invalid JSON triggers bounded self-correction retries.
- `/services/ai-runtime/tests/unit/test_streaming.py` — SSE event behavior.
- `/services/ai-runtime/tests/unit/test_tool_proxy.py` — backend tool proxy expectations.
- `/services/ai-runtime/tests/integration/test_chat_endpoint.py` — endpoint-level chat coverage.

## Shared packages (`/packages`)

Only a few packages are usually first-read material.

### Contracts and shared types

- `/packages/api-types`
  - `/src/index.ts` re-exports the shared TypeScript API surface.
  - `/src/decision-memory.ts` mirrors decision-memory session and preference payloads.
  - `/src/identity.ts` defines auth request/response schemas with Zod.
  - `/src/promotions.ts` defines promotion trigger/status payloads.
- `/packages/protocols/src/index.ts`
  - Main shared dynamic UI contract package.
  - Defines dynamic component/action types, `InteractionRequest`, and `applyUIOperations()` ordering/replacement semantics.
- `/packages/ui-protocol`
  - Schema-first protocol assets under `/src/schema`.
  - Open when you need canonical JSON schema artifacts rather than only TS helpers.

### UI and platform support

- `/packages/ui` — shared React UI primitives, hooks, and providers; used by web and likely future client surfaces.
- `/packages/env` — environment accessors for web/mobile.
- `/packages/config` — shared config/tooling helpers for TypeScript packages.
- `/packages/generated` — generated assets; read only when tracing codegen outputs.

## Specs worth reading before changing live flows

Specs are especially useful here because several current implementations are partial, evolving, or stricter in tests/contracts than in broad UI prose.

### Highest-signal specs

- `/specs/005-decision-memory/spec.md`
  - Expectations around autosave, branching, reasoning replay, and cross-device session restoration.
- `/specs/007-resume-shopping-session/spec.md`
  - Security and UX expectations for resume links, inactivity detection, provider abstraction, and invalid-token recovery.
- `/specs/008-llm-integration/spec.md`
  - AI runtime expectations for SSE streaming, provider abstraction, structured output, retries, and observability.
- `/specs/004-dynamic-ui-protocol/`
  - Read this when changing interaction payloads, component types, or schema expectations.
- `/specs/007-checkout-readiness/spec.md`
  - Checkout-readiness behavior and acceptance criteria.

### Contract artifacts inside specs

For implementation-detail questions, these are often faster than scanning all code first:

- `/specs/005-decision-memory/contracts/api-contracts.md`
- `/specs/005-decision-memory/contracts/ui-contracts.md`
- `/specs/007-resume-shopping-session/contracts/api-contracts.md`
- `/specs/008-llm-integration/contracts/api-contracts.md`

## Practical “open these first” bundles

### If you need to change conversational shopping behavior

1. `/apps/web/src/routes/dashboard.tsx`
2. `/apps/web/src/api/decision-memory.ts`
3. `/services/main-backend/internal/decision_memory/handler/chat.go`
4. `/services/ai-runtime/src/api/chat.py`
5. `/services/ai-runtime/src/graph/workflow.py`
6. `/services/ai-runtime/src/models/schemas.py`

### If you need to change resume-link behavior

1. `/apps/web/src/routes/resume.tsx`
2. `/apps/web/src/routes/resume.test.tsx`
3. `/services/main-backend/internal/resume_session/usecase/service.go`
4. `/services/main-backend/internal/resume_session/usecase/service_test.go`
5. `/specs/007-resume-shopping-session/spec.md`

### If you need to change session ownership or autosave behavior

1. `/apps/web/src/api/decision-memory.ts`
2. `/services/main-backend/internal/decision_memory/handler/handler.go`
3. `/services/main-backend/internal/decision_memory/usecase/service.go`
4. `/services/main-backend/internal/decision_memory/repository/postgres`
5. `/specs/005-decision-memory/spec.md`

### If you need to change dynamic UI interaction semantics

1. `/packages/protocols/src/index.ts`
2. `/services/main-backend/internal/decision_memory/handler/interaction.go`
3. `/apps/web/src/api/dynamic-ui-protocol.test.ts`
4. `/specs/004-dynamic-ui-protocol/`

### If you need to change promotions or checkout incentives

1. `/apps/web/src/features/checkout-incentive`
2. `/services/main-backend/internal/promotions`
3. `/packages/api-types/src/promotions.ts`
4. `/services/main-backend/internal/orders`

## Testing locations by stack

- **Web**: `/apps/web/src/**/*.test.ts` and `/apps/web/src/**/*.test.tsx`
- **Backend**: `/services/main-backend/internal/**/*_test.go`
- **AI runtime**: `/services/ai-runtime/tests/unit` and `/services/ai-runtime/tests/integration`
- **Spec/contract references**: `/specs/**/contracts`, plus backend OpenAPI checks under `/services/main-backend/internal/platform/testutil/contract`

## Quick cautions while navigating

- Do not assume the mobile app mirrors web behavior yet; check `/apps/mobile/app` before planning shared-flow changes.
- For decision-memory changes, inspect both **web helper validation** and **backend handler semantics**; many failures show up first in frontend envelope parsing.
- For AI changes, treat `workflow.py`, `schemas.py`, and runtime tests as one unit; changing only one usually breaks structured output expectations.
- For resume flows, tests in both web and backend are faster to trust than high-level assumptions; token lifecycle rules are stricter than the UI alone suggests.

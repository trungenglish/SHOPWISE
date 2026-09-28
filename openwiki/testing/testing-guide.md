---
type: guide
title: Testing Guide
description: Practical guide to the repository's current automated test coverage across the web app, Go backend, and AI runtime, including what each layer actually verifies and what to run for common changes.
tags: [testing, quality, frontend, backend, ai-runtime]
openwiki_generated: true
verified:
  - by: openwiki/0.6.0
    at: 2026-09-28T02:48:19.098Z
sources:
  - id: openwiki-source-c8f480dc0e39a3fb58a0a718
    resource: repo://apps/web/src/api/dynamic-ui-protocol.test.ts
  - id: openwiki-source-8c5e10b4fcf9b1ec4222179e
    resource: repo://apps/web/src/routes/resume.test.tsx
  - id: openwiki-source-6487f656f578243180fa4c8a
    resource: repo://apps/web/src/routes/resume.tsx
  - id: openwiki-source-6fe10cfb5ec054ca681eb262
    resource: repo://apps/web/vitest.config.ts
  - id: openwiki-source-35c57d7909eeadccbc068ddf
    resource: repo://conformance/retail-sdk/package.json
  - id: openwiki-source-261ca52f815b5dcb0e0dd689
    resource: repo://conformance/tool-sdk/package.json
  - id: openwiki-source-035c2e2ce3a9ed0b6155e32f
    resource: repo://conformance/ui-protocol/package.json
  - id: openwiki-source-18dad4774ed8bc0a70d31f57
    resource: repo://services/ai-runtime/pytest.ini
  - id: openwiki-source-85cedde5d4fbfc25098c14b5
    resource: repo://services/ai-runtime/src/api/chat.py
  - id: openwiki-source-95c5cf8058175b6f67241cb5
    resource: repo://services/ai-runtime/src/graph/workflow.py
  - id: openwiki-source-18e682d97b2e6bb5e338df78
    resource: repo://services/ai-runtime/tests/integration/test_chat_endpoint.py
  - id: openwiki-source-6dca1297d7084a2d4bb69450
    resource: repo://services/ai-runtime/tests/unit/test_graph.py
  - id: openwiki-source-9bbe67c04b1a109641046532
    resource: repo://services/main-backend/internal/decision_memory/handler/interaction_test.go
  - id: openwiki-source-4b8a5d433e7f00b922044ec6
    resource: repo://services/main-backend/internal/decision_memory/handler/interaction.go
  - id: openwiki-source-aab2acaa3eabb8d48be370be
    resource: repo://services/main-backend/internal/orders/usecase/service_test.go
  - id: openwiki-source-d91401aebb56a3c52d84389b
    resource: repo://services/main-backend/internal/resume_session/handler/handler.go
  - id: openwiki-source-831f6e54c15356919121aff1
    resource: repo://services/main-backend/internal/resume_session/usecase/service_test.go
  - id: openwiki-source-c107c02cfc1ca9c0522cab76
    resource: repo://services/main-backend/internal/resume_session/usecase/service.go
  - id: openwiki-source-b76a40124a4e52afe2d5f0f9
    resource: repo://services/main-backend/package.json
  - id: openwiki-source-b8ba99521c881c5787fce983
    resource: repo://services/main-backend/tests/e2e_test.go
generated: { by: "openwiki/0.6.0", at: "2026-09-28T02:48:19.098Z" }
---

# Testing Guide

## How testing is organized

The repository currently has three meaningful automated test runtimes, each validating a different slice of behavior:

- **TypeScript/Vitest in `apps/web`** for UI route behavior and shared protocol reducers.
- **Go `test` packages in `services/main-backend`** for backend use cases and request validation rules.
- **Python/Pytest in `services/ai-runtime`** for workflow orchestration and FastAPI endpoints.

This is not a top-heavy E2E suite. Most confidence comes from focused unit or service-level tests around high-risk boundaries: UI protocol application, resume-link recovery, checkout pricing and idempotency, decision-memory interactions, and AI response shaping.

## Runtime-specific strategy

### Web app: Vitest + jsdom

The web app uses Vitest with the `jsdom` environment and a shared setup file, so current tests behave like browser-side component and reducer tests rather than full browser automation. The config is intentionally lightweight: React plugin, tsconfig path resolution, globals enabled, and `./src/test/setup.ts` as the common setup entrypoint. (`apps/web/vitest.config.ts`)

Current visible tests emphasize two areas:

- **Dynamic UI protocol application** in `src/api/dynamic-ui-protocol.test.ts`
- **Resume route behavior** in `src/routes/resume.test.tsx`

#### Shared protocol invariants covered

`dynamic-ui-protocol.test.ts` verifies two important properties of `applyUIOperations` from `@shopwise/protocols`:

- operations are applied **deterministically by sequence**, even when provided out of order
- applying operations does **not mutate the previous UI state**
- a later `replace` can correctly target a nested component created by an earlier `append`

Those tests are high-signal for any change to chat rendering, streamed UI updates, or shared schema/protocol logic because they protect the client-side reducer semantics that consume AI-driven UI operations. (`apps/web/src/api/dynamic-ui-protocol.test.ts`)

#### Resume route invariants covered

`resume.test.tsx` exercises the route-level control flow of `ResumePage` with mocked React Query and router hooks. It verifies that:

- a successful resume removes the token from the URL, invalidates the session list and resumed-session queries, then navigates to `/dashboard` with the resumed `sessionId`
- an expired token renders the expired-link UI instead of navigating
- a generic network failure renders a connection-error state

This is meaningful because the production route does the same three-step sequence in a `useEffect`: sanitize the URL, invalidate relevant caches, and redirect to the dashboard. The route also maps token-specific backend errors to `ExpiredTokenUI`, exposes manual retry only for network-style failures, and sends users back to `/` to request a new link for expired, consumed, or revoked tokens. (`apps/web/src/routes/resume.test.tsx`, `apps/web/src/routes/resume.tsx`)

### Go backend: package tests around handlers and use cases

The backend package exposes these main verification commands:

- `test` and `test:unit`: `go test ./...`
- `test:integration`: a narrower tagged integration set covering selected postgres-backed repositories and test utilities
- `test:race`: `go test -race ./...`
- `check-types`: `go vet ./...`

These commands live in `services/main-backend/package.json`. In practice, most of the strongest current coverage is not end-to-end; it sits in package-local tests for domain services and handlers. (`services/main-backend/package.json`)

#### Decision memory interaction validation

`internal/decision_memory/handler/interaction_test.go` covers the backend's guardrails for interactive clarification responses before they are sent back to the AI runtime.

The tested behavior includes:

- mapping a selected option id back into the human-readable answer text sent to the model
- rejecting option ids that were not present in the active stored question
- refusing to stream an interaction for a session owned by another anonymous user, returning `404` without calling the AI runtime

The implementation under test in `interaction.go` adds more context for why this area matters. `InteractionStream`:

- bounds request bodies to 16 KiB
- only accepts schema version `1.0`, event `submit`, and action `question.answer`
- deduplicates or rejects conflicting repeated interactions using a persisted interaction record and payload hash
- validates that the interaction matches the active assistant question and one of its rendered component ids
- turns the structured answer back into plain text and forwards it to the AI runtime `/api/v1/chat`
- persists the returned assistant envelope and replays it as named SSE events: `token`, `ui_operation`, `envelope`, and `done`

This is a critical seam when changing chat/protocol behavior, because the backend is enforcing server-side trust boundaries around the same UI protocol the frontend reducer consumes. (`services/main-backend/internal/decision_memory/handler/interaction_test.go`, `services/main-backend/internal/decision_memory/handler/interaction.go`)

#### Checkout, promotions, and retailer-offer safety

`internal/orders/usecase/service_test.go` is one of the richest backend test files. It documents the invariants the checkout service is expected to preserve.

Representative behaviors covered include:

- **consumer prices are treated as VAT-inclusive**, so tax stays `0` and totals equal the summed item prices
- mixed orders can include official products plus retailer accessory offers while preserving idempotency metadata
- authoritative zero-priced ShopWise bundle lines are accepted without refreshing an external retailer
- if a retailer offer's refreshed snapshot changes, checkout fails with an `OfferChangedError` and persists the changed snapshot for an explicit retry
- an existing idempotent order is returned before retailer refresh work is retried
- product price snapshots and customer contact snapshots are copied into the persisted order
- active promotions apply normalized coupon codes and calculate percentage or fixed discounts against the checkout snapshot

These tests are the best local safety net after changing checkout totals, promotion math, bundle behavior, or retailer-offer refresh semantics. (`services/main-backend/internal/orders/usecase/service_test.go`)

#### Resume-session token lifecycle

`internal/resume_session/usecase/service_test.go` covers the token lifecycle more deeply than the frontend route tests. It verifies that:

- generating a new token allows a valid resume
- once consumed, the same token is rejected
- generating a newer token revokes earlier unconsumed tokens for the same session
- invalid JWTs are rejected
- a changed context hash is treated as an audit signal, not a hard block
- expired tokens are rejected even if their signature and stored record still match
- inactivity notifications require an authenticated session with a verified notification identity
- successful inactivity notifications write a notification log entry

The underlying service confirms the mechanics: token strings are JWT-backed, their SHA-256 hashes are stored, unconsumed prior tokens are revoked on issuance, successful resume marks the token consumed, and mismatched issuance/consumption context hashes only trigger audit logging. The HTTP handler turns those service errors into distinct client-facing codes such as `token_missing`, `token_expired`, `token_consumed`, `token_revoked`, and `token_invalid`. (`services/main-backend/internal/resume_session/usecase/service_test.go`, `services/main-backend/internal/resume_session/usecase/service.go`, `services/main-backend/internal/resume_session/handler/handler.go`)

#### Low-signal backend E2E coverage today

The current `services/main-backend/tests/e2e_test.go` file is only a placeholder scaffold. It defines named tests for recommendation, comparison, clarification, timeout recovery, out-of-stock recovery, and checkout readiness, but each body is empty except for a comment. Treat it as intent, not executable coverage. (`services/main-backend/tests/e2e_test.go`)

### AI runtime: Pytest for orchestration and API contracts

The AI runtime uses pytest with the repository root on `pythonpath` and function-scoped asyncio loops. (`services/ai-runtime/pytest.ini`)

Its tests are split in a useful way:

- `tests/unit/test_graph.py` focuses on workflow orchestration and model/tool contract handling
- `tests/integration/test_chat_endpoint.py` focuses on FastAPI endpoint behavior and response envelopes

#### Workflow invariants covered

`test_graph.py` validates the behavior of `create_workflow` in `src/graph/workflow.py`. The workflow itself loads catalog data and allowed offer comparisons from the backend tool proxy, prepends a generated system prompt, calls the LLM with a strict JSON-schema response format, hydrates product or offer references into a richer `AgentResponse`, and retries invalid model outputs up to `MAX_SELF_CORRECTION_RETRIES`. (`services/ai-runtime/src/graph/workflow.py`)

The tests cover representative invariants:

- the configured model name is passed through to the provider
- the full conversation history is preserved when invoking the LLM
- recommendation drafts are hydrated with catalog product data
- offer-comparison drafts are hydrated from backend `offer_comparison` tool responses and included in prompt context
- the system prompt instructs gaming flows to collect workload before budget
- the prompt tells the model to ask for exactly one missing dimension at a time
- replies like "flexible" or "no preference" count as satisfying a decision dimension
- when the first user message already supplies all dimensions, the workflow can proceed directly to a recommendation

These tests are unusually valuable because they protect not just syntax, but also prompt-level behavioral rules that the product depends on. (`services/ai-runtime/tests/unit/test_graph.py`)

#### Endpoint contract invariants covered

`test_chat_endpoint.py` verifies the FastAPI layer around that workflow.

Covered behaviors include:

- `/api/v1/chat` accepts multi-turn history and returns an agent envelope with `schema_version == "1.1"`
- question responses preserve structured question modes and options
- `/api/v1/chat/stream` emits named SSE events `token`, `ui_operation`, `envelope`, and `done`
- `/api/v1/greeting` falls back to a localized canned greeting when the provider fails instead of failing the dashboard request

Those assertions line up with the implementation in `src/api/chat.py`, where the synchronous chat endpoint invokes the workflow, the streaming endpoint converts the result into SSE, and the greeting endpoint catches all provider exceptions and returns locale-aware fallback text. (`services/ai-runtime/tests/integration/test_chat_endpoint.py`, `services/ai-runtime/src/api/chat.py`)

## Where coverage is strong vs. thin

### Stronger, change-sensitive areas

The best current automated coverage is around:

- **chat/protocol boundaries**: frontend reducer semantics, backend interaction validation, and AI runtime envelope/SSE contracts
- **resume session recovery**: frontend redirect behavior plus backend token lifecycle rules
- **checkout and promotions**: price snapshots, VAT assumptions, idempotency, coupon math, and retailer-offer refresh handling
- **AI workflow prompting and hydration**: history handling, tool hydration, and prompt invariants for recommendation flow

### Thin or placeholder areas

A few areas are visibly low-signal today:

- the backend `tests/e2e_test.go` file is placeholder-only
- the `conformance/retail-sdk`, `conformance/tool-sdk`, and `conformance/ui-protocol` packages currently contain only minimal private `package.json` files plus `.gitkeep`, so they should not be treated as active conformance suites
- frontend route/component coverage is still narrow relative to the size of the dashboard and surrounding UI

That means broad integration regressions are still more likely to be caught by targeted manual flows, type/build checks, or failures in neighboring service-level tests than by a single authoritative end-to-end suite. (`services/main-backend/tests/e2e_test.go`, `conformance/retail-sdk/package.json`, `conformance/tool-sdk/package.json`, `conformance/ui-protocol/package.json`)

## What to run for common changes

Choose the narrowest test set that exercises the boundary you changed, then widen only if the change crosses service or schema lines.

### After changing chat/protocol behavior

Run at least:

- `pnpm --filter web test` for `dynamic-ui-protocol.test.ts` and any route consumers
- from `services/ai-runtime`: `uv run pytest tests/unit/test_graph.py tests/integration/test_chat_endpoint.py`
- `pnpm --filter @shopwise/main-backend test` if the change affects backend interaction validation or SSE replay

Focus on these failure modes:

- operation ordering or immutability regressions in UI application
- envelope shape drift between AI runtime, backend SSE replay, and frontend consumption
- invalid option ids, component ids, or active-question matching in interaction submissions

### After changing checkout or promotions

Run at least:

- `pnpm --filter @shopwise/main-backend test`
- optionally `pnpm --filter @shopwise/main-backend test:race` if concurrency or idempotency behavior changed

Prioritize assertions around:

- total/subtotal/discount/tax calculations
- VAT-inclusive assumptions
- coupon normalization and maximum-discount behavior
- idempotent order reuse
- retailer-offer refresh and changed-offer retry behavior

### After changing resume session behavior

Run at least:

- `pnpm --filter web test`
- `pnpm --filter @shopwise/main-backend test`

Then manually verify one happy path and these failure states:

- missing token
- expired token
- consumed token
- revoked token
- transient network failure with retry

This matters because the frontend route and backend token service each enforce different parts of the resume flow contract.

### After changing shared schema or AI workflow logic

Run at least:

- from `services/ai-runtime`: `uv run pytest`
- `pnpm --filter web test`
- `pnpm --filter @shopwise/main-backend test` if the schema crosses the AI-runtime/backend boundary

Use this set for changes to:

- `AgentResponse`/question envelope shapes
- allowed comparison hydration
- system-prompt rules that influence recommendation flow
- SSE event naming or ordering

## Good default validation routine

For a non-trivial change with cross-service impact:

1. Run the smallest focused automated tests in the affected package.
2. Run package-level static or type checks (`go vet`, frontend type/build checks, or full pytest as appropriate).
3. Run neighboring boundary tests if the contract crosses runtimes.
4. Do one manual product flow through the UI when the change affects chat, resume, or checkout behavior.

In this repository, the safest habit is to trust **focused contract tests first**, then add **manual verification for end-to-end behavior** where the automated suite is still sparse.

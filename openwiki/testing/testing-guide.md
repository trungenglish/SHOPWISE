---
type: testing guide
title: Testing Guide
description: Repository-specific map of the web, Go backend, and AI runtime test layers, including the contracts they enforce, the strongest covered boundaries, and the minimum verification routine for risky changes.
tags: [testing, frontend, backend, ai-runtime, contracts]
openwiki_generated: true
verified:
  - by: openwiki/0.5.2
    at: 2026-09-21T02:34:46.016Z
sources:
  - id: openwiki-source-99de51df25f29bfc72caf823
    resource: repo://apps/web/package.json
  - id: openwiki-source-970773f08ee82250bf68a237
    resource: repo://apps/web/src/api/decision-memory.test.ts
  - id: openwiki-source-957fced29abbe93c444a6587
    resource: repo://apps/web/src/features/dashboard/agent-mapping.test.ts
  - id: openwiki-source-2856805214cd53457724cc33
    resource: repo://apps/web/src/features/dashboard/components/accessories-modal.test.tsx
  - id: openwiki-source-6cafd903240472bb28ca8ce8
    resource: repo://apps/web/src/features/dashboard/components/audit-trail.test.tsx
  - id: openwiki-source-f824dff2e7eabda6696fcfbf
    resource: repo://apps/web/src/features/dashboard/components/comparison-modal.test.tsx
  - id: openwiki-source-8c5e10b4fcf9b1ec4222179e
    resource: repo://apps/web/src/routes/resume.test.tsx
  - id: openwiki-source-85cedde5d4fbfc25098c14b5
    resource: repo://services/ai-runtime/src/api/chat.py
  - id: openwiki-source-95c5cf8058175b6f67241cb5
    resource: repo://services/ai-runtime/src/graph/workflow.py
  - id: openwiki-source-61d0d7812a01667276b208ef
    resource: repo://services/ai-runtime/src/models/schemas.py
  - id: openwiki-source-18e682d97b2e6bb5e338df78
    resource: repo://services/ai-runtime/tests/integration/test_chat_endpoint.py
  - id: openwiki-source-6dca1297d7084a2d4bb69450
    resource: repo://services/ai-runtime/tests/unit/test_graph.py
  - id: openwiki-source-8c46307bb52f1d31e366fb3c
    resource: repo://services/main-backend/internal/decision_memory/handler/chat_test.go
  - id: openwiki-source-2b135185faa7dcba48f4362c
    resource: repo://services/main-backend/internal/platform/testutil/contract/openapi_test.go
  - id: openwiki-source-f48110dd676e56211b3d548b
    resource: repo://services/main-backend/internal/platform/testutil/contract/openapi.go
  - id: openwiki-source-831f6e54c15356919121aff1
    resource: repo://services/main-backend/internal/resume_session/usecase/service_test.go
  - id: openwiki-source-b76a40124a4e52afe2d5f0f9
    resource: repo://services/main-backend/package.json
generated: { by: "openwiki/0.5.2", at: "2026-09-21T02:34:46.016Z" }
---

# Testing Guide

This repository has three distinct testing centers:

- the React web app, with Vitest and component/API-contract style tests
- the Go main backend, with package-level unit tests plus explicit contract checks against OpenAPI and generated Swagger
- the Python AI runtime, with unit tests around workflow/prompt invariants and integration tests around HTTP/SSE behavior

The strongest current coverage is around **decision-memory request/response boundaries**, **resume token lifecycle rules**, and **AI runtime response-shape invariants**. The biggest under-tested hotspot is still **dashboard orchestration**: there are useful component tests, but little evidence of one broad automated test that proves the whole dashboard flow stays correct when session state, streaming, agent envelopes, and UI interactions evolve together.

## Testing layers by subsystem

### Web app

The web package exposes a small but clear verification surface:

- `pnpm --filter web test` runs `vitest run`
- `pnpm --filter web check-types` runs `vite build && tsc --noEmit`

That means frontend verification is split between fast unit/component tests and a stricter build-plus-TypeScript pass that can catch route wiring, import drift, and type mismatches that narrow unit tests miss.

Current frontend tests cluster around three boundary types:

1. **API client contracts** in `src/api/*.test.ts`
2. **resume flow behavior** in `src/routes/resume.test.tsx`
3. **dashboard leaf components and mapping logic** under `src/features/dashboard/**`

This is useful coverage, but it is not the same as end-to-end dashboard orchestration coverage.

### Go main backend

The backend package scripts expose several useful test layers:

- `pnpm --filter @shopwise/main-backend test` runs `go test ./...`
- `pnpm --filter @shopwise/main-backend test:integration` runs a tagged subset focused on integration packages
- `pnpm --filter @shopwise/main-backend test:race` runs `go test -race ./...`
- `pnpm --filter @shopwise/main-backend check-types` runs `go vet ./...`

The most informative tests in the supplied areas are not generic unit tests; they define boundaries:

- `internal/decision_memory/handler/chat_test.go` fixes how session chat calls are forwarded to the AI runtime and what authorization rules apply for anonymous sessions.
- `internal/resume_session/usecase/service_test.go` fixes the resume-token lifecycle and failure cases.
- `internal/platform/testutil/contract/openapi_test.go` and helpers assert that the authored OpenAPI contract exists and that generated Swagger remains a superset of it.

### AI runtime

The AI runtime uses both unit and integration tests, with a strong emphasis on schema-safe behavior rather than subjective recommendation quality.

The most important current layers are:

- `tests/unit/test_graph.py` for workflow state, prompt invariants, catalog hydration, and session-only action boundaries
- `tests/integration/test_chat_endpoint.py` for HTTP request/response behavior, SSE event names, and graceful greeting fallback behavior

These tests establish that the runtime should return structurally valid agent envelopes and stream predictable event types, but they do **not** prove that recommendations are commercially or product-wise “best”. They are boundary tests, not product-judgment tests.

## Strongly covered boundaries and what they guarantee

### Decision-memory web API client: request shape and envelope validation

`apps/web/src/api/decision-memory.test.ts` gives the clearest frontend contract coverage.

It verifies that:

- chat requests include the active anonymous identity in `X-Anonymous-ID`
- `runAgentTurn` creates a session before the first chat turn when no session exists yet
- invalid agent envelopes are rejected in the client instead of silently accepted
- streamed chat parsing tolerates split SSE chunks and reconstructs the final envelope without depending on a separate SSE parser library

This matters because the dashboard depends on these helpers for session bootstrap and streaming UI updates. If you change session creation, SSE formatting, or envelope shape, this is one of the first tests that should fail.

### Resume flow: UI behavior after token consumption succeeds or fails

`apps/web/src/routes/resume.test.tsx` makes the resume route more than a visual page test. It locks down that:

- a successful resume clears the token from route search params, invalidates session queries, and navigates the user back into `/dashboard` with the resumed `sessionId`
- an expired token renders the expired-token UI instead of forwarding the user into the dashboard
- generic network failure renders a connection-error state

Together with backend resume-session tests, this makes the resume flow one of the better-defined end-user journeys in the repository.

### Dashboard component tests: useful leaf coverage, but not orchestration coverage

There is some real dashboard coverage, but it is concentrated in isolated components and data mapping:

- `agent-mapping.test.ts` verifies that recommendation products are mapped from authoritative catalog fields and do not invent synthetic performance metrics.
- `components/audit-trail.test.tsx` verifies that an active question envelope renders after completed logs and that user interaction payloads include the expected `question.answer` action and selected option IDs.
- `components/accessories-modal.test.tsx` verifies filter behavior, stale-offer checkout disabling, and pagination/reset behavior when filters change.
- `components/comparison-modal.test.tsx` checks a layout invariant for four-product comparison rendering.

This is valuable, but it mostly protects **presentation rules and interaction payloads inside components**. It does not yet prove the full dashboard continues to orchestrate:

- session loading
- agent turn submission
- streaming token updates
- envelope-to-surface updates
- comparison / checkout transitions
- accessory and audit-trail state together

So the dashboard remains a manual-QA-heavy subsystem even though it has several focused tests.

### AI runtime HTTP contract: agent envelope and SSE event names

`services/ai-runtime/tests/integration/test_chat_endpoint.py` defines two important external boundaries.

First, `/api/v1/chat` accepts a full message history and returns a structured agent envelope with `schema_version` `1.1` and question data.

Second, `/api/v1/chat/stream` emits named SSE events including:

- `token`
- `ui_operation`
- `envelope`
- `done`

Those names matter because the frontend streaming client parses named blocks, accumulates token text, and prefers the explicit `envelope` event when present.

The same integration file also establishes a resilience boundary for `/api/v1/greeting`: if the provider fails, the endpoint still returns a localized fallback instead of failing the dashboard greeting request.

### AI runtime workflow: prompt invariants and hydration rules

`services/ai-runtime/tests/unit/test_graph.py` is the strongest executable description of AI runtime behavior.

It verifies that the workflow:

- invokes the configured model name and forwards the complete user/assistant history into the provider call
- hydrates recommendation selections from catalog data so returned products carry real catalog IDs and prices
- hydrates offer-comparison responses from backend tool data when a session-approved product is supplied
- encodes prompt rules for gaming conversations: collect workload, ambition, and budget; ask exactly one missing dimension at a time; treat “flexible”/“no preference” as already answered; and prefer workload before budget
- can proceed directly to a recommendation when all three gaming dimensions are already present in the initial user message

These are some of the most important behavioral invariants in the repository because they constrain how prompt edits may change conversation flow.

### Session-only action boundaries inside the AI runtime

The workflow and schema code also enforce a security/consistency boundary around comparison-like actions.

`ChatRequest` includes `allowed_comparison_ids`, the workflow passes those IDs into both the system prompt and hydration logic, and `hydrate_agent_response` rejects `comparison`, `offer_comparison`, and `checkout_ready` selections that are not in that allowed session list.

That means these response types are not allowed to point at arbitrary catalog products; they must stay within the session-approved set forwarded by the caller.

### Main backend decision-memory handler: what is forwarded and when chat is blocked

`services/main-backend/internal/decision_memory/handler/chat_test.go` defines the backend boundary between saved session state and the AI runtime.

It proves that a chat request:

- loads prior session history
- appends the new user message before forwarding to the AI runtime
- extracts previously recommended product IDs from assistant reasoning graphs and forwards them as `allowed_comparison_ids`
- persists the assistant response back into session messages

The same test file also proves two important failure boundaries:

- an anonymous user cannot chat against another anonymous user’s session
- an anonymous user cannot read another anonymous user’s session history

In both cases the backend returns `404` and, for unauthorized chat, never calls the AI runtime.

### Resume-session service: lifecycle invariants for share/resume tokens

`services/main-backend/internal/resume_session/usecase/service_test.go` defines the minimum lifecycle rules for resume tokens.

The tested invariants are:

- a generated token can resume the target session
- consuming the same token twice returns `ErrTokenConsumed`
- generating a newer token revokes previously unconsumed tokens for that session, and revoked tokens return `ErrTokenRevoked`
- invalid JWT signatures return `ErrTokenInvalid`
- a different context hash does not hard-block resume
- an otherwise valid but expired token returns `ErrTokenExpired`
- inactivity notifications only succeed when the decision session still belongs to an authenticated user with a verified notification identity

This is the clearest source for what “resume links” actually guarantee today.

### OpenAPI contract tests: spec file and generated Swagger must stay aligned

`services/main-backend/internal/platform/testutil/contract/openapi_test.go` plus `openapi.go` define a lightweight but important contract-testing layer.

The tests assert that the onboarding OpenAPI spec contains required identity paths and response codes, and that generated Swagger documents every OpenAPI path/method from the source contract. The helper code resolves the canonical spec at `specs/002-get-started-onboarding/contracts/openapi.yaml` and compares it with `services/main-backend/docs/swagger.yaml`, allowing equivalent path parameters even when placeholder names differ.

This is the repository’s main evidence-backed guard against docs/spec drift for public backend endpoints.

## Under-tested hotspots

### Dashboard orchestration is still the biggest gap

The dashboard has meaningful tests around mapping, accessory display, question rendering, and comparison layout, but there is still no evidence here of one comprehensive automated test for the end-to-end dashboard flow.

Manual QA is still important whenever a change touches any of these together:

- `runAgentTurn` or `runAgentTurnStream`
- dashboard state machines or route-level loaders
- active-question handling and follow-up interactions
- envelope parsing or `ui_operations`
- transitions into comparison, offer comparison, or checkout-ready states

A component test passing does not prove the dashboard still orchestrates correctly across network, session, and streamed UI state.

### Greeting behavior is resilient, but not deeply tested in the web app

The AI runtime proves localized fallback behavior for greeting failures, but the supplied frontend evidence does not show a route-level or dashboard-level test that the greeting area degrades correctly in the browser. Manual verification still matters after greeting contract or locale changes.

### Contract tests cover identity paths, not every subsystem boundary

The OpenAPI contract tests are strong for the documented identity surface, but they do not imply equivalent formal contract coverage for decision memory, dashboard interactions, or resume flows. Those areas depend more on handler/use-case tests and frontend API-schema checks.

## Minimum verification routines by change area

### If you change dashboard orchestration or chat UX

Run at minimum:

- `pnpm --filter web test`
- `pnpm --filter web check-types`

Then manually verify:

- first-turn session creation from the dashboard
- follow-up question rendering and submission
- streamed responses updating progressively
- at least one recommendation flow
- if touched, comparison / offer-comparison / checkout-ready transitions

This is the area where manual QA is most important because current automation is broadest at the edges, not at the whole route orchestration layer.

### If you change the web decision-memory client or SSE parsing

Run at minimum:

- `pnpm --filter web test`
- focus on `src/api/decision-memory.test.ts`
- `pnpm --filter web check-types`

Also manually verify one real streamed conversation, because the client reconstructs envelopes from named SSE blocks and split chunks are only one subset of real browser/runtime behavior.

### If you change resume links, token handling, or resume-page behavior

Run at minimum:

- `pnpm --filter web test`
- `pnpm --filter @shopwise/main-backend test`

Manually verify:

- valid token resumes into the dashboard
- expired token shows the expired state
- invalid token shows a safe failure state
- a previously consumed token is rejected

If token issuance or revocation changed, also check that issuing a newer token invalidates older unconsumed links as the backend tests expect.

### If you change main-backend decision-memory chat handling

Run at minimum:

- `pnpm --filter @shopwise/main-backend test`
- `pnpm --filter @shopwise/main-backend test:race`

Manually verify:

- session history is preserved across turns
- anonymous session isolation still works
- comparison-eligible products still come only from prior assistant recommendations
- streamed chat still persists the final assistant message correctly

### If you change AI runtime prompts, workflow, or response schemas

Run at minimum from `services/ai-runtime`:

- `uv run pytest`

Pay special attention to:

- `tests/unit/test_graph.py`
- `tests/integration/test_chat_endpoint.py`

Then manually verify one end-to-end conversation through the backend/web path, because unit tests prove structural invariants, not whether the full stack still agrees on response semantics.

### If you change OpenAPI or generated backend docs

Run at minimum:

- `pnpm --filter @shopwise/main-backend test`

And verify:

- the contract tests still pass for required paths and status codes
- generated Swagger still documents every path/method defined in the source OpenAPI contract

## Good default verification routine

For any non-trivial change that crosses subsystem boundaries:

1. Run the nearest package tests first.
2. Run the package type/build checks (`vite build && tsc --noEmit` or `go vet`, depending on the area).
3. If backend chat/session logic changed, include `test:race` for the Go service.
4. If AI envelopes, SSE, or dashboard orchestration changed, execute one real manual flow in the browser.
5. Re-check the contract edges that tests already define well: anonymous session headers, resume-token behavior, agent envelope validity, and named SSE events.

That routine matches the actual repository risk profile better than a generic “run everything” rule: the most expensive regressions here usually happen at **service boundaries and orchestration seams**, not inside isolated pure functions.

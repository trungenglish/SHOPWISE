# Testing Guide

## Testing layers in this repository

### Web app
The web package (`/apps/web/package.json`) uses:
- `vitest run` for tests
- `vite build && tsc --noEmit` for type checks

Visible tests in the current tree include:
- `/apps/web/src/routes/resume.test.tsx`
- `/apps/web/src/components/expired-token-ui.test.tsx`

This suggests frontend test coverage currently concentrates on the resume-token/error-state flow rather than broad dashboard coverage.

### Go backend
The backend package (`/services/main-backend/package.json`) exposes:
- `test`
- `test:unit`
- `test:integration`
- `test:race`
- `check-types` via `go vet`

There are domain-local `_test.go` files under `internal/**`, including handler and service coverage for areas such as users, identity, orders, and resume session.

### AI runtime
The AI runtime has both unit and integration tests:
- `/services/ai-runtime/tests/unit/test_graph.py`
- `/services/ai-runtime/tests/unit/test_openai_provider.py`
- `/services/ai-runtime/tests/unit/test_self_correction.py`
- `/services/ai-runtime/tests/unit/test_streaming.py`
- `/services/ai-runtime/tests/integration/test_chat_endpoint.py`

This reflects the runtime's architecture: provider behavior, self-correction, streaming, and endpoint behavior are tested independently.

## Recommended commands by area

### Whole repo sanity checks
- `pnpm check-types`
- `pnpm build`
- `pnpm check`

### Web changes
- `pnpm --filter web test`
- `pnpm --filter web check-types`

### Backend changes
- `pnpm --filter @shopwise/main-backend test`
- `pnpm --filter @shopwise/main-backend test:race`
- `pnpm --filter @shopwise/main-backend check-types`

### AI runtime changes
- run from `services/ai-runtime`: `uv run pytest`

## What the current tests tell you

### Resume flow is comparatively well-defined
The presence of `resume.test.tsx`, `expired-token-ui.test.tsx`, and backend resume-session tests indicates the resume flow is one of the more intentionally exercised user journeys.

### AI runtime correctness focuses on structure, not product semantics
The AI runtime tests validate things like:
- graph invocation
- malformed JSON handling
- endpoint behavior when no API key is configured

That means they are good for contract safety, but they do not prove that recommendations are product-correct.

### Dashboard behavior is still under-tested
The dashboard is large and recently evolved, but there is little visible route/component test coverage for it in the files inspected. Changes there should be validated carefully with manual QA and, ideally, expanded tests.

## Practical testing guidance by feature

### If changing decision memory
Run:
- web tests touching resume/session UI
- backend tests for `decision_memory` and related handlers
- manual session CRUD checks in the UI

Also verify:
- anonymous session handling
- timestamp conflict behavior
- branch/restore behavior

### If changing resume links
Run:
- frontend resume tests
- backend resume-session tests
- manual checks for missing, expired, consumed, and invalid tokens

### If changing AI runtime schema or workflow
Run:
- AI runtime unit/integration tests
- manual end-to-end check through the backend proxy if applicable
- verify the frontend still understands the response shape

### If changing checkout/orders
Run:
- backend order tests
- manual checks in both auth and local bypass modes
- UI verification in order history components

## Known gaps and risks

- The dashboard route mixes orchestration, mock data, and UI state, but has limited visible automated test coverage.
- Root-level `pnpm build` and type checks may catch integration drift even when feature-local tests are sparse.
- Specs describe behavior that may not yet be enforced in tests; do not assume spec coverage implies implementation coverage.

## Good default verification routine

For non-trivial changes:
1. Run feature-local tests
2. Run package-level type checks
3. Run at least one end-to-end manual flow through the UI
4. Re-check any hardcoded ports/URLs involved in service-to-service integration

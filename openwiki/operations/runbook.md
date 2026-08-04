# Operations Runbook

## Monorepo commands

From the repository root (`/package.json`):
- `pnpm dev` — run Turbo dev tasks
- `pnpm build` — build all configured packages/apps
- `pnpm check-types` — run type-check pipelines
- `pnpm dev:web` — run only the web app
- `pnpm dev:server` — run only the Go backend
- `pnpm dev:server:docker` — run the backend stack via Docker Compose
- `pnpm check` / `pnpm fix` — run Ultracite quality checks/fixes

Turbo configuration is in `/turbo.json`.

## Web app local run

The root README documents the web app at `http://localhost:3001`.

Primary files:
- `/apps/web/package.json`
- `/apps/web/src/main.tsx`
- `/apps/web/src/routes/*`

## Main backend local run

Developer commands live in `/services/main-backend/package.json`:
- `pnpm --filter @shopwise/main-backend dev`
- `pnpm --filter @shopwise/main-backend dev:deps`
- `pnpm --filter @shopwise/main-backend dev:docker`
- `pnpm --filter @shopwise/main-backend seed`
- `pnpm --filter @shopwise/main-backend test`
- `pnpm --filter @shopwise/main-backend swagger:gen`

The backend entrypoint is `/services/main-backend/cmd/retail/main.go`.

### Backend environment shape
Use `/services/main-backend/.env.example` as the non-secret reference. Important variables include:
- `PORT` (default `18080`)
- `GIN_MODE`
- `CHECKOUT_AUTH_BYPASS`
- `ALLOWED_ORIGINS`
- `DATABASE_URL`
- `REDIS_URL`
- `JWT_SECRET`, `JWT_ACCESS_TTL`, `JWT_REFRESH_TTL`
- Google OAuth settings
- `LLM_PROVIDER`, `LLM_API_KEY`, `LLM_MODEL`
- `STORAGE_PATH`

Do not document or commit real secret values.

### Backend operational notes
- The backend expects Postgres and Redis.
- Bootstrap wires DB, cache, jobs, auth, orders, sessions, resume, catalog, and stores in one process.
- `CHECKOUT_AUTH_BYPASS` can relax auth for checkout/order endpoints in local development, but only when `GIN_MODE=debug`; this is enforced in `/services/main-backend/internal/platform/config/config.go` and covered by `config_test.go`.
- Swagger docs are generated from the Go service annotations.

### Migrations and seeding
`cmd/retail/main.go` supports flags handled through bootstrap:
- `-migrate`
- `-seed`

Bootstrap runs migrations for users, identity, orders, decision memory, and resume session repositories.

## AI runtime local run

The AI runtime README (`/services/ai-runtime/README.md`) documents:
- Python 3.11+
- `uv`
- `uv run uvicorn src.main:app --reload --port 8000`
- `uv run pytest`

Important source files:
- `/services/ai-runtime/src/main.py`
- `/services/ai-runtime/src/api/chat.py`
- `/services/ai-runtime/src/graph/workflow.py`

### AI runtime env expectations
The runtime expects provider config through environment variables and fails closed when the LLM API key is absent. This is reflected in `get_provider()` inside `/services/ai-runtime/src/api/chat.py`.

## Known local-development mismatches

These are likely sources of confusion during setup:

1. **Backend port mismatch**
   - docs/env default to `18080`
   - dashboard code still fetches products from `http://localhost:8080/api/v1/products`

2. **AI runtime path/port mismatch**
   - runtime exposes `/api/v1/chat` and `/api/v1/chat/stream`
   - backend proxy in `decision_memory/handler/chat.go` points to `http://localhost:8001/chat/stream`
   - runtime README documents port `8000`

3. **README drift**
   - root README still includes older scaffold descriptions that no longer match the active service layout exactly

When debugging startup issues, trust current source files more than old scaffold text.

## Suggested startup order for full-stack work

1. Start Postgres + Redis
2. Start `services/main-backend`
3. Start `services/ai-runtime`
4. Start `apps/web`
5. Verify `/health` on backend and runtime before testing UI integration

## Operational checkpoints before merging changes

- If you changed backend routes, re-check `bootstrap.go`
- If you changed env usage, verify `.env.example` still matches expectations without exposing secrets
- If you changed AI runtime schema or endpoints, check the backend proxy caller and web consumers
- If you changed checkout or auth behavior, verify whether bypass mode affects your test results
- If you changed docs generation or OpenWiki workflow, keep generated content confined to `/openwiki`

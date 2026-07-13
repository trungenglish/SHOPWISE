# Integration Points

## Overview

SHOPWISE has both internal and external integration boundaries. The important ones today are:
- web app ↔ Go backend
- Go backend ↔ AI runtime
- Go backend ↔ Postgres/Redis
- backend identity ↔ Google OAuth
- resume flow ↔ notification provider (currently Zalo-shaped)
- AI runtime ↔ OpenAI-compatible LLM provider

## Web to backend integration

The web app talks to the Go backend over REST.

Examples:
- `/apps/web/src/api/decision-memory.ts` → `/api/v1/sessions` and `/api/v1/session/resume`
- dashboard product fetch in `/apps/web/src/routes/dashboard.tsx` → `/api/v1/products`
- orders UI → checkout/order endpoints

### What to watch
- bearer auth is optional for some decision-memory routes because anonymous usage is supported
- anonymous session requests rely on `X-Anonymous-ID`
- hardcoded local URLs currently exist in some frontend code

## Backend to AI runtime integration

The most explicit caller is `/services/main-backend/internal/decision_memory/handler/chat.go`.

Responsibilities split this way:
- backend validates session identity and stores messages
- AI runtime generates structured/model output
- backend relays streaming output and persists the final assistant message

### Current mismatch
The backend proxy target does not currently match the runtime's documented path or route prefix. If you touch this area, verify all of the following together:
- runtime listening port
- runtime route prefix (`/api/v1`)
- backend proxy URL
- any frontend assumptions about the caller path

## Backend infrastructure integrations

### Postgres
The backend depends on Postgres for durable application data. Repositories under `internal/*/repository/postgres` are the main integration layer.

### Redis
Redis is used by backend infrastructure and job wiring through bootstrap. Even if a specific feature does not touch Redis directly, service startup expects it.

## Identity integrations

### JWT auth
The backend uses JWT services in both identity and resume flows.

There are two distinct JWT-related usages visible in the code inspected:
- regular authenticated user access (`identity` domain)
- single-use resume tokens (`resume_session` domain)

These should not be treated as the same trust level or lifecycle.

### Google OAuth
Bootstrap wires a Google OAuth service through the identity domain. Relevant config placeholders appear in `/services/main-backend/.env.example`.

## Notification provider integration

The resume session flow is intentionally pluggable. `/services/main-backend/internal/resume_session/usecase/service.go` depends on a `NotificationProvider` abstraction and currently uses a Zalo provider wiring in bootstrap.

Why this matters:
- product scope today is Zalo-focused
- architecture intends future provider replacement/expansion without rewriting resume logic

## LLM provider integration

The AI runtime currently supports an OpenAI provider path through `/services/ai-runtime/src/api/chat.py` and `/services/ai-runtime/src/llm/openai_provider.py`.

Important characteristics:
- API keys stay server-side
- provider is selected/configured through environment settings
- request/response behavior is wrapped behind provider methods
- structured output is enforced using a schema generated from Pydantic models

## Shared contract surfaces

### Structured AI response schema
`/services/ai-runtime/src/models/schemas.py` is one of the most important integration contracts in the repo. It defines the shape expected between model orchestration and downstream consumers.

If fields change here, verify:
- AI runtime workflow
- backend streaming/proxy expectations
- frontend rendering code for products, logs, agents, accessories, and reasoning

### Shared packages
Internal packages such as `@shopwise/env`, `@shopwise/ui`, and related type/schema packages are also integration points because they standardize conventions across apps.

## Git-history signals worth remembering

Recent history shows integration work moving in this order:
- resume-session security and Zalo notification foundations
- trust/readiness/dashboard UX expansion
- AI runtime modular refactor
- catalog/store handlers and decision-memory/backend wiring
- dashboard route consolidation around agent-assisted shopping

That progression helps explain why some integration surfaces are mature in design but still uneven in implementation detail.

## Safe-change checklist

When editing an integration point, verify the matching boundary on both sides:

- **Web API helper changed?** Check backend handler and route registration.
- **Backend proxy URL changed?** Check runtime route prefix and local run docs.
- **Schema changed?** Check frontend consumers and backend pass-through assumptions.
- **Auth behavior changed?** Check anonymous decision-memory flow and bypass flags.
- **Notification flow changed?** Check token issuance, revocation, and log persistence together.

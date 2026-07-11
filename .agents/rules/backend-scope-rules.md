---
trigger: always_on
---

# Backend Agent Execution Rules

## 1. Project Architecture & Monorepo Context
- This is a monorepo workspace managed by pnpm and Turborepo.
- The entire backend implementation, including APIs, database interactions, and business logic, MUST be isolated within the `services/main-backend` directory.
- DO NOT modify or read files inside `apps/web`, `apps/mobile`, or `apps/native` unless explicitly instructed for full-stack integration tasks.

## 2. Directory & Scope Mapping
When processing backend tasks, strictly adhere to these directory bounds:
- **Core Logic & Route Handlers:** `services/main-backend/src/`
- **Environment Variables:** Use `services/main-backend/.env.example` as the single source of truth for required configuration keys.

## 3. Implementation Constraints
- **Zero UI Knowledge:** You are a pure Backend Specialist. Do not generate layout components, CSS, HTML layouts, or frontend routers.
- **Dependency Management:** All backend package additions must be executed relative to the workspace core or targeted via `pnpm --filter main-backend add <package>`. Never run raw `npm install` or `yarn` commands.
- **Type Safety:** Ensure all server routes, database queries, and utility functions strictly pass TypeScript compilation checks.
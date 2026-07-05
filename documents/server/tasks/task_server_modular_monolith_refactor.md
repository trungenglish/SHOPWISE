# Task: Server modular monolith refactor

**Task ID:** `V1_MVP/01_Server/task_01_modular_monolith_refactor.md`  
**Version:** V1_MVP  
**Priority:** High  
**Status:** Done  
**Created Date:** 2026-06-14  
**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## 1. Detailed Description

Refactor `apps/server` from flat packages into a modular monolith aligned with General_diagram.md: six bounded modules, `platform/` shared infra, OpenAPI, and asynq worker.

---

## 2. Implementation Steps

- [x] 1. Extract `internal/platform/*` from legacy flat packages
- [x] 2. Scaffold identity, business, files, notification, administration modules
- [x] 3. Implement reference `users` module (CRUD + ports)
- [x] 4. Add `platform/job`, `cmd/worker`, compose worker service
- [x] 5. Integrate swaggo OpenAPI + Swagger UI (debug)
- [x] 6. Update Dockerfile, package.json scripts
- [x] 7. Add unit + integration tests
- [x] 8. Document under `documents/server/` and `documents/tests/`

---

## 3. Completion Criteria

- [x] `go build ./cmd/server` and `./cmd/worker` succeed
- [x] `go test ./internal/users/...` passes
- [x] `GET /health`, `GET/POST /api/v1/users` functional
- [x] Swagger UI available in debug mode
- [x] System design docs match implementation

---

## 4. Technical Specifications

### Package layout

See [system-design.md](../system-design.md) §3.

### Files created/modified

| Path | Action |
| --- | --- |
| `internal/platform/**` | Created (migrated infra) |
| `internal/{identity,users,business,files,notification,administration}/**` | Created |
| `cmd/worker/main.go` | Created |
| `docs/swagger.json` | Generated |
| `documents/server/**` | Created |
| `documents/tests/**` | Updated |

---

## 6. Related Documents

- [system-design.md](../system-design.md)
- [adr/001-modular-monolith-asynq.md](../adr/001-modular-monolith-asynq.md)
- [server-testing.md](../../tests/server-testing.md)

# ADR 001: Modular monolith with asynq background worker

**Status:** Accepted  
**Date:** 2026-06-14  
**Deciders:** Platform team

## Context

The boilerplate server started as a flat Gin layout (`handler/`, `repository/`) with two demo routes. [General_diagram.md](../../General_diagram.md) requires a modular monolith with REST/OpenAPI and a background worker for email, push, and similar tasks.

Team size is small (MVP). Traffic is low. Postgres and Redis already run in docker-compose.

## Decision

1. **Architecture:** Modular monolith with feature modules (`identity`, `users`, `business`, `files`, `notification`, `administration`) and shared `platform/` infrastructure.
2. **Layers:** Handler → Usecase → Repository per module.
3. **API contract:** OpenAPI via swaggo; Swagger UI in non-release mode.
4. **Background jobs:** Separate `cmd/worker` binary using **asynq** on existing Redis (not RabbitMQ).
5. **Reference module:** `users` fully wired with CRUD and welcome-email enqueue.

## Alternatives considered

| Option | Rejected because |
| --- | --- |
| Microservices | Complexity cost exceeds team/traffic; no measured scaling pain |
| RabbitMQ | Extra infra; Redis already present; architect growth tier allows Redis queue |
| In-process goroutine workers | Couples job processing to HTTP lifecycle; harder to scale independently |
| Flat handler layout | Does not scale to six bounded contexts |

## Consequences

**Positive**

- Clear module boundaries for future auth, files, and admin features
- Worker scales independently in compose/K8s
- OpenAPI enables web/native client generation

**Negative**

- More packages and wiring in `cmd/server/main.go`
- asynq has fewer routing features than RabbitMQ (acceptable for MVP)

## Compliance

- Go module path: `shopwise/apps/server`
- Docs link: [quanngynx tree URL](https://github.com/quanngynx/shopwise/tree/main/apps/server)

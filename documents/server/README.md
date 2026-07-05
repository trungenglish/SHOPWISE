# Server documentation

Documentation for the Go Gin modular monolith in `apps/server`.

**Source tree:** [apps/server on GitHub](https://github.com/quanngynx/shopwise/tree/main/apps/server)

## Index

| Document | Description |
| --- | --- |
| [system-design.md](./system-design.md) | C4 architecture, module map, sync/async flows |
| [adr/001-modular-monolith-asynq.md](./adr/001-modular-monolith-asynq.md) | ADR: modular monolith + asynq worker |
| [modules/users-prd.md](./modules/users-prd.md) | Users module PRD |
| [modules/user-create-welcome-email.md](./modules/user-create-welcome-email.md) | Async welcome-email business flow |
| [tasks/task_server_modular_monolith_refactor.md](./tasks/task_server_modular_monolith_refactor.md) | Refactor implementation task |

## Related

- [General_diagram.md](../../General_diagram.md) — system context diagram
- [OpenAPI spec](../../apps/server/docs/swagger.json) — generated REST contract
- [Testing guide](../tests/server-testing.md) — how to run server tests

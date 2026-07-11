# Project Structure

## Monorepo

```text
shopwise/

apps/
    web/
    mobile/

services/
    gateway/
    ai-runtime/
    retail/

packages/
    protocols/
    schemas/
    sdk/
    ui/
    types/

infra/

docs/

scripts/

deploy/
```

---

## Philosophy

```text
Business Logic

↓

Service

↓

Protocol

↓

Schema

↓

Generated Types
```

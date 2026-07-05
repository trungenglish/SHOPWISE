# Users - Product Requirements Document

**Document Version:** 1.0  
**Created Date:** 2026-06-14  
**Last Updated:** 2026-06-14  
**Status:** Approved  
**Repository:** [apps/server](https://github.com/quanngynx/shopwise/tree/main/apps/server)

---

## 1. Executive Summary

### 1.1 Purpose

The users module is the reference bounded context for the modular monolith. It manages user profiles (email, name) and demonstrates Handler → Usecase → Repository layering, OpenAPI documentation, and async notification enqueue on create.

### 1.2 Scope

**In scope:** CRUD API for users, welcome-email job enqueue, OpenAPI annotations.  
**Out of scope:** Authentication, roles, tenant isolation, profile images.

---

## 2. Functional Requirements

#### FR-001: List users

- **Priority:** High
- **Acceptance Criteria:**
  - [ ] `GET /api/v1/users` returns paginated list
  - [ ] Supports `limit` and `offset` query params

#### FR-002: Create user

- **Priority:** Critical
- **Acceptance Criteria:**
  - [ ] `POST /api/v1/users` creates user with email + name
  - [ ] Enqueues `notification:send_welcome_email` job
  - [ ] Returns 409 on duplicate email

#### FR-003: Get / update / delete user

- **Priority:** High
- **Acceptance Criteria:**
  - [ ] `GET /api/v1/users/{id}` returns user or 404
  - [ ] `PATCH /api/v1/users/{id}` updates email and/or name
  - [ ] `DELETE /api/v1/users/{id}` removes user

---

## 3. Data Model

```sql
Table users {
  id uuid [pk]
  email varchar(255) [unique, not null]
  name varchar(255) [not null]
  created_at timestamptz
  updated_at timestamptz
}
```

---

## 4. API Endpoints

| Method | Endpoint             | Description                    |
| ------ | -------------------- | ------------------------------ |
| GET    | `/api/v1/users`      | List users (`limit`, `offset`) |
| POST   | `/api/v1/users`      | Create user                    |
| GET    | `/api/v1/users/{id}` | Get user                       |
| PATCH  | `/api/v1/users/{id}` | Update user                    |
| DELETE | `/api/v1/users/{id}` | Delete user                    |

OpenAPI: `apps/server/docs/swagger.json`

---

## 5. Dependencies

| Dependency          | Type     | Description                |
| ------------------- | -------- | -------------------------- |
| `platform/job`      | Required | Enqueue welcome email      |
| `notification/jobs` | Required | Process welcome email task |
| `identity`          | Future   | Auth middleware            |

---

## Appendix B: References

- [system-design.md](../system-design.md)
- [user-create-welcome-email.md](./user-create-welcome-email.md)
- [OpenAPI spec](../../../apps/server/docs/swagger.json)

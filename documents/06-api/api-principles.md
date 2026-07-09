# API Principles

Version: 1.0

---

## Philosophy

APIs expose application capabilities.

They do not expose database tables.

They do not expose internal services.

---

## Design Rules

Resource-oriented.

Capability-first.

Versioned.

Stateless.

JSON-based.

Streaming capable.

Idempotent where appropriate.

---

## API Categories

Public APIs

Internal APIs

Streaming APIs

Provider APIs

Administration APIs

---

## Response Format

Every response contains

status

data

metadata

errors

---

## Error Format

code

message

details

correlation_id

# Phase 0: Research & Architecture Decisions

## Component Type Modeling

- **Decision**: Use discriminated unions with a strict `type` field.
- **Rationale**: Strongly typed languages like Go and TypeScript, and validation libraries like Pydantic in Python, have excellent support for discriminated unions. By including a `type` string field on every node, deserializers can easily map JSON objects to the correct class/struct.
- **Alternatives considered**: Duck typing (inferring type from fields) – rejected due to complexity in Go/Python.

## Event Routing & Component Identification

- **Decision**: Every interactive component requires a unique `id`. Events emitted back to the backend include the `componentId`, an `action` string, and a structured `payload`.
- **Rationale**: The backend can maintain a map of `componentId`s or just route based on the semantic `action`. This keeps the frontend entirely unaware of backend orchestration.
- **Alternatives considered**: Implicit routing based on the tree path – rejected as UI structures may shift during partial updates.

## UI Update Operations

- **Decision**: The root document includes an `operation` field (e.g., `replace`, `append`, `patch`, `remove`). For partial updates, it specifies a `targetId` to identify which existing component the operation applies to.
- **Rationale**: This allows the AI to stream independent chunks of UI (e.g., appending a new chat message) without re-sending the entire UI tree, saving bandwidth and rendering time.
- **Alternatives considered**: Always re-sending the entire UI state – rejected due to latency and bandwidth constraints on mobile.

## Validation Strategy

- **Decision**: Implement a dual-validation model. The JSON schema defines client-side constraints (like `required`, `pattern`, `minItems`), while the backend uses a dedicated `ErrorState` component to report business rule violations.
- **Rationale**: Provides immediate feedback for simple formatting errors but keeps complex pricing/inventory logic on the backend.
- **Alternatives considered**: Only backend validation – rejected due to poor UX (high latency for simple typo feedback).

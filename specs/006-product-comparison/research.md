# Phase 0: Research & Architecture Decisions

## Decision 1: State Management Separation

**Decision**: Separate Comparison State from Conversation State, and persist it in Decision Memory.
**Rationale**: The comparison workspace must survive across devices and channels (Zalo, Web). It should not be ephemeral to a single conversation thread. Decision Memory on the backend will own the persistence of `comparison_workspace_state`, while the frontend will merely render it.
**Alternatives considered**:
- *Client-side local storage*: Rejected because it breaks cross-device resumption.
- *Conversation state*: Rejected because comparisons might span multiple conversations or be resumed later.

## Decision 2: Dynamic UI Protocol Rendering

**Decision**: Implement the workspace as a composite UI schema containing nested components (Comparison Table, Difference Highlights, Action Bar).
**Rationale**: Allows the backend/AI to incrementally patch the UI (e.g., adding a 5th product triggers a replacement dialog schema patch, not a full page reload).
**Alternatives considered**:
- *Custom React components specifically for comparison*: Rejected as it violates the Declarative UI principle; the AI must generate schemas, not hardcoded React components.

## Decision 3: AI Explanation Generation

**Decision**: AI explanation generation is triggered as a semantic event after the product data is fully resolved by the Go backend.
**Rationale**: Adheres to the principle that AI must never fabricate data. The Go backend fetches the spec data, constructs the context, and passes it to the AI Runtime to stream the explanation.
**Alternatives considered**:
- *AI fetches specs directly*: Rejected because AI should use Tool Protocol and Backend owns retail logic.

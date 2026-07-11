# Research & Decisions: Decision Memory

All technical clarifications were resolved during the specification phase.

## Decisions Log
- **Branching Implementation**: Full-copy duplication of session up to the branch point.
  - *Rationale*: Simplifies query logic, reduces complex JOINs for the client, and ensures historical integrity.
  - *Alternatives considered*: Pointer/Delta method (rejected due to query complexity).
- **Conflict Resolution**: Timestamp-based Last-Write-Wins (LWW).
  - *Rationale*: Simplest to implement, aligns with typical e-commerce session patterns where concurrent multi-device rapid editing is rare.
  - *Alternatives considered*: CRDTs (overkill), explicit locking (bad UX).
- **Anonymous Storage Limits**: Rolling window of 10 sessions (oldest evicted).
  - *Rationale*: Prevents unbounded `localStorage` growth while keeping high utility for casual browsers.
  - *Alternatives considered*: Force login (adds friction).
- **Auto-save UI**: Subtle indicator (e.g., small checkmark).
  - *Rationale*: Reassures users without interrupting their shopping flow.

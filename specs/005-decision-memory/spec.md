# Decision Memory

**Version**: 1.0.0
**Status**: Draft
**Feature Directory**: specs/005-decision-memory
**Created**: 2026-07-11
**Last Updated**: 2026-07-11

---

## Overview

Decision Memory is an AI-driven feature that automatically saves and organizes a shopper's product discovery sessions, preferences, and AI reasoning. It empowers users to seamlessly revisit, resume, and branch their past shopping journeys across devices without losing context or having to restart the discovery process.

---

## Clarifications

### Session 2026-07-11

- Q: When an authenticated user has the same session open on two different devices concurrently, how should we handle state conflicts? → A: Use simple last-write-wins based on timestamp.
- Q: How should session branching be implemented at the data level? → A: Full Copy: Duplicate the session state up to the branching point into a new, independent session.
- Q: Should the UI visually indicate to the user that their session is being auto-saved? → A: Unobtrusive visual indicator (e.g., a small "Saved to Memory" status icon or subtle toast that doesn't interrupt flow).
- Q: For unauthenticated users storing sessions locally, what happens when they reach local storage limits? → A: Keep a rolling window of the last 10 anonymous sessions (oldest evicted automatically).

---

## Problem Statement

When shoppers leave the platform or start a new session, they currently lose their product comparison context, pinned items, and the AI's reasoning. This forces them to restart their discovery process from scratch in subsequent visits, leading to user frustration, repeated questions, and abandoned shopping journeys.

---

## Goals

- Enable shoppers to automatically save and resume shopping sessions across devices.
- Preserve full context including conversation state, pinned products, comparison workspaces, and AI reasoning.
- Remember explicit and implicit user preferences (budget, brand, usage) to personalize future recommendations.
- Allow users to review the AI's reasoning and evidence for previous recommendations.
- Provide tools to manage (search, rename, archive, delete) saved decision sessions.

## Non-Goals

- Checkout, order, or payment history tracking.
- Retail CRM synchronization.
- Marketing personalization campaigns.
- Changes to the underlying recommendation ranking algorithms.

---

## Target Users

| User Type | Description | Primary Need |
|---|---|---|
| Shopper | A user actively researching and comparing products | Needs to continue complex product comparisons over multiple sessions and devices without losing context. |

---

## User Scenarios & Testing

### Scenario 1: Auto-saving a session

**Given**: A shopper is actively comparing laptops.
**When**: The shopper closes the browser or app.
**Then**: The session is automatically saved with its current state.
**Acceptance**: Upon returning to the "Saved Decisions" area, the session appears with the correct timestamp and status.

### Scenario 2: Resuming a previous session

**Given**: A shopper has a saved session for "Gaming Laptops".
**When**: The shopper clicks to reopen the session.
**Then**: The workspace, pinned products, and conversation history are fully restored.
**Acceptance**: The user can immediately ask a follow-up question without re-explaining their requirements.

### Scenario 3: Replaying decision reasoning

**Given**: A shopper is reviewing a past recommendation for a specific monitor.
**When**: The shopper clicks to view the reasoning.
**Then**: The UI displays the AI's step-by-step logic and evidence used to make that recommendation.
**Acceptance**: The evidence references are accessible and clearly explain the "why".

### Scenario 4: Branching a decision

**Given**: A shopper wants to explore a different budget for an existing laptop comparison.
**When**: The shopper starts a new conversation from an earlier point in the session.
**Then**: A new branched session is created, preserving the original session's history intact.
**Acceptance**: Both the original and branched sessions are accessible in the history manager.

---

## Functional Requirements

### Core Memory Operations

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-01 | Session Auto-save | Must Have | State is persisted automatically without explicit user action; survives browser refresh; UI displays an unobtrusive visual indicator (e.g., "Saved to Memory"). |
| FR-02 | State Restoration | Must Have | Resuming a session restores chat history, pinned products, and comparison workspace. |
| FR-03 | Preference Extraction | Must Have | System extracts, stores, and applies user preferences (budget, brand, etc.) to new recommendations. |
| FR-04 | Preference Editing | Must Have | Users can view, edit, or delete stored preferences. |
| FR-05 | Reasoning Replay | Must Have | Users can inspect the AI reasoning and evidence for any recommendation. |
| FR-06 | Session Branching | Should Have | Continuing from an earlier point creates a new, independent session by fully duplicating the state up to the branching point, without overwriting the original. |
| FR-07 | Session Management | Must Have | Users can rename, archive, delete, search, and sort their saved sessions. |
| FR-08 | Cross-device Sync | Must Have | Authenticated users see the same sessions and preferences across all supported platforms. |

---

## Success Criteria

| Criterion | Metric | Target |
|---|---|---|
| Session Reuse | Percentage of returning users who resume a session | > 30% |
| Context Rebuild Time | Time spent re-establishing preferences | Reduced by 80% |
| Session Restoration Speed | Time to fully load and render a resumed session | < 1.5 seconds |
| Transparency Engagement | Percentage of recommendations where reasoning is viewed | > 15% |

---

## Assumptions

- Users must be authenticated to benefit from cross-device synchronization. Unauthenticated users may have sessions saved locally (keeping a rolling window of the last 10 sessions, oldest evicted automatically) which sync upon login.
- AI reasoning data can be serialized and stored efficiently alongside conversation history.

## Dependencies

- Existing authentication system.
- Database storage capable of handling large JSON blobs for conversation state and reasoning graphs.
- Decision sessions may expose resumable entry points for future external integrations.

## Risks

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Storage Costs | Medium | Medium | Implement data retention policies for abandoned or very old anonymous sessions; compress state blobs. |
| Sync Conflicts | Low | High | Use simple last-write-wins based on timestamp for concurrent session updates. |
| Privacy Concerns | Low | High | Ensure all stored preferences and chat histories comply with privacy principles; provide clear deletion controls. |

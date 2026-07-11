# Implementation Tasks: Decision Memory

**Feature**: Decision Memory (005-decision-memory)
**Strategy**: MVP-first incremental delivery grouped by user story.

## Phase 1: Foundational Database & Repositories

**Goal**: Establish the persistence layer in PostgreSQL and Go GORM.

- [x] T001 Create database migrations for decision memory entities in `apps/server/db/migrations/20260711_decision_memory.sql`
  - **Dependencies**: None
  - **Acceptance Criteria**: `decision_sessions`, `session_messages`, and `user_preferences` tables exist with correct schemas.
  - **Deliverable**: SQL migration file.

- [x] T002 [P] Implement GORM models in `apps/server/internal/decisionmemory/model.go`
  - **Dependencies**: T001
  - **Acceptance Criteria**: Go structs match DB schema.
  - **Deliverable**: `model.go` with `DecisionSession`, `SessionMessage`, `UserPreference`.

- [x] T003 Implement DecisionMemoryRepository in `apps/server/internal/decisionmemory/repository.go`
  - **Dependencies**: T002
  - **Acceptance Criteria**: CRUD functions exist for sessions, messages, and preferences.
  - **Deliverable**: `repository.go` with interface and Postgres implementation.

---

## Phase 2: Session Management MVP [US1] [US2]

**Goal**: Basic auto-save, list, and resume functionality including anonymous limits and conflict resolution.

- [x] T004 [US1] Implement Session REST handlers (Create, List, Get, Rename, Delete) in `apps/server/internal/decisionmemory/handler.go`
  - **Dependencies**: T003
  - **Acceptance Criteria**: Explicitly covers `GET /api/v1/sessions`, `GET /api/v1/sessions/:id`, `PATCH /api/v1/sessions/:id` for rename, and `DELETE /api/v1/sessions/:id`.
  - **Deliverable**: Gin route handlers.

- [x] T005 [US1] Implement Auto-save (PUT) endpoint with timestamp conflict resolution in `apps/server/internal/decisionmemory/handler.go`
  - **Dependencies**: T004
  - **Acceptance Criteria**: `PUT /api/v1/sessions/:id` rejects requests if `X-Client-Timestamp` is older than DB `updated_at` (409 Conflict).
  - **Deliverable**: PUT handler with conflict logic.

- [x] T006 [P] [US1] Implement Anonymous 10-session eviction limit logic in `apps/server/internal/decisionmemory/service.go`
  - **Dependencies**: T003
  - **Acceptance Criteria**: When saving an anonymous session, if count > 10, the oldest is deleted.
  - **Deliverable**: Service layer eviction logic.

- [x] T007 [P] [US1] Generate `@shopwise/api-types` contracts for sessions in `packages/api-types/src/decision-memory.ts`
  - **Dependencies**: None
  - **Acceptance Criteria**: TypeScript interfaces exist for all session payloads.
  - **Deliverable**: TS type definitions.

- [x] T008 [US2] Build `useSession` and `useSessionList` Query hooks in `apps/web/src/features/decision-memory/api/queries.ts`
  - **Dependencies**: T007
  - **Acceptance Criteria**: TanStack Query hooks successfully fetch from API.
  - **Deliverable**: API integration file.

- [x] T009 [US2] Build `SessionHistorySidebar` component in `apps/web/src/features/decision-memory/components/SessionHistorySidebar.tsx`
  - **Dependencies**: T008
  - **Acceptance Criteria**: Users can search, sort, rename, delete, and archive saved sessions.
  - **Deliverable**: React component.

- [x] T010 [P] [US1] Build `AutoSaveIndicator` component in `apps/web/src/features/decision-memory/components/AutoSaveIndicator.tsx`
  - **Dependencies**: None
  - **Acceptance Criteria**: Subtle toast indicates save state.
  - **Deliverable**: React component.

---

## Phase 3: Preferences & Reasoning [US3]

**Goal**: Extract, persist, and visualize AI reasoning and user preferences.

- [x] T011 [US3] Implement Preference REST handlers (List, Update, Delete) in `apps/server/internal/decisionmemory/preference_handler.go`
  - **Dependencies**: T003
  - **Acceptance Criteria**: Includes editing an existing stored preference and `DELETE /api/v1/preferences/:id` works for Undo.
  - **Deliverable**: Preference handlers.

- [x] T012 [P] [US3] Add Preference API contracts to `packages/api-types/src/decision-memory.ts`
  - **Dependencies**: None
  - **Acceptance Criteria**: TS types for preferences exist.
  - **Deliverable**: Updated TS types.

- [x] T013 [US3] Build `PreferenceUndoToast` component in `apps/web/src/features/decision-memory/components/PreferenceUndoToast.tsx`
  - **Dependencies**: T012
  - **Acceptance Criteria**: Notification allows 1-click preference deletion.
  - **Deliverable**: React component.

- [x] T014 [P] [US3] Build `ReasoningReplayModal` component in `apps/web/src/features/decision-memory/components/ReasoningReplayModal.tsx`
  - **Dependencies**: T007
  - **Acceptance Criteria**: Displays reasoning graph cleanly.
  - **Deliverable**: React component.

---

## Phase 4: Branching & Lifecycle [US4]

**Goal**: Deep-copy branching and long-term archiving logic.

- [x] T015 [US4] Implement full-copy Branch Session handler in `apps/server/internal/decisionmemory/handler.go`
  - **Dependencies**: T004
  - **Acceptance Criteria**: `POST /sessions/:id/branch` creates an independent deep copy storing `parent_session_id` and branch origin/checkpoint metadata.
  - **Deliverable**: Branch endpoint and service logic.

- [x] T016 [US4] Implement Archive Lifecycle Service in `apps/server/internal/decisionmemory/archive_service.go`
  - **Dependencies**: T003
  - **Acceptance Criteria**: Inactive sessions older than 90 days are marked archived (scheduling mechanism agnostic).
  - **Deliverable**: Archive worker code.

- [x] T017 [US4] Implement Session Restore handler in `apps/server/internal/decisionmemory/handler.go`
  - **Dependencies**: T004
  - **Acceptance Criteria**: `POST /sessions/:id/restore` sets status to active.
  - **Deliverable**: Restore endpoint.

- [x] T018 [US4] Build `BranchSessionButton` component in `apps/web/src/features/decision-memory/components/BranchSessionButton.tsx`
  - **Dependencies**: T008
  - **Acceptance Criteria**: Context menu action triggers branching.
  - **Deliverable**: React component.

---

## Phase 5: Testing & Future-Proofing [US5, Tech Debt]

**Goal**: Assure quality and prepare for cross-channel integrations.

- [x] T019 Implement contract tests in `apps/server/internal/decisionmemory/handler_test.go`
  - **Dependencies**: Phase 2, 4
  - **Acceptance Criteria**: Tests prove API payload shapes match specs.
  - **Deliverable**: Go test files.

- [x] T020 Implement integration tests in `apps/server/internal/decisionmemory/repository_test.go`
  - **Dependencies**: Phase 2, 4
  - **Acceptance Criteria**: Tests prove LWW logic and deep copy.
  - **Deliverable**: Go integration tests.

- [x] T021 Implement quickstart validation script in `specs/005-decision-memory/quickstart-runner.sh`
  - **Dependencies**: Phase 2, 4
  - **Acceptance Criteria**: Script runs cURL commands from `quickstart.md` automatically.
  - **Deliverable**: Bash script.

- [x] T022 [P] Expose a channel-agnostic resumable session entrypoint for future external integrations in `apps/server/internal/decisionmemory/external_integration.go`
  - **Dependencies**: T004
  - **Acceptance Criteria**: Generates or resolves a resumable session link/token; contains no Zalo-specific logic; sends no notification; can later be consumed by a separate Resume Conversation feature.
  - **Deliverable**: Go handler stub.

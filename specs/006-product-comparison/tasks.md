# Tasks: Product Comparison Workspace

**Input**: Design documents from `/specs/006-product-comparison/`

**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **Web app**: `backend/src/`, `frontend/src/`
- **AI**: `ai-runtime/src/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [ ] T001 Initialize product comparison workspace project context and structure in `backend/src/api/events/`
- [ ] T002 Initialize product comparison components directory in `frontend/src/components/`
- [ ] T003 [P] Add necessary tool registration skeleton in `ai-runtime/src/tools/fetch_comparison_data.py`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T004 Create `ComparisonWorkspace` model matching `data-model.md` in `backend/src/models/comparison_workspace.go`
- [ ] T005 Create `ComparisonProductDetail` model matching `data-model.md` in `backend/src/models/comparison_product.go`
- [ ] T006 [P] Update Decision Memory service in `backend/src/services/decision_memory.go` to handle `ComparisonWorkspace` persistence
- [ ] T007 Configure error mapping for Decision Memory failures in `backend/src/api/errors.go`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Core Comparison Flow (Priority: P1) 🎯 MVP

**Goal**: Users want to compare up to four products side-by-side to easily identify meaningful differences.

**Independent Test**: Can be tested by triggering the `ADD_TO_COMPARISON` event and verifying the backend returns the expected `ComparisonWorkspace` UI schema, and handling errors for max capacity or category mismatches.

### Implementation for User Story 1

- [ ] T008 [US1] Implement Backend category validation and 4-product limit in `backend/src/services/comparison_service.go`
- [ ] T009 [US1] Implement Backend 5th product handling to trigger `PENDING_REPLACEMENT` in `backend/src/services/comparison_service.go`
- [ ] T010 [US1] Implement Backend event handler for `ADD_TO_COMPARISON` in `backend/src/api/events/add_to_comparison.go`
- [ ] T011 [US1] Implement Backend event handler for `REPLACE_PRODUCT` in `backend/src/api/events/replace_product.go`
- [ ] T012 [US1] Implement Backend event handler for `REMOVE_PRODUCT` in `backend/src/api/events/remove_product.go`
- [ ] T013 [US1] Implement Backend logic to compute `HighlightType` differences in `backend/src/services/comparison_service.go`
- [ ] T014 [P] [US1] Create frontend `ComparisonWorkspace` Dynamic UI renderer in `frontend/src/components/ComparisonWorkspace.tsx`
- [ ] T015 [P] [US1] Create frontend `ErrorState` dialog component for cross-category errors in `frontend/src/components/ErrorState.tsx`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - AI Decision Support (Priority: P1)

**Goal**: Users want the AI to explain the trade-offs between compared products so they do not have to manually interpret technical specifications.

**Independent Test**: Can be tested by simulating an AI follow-up message about the workspace and verifying the AI correctly invokes the `fetch_comparison_data` tool without hallucinating.

### Implementation for User Story 2

- [ ] T016 [US2] Implement backend endpoint to serve Tool Protocol `fetch_comparison_data` requests in `backend/src/api/tools/comparison_data.go`
- [ ] T017 [P] [US2] Implement `fetch_comparison_data` Tool in AI Runtime `ai-runtime/src/tools/fetch_comparison_data.py`
- [ ] T018 [US2] Create system prompt context builder for comparison reasoning in `ai-runtime/src/prompts/comparison_reasoning.py`
- [ ] T019 [US2] Integrate Tool execution and streaming in AI reasoning flow in `ai-runtime/src/graph/workflow.py`

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Proceed to Checkout Readiness (Priority: P2)

**Goal**: Users want to seamlessly move a selected product from the comparison workspace into the checkout readiness state.

**Independent Test**: Trigger the `onProceedToCheckout` event and verify the backend transitions the user's state.

### Implementation for User Story 3

- [ ] T020 [US3] Implement Backend event handler for transitioning product to Checkout Readiness in `backend/src/api/events/checkout_readiness.go`
- [ ] T021 [US3] Update frontend workspace component to dispatch `CHECKOUT_READY` event in `frontend/src/components/ComparisonWorkspace.tsx`

**Checkpoint**: All user stories should now be independently functional

---

## Phase N: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [ ] T022 Run end-to-end `quickstart.md` validations to ensure complete integration
- [ ] T023 Code cleanup and review against the Ultracite Code Standards
- [ ] T024 Security hardening and privacy review (ensuring no unauthorized cross-tenant data access in comparisons)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed sequentially in priority order (P1 → P1 → P2)
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2)
- **User Story 2 (P1)**: Depends on `ComparisonWorkspace` state logic from US1 being implemented on the backend.
- **User Story 3 (P2)**: Depends on frontend components from US1 to fire the checkout event.

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- US1 UI rendering (`T012`, `T013`) can run in parallel with US1 Backend logic
- US2 AI Tool logic (`T015`) can run in parallel with US2 Backend API development (`T014`)

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently using `quickstart.md` validations 1-3.
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently → Deploy/Demo
4. Add User Story 3 → Test independently → Deploy/Demo
5. Each story adds value without breaking previous stories

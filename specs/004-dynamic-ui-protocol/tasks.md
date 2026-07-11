---
description: "Task list template for feature implementation"
---

# Tasks: Dynamic UI Protocol

**Input**: Design documents from `/specs/004-dynamic-ui-protocol/`

**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/

**Tests**: Tests are excluded as they were not explicitly requested.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Create `packages/ui-protocol/src/schema` directory structure for JSON schemas
- [x] T002 [P] Create `services/main-backend/internal/ui-protocol` directory for Go implementation
- [x] T003 [P] Create `apps/web/src/components/dynamic-ui` directory for React implementation
- [x] T004 [P] Create `services/ai-runtime/src/ui_protocol` directory for Python implementation

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T005 Create base JSON Schema for `DynamicUIDocument` and `ComponentNode` in `packages/ui-protocol/src/schema/base-schema.json`
- [x] T006 Create base JSON Schema for `InteractionEvent` in `packages/ui-protocol/src/schema/event-schema.json`
- [x] T007 [P] Create JSON Schemas for Layout Components (Page, Section, Card, Grid, Tabs, Modal, Drawer, Timeline) in `packages/ui-protocol/src/schema/layout-components.json`
- [x] T008 [P] Implement base TypeScript types generated from schemas in `packages/ui-protocol/src/types/index.ts`
- [x] T009 [P] Implement base Go structs for `DynamicUIDocument` and `InteractionEvent` in `services/main-backend/internal/ui-protocol/base.go`
- [x] T010 [P] Implement base Python Pydantic models for `DynamicUIDocument` and `InteractionEvent` in `services/ai-runtime/src/ui_protocol/base.py`
- [x] T011 Setup AJV validation configuration for the frontend React components in `apps/web/src/components/dynamic-ui/validation.ts`
- [x] T012 Implement VND Currency Standard validation logic (Principle XVI) across schemas, Go structs, and Pydantic models to ensure all price fields enforce VND formatting

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Structured Product Recommendation (Priority: P1) 🎯 MVP

**Goal**: Deliver structured interactive UI cards (Product Card, Carousel) directly from AI responses without frontend code generation.

**Independent Test**: AI Runtime generates a valid JSON payload containing Product Carousel and Card components which passes validation.

### Implementation for User Story 1

- [x] T013 [P] [US1] Create JSON Schemas for AI Components (Chat Message, Thinking Indicator, Recommendation Explanation, Decision Reasoning, Confidence Indicator) in `packages/ui-protocol/src/schema/ai-components.json`
- [x] T014 [P] [US1] Create JSON Schemas for Shopping Components (Product Card, Product Carousel, Product Comparison, Specification Table, Promotion Banner, Warranty Information, Inventory Status, Checkout Summary) in `packages/ui-protocol/src/schema/commerce-components.json`
- [x] T015 [P] [US1] Create JSON Schemas for Commerce Actions (Add to Comparison, Remove Product, Save Decision, Resume Conversation, Checkout Readiness) in `packages/ui-protocol/src/schema/actions.json`
- [x] T016 [P] [US1] Generate TypeScript types for AI and Commerce components in `packages/ui-protocol/src/types/commerce.ts`
- [x] T017 [P] [US1] Implement Go structs for AI and Commerce components in `services/main-backend/internal/ui-protocol/commerce.go`
- [x] T018 [P] [US1] Implement Pydantic models for AI and Commerce components in `services/ai-runtime/src/ui_protocol/commerce.py`
- [x] T019 [US1] Implement React components for AI UI (Chat Message, Thinking Indicator, etc.) in `apps/web/src/components/dynamic-ui/ai/`
- [x] T020 [US1] Implement React components for Shopping UI (Product Card, Carousel, Comparison Table, etc.) ensuring VND currency displays correctly in `apps/web/src/components/dynamic-ui/commerce/`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Dynamic Form Collection (Priority: P2)

**Goal**: Allow the AI to present interactive forms (Budget Sliders, Store Pickers) to gather user requirements structurally.

**Independent Test**: Frontend correctly renders budget sliders, validates input against the schema, and emits a structured InteractionEvent to the backend.

### Implementation for User Story 2

- [x] T021 [P] [US2] Create JSON Schemas for Interaction Components (Button, Quick Reply, Select, Radio Group, Checkbox, Text Input, Number Input, Budget Slider, Store Picker) in `packages/ui-protocol/src/schema/interaction-components.json`
- [x] T022 [P] [US2] Generate TypeScript types for Interaction components in `packages/ui-protocol/src/types/interaction.ts`
- [x] T023 [P] [US2] Implement Go structs for Interaction components in `services/main-backend/internal/ui-protocol/interaction.go`
- [x] T024 [P] [US2] Implement Pydantic models for Interaction components in `services/ai-runtime/src/ui_protocol/interaction.py`
- [x] T025 [US2] Implement React component renderers for all Interaction Components (Input, Sliders, Pickers) in `apps/web/src/components/dynamic-ui/interaction/`
- [x] T026 [US2] Integrate interaction components with the base event dispatcher in `apps/web/src/components/dynamic-ui/EventDispatcher.ts`

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Error Recovery and Fallbacks (Priority: P3)

**Goal**: Provide structured Error components with actionable recovery options (e.g., Retry) for graceful degradation.

**Independent Test**: A simulated backend failure generates a structured Error Component which the frontend renders and emits a Retry event.

### Implementation for User Story 3

- [x] T027 [P] [US3] Create JSON Schemas for Error Components (Error State, Retry, Continue with Cached Data, Modify Search, Alternative Recommendation) in `packages/ui-protocol/src/schema/error-components.json`
- [x] T028 [P] [US3] Generate TypeScript types for Error components in `packages/ui-protocol/src/types/error.ts`
- [x] T029 [P] [US3] Implement Go structs for Error components in `services/main-backend/internal/ui-protocol/error.go`
- [x] T030 [P] [US3] Implement Pydantic models for Error components in `services/ai-runtime/src/ui_protocol/error.py`
- [x] T031 [US3] Implement React component renderers for Error states and fallback actions in `apps/web/src/components/dynamic-ui/error/`
- [x] T032 [US3] Implement fallback UI handler for unsupported schema versions in `apps/web/src/components/dynamic-ui/ProtocolRenderer.tsx`

**Checkpoint**: All user stories should now be independently functional

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T033 [P] Run quickstart validation scenarios defined in `quickstart.md`
- [x] T034 [P] Implement continuous integration (CI) script to validate schemas against payloads
- [x] T035 Code cleanup and ensuring uniform event names across TypeScript, Go, and Python
- [x] T036 Verify backward compatibility strategy across all three languages

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can proceed sequentially or in parallel depending on staffing.
- **Polish (Final Phase)**: Depends on all user stories being complete.

### User Story Dependencies

- **User Story 1 (P1)**: Independent, but depends on Foundational phase.
- **User Story 2 (P2)**: Independent, but depends on Foundational phase.
- **User Story 3 (P3)**: Independent, but depends on Foundational phase.

### Parallel Opportunities

- All languages (Go, Python, TypeScript) can have their types/structs/models implemented in parallel for each phase.
- JSON Schema definitions can be parallelized from backend integration.

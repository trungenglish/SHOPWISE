# Tasks: AI Runtime LLM Integration

**Input**: Design documents from `/specs/008-llm-integration/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Test tasks are included per standard Python project structure (using `pytest`).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- All paths are relative to `services/ai-runtime/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Initialize Python project and dependencies (FastAPI, LangGraph, OpenAI, Pydantic, sse-starlette)
- [x] T002 [P] Configure linting and formatting (ruff, mypy)
- [x] T003 Create project structure (src/, tests/, etc.) in `services/ai-runtime/`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [x] T004 Create core configuration model (Provider, API Key, Base URL) in `services/ai-runtime/src/core/config.py`
- [x] T005 [P] Setup environment variable loading using Pydantic BaseSettings
- [x] T006 Configure error handling, structured logging, and observability (FR-009, no PII) in `services/ai-runtime/src/core/observability.py`
- [x] T007 Define core Provider Interface base class in `services/ai-runtime/src/llm/provider.py`
- [x] T008 Define structured AI Error State model in `services/ai-runtime/src/models/errors.py`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Secure LLM Invocation (Priority: P1) 🎯 MVP

**Goal**: The AI Runtime securely invokes an external LLM to generate conversational responses without exposing API keys to the frontend or Go Backend.

**Independent Test**: Can be fully tested by sending a simulated chat request to the AI Runtime and verifying a valid response is returned while API keys remain unexposed in logs or network inspection.

### Tests for User Story 1 (OPTIONAL - only if tests requested) ⚠️

- [x] T009 [P] [US1] Unit test for OpenAI Provider integration in `services/ai-runtime/tests/unit/test_openai_provider.py`
- [x] T010 [P] [US1] Integration test for base chat request in `services/ai-runtime/tests/integration/test_chat_endpoint.py`

### Implementation for User Story 1

- [x] T011 [P] [US1] Implement OpenAI provider wrapping `openai` SDK in `services/ai-runtime/src/llm/openai_provider.py`
- [x] T012 [P] [US1] Implement exponential backoff for transient errors in `services/ai-runtime/src/llm/openai_provider.py`
- [x] T013 [P] [US1] Implement Prompt Management module (composition, templates, context assembly, versioning) in `services/ai-runtime/src/core/prompts.py`
- [x] T014 [P] [US1] Implement Runtime Model Configuration parsing and injection from request context in `services/ai-runtime/src/api/chat.py`
- [x] T015 [US1] Create basic FastAPI router and Chat endpoint in `services/ai-runtime/src/api/chat.py`
- [x] T016 [US1] Connect router to main FastAPI app in `services/ai-runtime/src/main.py`
- [x] T017 [US1] Ensure logs emit only metadata and no message content (PII sanitation)

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Streaming Responses (Priority: P1)

**Goal**: The AI Runtime streams responses from the LLM back to the client for lower perceived latency and progressive dynamic UI updates.

**Independent Test**: Can be fully tested by monitoring the connection to the AI Runtime and receiving partial chunks of a message.

### Tests for User Story 2 (OPTIONAL - only if tests requested) ⚠️

- [x] T018 [P] [US2] Unit test for SSE streaming generator in `services/ai-runtime/tests/unit/test_streaming.py`

### Implementation for User Story 2

- [x] T019 [P] [US2] Implement Server-Sent Events (SSE) generator using `sse-starlette` in `services/ai-runtime/src/api/streaming.py`
- [x] T020 [US2] Update OpenAI provider to support streaming chunks in `services/ai-runtime/src/llm/openai_provider.py`
- [x] T021 [US2] Implement FastAPI endpoint to serve the SSE stream in `services/ai-runtime/src/api/chat.py`
- [x] T022 [US2] Verify TTFT (Time to First Token) meets < 1000ms SLA target

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Structured Output and Tool Calling (Priority: P2)

**Goal**: The AI Runtime forces the LLM to return structured JSON and can invoke tools based on LLM outputs to interact with the Retail Backend.

**Independent Test**: Can be fully tested by prompting the AI Runtime with a request that triggers a specific tool call and verifying the correct tool interface is invoked.

### Tests for User Story 3 (OPTIONAL - only if tests requested) ⚠️

- [x] T023 [P] [US3] Unit test for LangGraph tool orchestration in `services/ai-runtime/tests/unit/test_graph.py`
- [x] T024 [P] [US3] Unit test for JSON parsing and self-correction loop in `services/ai-runtime/tests/unit/test_self_correction.py`

### Implementation for User Story 3

- [x] T025 [P] [US3] Define Pydantic models for structured output in `services/ai-runtime/src/models/schemas.py`
- [x] T026 [P] [US3] Implement Tool SDK client/proxy to invoke existing Go Backend Tool APIs in `services/ai-runtime/src/tools/interfaces.py`
- [x] T027 [US3] Implement LangGraph workflow to orchestrate LLM calls and tool execution in `services/ai-runtime/src/graph/workflow.py`
- [x] T028 [US3] Implement bounded self-correction retry loop (max 2 attempts) for invalid JSON in `services/ai-runtime/src/graph/workflow.py`
- [x] T029 [US3] Wire LangGraph workflow into the main FastAPI chat/stream endpoints

**Checkpoint**: All user stories should now be independently functional

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T030 [P] Documentation updates in `services/ai-runtime/README.md`
- [x] T031 Code cleanup and refactoring
- [x] T032 Setup OpenTelemetry tracing for the LangGraph execution flow in `services/ai-runtime/src/core/observability.py`
- [x] T033 Run quickstart.md validation

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3+)**: All depend on Foundational phase completion
  - User stories can then proceed in parallel (if staffed)
  - Or sequentially in priority order (P1 → P2 → P3)
- **Polish (Final Phase)**: Depends on all desired user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational (Phase 2) - No dependencies on other stories
- **User Story 2 (P1)**: Can start after Foundational (Phase 2) - May integrate with US1 but should be independently testable
- **User Story 3 (P2)**: Can start after Foundational (Phase 2) - May integrate with US1/US2 but should be independently testable

### Within Each User Story

- Tests (if included) MUST be written and FAIL before implementation
- Models before services
- Services before endpoints
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- All tests for a user story marked [P] can run in parallel
- Models within a story marked [P] can run in parallel
- Different user stories can be worked on in parallel by different team members

---

## Parallel Example: User Story 3

```bash
# Launch all tests for User Story 3 together:
Task: "Unit test for LangGraph tool orchestration in services/ai-runtime/tests/unit/test_graph.py"
Task: "Unit test for JSON parsing and self-correction loop in services/ai-runtime/tests/unit/test_self_correction.py"

# Launch models/interfaces for User Story 3 together:
Task: "Define Pydantic models for structured output in services/ai-runtime/src/models/schemas.py"
Task: "Define tool interfaces/proxies in services/ai-runtime/src/tools/interfaces.py"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently → Deploy/Demo
4. Add User Story 3 → Test independently → Deploy/Demo
5. Each story adds value without breaking previous stories

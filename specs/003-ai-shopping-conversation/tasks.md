# Tasks: AI Shopping Conversation MVP

**Input**: Design documents from `/specs/003-ai-shopping-conversation/`
**Prerequisites**: plan.md, spec.md, data-model.md, contracts/ui-contracts.md

**Organization**: Tasks are grouped by the implementation milestones requested, while retaining traceability to user stories [US1..US4].

---

## Milestone 1 — Project Foundation

- [x] T001 [US1] Initialize repository structure in root
  - **Title**: Initialize repository structure
  - **Description**: Set up the basic directory structure for Go Backend, Python AI Runtime, and React Frontend.
  - **Dependencies**: None
  - **Acceptance Criteria**: Directories `apps/web`, `services/main-backend`, and `services/ai-runtime` exist.
  - **Expected Deliverables**: Folder structure initialized.

- [x] T002 [P] [US1] Configure development environment in root
  - **Title**: Configure development environment
  - **Description**: Configure pnpm workspace, Go modules, and Python `uv` virtual environments.
  - **Dependencies**: T001
  - **Acceptance Criteria**: `pnpm install`, `go mod download`, and `uv sync` succeed.
  - **Expected Deliverables**: `pnpm-workspace.yaml`, `go.mod`, `pyproject.toml`.

- [x] T003 [P] [US1] Setup shared schemas and contracts in `specs/003-ai-shopping-conversation/contracts/`
  - **Title**: Setup shared schemas and contracts
  - **Description**: Ensure JSON schemas for UI and Data Models are accessible for code generation or validation.
  - **Dependencies**: None
  - **Acceptance Criteria**: Types can be generated or referenced by all three languages.
  - **Expected Deliverables**: Centralized contract definitions.

---

## Milestone 2 — Retail Backend (Go)

- [x] T004 [US1] Implement Conversation API in `services/main-backend/api/conversation.go`
  - **Title**: Conversation API
  - **Description**: Expose HTTP endpoints for the frontend to start and continue a chat session.
  - **Dependencies**: T001
  - **Acceptance Criteria**: Frontend can send a message and receive a generic response.
  - **Expected Deliverables**: `POST /api/chat` endpoint.

- [x] T005 [P] [US2] Implement Product Catalog API in `services/main-backend/api/catalog.go`
  - **Title**: Product Catalog API
  - **Description**: Endpoint exposing product search capabilities (mocked).
  - **Dependencies**: T001
  - **Acceptance Criteria**: Tools can query for "Laptop" and get products in VND.
  - **Expected Deliverables**: `catalog.search` tool endpoint.

- [x] T006 [P] [US3] Implement Product Comparison API in `services/main-backend/api/compare.go`
  - **Title**: Product Comparison API
  - **Description**: Endpoint to retrieve specs for multiple items for comparison.
  - **Dependencies**: T005
  - **Acceptance Criteria**: Given two IDs, returns unified specs.
  - **Expected Deliverables**: `catalog.compare` tool endpoint.

- [x] T007 [US1] Implement Session Management in `services/main-backend/session/manager.go`
  - **Title**: Session Management
  - **Description**: In-memory or Redis-based session storage for short-term memory.
  - **Dependencies**: T004
  - **Acceptance Criteria**: Context persists across multiple turns of the same sessionId.
  - **Expected Deliverables**: Session store implementation.

- [x] T008 [US2] Implement Tool Registry in `services/main-backend/tools/registry.go`
  - **Title**: Tool Registry
  - **Description**: Central registry routing tool calls from the AI Runtime to the correct Go handler.
  - **Dependencies**: T005, T006
  - **Acceptance Criteria**: AI Runtime can dynamically discover and invoke tools.
  - **Expected Deliverables**: Tool routing middleware.

- [x] T009 [US2] Implement Mock Retail Provider in `services/main-backend/retail/mock.go`
  - **Title**: Mock Retail Provider
  - **Description**: Hardcoded JSON data provider for products, injecting VND formatting.
  - **Dependencies**: None
  - **Acceptance Criteria**: Returns stable mocked data for MVP testing.
  - **Expected Deliverables**: Mock data service.

- [x] T010 [US4] Implement Checkout Preparation API in `services/main-backend/api/checkout.go`
  - **Title**: Checkout Preparation API
  - **Description**: Expose the `checkout.prepare` tool endpoint.
  - **Dependencies**: T008
  - **Acceptance Criteria**: Generates a checkout ID and validates stock.
  - **Expected Deliverables**: `checkout.prepare` endpoint.

---

## Milestone 3 — AI Runtime (Python)

- [x] T011 [US1] Setup FastAPI server in `services/ai-runtime/main.py`
  - **Title**: FastAPI server
  - **Description**: Basic FastAPI app setup to receive requests from Go.
  - **Dependencies**: T002
  - **Acceptance Criteria**: Server runs on port 8000 and answers health checks.
  - **Expected Deliverables**: FastAPI application.

- [x] T012 [US1] Define LangGraph workflow in `services/ai-runtime/graph/workflow.py`
  - **Title**: LangGraph workflow
  - **Description**: Define the state graph for the agent (Plan -> Reason -> Act -> Respond).
  - **Dependencies**: T011
  - **Acceptance Criteria**: An empty graph compiles and executes.
  - **Expected Deliverables**: `StateGraph` definition.

- [x] T013 [US1] Implement Planner node in `services/ai-runtime/nodes/planner.py`
  - **Title**: Planner node
  - **Description**: Node that analyzes user input and determines if tools or clarification are needed.
  - **Dependencies**: T012
  - **Acceptance Criteria**: Successfully categorizes intent.
  - **Expected Deliverables**: Planner logic.

- [x] T014 [US1] Implement Reasoner node in `services/ai-runtime/nodes/reasoner.py`
  - **Title**: Reasoner node
  - **Description**: Node that formulates the final conversational response and trade-offs.
  - **Dependencies**: T012
  - **Acceptance Criteria**: Generates explainable rationale.
  - **Expected Deliverables**: Reasoner logic.

- [x] T015 [US2] Implement Tool Calling logic in `services/ai-runtime/tools/caller.py`
  - **Title**: Tool Calling
  - **Description**: HTTP client that formats and executes tool calls against the Go backend.
  - **Dependencies**: T008, T012
  - **Acceptance Criteria**: Successfully sends and receives tool responses.
  - **Expected Deliverables**: Tool execution node.

- [x] T016 [US2] Implement Recommendation Engine in `services/ai-runtime/engine/recommendation.py`
  - **Title**: Recommendation Engine
  - **Description**: Formats the retrieved product data into recommendations.
  - **Dependencies**: T015
  - **Acceptance Criteria**: Maps raw catalog data to the recommendation entity.
  - **Expected Deliverables**: Recommendation mapper.

- [x] T017 [P] [US3] Implement Product Comparison in `services/ai-runtime/engine/comparison.py`
  - **Title**: Product Comparison
  - **Description**: Logic to highlight differences and formulate comparison trade-offs.
  - **Dependencies**: T015
  - **Acceptance Criteria**: Extracts differences accurately.
  - **Expected Deliverables**: Comparison mapper.

- [x] T018 [US2] Implement Dynamic UI Schema Generator in `services/ai-runtime/ui/schema_builder.py`
  - **Title**: Dynamic UI Schema Generator
  - **Description**: Serializes AI output into the deterministic UI schemas (`carousel`, `comparison_table`, etc.).
  - **Dependencies**: T016, T017
  - **Acceptance Criteria**: Outputs valid JSON schemas per contracts.
  - **Expected Deliverables**: Schema builders.

---

## Milestone 4 — Frontend (React + Vite)

- [x] T019 [US1] Implement Chat Workspace in `apps/web/src/components/ChatWorkspace.tsx`
  - **Title**: Chat Workspace
  - **Description**: Main UI for displaying conversation history and text input.
  - **Dependencies**: T002
  - **Acceptance Criteria**: User can see messages and type text.
  - **Expected Deliverables**: Chat UI component.

- [x] T020 [US2] Implement Dynamic UI Renderer in `apps/web/src/components/DynamicUIRenderer.tsx`
  - **Title**: Dynamic UI Renderer
  - **Description**: Factory component that reads UI schema `type` and renders the correct sub-component.
  - **Dependencies**: T019
  - **Acceptance Criteria**: Parses JSON and mounts the correct React component.
  - **Expected Deliverables**: UI Renderer switch statement.

- [x] T021 [P] [US2] Implement Recommendation Cards in `apps/web/src/components/ui/ProductCarousel.tsx`
  - **Title**: Recommendation Cards
  - **Description**: Component mapping to the `carousel` schema.
  - **Dependencies**: T020
  - **Acceptance Criteria**: Renders product images, VND prices, and AI insights.
  - **Expected Deliverables**: Carousel component.

- [x] T022 [P] [US3] Implement Comparison Workspace in `apps/web/src/components/ui/ComparisonTable.tsx`
  - **Title**: Comparison Workspace
  - **Description**: Component mapping to the `comparison_table` schema.
  - **Dependencies**: T020
  - **Acceptance Criteria**: Renders side-by-side feature comparison.
  - **Expected Deliverables**: Comparison table component.

- [x] T023 [US4] Implement Checkout Readiness Panel in `apps/web/src/components/ui/CheckoutSummary.tsx`
  - **Title**: Checkout Readiness Panel
  - **Description**: Component mapping to the `checkout_summary` schema, including explicit confirmation button.
  - **Dependencies**: T020
  - **Acceptance Criteria**: Displays total in VND and requires click to confirm.
  - **Expected Deliverables**: Checkout summary component.

- [x] T024 [P] [US2] Implement Error State Components in `apps/web/src/components/ui/ErrorState.tsx`
  - **Title**: Error State Components
  - **Description**: Component mapping to the `error_state` schema for graceful recovery.
  - **Dependencies**: T020
  - **Acceptance Criteria**: Renders error messages with actionable retry buttons.
  - **Expected Deliverables**: Error state component.

---

## Milestone 5 — Integration

- [x] T025 [US1] Configure Go ↔ AI Runtime HTTP communication in `services/main-backend/client/ai_client.go`
  - **Title**: Go ↔ AI Runtime communication
  - **Description**: Wire the Go backend to forward user requests to the AI Runtime.
  - **Dependencies**: T004, T011
  - **Acceptance Criteria**: Go successfully acts as an API gateway for the AI.
  - **Expected Deliverables**: HTTP client wrapper.

- [x] T026 [US2] Integrate Tool execution flow between AI Runtime and Go Backend
  - **Title**: Tool execution
  - **Description**: End-to-end wiring of the AI invoking a Go tool and Go returning data.
  - **Dependencies**: T008, T015
  - **Acceptance Criteria**: Round-trip tool execution works.
  - **Expected Deliverables**: Integration wiring.

- [x] T027 [US1] Implement Server-Sent Events (SSE) streaming responses in Frontend and Go
  - **Title**: Streaming responses
  - **Description**: Stream partial chunks (text + UI schema) to the React frontend.
  - **Dependencies**: T025, T019
  - **Acceptance Criteria**: Text types out smoothly; UI components pop in instantly when schemas arrive.
  - **Expected Deliverables**: SSE integration.

- [x] T028 [US1] Implement Conversation state synchronization across all 3 tiers
  - **Title**: Conversation state synchronization
  - **Description**: Ensure the session ID binds the React, Go, and Python contexts together.
  - **Dependencies**: T007
  - **Acceptance Criteria**: State is shared correctly across multiple conversational turns.
  - **Expected Deliverables**: Session integration.

---

## Milestone 6 — Demo Scenarios

- [x] T029 [US2] End-to-end test: Laptop recommendation scenario
  - **Title**: Laptop recommendation
  - **Description**: Validate the "I need a gaming laptop under 20 million VND" flow.
  - **Dependencies**: Milestone 5
  - **Acceptance Criteria**: Successfully returns a Carousel schema with VND prices.
  - **Expected Deliverables**: Verified manual scenario.

- [x] T030 [US3] End-to-end test: Product comparison scenario
  - **Title**: Product comparison
  - **Description**: Validate the "Compare the first two laptops" flow.
  - **Dependencies**: Milestone 5
  - **Acceptance Criteria**: Successfully returns a Comparison schema.
  - **Expected Deliverables**: Verified manual scenario.

- [x] T031 [US1] End-to-end test: Requirement clarification scenario
  - **Title**: Requirement clarification
  - **Description**: Validate vague requirement handling.
  - **Dependencies**: Milestone 5
  - **Acceptance Criteria**: AI asks clarifying questions instead of hallucinating.
  - **Expected Deliverables**: Verified manual scenario.

- [x] T032 [US2] End-to-end test: Tool timeout recovery scenario
  - **Title**: Tool timeout recovery
  - **Description**: Validate behavior when Mock Catalog fails.
  - **Dependencies**: Milestone 5
  - **Acceptance Criteria**: AI renders Error State UI gracefully.
  - **Expected Deliverables**: Verified manual scenario.

- [x] T033 [US4] End-to-end test: Out-of-stock recovery scenario
  - **Title**: Out-of-stock recovery
  - **Description**: Validate checkout when an item is OOS.
  - **Dependencies**: Milestone 5
  - **Acceptance Criteria**: AI presents alternatives and does not auto-replace.
  - **Expected Deliverables**: Verified manual scenario.

- [x] T034 [US4] End-to-end test: Checkout readiness scenario
  - **Title**: Checkout readiness
  - **Description**: Validate final checkout confirmation.
  - **Dependencies**: Milestone 5
  - **Acceptance Criteria**: Requires explicit user confirmation button click.
  - **Expected Deliverables**: Verified manual scenario.

---

## Dependencies & Execution Order
1. **Milestone 1** (Foundation) must be completed first.
2. **Milestones 2, 3, and 4** can be executed in parallel by different developers.
3. **Milestone 5** integrates the three tiers and unblocks Milestone 6.
4. **Milestone 6** validates the success criteria.

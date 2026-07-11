# Feature Specification: Dynamic UI Protocol

**Feature Branch**: `004-dynamic-ui-protocol`

**Created**: 2026-07-11

**Status**: Draft

**Input**: User description: "Define a declarative UI protocol that allows the AI Runtime to dynamically describe interactive user interfaces without generating frontend code..."

## Clarifications

### Session 2026-07-11

- Q: Should the Dynamic UI Protocol define the specific transport mechanism for streaming and bidirectional communication, or should it remain strictly transport-agnostic? → A: Transport-agnostic. It defines only the structure, semantics, and lifecycle of UI schemas and interaction events. Transport mechanisms (e.g., SSE, WebSockets, HTTP, gRPC) are implementation concerns and may vary by platform without affecting protocol compatibility.
- Q: How should schema versioning be handled in the protocol payloads? → A: Defined once at the root payload level. Each document MUST include a root-level schemaVersion field (e.g., "1.0") using Semantic Versioning. Minor versions are backward compatible; Major versions require compatibility checks; unsupported versions SHOULD render a graceful fallback UI.
- Q: Should the protocol differentiate between AI-handled events and direct Backend Tool events? → A: No, routing is opaque to the frontend. Components declare semantic actions only and do not specify execution targets. The backend acts as the single orchestration layer for routing.
- Q: Should form validation rules be defined declaratively within the UI Protocol schema, or should the frontend submit the raw state and let the backend/AI handle validation and return error states? → A: Hybrid approach. The UI Protocol should support declarative validation rules for immediate client-side feedback (e.g., required fields, data types). However, the backend remains the source of truth for all business validation. Backend validation failures MUST be returned as structured Error State UI schemas.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Structured Product Recommendation (Priority: P1)

As a user chatting with the AI Sales Agent, I want to receive product recommendations as rich, interactive UI cards rather than just text, so that I can easily view specifications, prices (in VND), and compare options visually.

**Why this priority**: Displaying structured product data is the core value proposition of an AI sales agent over a generic text chatbot. It's essential for guiding the shopping journey.

**Independent Test**: Can be fully tested by sending a recommendation query and verifying the AI Runtime returns a valid JSON UI schema containing Product Card components that the frontend successfully renders without needing HTML/React code from the AI.

**Acceptance Scenarios**:

1. **Given** a user asks for laptop recommendations, **When** the AI responds, **Then** it returns a JSON payload containing Product Carousel and Product Card components.
2. **Given** a rendered Product Card, **When** the user clicks "Add to Comparison", **Then** the frontend emits a declarative interaction event with the component ID and payload to the backend.

---

### User Story 2 - Dynamic Form Collection (Priority: P2)

As a user looking for a specific type of device, I want the AI to present interactive forms (budget sliders, selections, checkboxes) to gather my requirements quickly instead of asking me open-ended text questions.

**Why this priority**: Requirement gathering is a critical step in retail commerce. Interactive forms reduce friction and improve conversion over free-text chat.

**Independent Test**: Can be tested by triggering a requirement collection flow, verifying the AI returns form components (Slider, Checkbox, Select) in JSON, and ensuring the frontend can emit the collected state back to the backend.

**Acceptance Scenarios**:

1. **Given** an incomplete user profile for a query, **When** the AI needs more details, **Then** it returns a JSON schema with a Budget Slider and Store Picker components.
2. **Given** an interactive form in the UI, **When** the user selects their budget and submits, **Then** the frontend sends an interaction event with the structured payload, and the AI acknowledges with an updated UI block.

---

### User Story 3 - Error Recovery and Fallbacks (Priority: P3)

As a user experiencing a network or backend issue during my shopping journey, I want to see clear error states with actionable recovery options (e.g., "Retry", "Continue with Cached Data") instead of a broken interface.

**Why this priority**: Graceful degradation is crucial for e-commerce. A broken UI can result in lost sales, whereas actionable error components can salvage the session.

**Independent Test**: Can be tested by simulating a backend failure and verifying that the frontend receives or displays Error Components defined by the protocol, allowing the user to trigger a Retry action.

**Acceptance Scenarios**:

1. **Given** a failure while fetching product inventory, **When** the backend times out, **Then** the AI (or backend) returns an Error Component schema with a "Retry" button.
2. **Given** an Error Component, **When** the user clicks "Retry", **Then** the frontend emits a retry event, and the UI dynamically updates to a Loading State.

### Edge Cases

- What happens if the AI Runtime returns an invalid JSON schema? (Frontend ignores and displays generic fallback error UI)
- How does the system handle unsupported or future schema version components? (Frontend ignores unknown components and renders the rest of the layout)
- What happens if a backend tool interaction takes too long? (Frontend displays a Loading State component until timeout, then switches to Error Component)

## Requirements *(mandatory)*

> **Note**: As per Principle I — Specification First, any interactions between components must be documented in separate API, UI, Tool, or Data Contracts.

### Functional Requirements

- **FR-001**: The protocol MUST define a JSON-serializable schema for Layout components including Page, Section, Card, Grid, Tabs, Modal, Drawer, and Timeline.
- **FR-002**: The protocol MUST define schemas for Shopping Components including Product Card, Product Carousel, Product Comparison, Specification Table, Promotion Banner, Warranty Information, Inventory Status, and Checkout Summary.
- **FR-003**: The protocol MUST define schemas for AI Components including Chat Message, Thinking Indicator, Recommendation Explanation, Decision Reasoning, and Confidence Indicator.
- **FR-004**: The protocol MUST define schemas for Interaction Components including Button, Quick Reply, Select, Radio Group, Checkbox, Text Input, Number Input, Budget Slider, and Store Picker.
- **FR-005**: The protocol MUST define schemas for Commerce Components including Add to Comparison, Remove Product, Save Decision, Resume Conversation, and Checkout Readiness.
- **FR-006**: The protocol MUST define schemas for Error Components including Retry, Continue with Cached Data, Modify Search, and Alternative Recommendation.
- **FR-007**: Every interactive component schema MUST declare fields for: Component ID, Action, Payload, Event, and State.
- **FR-008**: The protocol MUST allow the AI Runtime to specify whether a UI block should replace an existing block or be appended to the current conversation.
- **FR-009**: The protocol MUST clearly distinguish between Conversation State, UI State, Component State, Loading State, and Error State.
- **FR-010**: Components MUST declare semantic actions only; they MUST NOT specify execution targets. The frontend MUST NOT know whether an event is handled by the AI Runtime or a backend service. The backend acts as the single orchestration layer to route events.
- **FR-011**: The protocol MUST support declarative validation rules (e.g., required fields, data types, ranges, string length, patterns) for immediate client-side feedback. The backend remains the authoritative source for business validation, and any backend validation failures MUST be returned as structured Error State UI schemas.

### Non-Functional Requirements
- **NFR-001**: The schema MUST NOT contain any HTML, CSS, React Components, JavaScript, or Tailwind classes.
- **NFR-002**: The protocol MUST remain frontend-framework independent.
- **NFR-003**: The protocol MUST support streaming updates, partial UI replacement, and incremental rendering.
- **NFR-004**: The protocol MUST be strictly transport-agnostic. It defines only the structure, semantics, and lifecycle of UI schemas and interaction events. Transport mechanisms (e.g., SSE, WebSockets, HTTP, gRPC) are implementation concerns and may vary by platform without affecting protocol compatibility.
- **NFR-005**: The protocol MUST support schema versioning and backward compatibility. Each Dynamic UI document MUST include a root-level schemaVersion field (e.g., "1.0"). Versioning MUST follow Semantic Versioning (MAJOR.MINOR). Minor versions (1.0 → 1.1) may add optional fields or new component types while remaining backward compatible. Major versions (1.x → 2.0) may introduce breaking changes and require frontend compatibility checks before rendering. Frontends encountering an unsupported schemaVersion SHOULD render a graceful fallback UI rather than failing completely.
- **NFR-006**: The protocol MUST be extensible to allow adding new components in the future.

### Key Entities

- **UI Schema / Component Node**: The fundamental building block of the protocol, representing a single UI element (e.g., a Card, a Button). Contains type information, properties, children nodes, and interaction metadata.
- **Interaction Event**: The payload emitted by the frontend when a user interacts with a component. Contains Component ID, Action type, and the associated data Payload.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The protocol can express an entire, end-to-end shopping journey (from discovery to checkout readiness) using structured JSON only, without any text-only fallbacks.
- **SC-002**: The identical JSON UI schema can be successfully rendered by at least two different frontend implementations (e.g., React and a CLI/test renderer) without requiring modifications to the payload.
- **SC-003**: 100% of user interactions (clicks, inputs, selections) defined in the protocol are represented as declarative events emitted to the backend, rather than executing frontend-specific business logic.
- **SC-004**: AI generated schema payloads are strictly validated against the JSON schema definitions with a 0% failure rate for invalid schema structure (e.g. no embedded HTML).

## Assumptions

- The frontend application has a component library capable of interpreting and rendering the declarative schema components.
- The backend infrastructure is capable of routing interaction events from the frontend to the appropriate tools or the AI Runtime.
- The AI Runtime has the necessary context to decide which UI components are appropriate for the current state of the conversation.

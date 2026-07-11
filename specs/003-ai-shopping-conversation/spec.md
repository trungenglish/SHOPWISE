# Feature Specification: AI Shopping Conversation

**Feature Branch**: `003-ai-shopping-conversation`

**Created**: 2026-07-11

**Status**: Draft

**Input**: User description: "Feature: AI Shopping Conversation Goal:Create a complete specification for the conversational shopping experience of SHOPWISE..."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Initial Requirement Gathering (Priority: P1)

Customer starts a conversation without knowing exactly what they want. AI prompts using a hybrid onboarding flow (quick replies + free-form text) to determine intent.

**Why this priority**: Essential entry point to the shopping experience.

**Independent Test**: User can open the interface, see initial quick-reply suggestions, and start a conversation successfully.

**Acceptance Scenarios**:

1. **Given** a new shopping session, **When** the user opens the chat, **Then** the AI greets them naturally and presents optional quick-reply suggestions (e.g., "Laptop", "Gaming").
2. **Given** the chat is open, **When** the user types a free-form request, **Then** the AI extracts intent and follows up to gather missing details (budget, purpose) without presenting a rigid questionnaire.

---

### User Story 2 - Product Search and Recommendation (Priority: P1)

Customer provides their shopping requirements. AI searches the catalog and returns a product carousel with options along with explainable rationale.

**Why this priority**: Core value proposition of the conversational shopping feature.

**Independent Test**: User asks for a laptop under $1000, AI responds with a UI carousel of matching products.

**Acceptance Scenarios**:

1. **Given** sufficient user requirements, **When** the AI queries the catalog, **Then** the AI renders a declarative UI schema (Product Carousel).
2. **Given** a recommended product, **When** the AI responds, **Then** it explains why this product fits the user's needs, including trade-offs and confidence.

---

### User Story 3 - Product Comparison (Priority: P2)

Customer asks to compare two or more recommended products. AI retrieves specifications, compares them, and renders a Comparison Table with trade-offs.

**Why this priority**: Enhances the decision-making process for the user.

**Independent Test**: User selects two products and asks to compare them; AI returns a comparison table UI schema.

**Acceptance Scenarios**:

1. **Given** multiple products in context, **When** the user asks to compare them, **Then** the AI generates a Comparison Table UI schema highlighting the differences.
2. **Given** a comparison, **When** the AI describes it, **Then** the trade-offs are clearly explained.

---

### User Story 4 - Checkout Handoff (Priority: P1)

Customer decides to buy a product. AI prepares checkout via the Tool Protocol, maintains human-in-the-loop validation, and hands off to the checkout service.

**Why this priority**: Connects the conversational journey to the final conversion goal.

**Independent Test**: User confirms intent to purchase a product; AI triggers the checkout preparation tool and presents a Checkout Summary.

**Acceptance Scenarios**:

1. **Given** the user's intent to purchase, **When** the AI is invoked, **Then** it uses the `checkout.prepare` tool to generate a Checkout Summary UI schema.
2. **Given** a completed checkout preparation, **When** the retailer supports direct checkout, **Then** the AI awaits explicit user confirmation before invoking `checkout.submit`.
3. **Given** a completed checkout preparation, **When** the retailer only supports guest/app checkout, **Then** the AI provides a deep link rather than an in-app panel.

### Edge Cases & Error Handling

- **Tool Execution Failures**: When a backend tool (e.g., catalog search) fails or times out, the system MUST attempt intelligent recovery before interrupting the UX:
  - Retry transient failures automatically.
  - Continue with cached or previously retrieved information if appropriate.
  - Fall back to alternative reasoning strategies if live data is unavailable, with clear indication to the user that data may not be current.
  - Render a structured Error State UI with recovery actions (Retry, Modify Search, Continue with Available Information) *only* when background recovery fails.
- **Out of Stock Before Checkout**: If a product becomes unavailable during checkout preparation, the AI MUST execute a graceful stock recovery:
  - Inform the user of the unavailability and the reason (e.g., recently sold out).
  - Present 2-3 similar alternatives ranked using the user's original requirements.
  - Highlight the differences (price, specs, availability).
  - Let the user continue with an alternative or modify requirements.
  - NEVER auto-replace the product without explicit user confirmation.
- **Contradictory Requirements**: If the user's requirements conflict (e.g., "gaming laptop under $100"), the AI MUST perform intelligent requirement negotiation:
  - Identify and explain the conflicting constraints conversationally in non-technical language.
  - Suggest 2-3 realistic alternatives by relaxing one constraint at a time (e.g., higher budget, lower performance).
  - Ask the user which trade-off they prefer.
  - Preserve original requirements in history so the user can easily revise.

## Clarifications

### Session 2026-07-11

- Q: How should the UI respond if a backend tool times out or fails during the conversation? → A: Attempt intelligent recovery (auto-retry, use cache, alternative reasoning with warning) before rendering a structured Error State UI with explicit recovery actions.
- Q: What happens if the AI recommends a product that goes out of stock immediately before checkout? → A: Graceful stock recovery: inform user, explain reason, present 2-3 similar alternatives based on original requirements, and require explicit confirmation (never auto-replace).
- Q: How does the system handle vague or contradictory user requirements (e.g., 'gaming laptop under $100')? → A: Intelligent requirement negotiation: explain conflicts, suggest 2-3 trade-off alternatives, ask for preference, and preserve original context. All prices MUST be in VND.

## Requirements *(mandatory)*

> **Note**: As per Principle I — Specification First, any interactions between components must be documented in separate API, UI, Tool, or Data Contracts.

### Functional Requirements

- **FR-001**: The system MUST allow users to interact via free-form text and a hybrid conversational onboarding with quick replies.
- **FR-002**: The AI MUST invoke backend interactions strictly through the registered Tool Protocol (e.g., `catalog.search`, `checkout.prepare`).
- **FR-003**: The AI MUST return declarative UI schema (Product Carousel, Comparison Table, Checkout Summary) instead of raw HTML/React components.
- **FR-004**: The AI MUST explain all recommendations, explicitly detailing trade-offs, alternatives, and confidence.
- **FR-005**: The AI MUST maintain conversation context across turns within a session (short-term memory).
- **FR-006**: The system MUST require explicit user confirmation for any state-changing operations like checkout.
- **FR-007**: The system MUST use Vietnamese Dong (VND) for all prices and budget-related interactions, converting or normalizing any other currency before presentation.

### Key Entities

- **Customer Session**: Manages the state and short-term memory of the ongoing conversation.
- **Product**: The item recommended, compared, or added to the cart.
- **Tool**: The backend function exposed via Tool Protocol to the AI (e.g., Search, Compare, Checkout).
- **UI Schema**: The deterministic JSON structure dictating the frontend rendering layout.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A user can complete an end-to-end shopping journey (from greeting to checkout handoff) entirely via conversation.
- **SC-002**: 100% of AI recommendations include a structured explanation of why the product was chosen.
- **SC-003**: All frontend rendering is driven by UI schema events, with zero visual components generated directly by the LLM.
- **SC-004**: All backend interactions initiated by the AI occur strictly via registered Tool Protocol tools.

## Assumptions

- Backend Tool Protocol implementations (`catalog.search`, `cart.add`, `checkout.prepare`, etc.) exist or will be implemented concurrently under a separate feature.
- A frontend renderer capable of interpreting UI schemas (Carousel, Comparison Table, Checkout Summary) exists.
- The retail integration layer (e.g., Phong Vu, Shopee) handles actual payment and inventory validation (Out of Scope for the AI Runtime).
- Long-term memory is explicitly managed by the user; the AI Runtime handles short-term context which is discarded after the session ends.

# Phase 0: Outline & Research

**Feature**: AI Shopping Conversation (003-ai-shopping-conversation)

## Architecture Decisions

The following technical decisions were established by the initial proposal and verified against the SHOPWISE Constitution:

### 1. Separation of Concerns (Three-Tier Architecture)
- **Decision**: Split the MVP into React (Frontend), Go 1.24 (Retail Backend), and Python (AI Runtime).
- **Rationale**: Enforces Principle IX (Language Responsibilities) and Principle III (AI is Stateless). The backend manages sessions and business logic, while Python handles reasoning via LangGraph.
- **Alternatives**: A single monolith (e.g., all Next.js or all Python) was rejected because it violates the constitution's strict language responsibility boundaries.

### 2. UI Generation Protocol
- **Decision**: The AI Runtime will output declarative UI schemas (e.g., Carousel, Comparison Table) via Server-Sent Events (SSE) or REST.
- **Rationale**: Complies with Principle VIII (Declarative UI). It ensures the frontend remains fully responsible for the visual presentation and React component lifecycle.
- **Alternatives**: AI generating raw HTML or React components was explicitly forbidden.

### 3. Tool Communication
- **Decision**: The AI Runtime must fetch product data and prepare checkout via a defined Tool Protocol exposed by the Go Backend.
- **Rationale**: Complies with Principle VI (Tool First). AI cannot directly query the database.

### 4. Data Storage
- **Decision**: Mocked JSON files for the MVP.
- **Rationale**: The focus is on the AI and conversational UX. Real retail integration is Out of Scope for this specific MVP milestone.

## Conclusion
All "NEEDS CLARIFICATION" points have been resolved. The architecture is sound and complies with all 15 established engineering principles. Proceed to Phase 1 (Data Model & Contracts).

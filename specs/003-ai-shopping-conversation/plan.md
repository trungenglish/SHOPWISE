# Implementation Plan: AI Shopping Conversation

**Branch**: `003-ai-shopping-conversation` | **Date**: 2026-07-11 | **Spec**: [spec.md](file:///d:/Github/SHOPWISE/specs/003-ai-shopping-conversation/spec.md)

**Input**: Feature specification from `/specs/003-ai-shopping-conversation/spec.md` and provided proposal `2.PLAN.md`.

## Summary

Implement an end-to-end AI Sales Agent MVP demonstrating conversational shopping, product recommendation, comparison, explainable AI, dynamic UI, and checkout readiness using a 3-tier architecture (React Frontend, Go Retail Backend, Python AI Runtime).

## Technical Context

**Language/Version**: React (TypeScript), Go 1.24, Python 3.x

**Primary Dependencies**: Vite, TailwindCSS, shadcn/ui, TanStack Query, Zustand, Gin, FastAPI, LangGraph

**Storage**: Mocked JSON (Products, Inventory, Promotions)

**Testing**: standard test runners (Vitest/Jest, `go test`, `pytest`)

**Target Platform**: Web application

**Project Type**: Full-stack web service (Frontend + API Gateway + AI Microservice)

**Performance Goals**: MVP optimized for fast iteration and demo readiness within 1 day.

**Constraints**:
- Backend owns all business logic.
- AI Runtime is purely stateless.
- AI communicates via Tool API only.
- AI generates only declarative UI schemas (no raw HTML/React components).
- All prices MUST use Vietnamese Dong (VND).

**Scale/Scope**: End-to-end shopping journey MVP. Excludes Payment Gateway, OMS, CRM, Auth, and distributed microservices infrastructure.

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] **Specification First**: Are all necessary contracts (API, UI, Tool, Data) identified? (Yes, moving to Phase 1 generation).
- [x] **AI is Stateless**: Does the plan ensure AI Runtime has no persistent state? (Yes, explicitly constrained).
- [x] **Backend Owns Business Logic**: Is all business logic delegated to backend services? (Yes, Go Backend handles catalog/session).
- [x] **Tool First**: Is the AI Runtime communicating via Tool Protocol only? (Yes).
- [x] **Declarative UI**: Is the AI only generating UI schema (no HTML/CSS/components)? (Yes).
- [x] **Language Responsibilities**: Are languages (Go, Python, TypeScript) used for their designated responsibilities only? (Yes).
- [x] **Testable Architecture**: Are unit, contract, and integration tests planned? (Yes).

## Project Structure

### Documentation (this feature)

```text
specs/003-ai-shopping-conversation/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
└── tasks.md             # Phase 2 output
```

### Source Code (repository root)

```text
apps/web/
├── src/
│   ├── components/
│   ├── hooks/
│   ├── services/
│   └── store/ (Zustand)

services/main-backend/ (Go)
├── src/
│   ├── api/
│   ├── catalog/
│   ├── session/
│   └── tools/

services/ai-runtime/ (Python)
├── src/
│   ├── agent/ (LangGraph)
│   ├── intent/
│   └── tools_client/
```

**Structure Decision**: The application is split into three runtimes to strictly adhere to Language Responsibilities and AI Statelessness constraints as defined in the Constitution.

## Complexity Tracking

No constitution violations detected. All constraints satisfied.

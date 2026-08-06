# Implementation Plan: [FEATURE]

**Branch**: `[###-feature-name]` | **Date**: [DATE] | **Spec**: [link]

**Input**: Feature specification from `/specs/[###-feature-name]/spec.md`

**Note**: This template is filled in by the `/speckit-plan` command. See `.specify/templates/plan-template.md` for the execution workflow.

## Summary

[Extract from feature spec: primary requirement + technical approach from research]

## Technical Context

**Language/Version**: TypeScript (Frontend), Go (Backend), Python (AI Runtime)

**Primary Dependencies**: React/Vite (Frontend), FastAPI/LangGraph (AI Runtime)

**Storage**: N/A (Protocol Definition only)

**Testing**: JSON Schema validation, contract tests across languages

**Target Platform**: Web (React), extensible to Mobile

**Project Type**: Protocol Definition (JSON Schemas) + Shared Types/SDKs

**Performance Goals**: Low latency parsing, compact JSON payloads, strict validation < 10ms

**Constraints**: Strictly declarative, no HTML/CSS, transport-agnostic, framework-independent

**Scale/Scope**: ~30 components (Core, Commerce, Conversation, Interaction, Error)

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] **Specification First**: Are all necessary contracts (API, UI, Tool, Data) identified?
- [x] **AI is Stateless**: Does the plan ensure AI Runtime has no persistent state?
- [x] **Backend Owns Business Logic**: Is all business logic delegated to backend services?
- [x] **Tool First**: Is the AI Runtime communicating via Tool Protocol only?
- [x] **Declarative UI**: Is the AI only generating UI schema (no HTML/CSS/components)?
- [x] **Language Responsibilities**: Are languages (Go, Python, TypeScript) used for their designated responsibilities only?
- [x] **Testable Architecture**: Are unit, contract, and integration tests planned?

## Project Structure

### Documentation (this feature)

```text
specs/[###-feature]/
├── plan.md              # This file (/speckit-plan command output)
├── research.md          # Phase 0 output (/speckit-plan command)
├── data-model.md        # Phase 1 output (/speckit-plan command)
├── quickstart.md        # Phase 1 output (/speckit-plan command)
├── contracts/           # Phase 1 output (/speckit-plan command)
└── tasks.md             # Phase 2 output (/speckit-tasks command - NOT created by /speckit-plan)
```

```text
packages/ui-protocol/
├── src/
│   ├── schema/           # JSON Schemas
│   ├── types/            # Generated TypeScript types
│   └── models/           # Shared models
└── tests/

services/main-backend/
├── internal/
│   └── ui-protocol/      # Go structs and serializers for the protocol

apps/web/
├── src/
│   └── components/
│       └── dynamic-ui/   # React components implementing the protocol renderer

services/ai-runtime/
├── src/
│   └── ui_protocol/      # Python Pydantic models for the protocol
```

**Structure Decision**: The protocol itself will be defined as language-agnostic JSON Schemas in a shared package (or directory), which will be used to generate or validate bindings in TypeScript (frontend), Go (backend), and Python (AI Runtime).

## Complexity Tracking

> **Fill ONLY if Constitution Check has violations that must be justified**

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| [e.g., 4th project] | [current need] | [why 3 projects insufficient] |
| [e.g., Repository pattern] | [specific problem] | [why direct DB access insufficient] |

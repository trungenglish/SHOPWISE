---
description: "Task list template for feature implementation"
---

# Tasks: Checkout Readiness

**Input**: Design documents from `/specs/007-checkout-readiness/`

**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are excluded as they were not explicitly requested in the feature specification. Independent verification steps are included in the quickstart.md.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [X] T001 Create `CheckoutReadiness` UI component scaffolding in `apps/web/src/components/checkout-readiness/`
- [X] T002 [P] Create `checkout-readiness` REST API router scaffolding in `services/main-backend/src/api/`
- [X] T003 [P] Create `checkout-readiness` AI workflow scaffolding in `services/ai-runtime/src/workflows/`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T004 Define `CheckoutReadinessState`, `ValidationItem`, and `CustomerProfile` structs in `services/main-backend/src/models/checkout.go`
- [X] T005 [P] Define `PromotionContext` and `ProductSnapshot` structs in `services/main-backend/src/models/promotion.go`
- [X] T006 Implement semantic event handlers structure (`/api/events/`) in `services/main-backend/src/api/events.go`
- [X] T007 Implement the `ValidateCheckoutReadiness` Tool definition for the AI Runtime in `services/ai-runtime/src/tools/checkout_tools.py`
- [X] T008 [P] Implement `ErrorStateUI` dynamic UI component in `apps/web/src/components/checkout-readiness/ErrorStateUI.tsx`

**Checkpoint**: Foundation ready - user story implementation can now begin in parallel

---

## Phase 3: User Story 1 - Successful Checkout Readiness Assessment (Priority: P1) 🎯 MVP

**Goal**: As a customer who has selected a product and provided all necessary information, I want the system to confirm I am ready to check out so that I can proceed to the retailer's payment flow with confidence.

**Independent Test**: Use quickstart.md Scenario 1 (Happy Path). Can be fully tested by simulating a user with a complete profile selecting an in-stock product and seeing the "Ready" status and action to continue.

### Implementation for User Story 1

- [X] T009 [US1] Implement `CheckoutSummary` dynamic UI component in `apps/web/src/components/checkout-readiness/CheckoutSummary.tsx`
- [X] T010 [P] [US1] Implement `ActionBar` dynamic UI component supporting Continue to Checkout, Request Similar Products, Change Store, Retry Validation, Save Session, and Resume Session actions in `apps/web/src/components/checkout-readiness/ActionBar.tsx`
- [X] T011 [US1] Implement backend validation logic (product available, price confirmed, warranty verified) and persist `CheckoutReadinessState` to Decision Memory in `services/main-backend/src/services/retail/validation_service.go`
- [X] T012 [US1] Implement the `POST /api/events/continue-checkout` handler in `services/main-backend/src/api/events.go`
- [X] T013 [US1] Update `services/ai-runtime/src/workflows/checkout_workflow.py` to call `ValidateCheckoutReadiness` and return `CheckoutSummary` when `status` is `READY`

**Checkpoint**: At this point, User Story 1 should be fully functional and testable independently

---

## Phase 4: User Story 2 - Resolving Missing Customer Information (Priority: P1)

**Goal**: As a customer with incomplete profile information, I want the AI to tell me exactly what is missing and allow me to provide it before checking out, so that my checkout process is not interrupted by basic data entry errors.

**Independent Test**: Use quickstart.md Scenario 2. Can be fully tested by initiating checkout readiness without an address/phone number and verifying the system prompts for it and updates the state.

### Implementation for User Story 2

- [X] T014 [US2] Implement `CustomerInformationForm` dynamic UI component in `apps/web/src/components/checkout-readiness/CustomerInformationForm.tsx`
- [X] T015 [P] [US2] Implement `StorePicker` dynamic UI component with geolocation fallback in `apps/web/src/components/checkout-readiness/StorePicker.tsx`
- [X] T016 [US2] Implement the `POST /api/checkout-readiness/customer-info` endpoint to securely collect PII in `services/main-backend/src/api/checkout.go`
- [X] T017 [P] [US2] Implement the `POST /api/events/store-selected` endpoint in `services/main-backend/src/api/events.go`
- [X] T018 [US2] Update `services/ai-runtime/src/workflows/checkout_workflow.py` to generate the `CustomerInformationForm` or `StorePicker` when `status` is `MISSING_INFORMATION`

**Checkpoint**: At this point, User Stories 1 AND 2 should both work independently

---

## Phase 5: User Story 3 - Handling Inventory and Price Changes (Priority: P2)

**Goal**: As a customer attempting to buy a product, I want to be notified if the product goes out of stock or if the price/promotion changes before I check out, so that I can make an informed decision or choose an alternative.

**Independent Test**: Can be fully tested by mocking backend responses for out-of-stock or changed price conditions and verifying the Error State UI and AI explanations.

### Implementation for User Story 3

- [X] T019 [US3] Extend `services/main-backend/src/services/retail/validation_service.go` to handle real-time inventory and price validations against the Retail SDK
- [X] T020 [US3] Implement logic to handle `OUT_OF_STOCK`, `STORE_UNAVAILABLE`, `INVENTORY_CHANGED`, and `REQUIRES_USER_ACTION` states in `services/main-backend/src/services/retail/validation_service.go`
- [X] T021 [US3] Update `services/ai-runtime/src/workflows/checkout_workflow.py` to explain inventory changes to the user and suggest alternatives using AI streaming text

**Checkpoint**: All user stories up to P2 should now be independently functional

---

## Phase 6: User Story 4 - Handling Promotional Countdowns (Priority: P2)

**Goal**: As a customer with a limited-time promotional offer in my session, I want to see a live countdown of how much time I have left to complete checkout, so that I don't miss out on the discount.

**Independent Test**: Use quickstart.md Scenario 3. Can be tested by returning an active promotion with an absolute expiration timestamp from the backend, verifying the live countdown in the Promotion Summary, and verifying the auto-revalidation when it expires.

### Implementation for User Story 4

- [X] T022 [US4] Implement `PromotionSummary` dynamic UI component with countdown timer in `apps/web/src/components/checkout-readiness/PromotionSummary.tsx`
- [X] T023 [US4] Implement `POST /api/events/promotion-expired` endpoint to force re-validation in `services/main-backend/src/api/events.go`
- [X] T024 [US4] Implement logic to enforce explicit price change tolerances and handle the race condition upon checkout in `services/main-backend/src/services/retail/validation_service.go`
- [X] T025 [US4] Update `services/ai-runtime/src/workflows/checkout_workflow.py` to handle the `PROMOTION_CHANGED` state when the countdown reaches zero and notify the user

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [X] T026 [P] Add detailed logging for validation failures and timeout metrics in `services/main-backend/src/services/retail/validation_service.go`
- [X] T027 Code cleanup and refactoring in frontend React components to ensure UI consistency
- [X] T028 Security hardening: verify PII bypass mechanisms (Constitution Principle XV) in `services/main-backend/src/api/checkout.go`
- [X] T029 Verify Explainable AI requirements (Constitution Principle XI) across all `checkout_workflow.py` prompts
- [X] T030 Add concurrent update protection using `cartVersion` hash across all endpoints in `services/main-backend/src/api/events.go`
- [X] T031 Enforce Vietnamese Dong (VND) normalization in the Go Backend and formatting in UI components (Constitution Principle XVI)

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
- **User Story 2 (P1)**: Can start after Foundational (Phase 2) - Depends on US1 completion for the base `CheckoutSummary` fallback
- **User Story 3 (P2)**: Can start after Foundational (Phase 2) - Expands on the validation service built in US1
- **User Story 4 (P2)**: Can start after Foundational (Phase 2) - Requires US1 frontend structure to inject the `PromotionSummary`

### Within Each User Story

- Models before services
- Services before endpoints
- Core implementation before integration
- Story complete before moving to next priority

### Parallel Opportunities

- All Setup tasks marked [P] can run in parallel
- All Foundational tasks marked [P] can run in parallel (within Phase 2)
- Once Foundational phase completes, all user stories can start in parallel (if team capacity allows)
- Different user stories can be worked on in parallel by different team members

---

## Parallel Example: User Story 1

```bash
# Launch backend UI components together:
Task: "Implement CheckoutSummary dynamic UI component in apps/web/src/components/checkout-readiness/CheckoutSummary.tsx"
Task: "Implement ActionBar dynamic UI component in apps/web/src/components/checkout-readiness/ActionBar.tsx"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1
4. **STOP and VALIDATE**: Test User Story 1 independently using quickstart.md Scenario 1
5. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 → Test independently → Deploy/Demo (MVP!)
3. Add User Story 2 → Test independently using quickstart.md Scenario 2 → Deploy/Demo
4. Add User Story 3 & 4 → Test independently using quickstart.md Scenario 3 → Deploy/Demo
5. Each story adds value without breaking previous stories

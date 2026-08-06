# Implementation Tasks: Checkout Incentive Countdown

## Implementation Strategy
**MVP First Delivery**: Focus on building the backend foundation and the state machine first. The backend handles all authoritative state (expiration, eligibility, cooldowns), ensuring robust rules enforcement before moving to frontend integration.

## Dependencies
- **Phase 1** blocks all backend tasks in subsequent phases.
- **Phase 2** blocks the main integration of the frontend components.
- **Phase 3** can be executed parallel to **Phase 2**.
- **Phase 4** integration tasks depend on the completion of the prior phases.

---

## Phase 1: Database & Data Model

- [x] T001 Create PostgreSQL migration for `checkout_promotions` table (including reward metadata: `campaign_code`, `reward_type`, `reward_value`, `reward_scope`) with unique session constraints in `apps/server/internal/db/migrations/000004_create_checkout_promotions.up.sql`
- [x] T002 Create GORM `CheckoutPromotion` model with state enum mapping and reward metadata defaults in `apps/server/internal/promotions/model.go`
- [x] T003 Implement `Create`, `FindBySession`, `FindByEligibilityKey`, `UpdateStatus`, and `SaveVoucherResult` repository methods in `apps/server/internal/promotions/repository/postgres.go`
- [x] T004 [P] Define shared API types for trigger payloads and status responses in `packages/api-types/src/promotions.ts`

**Acceptance Criteria:**
- 15-minute `expires_at` is persisted correctly.
- Authenticated users use `user_id`; anonymous users use a privacy-preserving device identifier.
- Promotion records cannot issue duplicate vouchers.
- Expired promotions cannot be restarted in the same session.

---

## Phase 2: Backend API & State Logic

- [x] T005 [Backend API] Implement `StartPromotion` service enforcing 30-day configurable cooldown and idempotent activation in `apps/server/internal/promotions/service.go`
- [x] T006 [Backend API] Implement POST `/api/v1/checkout-incentive/start` handler for Save Product / Checkout Started triggers in `apps/server/internal/promotions/handler.go`
- [x] T007 [Backend API] Implement GET `/api/v1/checkout-incentive/status` with lazy ACTIVE to EXPIRED transition based on backend time in `apps/server/internal/promotions/handler.go`
- [x] T008 [Backend API] Implement Payment Completed integration contract handling transition to `COMPLETED_PENDING_VOUCHER` in `apps/server/internal/promotions/service.go`
- [x] T009 [Backend API] Implement local voucher issuance logic (generate locally, persist result, exactly-once enforcement, clean internal interface), `ISSUE_RETRY_REQUIRED` state, and retry handling in `apps/server/internal/promotions/voucher.go`
- [x] T010 [P] [Backend API] Implement POST `/api/v1/checkout-incentive/dev/simulate-payment` development-only endpoint in `apps/server/internal/promotions/handler_dev.go`
- [x] T011 [P] [Backend API] Write API contract tests and exact 15-minute deadline boundary tests in `apps/server/internal/promotions/handler_test.go`
- [x] T012 [P] [Backend API] Write state-machine unit tests in `apps/server/internal/promotions/service_test.go`

**Acceptance Criteria:**
- Repeated Save Product or Checkout Started actions do not reset or extend the timer.
- Page refresh does not restart the timer.
- Backend time is authoritative.
- Payment completed before deadline may issue one voucher.
- Payment completed after deadline must not issue a voucher.
- Duplicate payment events do not issue duplicate vouchers.
- Temporary voucher failure preserves eligibility.
- A new session is required after expiration.

---

## Phase 3: Frontend UI Components

- [ ] T013 [Frontend] Set up TanStack Query hooks and `CheckoutIncentiveProvider` context allowing nested components to trigger mutations and query status in `apps/web/src/features/checkout-incentive/api/queries.ts`
- [ ] T014 [Frontend] Create `CountdownDisplay` component with backend clock drift reconciliation in `apps/web/src/features/checkout-incentive/components/CountdownDisplay.tsx`
- [ ] T015 [Frontend] Create `PromotionBadge` component for the collapsed state with expand/collapse behavior in `apps/web/src/features/checkout-incentive/components/PromotionBadge.tsx`
- [ ] T016 [Frontend] Create `VoucherIssuedState`, `PromotionExpiredState`, `PromotionErrorState`, and `ISSUE_RETRY_REQUIRED` state components in `apps/web/src/features/checkout-incentive/components/States.tsx`
- [ ] T017 [Frontend] Create the main `CheckoutIncentivePanel` supporting desktop right-side floating panel and mobile compact bottom sheet/badge layouts in `apps/web/src/features/checkout-incentive/components/CheckoutIncentivePanel.tsx`
- [ ] T018 [P] [Frontend] Write frontend component tests in `apps/web/src/features/checkout-incentive/components/CheckoutIncentivePanel.test.tsx`

**Acceptance Criteria:**
- Component appears after first eligible trigger.
- Panel does not cover recommendation, comparison, or checkout content.
- Timer display remains accurate after tab sleep and refresh.
- Active, expired, success, retry, and error states are visually distinct.
- Keyboard and screen-reader users can operate the component (WCAG accessibility).
- Frontend does not determine promotion eligibility independently.

---

## Phase 4: Integration & End-to-End Validation

- [x] T019 [Integration] Implement `CheckoutIncentiveProvider` and `useCheckoutIncentive` hook, rendering the panel at the layout level and allowing nested components to trigger promotion activation.
- [x] T020 [Integration] Wire Payment Completed event to backend promotion handling in `apps/server/internal/checkout/handler.go` (or applicable internal event bus)
- [x] T021 [Integration] Validate end-to-end quickstart locally and ensure monorepo build/type-checks pass.
- [x] T022 [Integration] Validate refresh continuity, cross-session behavior, cooldown, expired-session lockout, and voucher duplication prevention.

**Acceptance Criteria:**
- Monorepo build and type-check pass.
- All integration and end-to-end tests meet the specified constraints without manual interference.
- Tasks in tasks.md are not marked complete until acceptance criteria pass.

# Checkout Incentive Countdown

**Version**: 1.0.0
**Status**: Draft
**Feature Directory**: specs/009-checkout-incentive
**Created**: 2026-08-04
**Last Updated**: 2026-08-04

---

## Overview

The Checkout Incentive Countdown introduces a small promotional countdown panel on the right side of the shopping workspace. When a shopper saves a product or starts checkout, a 15-minute timer begins. If the user completes their payment within this window, they receive a 15% discount voucher for accessories on their next purchase, creating urgency without interrupting the core shopping experience.

---

## Clarifications

### Session 2026-08-04
- Q: Can a shopper receive this 15% voucher multiple times across different sessions, or is it strictly a one-time reward per user account? → A: Once per shopping session; they can receive it again in future sessions.
- Q: If the promotion expires, when can the user trigger a new 15-minute countdown? → A: Never within the same session. They must start a new session to become eligible again.

---

## Problem Statement

Shoppers often evaluate products or begin the checkout process but hesitate or abandon the transaction before completion. The current flow lacks a time-sensitive incentive to encourage immediate action. By offering a time-limited 15% accessory voucher, we can drive urgency and increase checkout completion rates.

---

## Goals

- Encourage shoppers to complete checkout within 15 minutes of expressing strong intent.
- Start the 15-minute countdown exactly once per session upon saving a product or starting checkout.
- Display a non-blocking UI panel with the remaining time, reward explanation, and current promotion state.
- Accurately track time using the backend as the source of truth to ensure consistency across page refreshes and devices.
- Issue a 15% accessory voucher for a future purchase upon successful payment within the active countdown window.

## Non-Goals

- Implementation of the payment gateway itself.
- Modifications to the core checkout business logic.
- Building the accessory catalog.
- Handling the voucher redemption flow on future purchases.
- Marketing campaign management or administration UI.
- Notification delivery via external channels (Zalo, SMS, Email).
- AI-generated promotion rules.

---

## Target Users

| User Type | Description | Primary Need |
|---|---|---|
| Authenticated Shopper | A logged-in user evaluating or buying products | To receive a promotional incentive for checking out quickly that persists across their authenticated devices. |
| Anonymous Shopper | An unauthenticated user browsing or checking out | To receive the same incentive reliably within their current browser session. |

---

## User Scenarios & Testing

### Scenario 1: Countdown Starts
**Given**: The shopper has no active promotion in their current session.
**When**: The shopper saves a product or selects to start checkout.
**Then**: A 15-minute promotional countdown starts on the backend.
**Acceptance**: A non-blocking panel appears on the right side of the workspace, displaying the remaining time, the 15% voucher reward, and an explanation of eligibility.

### Scenario 2: Successful Completion
**Given**: The shopper has an active 15-minute promotional countdown.
**When**: The shopper successfully completes the payment before the timer expires.
**Then**: The promotion is marked as completed and a 15% accessory voucher is issued for their next purchase.
**Acceptance**: The UI updates to a success state, confirming the voucher was earned. Exactly one voucher is issued.

### Scenario 3: Expiration
**Given**: The shopper has an active 15-minute promotional countdown.
**When**: 15 minutes pass without a confirmed payment completion.
**Then**: The promotion becomes expired and the timer stops at zero.
**Acceptance**: No voucher is issued, the UI explains the window has ended, and repeating trigger actions (like starting checkout again) does not restart the same expired promotion.

### Scenario 4: Refresh and Device Continuity
**Given**: An authenticated shopper starts the countdown on their desktop device.
**When**: The shopper refreshes the page or logs into the same session on a mobile device.
**Then**: The countdown continues uninterrupted based on the backend time.
**Acceptance**: The remaining time displayed is accurate, not reset to 15 minutes, and the panel state remains correct.

### Scenario 5: Error Recovery
**Given**: A shopper successfully completes payment within the time window.
**When**: The voucher issuance API temporarily fails.
**Then**: The shopper's eligibility is preserved.
**Acceptance**: The UI displays a recoverable error state and allows for a retry, ensuring the earned voucher is not lost due to temporary system failures.

---

## Functional Requirements

### Trigger and State Management

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-01 | Trigger Conditions | Must Have | Promotion activates only on the first instance of "save product" or "start checkout". Subsequent triggers do not extend or reset the timer. If the promotion expires, the user cannot trigger a new countdown within the same session. |
| FR-02 | Backend Time Truth | Must Have | The timer duration is governed by the backend. The frontend must not determine eligibility using only a local timer. |
| FR-03 | Promotion States | Must Have | The system must correctly reflect Not started, Active, Completed successfully, Expired, Voucher issued, Ineligible, and Error states. |

### Presentation and UI

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-04 | Desktop Layout | Must Have | Display as a small right-side panel that does not cover recommendation, comparison, or checkout content. |
| FR-05 | Mobile Layout | Must Have | Display an equivalent compact presentation that does not block the checkout flow. |
| FR-06 | Collapsible State | Must Have | The user can dismiss or collapse the panel into a small badge without canceling the promotion. |
| FR-07 | Accessibility | Must Have | The countdown and panel states must be accessible to keyboard navigation and screen-reader users. |

### Voucher Issuance

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-08 | Issuance Condition | Must Have | Voucher is issued if and only if payment completion is confirmed before the backend deadline. |
| FR-09 | Deduplication | Must Have | A maximum of one voucher is issued per shopping session/promotion. The shopper is eligible to receive this incentive again in future shopping sessions. |

---

## Success Criteria

| Criterion | Metric | Target |
|---|---|---|
| Reliability | Timer persistence across page reloads | 100% (Timer does not reset on refresh) |
| Voucher Integrity | Duplicate vouchers issued per promotion | 0 duplicates |
| Trigger Accuracy | Timer extensions from repeated actions | 0 extensions (Timer remains strict to the original 15 minutes) |

---

## Assumptions

- Integration endpoints for "Saved Product", "Checkout Started", and "Payment Completed" events are stable and available.
- A Voucher Issuance API exists and handles the actual generation and storage of the voucher codes.
- The shopping session architecture securely identifies returning authenticated users and local anonymous browser sessions.

## Dependencies

- Shopping Session API
- Voucher Issuance API
- Event bus or API hooks for checkout/save events

## Risks

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Local timer drift | Medium | High | The frontend will use the backend time as the source of truth, and the backend will ultimately enforce the payment completion deadline. |
| Temporary voucher issuance failure | Medium | Medium | The backend will preserve the "Completed successfully" state and allow retrying the issuance without failing the promotion. |

---

## Open Questions

- None at this time.

## Remediation Edits

- **Reward Modeling**: The MVP uses the default campaign `ACCESSORY_NEXT_PURCHASE_15` with a 15% discount for accessories. This metadata is stored on the promotion model.
- **Voucher Architecture**: For MVP, voucher issuance is handled locally by the backend without external services.
- **Frontend Architecture**: A feature-scoped integration layer (`CheckoutIncentiveProvider` and `useCheckoutIncentive`) is used instead of layout-level direct tracking.

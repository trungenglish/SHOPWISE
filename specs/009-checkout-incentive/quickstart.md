# Quickstart: Checkout Incentive Countdown Validation

This guide explains how to validate the Checkout Incentive feature locally without needing a real payment gateway.

## Prerequisites
- Both `apps/web` and `apps/server` running (`pnpm dev`)
- Postgres running (`docker compose up postgres`)

## Validation Scenario 1: Successful Checkout within 15 Minutes

1. **Trigger Promotion**:
   Open the frontend (e.g. `localhost:5173`), ensure you are logged in (or using an anonymous session), and click "Save" on any product, or click "Checkout".
   *Expected*: The `CheckoutIncentivePanel` appears on the right side with a 15:00 countdown.

2. **Simulate Payment**:
   Instead of using a real payment gateway, invoke the developer simulation endpoint using your session ID:
   ```bash
   curl -X POST http://localhost:8080/api/v1/checkout-incentive/dev/simulate-payment -d '{"session_id": "your-session-id"}' -H "Content-Type: application/json"
   ```
   *Expected*: The timer panel updates to the `PromotionSuccessState` showing the 15% voucher.

## Validation Scenario 2: Expiration

1. **Trigger Promotion**:
   Save a product or click Checkout.
   *Expected*: The timer starts.

2. **Wait or Simulate Expiration**:
   Wait 15 minutes, or manually update the `expires_at` in the database to be in the past.
   *Expected*: The timer stops at zero. The panel switches to `PromotionExpiredState`.

3. **Verify Constraints**:
   Click "Checkout" again.
   *Expected*: The promotion does not restart. You remain in the expired state.

## Validation Scenario 3: Page Refresh Continuity

1. **Trigger Promotion**:
   Start the promotion. Observe the remaining time (e.g. 14:30).

2. **Refresh Page**:
   Refresh the browser page (`F5`).
   *Expected*: The timer loads and resumes from exactly the correct backend time (e.g. ~14:28), and does not reset to 15:00.

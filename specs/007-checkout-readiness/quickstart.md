# Quickstart Validation Guide: Checkout Readiness

## Overview
This guide provides instructions to validate the end-to-end Checkout Readiness flow without requiring an actual retailer checkout integration.

## Prerequisites
- React Frontend running locally (`apps/web`)
- Go Backend running locally (`services/main-backend`)
- Python AI Runtime running locally (`services/ai-runtime`)
- Mock Retail API server active (returns dummy inventory/prices)

## Validation Scenario 1: Happy Path (Ready to Checkout)
1. Open the Product Comparison workspace and add a mock product (e.g., `mock_iphone`) to the checkout queue.
2. Ensure your user session has a valid Name, Phone, and Delivery Address set.
3. Click "Proceed to Checkout".
4. **Expected Outcome**:
   - The AI Runtime calls `ValidateCheckoutReadiness`.
   - The Go backend verifies `mock_iphone` is in stock.
   - The UI renders `CheckoutSummary` indicating "READY".
   - The `Continue to Retail Checkout` action is enabled.

## Validation Scenario 2: Missing Information (PII Collection)
1. Clear the delivery preference and phone number from the session.
2. Click "Proceed to Checkout".
3. **Expected Outcome**:
   - The UI renders the `CustomerInformationForm`.
   - The AI asks, "To proceed, please provide your phone number and delivery preference."
4. **Action**: Fill out the form in the UI and click submit.
5. **Expected Outcome**:
   - The form posts to `/api/checkout-readiness/customer-info`.
   - The AI responds, "Thank you, your information is saved. You are ready to check out."
   - The UI transitions to the "READY" state.

## Validation Scenario 3: Promotion Expiration
1. Configure the mock Retail API to return a promotion for `mock_iphone` that expires in 10 seconds.
2. Click "Proceed to Checkout".
3. **Expected Outcome**:
   - The UI renders `PromotionSummary` with a live 10-second countdown.
4. **Wait 10 seconds**.
5. **Expected Outcome**:
   - The countdown hits 0.
   - The Frontend emits `PROMOTION_EXPIRED`.
   - The Backend automatically fetches the new price (without the discount).
   - The AI streams a message: "The flash sale has ended. The new price is 20,000,000 VND. Would you still like to proceed?"
   - The UI updates to `Promotion Changed` state.

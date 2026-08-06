# UI Contracts: Checkout Incentive

## Components (React / TanStack Query)

### `CheckoutIncentiveProvider`
A feature-scoped Context Provider to wrap the layout.
**Exports**: `useCheckoutIncentive()` hook providing `startPromotion(triggerType)` and the current promotion state.

### `CheckoutIncentivePanel`
The main container component rendered conditionally on the right side of the screen when a promotion is active, completed, expired, or in an error state.
**Props**:
- `sessionId`: Current shopping session context identifier.

### `CountdownDisplay`
Displays the remaining minutes and seconds. Syncs strictly with the backend `expiresAt` timestamp. Periodically checks backend status to ensure the local timer does not drift or spoof the time.
**Props**:
- `expiresAt`: `string` (ISO datetime)

### `PromotionBadge`
The collapsed state of the panel. Shows a small icon and remaining time. Clicking this expands back to the full `CheckoutIncentivePanel`.

### `PromotionSuccessState`
Rendered when status is `VOUCHER_ISSUED`. Shows the voucher code and next steps for the user.

### `PromotionExpiredState`
Rendered when status is `EXPIRED`. Informs the user they missed the window and encourages them to complete the checkout regardless.

### `PromotionErrorState`
Rendered when status is `ERROR_RECOVERABLE` or `ISSUE_RETRY_REQUIRED`. Shows a "Retry" button.

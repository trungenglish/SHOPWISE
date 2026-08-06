# Feature Specification: Checkout Readiness

**Feature Branch**: `[007-checkout-readiness]`

**Created**: 2026-07-12

**Status**: Draft

**Input**: User description: "Feature: Checkout Readiness..."

## Clarifications

### Session 2026-07-12

- Q: How should the AI Sales Agent handle the Personally Identifiable Information (PII) like name, phone, and address it collects during Checkout Readiness? → A: PII must be collected via secure Dynamic UI form components directly to the Retail Backend APIs. AI Runtime uses non-sensitive placeholders (e.g., "phone number verified") and does not persist raw PII.
- Q: When a user selects 'Store pickup,' how does the system determine which stores to display as available? → A: Hybrid approach using device geolocation API with user consent to fetch nearby stores from backend, falling back to manual search/filter in the Dynamic UI if denied.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Successful Checkout Readiness Assessment (Priority: P1)

As a customer who has selected a product and provided all necessary information, I want the system to confirm I am ready to check out so that I can proceed to the retailer's payment flow with confidence.

**Why this priority**: This is the happy path and represents the primary value of the feature—bridging the gap between product selection and actual purchase smoothly.

**Independent Test**: Can be fully tested by simulating a user with a complete profile selecting an in-stock product and seeing the "Ready" status and action to continue.

**Acceptance Scenarios**:

1. **Given** a selected product is in stock and price is unchanged, and the customer has provided name, phone, and delivery preference, **When** the checkout readiness assessment runs, **Then** the system displays a "Ready" state and a "Continue to Retail Checkout" action.
2. **Given** the customer selects store pickup, **When** they select a valid store location, **Then** the system confirms readiness and enables the checkout action.

---

### User Story 2 - Resolving Missing Customer Information (Priority: P1)

As a customer with incomplete profile information, I want the AI to tell me exactly what is missing and allow me to provide it before checking out, so that my checkout process is not interrupted by basic data entry errors.

**Why this priority**: Missing information is a common cause of checkout abandonment. Resolving it conversationally before handoff is a core requirement.

**Independent Test**: Can be fully tested by initiating checkout readiness without an address/phone number and verifying the system prompts for it and updates the state.

**Acceptance Scenarios**:

1. **Given** the customer has not provided a phone number, **When** the checkout readiness assessment runs, **Then** the system enters a "Missing Information" state and the AI prompts the user for their phone number.
2. **Given** the system is in a "Missing Information" state, **When** the user provides the missing details, **Then** the system updates the validation checklist and transitions to the "Ready" state.

---

### User Story 3 - Handling Inventory and Price Changes (Priority: P2)

As a customer attempting to buy a product, I want to be notified if the product goes out of stock or if the price/promotion changes before I check out, so that I can make an informed decision or choose an alternative.

**Why this priority**: Inventory and pricing are dynamic. Handling these gracefully prevents poor user experiences and customer service issues post-handoff.

**Independent Test**: Can be fully tested by mocking backend responses for out-of-stock or changed price conditions and verifying the Error State UI and AI explanations.

**Acceptance Scenarios**:

1. **Given** the selected product goes out of stock, **When** the checkout readiness assessment runs, **Then** the system displays an "Out of Stock" state, explains the issue, and suggests similar products.
2. **Given** the product price has changed (e.g., a promotion expired), **When** the assessment runs, **Then** the system displays a "Promotion Changed" state with the new VND price and asks if the user still wants to proceed.

### User Story 4 - Handling Promotional Countdowns (Priority: P2)

As a customer with a limited-time promotional offer in my session, I want to see a live countdown of how much time I have left to complete checkout, so that I don't miss out on the discount.

**Why this priority**: Drives urgency and conversion, but is secondary to the core checkout readiness logic.

**Independent Test**: Can be tested by returning an active promotion with an absolute expiration timestamp from the backend, verifying the live countdown in the Promotion Summary, and verifying the auto-revalidation when it expires.

**Acceptance Scenarios**:

1. **Given** an active promotion with an expiration timestamp, **When** the customer views Checkout Readiness, **Then** the Promotion Summary displays a live countdown.
2. **Given** the user leaves the session and returns, **When** they view Checkout Readiness, **Then** the countdown accurately reflects the remaining time based on the absolute backend timestamp.
3. **Given** the countdown reaches zero while the user is active, **Then** the system automatically re-validates, updates the price, and the AI proactively informs the user that the promotion has expired.

### Edge Cases

- **Promotion Expiration Race Condition**: If the promotional countdown expires exactly at the moment the user clicks "Continue to Retail Checkout", the Backend MUST reject the transition, revert to `Promotion Changed` state, and return the updated state to the AI.
- **Backend Timeouts**: When the Retail Backend times out or is temporarily unavailable during validation, the system enters an `Error State` and renders a Structured Error State UI with retry options.
- **Store Availability Changes**: If a selected store for pickup suddenly closes or runs out of inventory during validation, the system MUST transition to `Store Unavailable` state and prompt the user to select a new store.
- **Invalid State Transitions**: If the user tries to proceed to checkout while in an "Error" or "Missing Information" state, the backend rejects the request with a 400 Bad Request, returning the specific missing checklist items to flash in the UI.
- **Concurrent Session Updates**: If a user updates their cart in another tab while validation is running, the current Readiness State is invalidated. The backend tracks a `cartVersion` hash and forces re-validation if mismatched.

## Requirements *(mandatory)*

> **Note**: As per Principle I — Specification First, any interactions between components must be documented in separate API, UI, Tool, or Data Contracts.

### Functional Requirements

- **FR-001**: The system MUST verify product availability, inventory status, latest pricing (in VND), promotions, and warranty information against the Retail Backend.
- **FR-002**: The system MUST evaluate if all required customer information (e.g., name, phone number, delivery preference, store selection) is present.
- **FR-003**: The AI MUST prompt for missing customer information, which MUST be collected via secure Dynamic UI form components submitted directly to the Retail Backend APIs. The AI Runtime MUST NOT persist raw PII in its conversation memory, instead receiving semantic events with non-sensitive placeholders (e.g., `<phone_verified>`, `<address_collected>`).
- **FR-004**: The system MUST present a dynamic UI displaying a Checkout Summary, Validation Checklist, Delivery Options, and Readiness Status.
- **FR-005**: The system MUST support selection between "Home delivery" and "Store pickup" using a Store Picker that attempts device geolocation (with consent) to display nearby stores. It MUST fallback to manual location search if: 1) The user denies location permission, 2) The Geolocation API times out after 5 seconds, or 3) The Backend returns 0 stores within a 50km radius.
- **FR-006**: The system MUST classify the readiness state as one of: Ready, Missing Information, Inventory Changed, Promotion Changed, Out of Stock, or Requires User Action.
- **FR-007**: The AI MUST explain validation failures clearly and recommend corrective actions (e.g., suggesting similar products if out of stock).
- **FR-008**: The system MUST provide commerce actions including "Continue to Retail Checkout", "Modify Product Selection", "Select Another Store", "Request Similar Products", "Save Session", and "Resume Later".
- **FR-009**: The system MUST gracefully handle backend timeouts (threshold: 2000ms) or service unavailability by displaying a structured Error State UI with recovery actions.
- **FR-010**: The system MUST persist the checkout readiness information as part of Decision Memory.
- **FR-011**: The system MUST support an informational promotional countdown managed by an absolute expiration timestamp (UTC ISO-8601 format) provided by the Retail Backend, integrated directly into the Promotion Summary component. The Frontend live countdown is purely presentational; the Go backend is the ultimate arbiter of whether the checkout occurred before expiration.
- **FR-012**: The system MUST automatically re-validate the checkout readiness state, update the price, and proactively inform the user if the promotional countdown expires while the user is active.
- **FR-013**: The system MUST enforce explicit price change tolerances. If the price increases by any amount, the user MUST explicitly accept the new price. If the price decreases, the system auto-accepts and notifies the user.

### Key Entities *(include if feature involves data)*

- **Checkout Readiness State**: Represents the current validation status (Ready, Missing Information, etc.) and the checklist of validated items.
- **Customer Profile**: Contains the accumulated user information (Name, Phone, Delivery Address, Preferred Store).
- **Validation Checklist**: A list of required conditions (Product Available, Price Confirmed, Customer Info Complete) and their pass/fail status.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 90% of users who reach the "Ready" state successfully transition to the retailer's checkout flow.
- **SC-002**: Checkout abandonment due to missing customer information or inventory issues drops by 30%.
- **SC-003**: The checkout readiness validation process completes within 2 seconds.
- **SC-004**: 80% of users encountering a "Missing Information" state successfully provide the information and proceed.

## Assumptions

- Standard customer information required for delivery is Name, Phone Number, and Delivery Address.
- Standard customer information required for pickup is Name, Phone Number, and Selected Store.
- The Retail Backend provides APIs with acceptable latency to perform real-time inventory and price checks.
- The retailer's existing checkout flow can accept a handoff payload containing the collected product and customer data.

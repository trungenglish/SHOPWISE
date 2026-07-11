# ShopWise Get Started Onboarding Flow

**Version**: 1.0.0 **Status**: Draft **Feature Directory**: specs/002-get-started-onboarding **Created**: 2026-07-11 **Last Updated**: 2026-07-11

---

## Overview

The ShopWise "Get Started" onboarding flow is a guided entry experience that bridges the gap between landing page discovery and active platform usage inside the dashboard. It routes existing users through sign-in and guides new users through profile setup with custom validation, maintaining visual alignment with the dark-first design system.

---

## Clarifications

### Session 2026-07-11

- Q: Should we build a minimal, custom sign-in component as part of the new /get-started onboarding workspace, or should we build it under the /auth layout path? → A: Build a new, minimal sign-in form directly inside the /get-started multi-step flow to maintain one continuous onboarding experience and avoid redirecting users between layouts.
- Q: What credentials should returning users enter to sign in (Email, Phone, or both)? → A: Support both email and phone number with a single "Email or Phone Number" input plus password.
- Q: Should this form be presented as a single page or split into multiple sequential step screens? → A: Use a multi-step onboarding wizard. For new users, guide them through Step 1 (Full Name), Step 2 (Phone Number), Step 3 (Email Address), Step 4 (Review), and Step 5 (Submit), with a progress indicator and state preservation on back navigation.
- Q: Should the phone number validation specifically target Vietnamese phone numbers or global format? → A: Validate standard Vietnamese phone numbers: exactly 10 digits starting with standard mobile prefixes (03, 05, 07, 08, 09), trimming whitespace before checking, and displaying inline validation errors.
- Q: Can users navigate back to the initial account-status choice step from subsequent step views in the onboarding wizard? → A: Yes, support back navigation to the choice screen from all subsequent steps, preserving entered form data.
- Q: Should onboarding submission trigger a real API endpoint request, or operate purely as a local mock? → A: Client-side simulated submission (network latency ~800-1500ms, loading state display, success/failure UI, navigate to /dashboard after success, isolated submission logic).
- Q: Does the /dashboard route exist in the codebase, or should we create a minimal route placeholder? → A: Do not create a placeholder /dashboard route (developed in a separate branch). Navigate to "/dashboard" on success without modifying or creating dashboard files.

---

## Problem Statement

Prospective users and returning customers need a friction-free, intuitive entry path when transitioning from the landing page. Without a structured entry flow, users cannot sign in or set up their profile, leading to high drop-off before reaching the product search dashboard.

---

## Goals

- Provide a single, accessible `/get-started` entry point.
- Implement a clean branching flow based on account status ("Yes, sign in" vs "No, create my profile").
- Collect and validate profile fields (Full Name, Phone Number, Email) for new users with real-time feedback.
- Navigate successfully authenticated or registered users to `/dashboard` using TanStack Router.

## Non-Goals

- Implementing backend authentication logic, session cookies, database persistence, OTP validation, SMS dispatch, or social auth providers.
- Redesigning or implementing the actual `/dashboard` workspace.

---

## Target Users

| User Type | Description | Primary Need |
| --- | --- | --- |
| Online Shopper | A consumer searching for products on e-commerce sites (e.g. Phong Vu) | Quickly setup a temporary profile to start searching and comparing products |
| Sales Team | ShopWise operators assisting users | Demonstrate tool flow to buyers without backend delays |
| Support Team | ShopWise administrators | Simple authentication to manage client queries |

---

## User Scenarios & Testing

### Scenario 1: Flow Selection (Account Status Choice)

**Given**: A user navigates to `/get-started` on any browser. **When**: They view the account selection screen. **Then**: They see clear options: "Yes, sign in" and "No, create my profile". **Acceptance**: Buttons are keyboard navigable with focus rings.

### Scenario 2: Existing User Sign-In

**Given**: A user selects "Yes, sign in" on the `/get-started` route. **When**: They enter their login information and click submit. **Then**: The system shows a loading indicator, then redirects them to `/dashboard`. **Acceptance**: Validation errors appear if fields are empty.

### Scenario 3: New User Registration

**Given**: A user selects "No, create my profile" on the `/get-started` route. **When**: They fill out their Full name, Phone number, and Email, and click submit. **Then**: The system validates fields, shows a loading indicator, and redirects to `/dashboard`. **Acceptance**: Invalid inputs (e.g. malformed email/phone) block submission and display inline errors.

---

## Functional Requirements

### Account Status Choice

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-GET-01 | Clear step heading asking "Do you already have an account?" | Must Have | Rendered as high-contrast heading element. |
| FR-GET-02 | Choice buttons: "Yes, sign in" and "No, create my profile". | Must Have | Accessible button targets (>= 44px height). Keyboard selection supported. |
| FR-GET-03 | Preserve entered form data when moving back and forth between steps. | Should Have | Form field values are cached locally during the session. |

### Onboarding Steps & Authentication Form

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-GET-04 | Existing user sign-in step displaying credential placeholder inputs (Email or Phone). | Must Have | User inputs are validated. Temporary "Sign In" button continues to /dashboard. |
| FR-GET-05 | New user profile form containing Full Name, Phone Number, and Email fields. | Must Have | Inputs use custom styled `@shopwise/ui` input primitives matching the dark-first design system. |
| FR-GET-06 | Client-side validation: Name is required; Phone is required (vietnamese format check); Email is required and valid. | Must Have | Uses schema-based validation (Zod/React Hook Form) with instant inline validation text. |
| FR-GET-07 | Submit button with loading state spinner. Disable double submissions. | Must Have | Button disabled during submit, loading indicator visible. |
| FR-GET-08 | Error banner display when mocked submission fails. | Should Have | User-friendly warning banner appears at the top of the card. |
| FR-GET-09 | Successful submission routes user to `/dashboard` using TanStack Router. | Must Have | Router transition occurs immediately upon mock success. |

---

## Success Criteria

| Criterion | Metric | Target |
| --- | --- | --- |
| Onboarding Form Submission | Form validation success | 100% block on invalid formats. Clear inline errors. |
| Speed of Routing | Time to transition to /dashboard | < 1.5 seconds after mock API response. |
| Keyboard Accessibility | Logical tab index order | TAB key cycles through fields/buttons sequentially with focus outlines. |
| Mobile Optimization | Single-column form flow | Zero horizontal overflow at 375px. Input target heights >= 44px. |

---

## Assumptions

- A placeholder `/dashboard` route is available or can be targeted in TanStack Router.
- Onboarding simulation uses local states without backend integration.
- Icons use `lucide-react`.

## Dependencies

- TanStack Router for route configuration.
- Tailwind CSS v4 and `@shopwise/ui` packages.

## Risks

| Risk | Likelihood | Impact | Mitigation |
| --- | --- | --- | --- |
| User loses form data on browser refresh | Low | Medium | Cache form progress in session storage if possible. |
| Navigation back button resets status state | Medium | Low | Ensure history stack is correctly managed. |

---

## Open Questions

- Should we cache the completed profile data in `localStorage` to simulate persistence across page reloads?
- Should the phone input format specifically support Vietnamese carrier prefixes?

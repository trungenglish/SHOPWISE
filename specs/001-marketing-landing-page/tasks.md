# tasks.md — Marketing Landing Page Repositioning

**Feature**: specs/001-marketing-landing-page **Plan**: specs/001-marketing-landing-page/plan.md **Generated**: 2026-07-11 **Total tasks**: 38

---

## Dependency Graph

```
Phase 1 (Audit & Mapping)
  └── Phase 2 (Nav, IDs, Data Updates)
        ├── Phase 3 (Hero & Trusted By)            ← [US1]
        ├── Phase 4 (The Search Crisis)             ← [US2]
        ├── Phase 5 (AI Decision Workspace)         ← [US3]
        ├── Phase 6 (Meet the Workspace)            ← [US4]
        ├── Phase 7 (Why ShopWise OS Verification)   ← [US5]
        ├── Phase 8 (Intelligence Layer)            ← [US6]
        └── Phase 9 (Confidence & Enterprise CTA)    ← [US7]
              └── Phase 10 (Remove Obsolete Sections)
                    └── Phase 11 (Visual Alignment)
                          └── Phase 12 (Validation & QA)
```

---

## Phase 1: Specification-to-code audit and component mapping

> Audit the existing v1.0.0 implementation components and map dependencies.

- [x] T001 [P] Audit existing components in `apps/web/src/features/landing/components/` against new v2.1.0 specifications
- [x] T002 [P] Inspect design token configurations and verify `--landing-*` variables in `packages/ui/src/styles/globals.css`
- [x] T003 [P] Map layout composition and dependencies in `apps/web/src/features/landing/LandingPage.tsx`

---

## Phase 2: Navigation, section IDs, content model, and static data updates

> Update static data schemas and rename navigation targets.

- [x] T004 [P] Refactor `FEATURES` static capabilities array in `apps/web/src/features/landing/data/features.ts` to define the 6 capabilities
- [x] T005 [P] Refactor `WORKFLOW_STEPS` static pipeline array in `apps/web/src/features/landing/data/workflow-steps.ts` to define the 7 stages
- [x] T006 Refactor navbar section mapping and `useScrollSpy` IDs in `apps/web/src/features/landing/components/navbar/LandingNavbar.tsx`
- [x] T007 [P] Refactor mobile nav anchors array in `apps/web/src/features/landing/components/navbar/MobileNavDrawer.tsx`
- [x] T008 [P] Refactor platform links and description text in `apps/web/src/features/landing/components/footer/LandingFooter.tsx`

---

## Phase 3: Hero and Trusted By alignment

> Re-align Hero copy/CTAs to enterprise positioning and make Trusted By optional.

- [x] T009 [US1] Update titles and change CTAs to "Launch Decision Workspace" / "See How ShopWise Reasons" in `apps/web/src/features/landing/components/hero/HeroSection.tsx`
- [x] T010 [P] [US1] Set `TRUSTED_BY_LOGOS` static array to empty `[]` by default for launch in `apps/web/src/features/landing/data/trusted-by.ts`
- [x] T011 [US1] Add a conditional length check to return `null` if the logo array is empty in `apps/web/src/features/landing/components/trusted-by/TrustedBySection.tsx`

---

## Phase 4: The Search Crisis section

> Create "The Search Crisis" problem comparison and differentiation workspace section.

- [x] T012 [US2] Create problem/solution comparison list layouts in `apps/web/src/features/landing/components/problem/ProblemStatementSection.tsx`
- [x] T013 [US2] Embed the chatbot/search engine differentiation statement at the bottom of `apps/web/src/features/landing/components/problem/ProblemStatementSection.tsx`
- [x] T014 [US2] Wire `<ProblemStatementSection />` into the page assembly `apps/web/src/features/landing/LandingPage.tsx`

---

## Phase 5: AI Decision Workspace

> Refactor the AI Decision Section to represent the non-interactive workspace panel.

- [x] T015 [US3] Rename headers and update description text to match the AI Decision Workspace specification in `apps/web/src/features/landing/components/ai-viz/AiDecisionSection.tsx`
- [x] T016 [P] [US3] Disable all click event handlers and configure decorative scroll-triggered animations in `apps/web/src/features/landing/components/ai-viz/AiDecisionPanel.tsx`

---

## Phase 6: Meet the Workspace

> Refactor the capability cards to represent the 6 workspace capabilities.

- [x] T017 [US4] Update component headers and map updated capability data in `apps/web/src/features/landing/components/features/FeaturesSection.tsx`

---

## Phase 7: Why ShopWise OS

> Verify the OS branding removal and structural merger.

- [x] T018 [US5] Verify that no standalone "Why ShopWise OS" section remains and its differentiation points are correctly merged into Phase 4 in `apps/web/src/features/landing/components/`

---

## Phase 8: Intelligence Layer

> Refactor the workflow section into a 7-stage sequential intelligence timeline.

- [x] T019 [US6] Rename headers and assign `id="intelligence"` in `apps/web/src/features/landing/components/workflow/WorkflowSection.tsx`
- [x] T020 [US6] Refactor badges, labels, and connector lines for 7 steps in `apps/web/src/features/landing/components/workflow/WorkflowTimeline.tsx`

---

## Phase 9: Decision Confidence and Final Enterprise CTA

> Add the Decision Confidence signals panel and refactor the final conversion CTA.

- [x] T021 [US7] Create the visual layout and copy for `apps/web/src/features/landing/components/confidence/DecisionConfidenceSection.tsx`
- [x] T022 [P] [US7] Create the decorative confidence signals mockup in `apps/web/src/features/landing/components/confidence/DecisionConfidencePanel.tsx`
- [x] T023 [US7] Refactor copy and change buttons to "Launch Decision Workspace" / "Request Enterprise Demo" in `apps/web/src/features/landing/components/final-cta/FinalCtaSection.tsx`
- [x] T024 [US7] Wire `<DecisionConfidenceSection />` into the page assembly `apps/web/src/features/landing/LandingPage.tsx`

---

## Phase 10: Remove Testimonials, Pricing, FAQ, obsolete data, and unused imports

> Delete obsolete testimonials, pricing, and FAQ assets to avoid unverified customer or pricing data in production.

- [x] T025 Remove stubs and rendering slots for `TestimonialsSection` in `apps/web/src/features/landing/LandingPage.tsx`
- [x] T026 Remove stubs and rendering slots for `PricingSection` in `apps/web/src/features/landing/LandingPage.tsx`
- [x] T027 Remove stubs and rendering slots for `FaqSection` in `apps/web/src/features/landing/LandingPage.tsx`
- [x] T028 Remove testimonials data file `apps/web/src/features/landing/data/testimonials.ts`
- [x] T029 Remove pricing data file `apps/web/src/features/landing/data/pricing.ts`
- [x] T030 Remove FAQ data file `apps/web/src/features/landing/data/faq.ts`
- [x] T031 Delete obsolete components directories `testimonials/`, `pricing/`, and `faq/` under `apps/web/src/features/landing/components/`

---

## Phase 11: Full-page visual alignment to the approved dark enterprise design

> Polish page-level typography, spacing, and dark-first color variables.

- [x] T032 Verify consistent margin, padding, and vertical spacing between all active sections in `apps/web/src/features/landing/LandingPage.tsx`
- [x] T033 Verify that all fonts and color specificity meet the premium dark-first aesthetic in `packages/ui/src/styles/globals.css`

---

## Phase 12: Accessibility, responsive, reduced-motion, type-check, test, lint, and build validation

> Confirm that the repositioned landing page passes all quality gates.

- [x] T034 Verify responsive layout reflows and tap targets at 375px, 768px, 1024px, and 1440px in `apps/web/src/features/landing/`
- [x] T035 Verify prefers-reduced-motion media query overrides and GSAP Snap-to-end setups in all visual components in `apps/web/src/features/landing/components/`
- [x] T036 Run TypeScript type safety compilation checks using `pnpm --filter web check-types`
- [x] T037 Run linter checks and apply biome/prettier formats using `pnpm dlx ultracite check`
- [x] T038 Execute monorepo production build command `pnpm build`

---

## Parallel Execution Examples

- **Parallel Stream A**: T004 (features data), T005 (workflow-steps data)
- **Parallel Stream B**: T007 (mobile drawer), T008 (footer layout), T010 (trusted logos configuration)
- **Parallel Stream C**: T016 (AI workspace panel details), T022 (decision confidence panel details)
- **Parallel Stream D**: T028 (delete testimonials data), T029 (delete pricing data), T030 (delete FAQ data)

---

## Implementation Strategy

1. **Incremental Updates**: Update static data models (Phase 2) first, ensuring no code references break.
2. **Sequential Refactoring**: Refactor sections one-by-one from top to bottom (Hero → Trusted By → Search Crisis → Workspace → Capabilities → Timeline → Confidence → CTA).
3. **No Breakage**: Do not delete stubs or folders for testimonials, pricing, and FAQ until Phase 10, when the final page layouts are fully wired and stable.
4. **Validation Gate**: Run linting, type-checks, and builds at the end of each phase checkpoint. /speckit-specify

Create a new feature specification for the ShopWise Get Started onboarding flow.

Target application: apps/web

Route:

- /get-started

Context

The ShopWise dashboard already exists and is outside the scope of this feature.

This feature only handles the entry flow before the user reaches the dashboard.

Primary flow

1. User navigates to /get-started.
2. The page asks: "Do you already have an account?"
3. If the user selects "Yes":
   - Show the existing-user sign-in step.
   - After successful sign-in or temporary mock completion, navigate to /dashboard.
4. If the user selects "No":
   - Guide the user through a short onboarding form:
     - Full name
     - Phone number
     - Email address
   - Validate each field.
   - Submit the form.
   - After successful submission or temporary mock completion, navigate to /dashboard.

Scope

Build only:

- /get-started route
- account-status choice step
- existing-user sign-in step or placeholder
- new-user onboarding form
- client-side validation
- loading state
- error state
- success redirect to /dashboard

Out of scope

- Dashboard implementation
- Dashboard redesign
- Backend authentication
- OTP verification
- Password reset
- Social login
- Database persistence
- Email verification
- Phone verification
- Real session management

Account Status Step

Question: "Do you already have an account?"

Actions:

- "Yes, sign in"
- "No, create my profile"

Requirements:

- Clear step heading
- Accessible button labels
- Keyboard navigation
- Back navigation where applicable
- Preserve entered form data when moving between steps if reasonable

Existing User Step

If the current project already has an authentication form or route:

- Reuse it.
- Do not duplicate authentication UI.

If no authentication UI exists:

- Provide a minimal placeholder step.
- Include email or phone input only if required by the approved design.
- Add a clearly marked temporary action that continues to /dashboard.
- Do not present mock authentication as production-ready.

New User Step

Fields:

- Full name
- Phone number
- Email address

Validation:

- Full name is required.
- Phone number is required and must use a reasonable format.
- Email is required and must be valid.
- Show inline validation messages.
- Disable duplicate submissions.
- Preserve accessibility with labels, descriptions, aria-invalid, and error associations.

Submission

- Show a loading state.
- Show a user-friendly error state if submission fails.
- On success, navigate to /dashboard using TanStack Router.
- Do not implement the dashboard.

Responsive Design

- Mobile-first
- Single-column form
- Comfortable input sizes
- Minimum 44px interactive targets
- Desktop layout may use a centered card or split layout
- Avoid unnecessary multi-column complexity

Visual Direction

- Match the existing ShopWise dark-first design system
- Premium
- Minimal
- Clear
- Low cognitive load
- Strong progress indication
- Do not copy the landing page layout directly

Architecture

- React 19
- Vite
- TypeScript
- TanStack Router
- Tailwind CSS v4
- @shopwise/ui
- React Hook Form and Zod may be used because they already exist in the project
- Feature code should live under: apps/web/src/features/get-started/

Suggested structure:

apps/web/src/features/get-started/ components/ schemas/ hooks/ GetStartedPage.tsx

Success Criteria

- /get-started renders correctly
- User can choose existing or new account flow
- New-user form validates name, phone, and email
- Successful completion navigates to /dashboard
- Dashboard code is not modified
- Type checking passes

Do not write implementation code. Create only the feature specification.

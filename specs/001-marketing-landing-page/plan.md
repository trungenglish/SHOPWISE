# Implementation Plan: Reposition ShopWise as Enterprise AI Decision Intelligence Platform

Reposition the ShopWise marketing landing page from a generic consumer/SaaS subscription surface to a premium, dark-first enterprise AI Decision Intelligence Platform for procurement teams. This requires removing pricing, testimonials, and FAQ sections, refactoring existing sections (Hero, Trusted By, Navbar, Workspace visual, Capabilities, Workflow steps, Final CTA, Footer), and adding new problem statement and decision confidence sections.

---

## User Review Required

> [!IMPORTANT] **Primary Conversion Target**: All free-trial and pricing conversion hooks (e.g. "Start for Free") are replaced with a single, clear enterprise action: "Launch Decision Workspace" (linking to the app's `/auth` route). The secondary CTA links to a sales request ("Request Enterprise Demo") which will temporarily use a mailto link.

> [!WARNING] **Fictional Testimonials and Pricing Removed**: The Testimonials, Pricing, and FAQ sections are fully removed from the landing page. Their data files and code assets are deleted to prevent fictional customer or pricing data from displaying in production.

> [!NOTE] **Conditional Trusted By Section**: The Trusted By section is refactored to be optional. If no verified, brand-approved logos are configured in `trusted-by.ts`, the section returns `null` and is omitted entirely from the page layout without breaking.

---

## Open Questions

> [!NOTE] None. All primary ambiguities (target audience hierarchy, OS branding usage, simulation interactivity, Trusted By conditional gating, and narrative order) have been fully resolved in the v2.1.0 specification clarifications.

---

## Proposed Changes

### Configuration and Data Model

#### [MODIFY] [trusted-by.ts](file:///c:/D/SHOPWISE/apps/web/src/features/landing/data/trusted-by.ts)

- Clean up placeholder logos. The array defaults to empty (`[]`) for production gating.

#### [MODIFY] [features.ts](file:///c:/D/SHOPWISE/apps/web/src/features/landing/data/features.ts)

- Update static list to represent the 6 workspace capabilities (Product & Supplier Canvas, Requirement Matching, Trust & Evidence Centre, Audit Trail, Team Collaboration, Enterprise Integrations).

#### [MODIFY] [workflow-steps.ts](file:///c:/D/SHOPWISE/apps/web/src/features/landing/data/workflow-steps.ts)

- Update static list to represent the 7 intelligence pipeline stages (Business Goal, Planning, Data & Evidence, Specialist Analysis, Comparison Logic, Recommendation, Human Approval).

#### [DELETE] [pricing.ts](file:///c:/D/SHOPWISE/apps/web/src/features/landing/data/pricing.ts)

- Delete pricing data file.

#### [DELETE] [testimonials.ts](file:///c:/D/SHOPWISE/apps/web/src/features/landing/data/testimonials.ts)

- Delete testimonials data file.

#### [DELETE] [faq.ts](file:///c:/D/SHOPWISE/apps/web/src/features/landing/data/faq.ts)

- Delete FAQ data file.

---

### Layout & Navigation Components

#### [MODIFY] [LandingNavbar.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/navbar/LandingNavbar.tsx)

- Update navigation links: Problem (`#problem`), Workspace (`#workspace`), Intelligence (`#intelligence`), Confidence (`#confidence`), Contact (`#contact`).
- Update `useScrollSpy` parameters: `sectionIds: ["problem", "workspace", "intelligence", "confidence", "contact"]`.
- Update CTA button labels and paths.

#### [MODIFY] [MobileNavDrawer.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/navbar/MobileNavDrawer.tsx)

- Align links array and CTAs with updated navbar items.

#### [MODIFY] [LandingFooter.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/footer/LandingFooter.tsx)

- Update footer descriptor to: "AI Decision Intelligence for Enterprise Procurement."
- Align product links with new page anchors.
- Update copyright notice current year logic if required.

---

### Hero & Credibility Components

#### [MODIFY] [HeroSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/hero/HeroSection.tsx)

- Update subtitle and headline to reflect enterprise positioning (From Fragmented Research to Confident Purchasing Decisions).
- Change primary CTA: "Launch Decision Workspace" (`/auth`).
- Change secondary CTA: "See How ShopWise Reasons" (scrolls to `#problem`).

#### [MODIFY] [TrustedBySection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/trusted-by/TrustedBySection.tsx)

- Refactor to conditionally render: return `null` if `TRUSTED_BY_LOGOS` is empty.

---

### Problem & Demonstration Components

#### [NEW] [ProblemStatementSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/problem/ProblemStatementSection.tsx)

- Implement "The Search Crisis" section.
- Layout: split layout with Legacy Approach (left, muted) vs ShopWise Approach (right, highlighted) using 4 paired procurement pain points and outcomes.
- Include the merged product differentiation copy at the bottom of the section (Not a search engine, chatbot, or marketplace — but an AI Decision Intelligence Platform).
- Assign `id="problem"`.

#### [MODIFY] [AiDecisionSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/ai-viz/AiDecisionSection.tsx)

- Refactor copy to align with "AI Decision Workspace" spec.
- Assign `id="workspace"`.

#### [MODIFY] [AiDecisionPanel.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/ai-viz/AiDecisionPanel.tsx)

- Ensure visual mockup is fully decorative. Ensure all scores and indicator bars populate via scroll-triggered GSAP timelines once per entry.
- Verify clicks are ignored; hover triggers CSS-only visual highlight.

#### [MODIFY] [FeaturesSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/features/FeaturesSection.tsx)

- Refactor copy to represent "Meet the Workspace" capabilities.
- Renders the 6 capabilities from updated `features.ts`.
- Keep `id="features"` internally (not explicitly in navbar, but used for layout consistency).

---

### Intelligence & Executive Components

#### [MODIFY] [WorkflowSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/workflow/WorkflowSection.tsx)

- Refactor timeline to show 7 sequential stages from `workflow-steps.ts`.
- Assign `id="intelligence"`.

#### [MODIFY] [WorkflowTimeline.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/workflow/WorkflowTimeline.tsx)

- Reframe badge numbers and connections to support 7 stages in horizontal desktop and vertical mobile flows.

#### [NEW] [DecisionConfidenceSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/confidence/DecisionConfidenceSection.tsx)

- Create "Decision Confidence" section.
- Focus on secondary executive audience: explain auditability, risk mitigation, and justified ROI.
- Render `<DecisionConfidencePanel />`.
- Assign `id="confidence"`.

#### [NEW] [DecisionConfidencePanel.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/confidence/DecisionConfidencePanel.tsx)

- Implement the decorative simulation panel displaying 4 signals: Evidence Status, Confidence Indicator, Risk Status, and Auditability Signal.
- GSAP timeline animates signals sequentially on scroll entry. Snaps immediately to completed final state on prefers-reduced-motion.

---

### Final CTA & Page Assembly

#### [MODIFY] [FinalCtaSection.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/final-cta/FinalCtaSection.tsx)

- Refactor to match "Enterprise CTA" spec.
- Headline: "Start Making Confident Purchasing Decisions."
- Primary CTA: "Launch Decision Workspace" (navigates to `/auth`).
- Secondary CTA: "Request Enterprise Demo" (mailto placeholder link).
- Remove pricing and trial copy.
- Assign `id="contact"`.

#### [MODIFY] [LandingPage.tsx](file:///c:/D/SHOPWISE/apps/web/src/features/landing/LandingPage.tsx)

- Remove imports and rendering slots for: `TestimonialsSection`, `PricingSection`, `FaqSection`.
- Import and render: `ProblemStatementSection` (after `TrustedBySection`), `DecisionConfidenceSection` (before `FinalCtaSection`).

#### [DELETE] [testimonials/](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/testimonials/)

- Delete testimonials components folder.

#### [DELETE] [pricing/](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/pricing/)

- Delete pricing components folder.

#### [DELETE] [faq/](file:///c:/D/SHOPWISE/apps/web/src/features/landing/components/faq/)

- Delete FAQ components folder.

---

## Verification Plan

### Automated Tests

- Run check-types to verify TS compilation:
  ```bash
  pnpm --filter web check-types
  ```
- Run linter to verify code standards:
  ```bash
  pnpm dlx ultracite check
  ```
- Run monorepo build to confirm page builds cleanly:
  ```bash
  pnpm build
  ```

### Manual Verification

- Visual inspection at 1440px viewport to confirm the new 8 sections render in sequence.
- Inspect at 375px viewport to confirm zero horizontal scroll, single-column reflow, and mobile nav drawer links.
- Test keyboard navigation (Tab/Shift+Tab) to ensure focus ring is visible on all interactive elements.
- Toggle prefers-reduced-motion in browser to verify that all animations cease and visual panels immediately display in their complete final states.

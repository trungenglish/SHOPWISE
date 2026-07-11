# UI Contracts: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page **Created**: 2026-07-10 **Last Updated**: 2026-07-11

> These contracts define the public interface (props) of each custom component built for this feature. They are the implementation contract between the page assembly (`LandingPage.tsx`) and the individual section components. All components live in `apps/web/src/features/landing/components/`.

---

## `LandingPage`

**File**: `apps/web/src/features/landing/LandingPage.tsx`  
**Role**: Page root — composes all sections in order. Mounted by the index route.

```ts
// No props — self-contained page root
type LandingPageProps = Record<string, never>;
```

**Composition order**:

1. `<LandingNavbar />`
2. `<HeroSection />`
3. `<TrustedBySection />` (Optional)
4. `<ProblemStatementSection />` [NEW]
5. `<AiDecisionSection />` (Refactored to "AI Decision Workspace")
6. `<FeaturesSection />` (Refactored to "Meet the Workspace")
7. `<WorkflowSection />` (Refactored to "Intelligence Layer")
8. `<DecisionConfidenceSection />` [NEW]
9. `<FinalCtaSection />` (Refactored to "Enterprise CTA")
10. `<LandingFooter />`

---

## `LandingNavbar`

**File**: `components/navbar/LandingNavbar.tsx`

```ts
// No external props
type LandingNavbarProps = Record<string, never>;
```

**Internal state**: `isScrolled: boolean`, `isMobileMenuOpen: boolean`  
**Updates**: Nav links are updated to:

- Problem (`#problem`)
- Workspace (`#workspace`)
- Intelligence (`#intelligence`)
- Confidence (`#confidence`)
- Contact (`#contact` - scrolls to Final CTA section) Active section is tracked via `useScrollSpy` with `sectionIds: ["problem", "workspace", "intelligence", "confidence", "contact"]`.

---

## `MobileNavDrawer`

**File**: `components/navbar/MobileNavDrawer.tsx`

```ts
type MobileNavDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
};
```

**Updates**: Links are updated to map the new navbar anchors.

---

## `HeroSection`

**File**: `components/hero/HeroSection.tsx`

```ts
// No external props
type HeroSectionProps = Record<string, never>;
```

**Updates**:

- Headline communicates positioning: "From Fragmented Research to Confident Purchasing Decisions."
- Primary CTA: "Launch Decision Workspace" (navigates to `/auth`).
- Secondary CTA: "See How ShopWise Reasons" (scrolls to `#problem` or `#intelligence`).

---

## `TrustedBySection`

**File**: `components/trusted-by/TrustedBySection.tsx`

```ts
// No external props
type TrustedBySectionProps = Record<string, never>;
```

**Updates**:

- Reads `TRUSTED_BY_LOGOS`.
- If `TRUSTED_BY_LOGOS` is empty (`[]`), returns `null` immediately.

---

## `ProblemStatementSection` [NEW]

**File**: `components/problem/ProblemStatementSection.tsx`

```ts
// No external props
type ProblemStatementSectionProps = Record<string, never>;
```

**Role**: Implements "The Search Crisis" (section 4 of spec). Displays two columns (Legacy approach vs ShopWise approach) plus the merged chatbot/search engine differentiation statement.  
**Layout**: Two columns on desktop, single-column vertical stack on mobile.

---

## `AiDecisionSection` (AI Decision Workspace)

**File**: `components/ai-viz/AiDecisionSection.tsx`

```ts
// No external props
type AiDecisionSectionProps = Record<string, never>;
```

**Updates**:

- Renamed internally/refactored to "AI Decision Workspace" section.
- Section ID updated to `id="workspace"`.
- Headline: "One Workspace. Every Signal. One Decision."

---

## `AiDecisionPanel`

**File**: `components/ai-viz/AiDecisionPanel.tsx`

```ts
// No external props
type AiDecisionPanelProps = Record<string, never>;
```

**Updates**:

- Purely decorative simulation.
- Pulsing live analysis indicator and score bars animating width 0 → final value on scroll trigger.
- Click events are fully disabled. No mock state machine.

---

## `FeaturesSection` (Meet the Workspace)

**File**: `components/features/FeaturesSection.tsx`

```ts
// No external props
type FeaturesSectionProps = Record<string, never>;
```

**Updates**:

- Renamed internally/refactored to "Meet the Workspace".
- Headline: "Meet the Workspace" + capability cards mapped from features list.
- Component ID: `id="features"`.

---

## `WorkflowSection` (Intelligence Layer)

**File**: `components/workflow/WorkflowSection.tsx`

```ts
// No external props
type WorkflowSectionProps = Record<string, never>;
```

**Updates**:

- Renamed internally/refactored to "Intelligence Layer".
- Section ID updated to `id="intelligence"`.
- Headline: "Intelligence Layer" + sequential timeline.

---

## `WorkflowTimeline`

**File**: `components/workflow/WorkflowTimeline.tsx`

```ts
type WorkflowTimelineProps = {
  steps: WorkflowStep[];
};
```

**Updates**:

- Renders exactly 7 stages (Business Goal, Planning, Data and Evidence, Specialist Analysis, Comparison Logic, Recommendation, Human Approval).
- Animates connecting lines and step badges on scroll entry.

---

## `DecisionConfidenceSection` [NEW]

**File**: `components/confidence/DecisionConfidenceSection.tsx`

```ts
// No external props
type DecisionConfidenceSectionProps = Record<string, never>;
```

**Role**: Implements "Decision Confidence" (section 8 of spec). Focuses on secondary executive audience (audit trail, governance, ROI). Renders `<DecisionConfidencePanel />`.  
**Layout**: Two columns on desktop, single-column stack on mobile. Section ID: `id="confidence"`.

---

## `DecisionConfidencePanel` [NEW]

**File**: `components/confidence/DecisionConfidencePanel.tsx`

```ts
// No external props
type DecisionConfidencePanelProps = Record<string, never>;
```

**Role**: Illustrates a confident explainable recommendation with 4 signals: Evidence Status, Confidence Indicator, Risk Status, and Auditability Signal.  
**UX Behaviour**: GSAP ScrollTrigger timeline animates each signal sequentially. No interactivity. Under prefers-reduced-motion, snaps immediately to final state.

---

## `FinalCtaSection` (Enterprise CTA)

**File**: `components/final-cta/FinalCtaSection.tsx`

```ts
// No external props
type FinalCtaSectionProps = Record<string, never>;
```

**Updates**:

- Renamed internally/refactored to "Enterprise CTA".
- Section ID updated to `id="contact"`.
- Headline: "Start Making Confident Purchasing Decisions."
- Primary CTA: "Launch Decision Workspace" (navigates to `/auth`).
- Secondary CTA: "Request Enterprise Demo" / "Talk to Sales" (mailto link or anchor).
- Pricing/trial terms removed.

---

## `LandingFooter`

**File**: `components/footer/LandingFooter.tsx`

```ts
// No external props
type LandingFooterProps = Record<string, never>;
```

**Updates**:

- Updates platform links to match new section anchors.
- Updates copyright notice.

---

## Removed Components

The following components are **fully removed** from the codebase:

- `TestimonialsSection.tsx` (`components/testimonials/`)
- `PricingSection.tsx`, `PricingCard.tsx`, `BillingToggle.tsx` (`components/pricing/`)
- `FaqSection.tsx`, `FaqItem.tsx` (`components/faq/`)

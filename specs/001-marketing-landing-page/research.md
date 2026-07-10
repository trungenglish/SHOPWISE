# Research: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page
**Completed**: 2026-07-10
**Status**: All NEEDS CLARIFICATION resolved

---

## Decision Log

### D-01: Route placement

**Decision**: The landing page replaces the placeholder at `apps/web/src/routes/index.tsx` (route `/`).
**Rationale**: The index route is already registered in the TanStack Router tree and rendered by `LandingPageComponent`. The existing placeholder content will be fully replaced. No new route file is needed.
**Alternatives considered**: A dedicated `/landing` route. Rejected — index (`/`) is the canonical public entry point for a marketing page and avoids a redirect.

---

### D-02: Lenis smooth scroll

**Decision**: Lenis IS available and can be used. Mount `<Lenis>` inside `LandingLayout` (or inside the landing route's component) rather than at the app root, so smooth scroll is scoped to the public landing surface and does not affect authenticated app routes.
**Rationale**: `apps/web/src/components/lenis/index.tsx` exports a `<Lenis>` wrapper that uses `ReactLenis` with `root` prop. It is not currently mounted anywhere. Mounting it in `LandingLayout` is safe and constitutionally correct ("use Lenis only if already configured").
**Alternatives considered**: Mounting at `__root.tsx`. Rejected — Lenis's `root` prop replaces the native scroll container globally; this could break app routes that don't want smooth scroll.

---

### D-03: CTA route paths

**Decision**: Use the following routes (confirmed from `landing-layout.tsx` and route tree):
- "Sign In" → `/auth`
- "Get Started" / "Start for Free" → `/auth` (registration flow, same auth route with a mode param or redirect)
- "AI Workspace" → `/auth` (post-login redirect; no dedicated workspace route exists yet — link to `/auth` with a `redirect` search param intention)
- "Talk to Sales" → external `mailto:sales@shopwise.io` or a `#contact` anchor (placeholder)
**Rationale**: Only `/` and `/auth` exist in the route tree. Workspace route is a future feature.
**Assumption documented**: Final CTA destinations will be updated when the workspace route is created.

---

### D-04: Design tokens — new indigo/violet gradient system

**Decision**: Add landing-page-specific CSS tokens to `packages/ui/src/styles/globals.css` rather than `apps/web/src/index.css`, following the Constitution's requirement that HSL/OKLCH tokens live in the shared styles file.

New tokens to add (scoped to `.dark` and optionally `:root`):
- `--landing-gradient-start`: deep navy-indigo (e.g., `oklch(12% 0.03 260deg)`)
- `--landing-gradient-mid`: indigo (`oklch(30% 0.12 265deg)`)
- `--landing-gradient-accent`: violet (`oklch(55% 0.18 290deg)`)
- `--landing-glow`: soft violet glow for CTA pulse (`oklch(55% 0.18 290deg / 25%)`)
- `--landing-glass-bg`: `oklch(100% 0 0deg / 6%)`
- `--landing-glass-border`: `oklch(100% 0 0deg / 12%)`

**Rationale**: Keeps token definition in the canonical location. Components reference them via CSS variables, never hard-coded.

---

### D-05: @shopwise/ui components to reuse

| Component | Used in |
|---|---|
| `Button` | Navbar CTA, Hero CTAs, Pricing CTAs, Final CTA |
| `Card` | Feature cards, Pricing cards |
| `Badge` | "Most Popular" badge on Pricing |
| `Carousel` | Testimonials section |
| `Avatar` | Testimonial author avatar |
| `Sheet` | Mobile nav drawer |
| `Separator` | Footer divider |
| `collapsible.tsx` (Collapsible primitives) | FAQ accordion items |
| `NavigationMenu` | Desktop nav links |

Components that do NOT exist in `@shopwise/ui` and must be built in `features/landing/`:
- `HeroVisual` — animated AI data graphic
- `AiDecisionPanel` — the AI visualisation card
- `WorkflowTimeline` — step progress layout
- `TrustedByMarquee` — logo marquee
- `PricingToggle` — billing period switch

---

### D-06: GSAP usage pattern

**Decision**: Use `useGSAP` from `@gsap/react` for all GSAP animations. Register `ScrollTrigger` plugin once in a `gsap-config.ts` utility module at `apps/web/src/lib/gsap-config.ts`. Components import from there to guarantee single registration.
**Rationale**: GSAP plugin registration must be idempotent. Centralising it avoids double-registration bugs across components.
**Reduced-motion guard**: All `useGSAP` hooks check `window.matchMedia('(prefers-reduced-motion: reduce)').matches` before creating timelines. If true, elements are set to their end-state immediately via `gsap.set()`.

---

### D-07: Typed static data files

**Decision**: All repeated content (features list, FAQ items, pricing tiers, testimonials, workflow steps, trusted-by logos) lives in typed `.ts` data files under `apps/web/src/features/landing/data/`.
**Rationale**: Separates content from structure, makes copy updates trivial, and ensures TypeScript enforces shape consistency.
**File structure**:
```
apps/web/src/features/landing/data/
  features.ts
  faq.ts
  pricing.ts
  testimonials.ts
  workflow-steps.ts
  trusted-by.ts
```

---

### D-08: Component file structure

**Decision**: All landing-page-specific components are scoped to `apps/web/src/features/landing/`. The route file (`apps/web/src/routes/index.tsx`) imports only the top-level page component.

```
apps/web/src/features/landing/
  components/
    navbar/
      LandingNavbar.tsx
      MobileNavDrawer.tsx
    hero/
      HeroSection.tsx
      HeroVisual.tsx
    trusted-by/
      TrustedBySection.tsx
      LogoMarquee.tsx
    ai-viz/
      AiDecisionSection.tsx
      AiDecisionPanel.tsx
    features/
      FeaturesSection.tsx
      FeatureCard.tsx
    workflow/
      WorkflowSection.tsx
      WorkflowTimeline.tsx
    testimonials/
      TestimonialsSection.tsx
    pricing/
      PricingSection.tsx
      PricingCard.tsx
      BillingToggle.tsx
    faq/
      FaqSection.tsx
      FaqItem.tsx
    final-cta/
      FinalCtaSection.tsx
    footer/
      LandingFooter.tsx
  data/
    features.ts
    faq.ts
    pricing.ts
    testimonials.ts
    workflow-steps.ts
    trusted-by.ts
  hooks/
    useReducedMotion.ts
    useScrollSpy.ts
  LandingPage.tsx         ← page root, composes all sections
```

---

### D-09: LandingLayout disposition

**Decision**: The existing `LandingLayout` in `apps/web/src/components/layouts/landing-layout.tsx` will be **replaced** by the new `LandingNavbar` and `LandingFooter` components composed directly inside `LandingPage.tsx`. The old layout contains a minimal placeholder navbar/footer that does not meet the spec.
**Rationale**: The old layout was scaffold-grade; the new components are spec-grade. Removing the wrapper keeps responsibility explicit and avoids a layout nesting mismatch.
**Impact**: `apps/web/src/routes/index.tsx` is updated to render `<LandingPage />` directly without `<LandingLayout>`.

---

### D-10: CSS theme — current palette mismatch

**Decision**: The current `--primary` token (green/teal) is NOT used on the landing page gradient surfaces. Landing page sections use the new `--landing-*` tokens exclusively for gradient/glow effects. The `@shopwise/ui` Button component will use a `variant="landing"` or we pass the `className` override approach to style buttons in indigo/violet on the landing page without modifying the shared Button component.

**Rationale**: Modifying `--primary` globally would break the app dashboard. The landing page uses scoped CSS tokens.

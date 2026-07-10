# Implementation Plan: ShopWise Marketing Landing Page

**Feature**: specs/001-marketing-landing-page
**Version**: 1.0.0
**Created**: 2026-07-10
**Status**: Ready for implementation

---

## Technical Context

| Item | Resolved Value |
|---|---|
| Target workspace | `apps/web` |
| Route | `/` — replaces `apps/web/src/routes/index.tsx` placeholder |
| Framework | React 19 + Vite + TypeScript (strict mode) |
| Styling | Tailwind CSS v4, tokens in `packages/ui/src/styles/globals.css` |
| Shared UI | `@shopwise/ui` (Button, Card, Badge, Carousel, Avatar, Sheet, Collapsible, Separator, Toggle) |
| Routing | TanStack Router file-based (`apps/web/src/routes/index.tsx`) |
| Animation | GSAP + `@gsap/react`, registered in `apps/web/src/lib/gsap-config.ts` |
| Smooth scroll | Lenis — available, mount inside `LandingPage` only |
| Icons | `lucide-react` |
| Existing layout | `LandingLayout` replaced by new `LandingNavbar` + `LandingFooter` |
| Auth route | `/auth` (Sign In + Get Started destination) |
| No SSR | CSR only; no Next.js patterns |
| Lenis status | Component exists, **not yet mounted** — must be added to `LandingPage` |

---

## Constitution Check

| Principle | Status | Notes |
|---|---|---|
| Component-First Architecture | ✅ | Each section is a standalone component; logic extracted to hooks |
| Shared UI Package First | ✅ | 9 @shopwise/ui components reused; custom components only where no shared equivalent exists |
| Feature-Based Organisation | ✅ | All code in `apps/web/src/features/landing/` |
| Accessibility First | ✅ | WCAG 2.1 AA required; ARIA roles, keyboard nav, reduced-motion — see VS-06 to VS-10 |
| Responsive First | ✅ | Mobile-first Tailwind breakpoints; 375 px minimum tested |
| Performance First | ✅ | Route-level code split; GSAP loaded only for landing route; logo images use `loading="lazy"` |
| Type-Safe Code | ✅ | All data files are typed; `any` is forbidden; `pnpm --filter web check-types` is a gate |
| Composition Over Abstraction | ✅ | No HOCs or generic abstractions; each component has one job; data files ≥ 2 consumers |

**Routing rules**: TanStack Router only. No Next.js patterns. ✅
**Styling rules**: Tailwind v4, @shopwise/ui first, lucide-react, dark-first tokens. ✅
**Motion rules**: GSAP for complex timelines, CSS for hover, Lenis scoped to landing, prefers-reduced-motion respected. ✅
**Quality gates**: `check-types` + `ultracite check` + `vitest` run before merge. ✅

---

## Architecture Overview

```
apps/web/src/
  lib/
    gsap-config.ts                  ← GSAP plugin registration (new)
  features/
    landing/
      LandingPage.tsx               ← Page root (new)
      data/
        features.ts                 ← FeatureItem[]
        faq.ts                      ← FaqItem[]
        pricing.ts                  ← PricingTier[]
        testimonials.ts             ← Testimonial[]
        workflow-steps.ts           ← WorkflowStep[]
        trusted-by.ts               ← TrustedByLogo[]
      hooks/
        useReducedMotion.ts         ← prefers-reduced-motion hook
        useScrollSpy.ts             ← active section detection
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
  routes/
    index.tsx                       ← Updated to render <LandingPage />

packages/ui/src/styles/
  globals.css                       ← New --landing-* CSS tokens added
```

---

## Milestones

### Milestone 0 — Foundation (Prerequisite for all other milestones)

**Goal**: Establish the infrastructure all section components depend on. Nothing renders without this.

**Tasks**:

**M0-T1**: Add landing-page design tokens to `packages/ui/src/styles/globals.css`

Add to `.dark` block (and optional `:root` fallbacks):
```css
--landing-gradient-start: oklch(10% 0.025 260deg);
--landing-gradient-mid: oklch(18% 0.06 265deg);
--landing-gradient-accent: oklch(55% 0.2 290deg);
--landing-glow: oklch(55% 0.2 290deg / 20%);
--landing-glass-bg: oklch(100% 0 0deg / 6%);
--landing-glass-border: oklch(100% 0 0deg / 10%);
```

**M0-T2**: Create `apps/web/src/lib/gsap-config.ts`
- Import `gsap` from `gsap`
- Import `ScrollTrigger` from `gsap/ScrollTrigger`
- Call `gsap.registerPlugin(ScrollTrigger)`
- Export `{ gsap, ScrollTrigger }`
- All GSAP-using components import from this file.

**M0-T3**: Create `useReducedMotion` hook
- File: `apps/web/src/features/landing/hooks/useReducedMotion.ts`
- Wraps `window.matchMedia('(prefers-reduced-motion: reduce)')` with a `useState` + `useEffect` listener.
- Returns `boolean`.

**M0-T4**: Create `useScrollSpy` hook
- File: `apps/web/src/features/landing/hooks/useScrollSpy.ts`
- Accepts `{ sectionIds: string[]; offset?: number }`.
- Uses `IntersectionObserver` to return the `id` of the currently visible section.

**M0-T5**: Create all typed data files (empty arrays as starting point):
- `data/features.ts` — exports `FEATURES: FeatureItem[]`
- `data/faq.ts` — exports `FAQ_ITEMS: FaqItem[]`
- `data/pricing.ts` — exports `PRICING_TIERS: PricingTier[]`
- `data/testimonials.ts` — exports `TESTIMONIALS: Testimonial[]`
- `data/workflow-steps.ts` — exports `WORKFLOW_STEPS: WorkflowStep[]`
- `data/trusted-by.ts` — exports `TRUSTED_BY_LOGOS: TrustedByLogo[]`
- Each file includes the full type definition and populated placeholder content.

**M0-T6**: Create `LandingPage.tsx` (scaffold only — no sections yet)
- Mounts `<ReactLenis root>` from `lenis/react`.
- Renders a single `<main>` with section slot placeholders (comments).

**M0-T7**: Update `apps/web/src/routes/index.tsx`
- Replace `LandingPageComponent` body to render `<LandingPage />`.
- Remove the old `LandingLayout` import.
- Update `head()` meta: title "ShopWise — AI-Powered Procurement Intelligence", description updated.

**Verification**: `pnpm --filter web check-types` passes. Dev server loads `/` without errors.

---

### Milestone 1 — Navigation (Navbar + Footer)

**Goal**: Persistent, accessible, responsive navigation surfaces that bookend the page. Implements VS-02, VS-03, VS-04, VS-05.

**Dependencies**: Milestone 0 complete.

**Tasks**:

**M1-T1**: `LandingNavbar.tsx`
- Sticky header with `position: sticky; top: 0; z-50`.
- `isScrolled` state: `useEffect` listening to `scroll` event; sets true when `scrollY > 100`.
- When `isScrolled`: apply `bg-background/80 backdrop-blur-md border-b border-border/40` via conditional class.
- Left: ShopWise wordmark as `<Link to="/">`.
- Centre (desktop, hidden on mobile): nav links to section anchors (`#features`, `#workflow`, `#pricing`, `#faq`). Active link highlighted via `useScrollSpy`.
- Right: "Sign In" ghost `Button` + "Get Started" filled `Button`, both `asChild` wrapping `<Link to="/auth">`.
- Mobile: `isMobileMenuOpen` state drives `<MobileNavDrawer>` via `<Sheet>`.
- Hamburger button: `aria-expanded={isMobileMenuOpen}`, `aria-controls="mobile-nav"`.
- Wraps in `<nav aria-label="Main navigation">`.

**M1-T2**: `MobileNavDrawer.tsx`
- Uses `@shopwise/ui` `Sheet` + `SheetContent` (side="right").
- Lists all nav links as `<a>` anchor links.
- Each link calls `onClose` on click.
- "Get Started" CTA inside drawer.
- `Sheet` handles focus trap and Escape key natively.

**M1-T3**: `LandingFooter.tsx`
- `<footer>` landmark element.
- Multi-column desktop layout: brand col + 3 link group cols.
- Link groups: Product, Company, Legal (inline `FooterLinkGroup[]` data).
- Social links (LinkedIn, GitHub): `target="_blank" rel="noopener noreferrer"`, `<span className="sr-only">ShopWise on LinkedIn</span>`.
- Copyright: `© {new Date().getFullYear()} ShopWise. All rights reserved.`
- `<Separator />` from `@shopwise/ui` above copyright row.
- Single-column stack on mobile.

**M1-T4**: Wire Navbar and Footer into `LandingPage.tsx`.

**Verification**: VS-02, VS-03, VS-04, VS-05 pass. VS-10 (keyboard nav) passes for Navbar and Footer elements.

---

### Milestone 2 — Hero Section

**Goal**: The first-impression section — full viewport height, gradient background, headline, CTAs, and animated visual. Implements VS-01 (partial), VS-06 (reduced-motion), VS-09 (reduced-motion).

**Dependencies**: Milestone 0 complete.

**Tasks**:

**M2-T1**: `HeroSection.tsx`
- `<section id="hero">` with `min-h-dvh` and gradient background using `--landing-gradient-start`, `--landing-gradient-mid` via inline CSS variables or a custom Tailwind class.
- Two-column layout on ≥ 1024 px (content 55%, visual 45%); single column on mobile.
- Content column:
  - Eyebrow `<Badge>` (outline variant): "AI-Powered Procurement"
  - `<h1>` headline: "Make Smarter Buying Decisions, Powered by AI."
  - `<p>` subtitle.
  - Two `<Button>` elements:
    - Primary (filled): "Start for Free" → `asChild <Link to="/auth">`
    - Secondary (outline): "See How It Works" → anchor scroll to `#ai-decision`
- Visual column: `<HeroVisual />` (hidden on mobile with `hidden lg:block`).

**M2-T2**: `HeroVisual.tsx`
- `aria-hidden="true"` — purely decorative.
- Abstract SVG or CSS composition: floating geometric shapes in indigo/violet gradients.
- Subtle CSS keyframe float animation (translate-y loop, 4 s ease-in-out infinite).
- Animation `animation-play-state: paused` when `prefers-reduced-motion: reduce`.

**M2-T3**: GSAP entry animation in `HeroSection.tsx`
- `useGSAP` hook.
- Check `useReducedMotion()` — if true, `gsap.set()` all targets to end-state and return early.
- Timeline: eyebrow badge → h1 words (staggered by 0.05 s) → subtitle → buttons (stagger).
- Each: `opacity: 0 → 1`, `y: 30 → 0`, `duration: 0.6`, `ease: 'power2.out'`.

**Verification**: VS-01 (hero value proposition legible), VS-09 (no animation under reduced-motion). `<h1>` is the only `<h1>` on the page.

---

### Milestone 3 — Trusted By Section

**Goal**: Credibility band of company logos. Implements VS-01 (trust signals visible early).

**Dependencies**: Milestone 0, trusted-by data file complete.

**Tasks**:

**M3-T1**: Acquire/create placeholder SVG logo assets (6 placeholder brands).
Place in `apps/web/src/assets/logos/`. Monochrome SVGs preferred.

**M3-T2**: Populate `data/trusted-by.ts` with 6 placeholder entries.

**M3-T3**: `TrustedBySection.tsx`
- `<section id="trusted-by">` — compact vertical padding.
- Muted label: "Trusted by leading enterprises worldwide".
- Renders `<LogoMarquee logos={TRUSTED_BY_LOGOS} />` on desktop.
- On mobile: static 2-column grid of logos (`grid-cols-2 sm:grid-cols-3`).

**M3-T4**: `LogoMarquee.tsx`
- CSS `@keyframes` scroll animation (translate-x: 0 → -50%, looping) on a doubled logo list for seamless loop.
- `pauseOnHover`: adds `[animation-play-state:paused]` class on `hover:` modifier.
- `@media (prefers-reduced-motion: reduce)`: `animation: none` — static row.
- `role="region"`, `aria-label="Trusted by — company logos"`.
- Each logo `<img>` has `alt={logo.name}`, `loading="lazy"`, explicit `width` and `height`.

**Verification**: No horizontal overflow. Logo alt texts announced by screen reader. Marquee pauses on hover.

---

### Milestone 4 — AI Decision Visualisation Section

**Goal**: The product differentiator made tangible. Implements VS-02.

**Dependencies**: Milestone 0 complete.

**Tasks**:

**M4-T1**: `AiDecisionSection.tsx`
- `<section id="ai-decision">`.
- Two-column on desktop (copy left, panel right), single column on mobile.
- Copy column: eyebrow → `<h2>` headline → supporting paragraph → "Try the AI Workspace" `<Button>` (link to `/auth`).
- Panel column: `<AiDecisionPanel />`.
- `<p className="sr-only">` describing the visualisation for screen readers.

**M4-T2**: `AiDecisionPanel.tsx`
- `aria-hidden="true"`.
- Styled as a glass card: `bg-[var(--landing-glass-bg)] border border-[var(--landing-glass-border)] backdrop-blur-sm rounded-xl`.
- Inner layout simulates an AI workspace:
  - Input row: a search/query string that types in via GSAP.
  - Processing indicator: pulsing ring animation (CSS).
  - Output list: 3 product/vendor rows with score bars that animate width 0 → final value.
- `useGSAP` with ScrollTrigger: timeline plays when panel enters viewport.
- `useReducedMotion()`: if true, display completed/output state statically.

**M4-T3**: GSAP ScrollTrigger wiring on `AiDecisionSection`
- Panel and copy column scroll-reveal (opacity 0 → 1, y 40 → 0) when section enters viewport.
- `start: 'top 75%'`.

**Verification**: VS-02 (AI viz self-explanatory). Panel is `aria-hidden`. `sr-only` description present. Reduced-motion shows static output state.

---

### Milestone 5 — Features Section

**Goal**: Scannable product capability grid. Implements VS-01 (feature evaluation).

**Dependencies**: Milestone 0, features data file complete.

**Tasks**:

**M5-T1**: Populate `data/features.ts` with 6 `FeatureItem` entries (see spec section 5 for proposed list). Assign a `LucideIcon` to each.

**M5-T2**: `FeaturesSection.tsx`
- `<section id="features">`.
- Centred `<h2>` + optional sub-headline.
- `<ul>` with CSS grid: `grid-cols-1 sm:grid-cols-2 lg:grid-cols-3`.
- Maps `FEATURES` to `<FeatureCard item={f} key={f.id} />`.
- GSAP ScrollTrigger stagger on card reveal: `stagger: 0.08`, `opacity 0→1`, `y 30→0`.
- Reduced-motion: all cards visible statically.

**M5-T3**: `FeatureCard.tsx`
- Renders as `<li>`.
- Uses `@shopwise/ui` `Card`, `CardContent`.
- Custom glass surface override: `bg-[var(--landing-glass-bg)] border-[var(--landing-glass-border)]`.
- Icon: `<item.icon className="..." aria-hidden="true" />`.
- `<h3>` title + `<p>` description.
- Hover: CSS `transition-shadow` + `transition-colors` for border glow. No GSAP.

**Verification**: Grid reflows at all breakpoints. Icons are `aria-hidden`. Hover state is CSS-only. Reduced-motion shows all cards statically.

---

### Milestone 6 — Workflow Section

**Goal**: Sequential "how it works" steps. Implements VS-01 (onboarding confidence).

**Dependencies**: Milestone 0, workflow-steps data file complete.

**Tasks**:

**M6-T1**: Populate `data/workflow-steps.ts` with 4 `WorkflowStep` entries.

**M6-T2**: `WorkflowSection.tsx`
- `<section id="workflow">`.
- `<h2>` headline + sub-headline.
- Renders `<WorkflowTimeline steps={WORKFLOW_STEPS} />`.

**M6-T3**: `WorkflowTimeline.tsx`
- Desktop (≥ 1024 px): horizontal `<ol>` with steps in a row; connector line (CSS pseudo-element `::before` on each `<li>` except last).
- Mobile (< 1024 px): vertical `<ol>`; vertical connector line left of step badges.
- Each step: numbered badge (`<span>` with step number) + `<h3>` title + `<p>` description.
- GSAP ScrollTrigger:
  - Connector line draws (clip-path or scaleX from 0 → 1).
  - Steps stagger in (opacity, y) as line reaches them.
  - Reduced-motion: all steps visible, line fully rendered.

**Verification**: Steps announced as ordered list by screen reader. Step numbers visible in DOM. Connector simplified on mobile.

---

### Milestone 7 — Testimonials Section

**Goal**: Social proof carousel. Implements VS-01 (credibility), VS-08 (carousel UX).

**Dependencies**: Milestone 0, testimonials data file complete.

**Tasks**:

**M7-T1**: Populate `data/testimonials.ts` with 4 `Testimonial` entries (placeholder copy).

**M7-T2**: `TestimonialsSection.tsx`
- `<section id="testimonials">`.
- `<h2>` headline.
- Uses `@shopwise/ui` `Carousel` component:
  - `opts={{ align: 'start', loop: true }}` for continuous looping.
  - Desktop: `basis-1/2` per `CarouselItem` (2 visible).
  - Mobile: full-width `CarouselItem` (1 visible).
- Each carousel item: quote, author, role, company, optional `<Avatar>`.
- Auto-advance: `useEffect` with `setInterval(5000)`. Calls `api.scrollNext()` from Carousel `setApi` callback. Clears on unmount.
- Pause on hover: `onMouseEnter` clears interval; `onMouseLeave` restarts.
- Pause on focus: `onFocus` clears interval.
- Reduced-motion: no auto-advance.
- `role="region"`, `aria-label="Customer testimonials"` on wrapping `<div>`.
- `aria-live="polite"` on visible slide area.

**Verification**: VS-08 (carousel controls keyboard-operable, auto-advance, pauses). Reduced-motion: no auto-advance.

---

### Milestone 8 — Pricing Section

**Goal**: Self-service plan comparison. Implements VS-03 (pricing self-sufficiency).

**Dependencies**: Milestone 0, pricing data file complete.

**Tasks**:

**M8-T1**: Populate `data/pricing.ts` with 3 `PricingTier` entries (Starter, Pro, Enterprise). Pro is `highlighted: true`.

**M8-T2**: `PricingSection.tsx`
- `<section id="pricing">`.
- `<h2>` headline.
- `billingPeriod` state (`'monthly' | 'annual'`), default `'monthly'`.
- Renders `<BillingToggle value={billingPeriod} onChange={setBillingPeriod} />`.
- Renders `<PricingCard tier={t} billingPeriod={billingPeriod} key={t.id} />` for each tier.
- Card layout: `grid-cols-1 md:grid-cols-2 lg:grid-cols-3`.

**M8-T3**: `BillingToggle.tsx`
- Two labels ("Monthly" / "Annual") with a toggle track between them.
- Uses `@shopwise/ui` `Switch` or a custom two-button toggle group.
- Annual label shows discount badge ("Save 20%") when in annual mode.
- `role="switch"`, `aria-checked={billingPeriod === 'annual'}`.

**M8-T4**: `PricingCard.tsx`
- Uses `@shopwise/ui` `Card`, `CardHeader`, `CardContent`, `CardFooter`, `Badge`, `Button`, `Separator`.
- When `tier.highlighted`: elevated border using `--landing-gradient-accent` colour, "Most Popular" `<Badge>` visually and in DOM.
- Price display: `tier.monthlyPrice ?? 'Custom'` / `tier.annualPrice ?? 'Custom'` based on `billingPeriod`.
- Feature list: `<ul>` with checkmark icons (lucide `Check`).
- CTA: `<Button>` → `asChild <Link to={tier.ctaHref}>` or `<a>` for external.
- Hover: `translateY(-4px)` + shadow increase (CSS transition).

**Verification**: VS-03 (pricing self-sufficiency). Billing toggle updates prices. Recommended badge in DOM. All CTAs navigate correctly. Single-column on mobile.

---

### Milestone 9 — FAQ Section

**Goal**: Objection handling accordion. Implements VS-06 (FAQ keyboard + ARIA).

**Dependencies**: Milestone 0, faq data file complete.

**Tasks**:

**M9-T1**: Populate `data/faq.ts` with 7 `FaqItem` entries (see spec section 9 for proposed questions).

**M9-T2**: `FaqSection.tsx`
- `<section id="faq">`.
- `<h2>` headline.
- Maps `FAQ_ITEMS` to `<FaqItem item={f} key={f.id} />`.
- Single-open behaviour: track `openId: string | null` in state, pass to each `FaqItem`.

**M9-T3**: `FaqItem.tsx`
- Uses `@shopwise/ui` `Collapsible`, `CollapsibleTrigger`, `CollapsibleContent`.
- Trigger: `<button>` (via `CollapsibleTrigger`) with question text + `ChevronDown` icon (rotates 180° when open via CSS transition).
- Content panel: `id={item.id}`, `role="region"`, `aria-labelledby={triggerId}`.
- Trigger: `aria-expanded={isOpen}`, `aria-controls={item.id}`.
- CSS height transition on content. Transition removed under `prefers-reduced-motion`.
- Reduced-motion: instant show/hide (no height animation).

**Verification**: VS-06 (keyboard operable, single-open, ARIA states). Chevron rotates. Reduced-motion: instant toggle.

---

### Milestone 10 — Final CTA Section + Polish

**Goal**: Conversion capture for motivated visitors. Page-level polish pass. Implements VS-01 (conversion), VS-09.

**Dependencies**: Milestones 1–9 complete.

**Tasks**:

**M10-T1**: `FinalCtaSection.tsx`
- `<section id="cta">` with high-contrast gradient background (`--landing-gradient-accent`-based).
- Centred single-column content.
- Eyebrow text → `<h2>` headline → supporting statement → primary `<Button>` + secondary `<Button>`.
- Primary: "Start for Free" → `/auth`. Secondary: "Talk to Sales" → `mailto:sales@shopwise.io` (external `<a>`).
- CSS pulse animation on primary button: `box-shadow` keyframe using `--landing-glow`. Paused under `prefers-reduced-motion`.
- GSAP scroll-entry: opacity 0→1, scale 0.96→1 when section enters viewport. Reduced-motion: static.

**M10-T2**: Wire `FinalCtaSection` into `LandingPage.tsx`.

**M10-T3**: Full polish pass:
- Verify section spacing is consistent (no collisions, no excessive gaps).
- Check dark-mode gradient rendering in all sections.
- Ensure `--landing-*` tokens are correctly applied everywhere.
- Verify no `any` types across all landing feature files.
- Remove any `console.log` or `debugger` statements.

**M10-T4**: Run all quality gates:
```bash
pnpm --filter web check-types
pnpm dlx ultracite check
pnpm --filter web test
```

**M10-T5**: Update `AGENTS.md` SPECKIT context section (see Step 4 below).

**Verification**: All 12 quickstart scenarios (VS-01 through VS-12) pass.

---

## Quality Gates (Must Pass Before Merge)

| Gate | Command | Expected Result |
|---|---|---|
| TypeScript | `pnpm --filter web check-types` | Zero errors |
| Lint | `pnpm dlx ultracite check` | Zero errors |
| Unit tests | `pnpm --filter web test` | All pass |
| Build | `pnpm build` | Zero build errors |

---

## Implementation Order (Recommended)

```
M0 → M1 → M2 → M3 → M4 → M5 → M6 → M7 → M8 → M9 → M10
```

M1, M2, M3 can be parallelised after M0 is complete (they share no component dependencies).
M4, M5, M6, M7, M8, M9 can be parallelised after M1 is done (Navbar/Footer present).
M10 requires all prior milestones complete.

---

## Risk Register

| Risk | Milestone | Mitigation |
|---|---|---|
| GSAP ScrollTrigger + Lenis conflict | M2, M4, M5, M6 | Use `lenis.on('scroll', ScrollTrigger.update)` in `gsap-config.ts` to keep ScrollTrigger in sync with Lenis. |
| `--landing-*` token specificity clash with existing tokens | M0 | Use unique `--landing-` prefix; never override existing `--primary` or `--background` tokens. |
| CLS from GSAP initial `opacity: 0` state | M2 | Set initial states in CSS (not JS); GSAP only transitions them. Alternatively, use GSAP's `immediateRender: false`. |
| Carousel auto-advance interfering with focus | M7 | Pause on `onFocus` as well as `onMouseEnter`. |
| LandingLayout removal breaking other routes | M0-T7 | Confirm no other route imports `LandingLayout` before removal. Currently only `index.tsx` uses it. |
| Content not ready (logos, testimonials) | M3, M7 | High-quality placeholder content designed to be swap-in-ready. |

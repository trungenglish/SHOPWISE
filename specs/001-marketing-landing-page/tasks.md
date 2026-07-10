# tasks.md — Marketing Landing Page

**Feature**: specs/001-marketing-landing-page
**Plan**: specs/001-marketing-landing-page/plan.md
**Generated**: 2026-07-10
**Total tasks**: 63

---

## Dependency Graph

```
Phase 1 (Setup)
  └── Phase 2 (Foundation)
        ├── Phase 3 (Navbar + Footer)
        ├── Phase 4 (Hero)          ← can run parallel with P3
        ├── Phase 5 (Trusted By)    ← can run parallel with P3, P4
        └── after Phase 3 completes:
              ├── Phase 6  (AI Decision Viz)
              ├── Phase 7  (Features)
              ├── Phase 8  (Workflow)
              ├── Phase 9  (Testimonials)
              ├── Phase 10 (Pricing)
              └── Phase 11 (FAQ)
                    └── Phase 12 (Final CTA + Assembly + QA)
```

---

## Phase 1 — Setup & Shared Infrastructure

> GSAP config, motion hook, scroll-spy hook, design tokens. No component renders without these.

- [x] T001 Add `--landing-*` OKLCH design tokens to `.dark` block in `packages/ui/src/styles/globals.css` (`--landing-gradient-start`, `--landing-gradient-mid`, `--landing-gradient-accent`, `--landing-glow`, `--landing-glass-bg`, `--landing-glass-border`)
- [x] T002 Create `apps/web/src/lib/gsap-config.ts` — import `gsap` and `ScrollTrigger` from `gsap`, call `gsap.registerPlugin(ScrollTrigger)`, and export `{ gsap, ScrollTrigger }`
- [x] T003 [P] Create `apps/web/src/features/landing/hooks/useReducedMotion.ts` — wraps `window.matchMedia('(prefers-reduced-motion: reduce)')` in `useState` + `useEffect` listener; returns `boolean`
- [x] T004 [P] Create `apps/web/src/features/landing/hooks/useScrollSpy.ts` — accepts `{ sectionIds: string[]; offset?: number }`; uses `IntersectionObserver` to return the `id` of the currently visible section or `null`

---

## Phase 2 — Typed Static Data Files

> All section content in typed `.ts` files. Depends on Phase 1 (type definitions reference `LucideIcon`).

- [x] T005 Create `apps/web/src/features/landing/data/trusted-by.ts` — define `TrustedByLogo` type and export `TRUSTED_BY_LOGOS: TrustedByLogo[]` with 6 placeholder company logo entries (name, src pointing to `src/assets/logos/`, width, height)
- [x] T006 [P] Create `apps/web/src/features/landing/data/features.ts` — define `FeatureItem` type and export `FEATURES: FeatureItem[]` with 6 entries: AI Vendor Scoring, Real-Time Market Intelligence, Procurement Workflow Automation, Spend Analytics, Collaborative Decision Workspace, Enterprise Integrations — each with a `lucide-react` icon assignment
- [x] T007 [P] Create `apps/web/src/features/landing/data/workflow-steps.ts` — define `WorkflowStep` type and export `WORKFLOW_STEPS: WorkflowStep[]` with 4 entries: Connect Your Data, Define Your Goals, Run AI Analysis, Decide with Confidence
- [x] T008 [P] Create `apps/web/src/features/landing/data/testimonials.ts` — define `Testimonial` type and export `TESTIMONIALS: Testimonial[]` with 4 placeholder entries (VP Procurement, Head of Category Management, CTO, Director of Finance); each with `id`, `quote`, `author`, `role`, `company`; `avatarSrc` optional
- [x] T009 [P] Create `apps/web/src/features/landing/data/pricing.ts` — define `PricingTier` type and export `PRICING_TIERS: PricingTier[]` with 3 tiers: Starter (`monthlyPrice: 49`, `annualPrice: 39`, `highlighted: false`), Pro (`monthlyPrice: 149`, `annualPrice: 119`, `highlighted: true`, `badge: 'Most Popular'`), Enterprise (`monthlyPrice: null`, `annualPrice: null`, `cta: 'Talk to Sales'`, `ctaHref: 'mailto:sales@shopwise.io'`)
- [x] T010 [P] Create `apps/web/src/features/landing/data/faq.ts` — define `FaqItem` type and export `FAQ_ITEMS: FaqItem[]` with 7 entries covering: AI accuracy, ERP integration, security/GDPR/SOC2, onboarding time, support SLA, customisation, free trial

---

## Phase 3 — Page Shell & Layout

> `LandingPage` root, Lenis mount, route update, placeholder logo assets.

- [x] T011 Add 6 placeholder SVG logo assets to `apps/web/src/assets/logos/` (monochrome SVGs for the Trusted By section; one per `TrustedByLogo` entry in `trusted-by.ts`)
- [x] T012 Create `apps/web/src/features/landing/LandingPage.tsx` — mount `<ReactLenis root>` from `lenis/react`; render a `<div>` containing ordered section slot comments; import and compose: `LandingNavbar`, `HeroSection`, `TrustedBySection`, `AiDecisionSection`, `FeaturesSection`, `WorkflowSection`, `TestimonialsSection`, `PricingSection`, `FaqSection`, `FinalCtaSection`, `LandingFooter` (import stubs initially)
- [x] T013 Update `apps/web/src/routes/index.tsx` — replace `LandingPageComponent` body to render `<LandingPage />`; remove `LandingLayout` import; update `head()` meta: `title: 'ShopWise — AI-Powered Procurement Intelligence'`, `description: 'ShopWise is an AI-powered shopping decision platform for enterprise procurement teams.'`

---

## Phase 4 — Navbar & Footer

> Sticky nav with scroll-spy + active state, mobile drawer, accessible footer.

- [x] T014 Create `apps/web/src/features/landing/components/navbar/LandingNavbar.tsx` — sticky `<nav aria-label="Main navigation">` with `isScrolled` state (scroll listener, true when `scrollY > 100`); on scroll: add `bg-background/80 backdrop-blur-md border-b border-border/40` classes; left: ShopWise wordmark as `<Link to="/">`; centre (desktop, hidden on mobile): nav links to `#features`, `#workflow`, `#pricing`, `#faq` with `useScrollSpy` active highlight; right: "Sign In" ghost `<Button asChild><Link to="/auth">` + "Get Started" filled `<Button asChild><Link to="/auth">`; mobile: hamburger icon button with `aria-expanded` and `aria-controls="mobile-nav"` driving `<MobileNavDrawer>`
- [x] T015 Create `apps/web/src/features/landing/components/navbar/MobileNavDrawer.tsx` — uses `@shopwise/ui` `Sheet`, `SheetContent` (side: "right"); props: `isOpen: boolean`, `onClose: () => void`; lists nav anchor links; each link calls `onClose` on click; includes "Get Started" `<Button>` at bottom; `id="mobile-nav"`
- [x] T016 Create `apps/web/src/features/landing/components/footer/LandingFooter.tsx` — `<footer>` landmark; multi-column desktop layout (brand col + Product, Company, Legal link groups); brand col: ShopWise wordmark + descriptor "AI-powered procurement intelligence for modern enterprises."; social links (LinkedIn, GitHub) with `target="_blank" rel="noopener noreferrer"` and `<span className="sr-only">` labels; uses `@shopwise/ui` `Separator` above copyright row; copyright: `© {new Date().getFullYear()} ShopWise. All rights reserved.`; single-column stack on mobile (`flex-col` below `lg:`)
- [x] T017 Wire `LandingNavbar` and `LandingFooter` into `LandingPage.tsx` (replace import stubs with real imports)

---

## Phase 5 — Hero Section

> Full-viewport hero with gradient, GSAP entry animation, two CTAs.

- [ ] T018 Create `apps/web/src/features/landing/components/hero/HeroVisual.tsx` — `aria-hidden="true"` decorative graphic; abstract SVG or CSS composition of floating geometric shapes in indigo/violet; CSS `@keyframes` float animation (`translateY` loop, 4 s ease-in-out infinite); apply `@media (prefers-reduced-motion: reduce) { animation-play-state: paused }` in Tailwind arbitrary or a `<style>` tag
- [ ] T019 Create `apps/web/src/features/landing/components/hero/HeroSection.tsx` — `<section id="hero">` with `min-h-dvh`; gradient background via `style={{ background: 'linear-gradient(...)' }}` referencing `--landing-gradient-start` and `--landing-gradient-mid` CSS vars; two-column on ≥ 1024 px (content 55 %, visual 45 %); content column: `<Badge>` eyebrow "AI-Powered Procurement", `<h1>` "Make Smarter Buying Decisions, Powered by AI.", `<p>` subtitle, "Start for Free" filled `<Button asChild><Link to="/auth">`, "See How It Works" outline `<Button>` as `<a href="#ai-decision">`; visual column: `<HeroVisual />` with `hidden lg:block`; GSAP entry via `useGSAP` using `gsap-config.ts`: check `useReducedMotion()` — if true call `gsap.set()` targets to end-state and return; otherwise stagger: eyebrow → h1 words → subtitle → buttons (`opacity 0→1`, `y 30→0`, `duration 0.6`, `ease: 'power2.out'`)
- [ ] T020 Wire `HeroSection` and `HeroVisual` into `LandingPage.tsx` (replace import stubs)

---

## Phase 6 — Trusted By Section

> Credibility logo band with auto-scrolling marquee on desktop.

- [ ] T021 Create `apps/web/src/features/landing/components/trusted-by/LogoMarquee.tsx` — props: `logos: TrustedByLogo[]`, `pauseOnHover?: boolean` (default `true`); duplicate logo list for seamless CSS loop; CSS `@keyframes` scroll (`translateX: 0 → -50%`); `role="region"` `aria-label="Trusted by — company logos"`; each `<img>` has `alt={logo.name}`, `loading="lazy"`, explicit `width`/`height`; pause class on `hover:` when `pauseOnHover`; `@media (prefers-reduced-motion: reduce) { animation: none }`
- [ ] T022 Create `apps/web/src/features/landing/components/trusted-by/TrustedBySection.tsx` — `<section id="trusted-by">`; muted label "Trusted by leading enterprises worldwide"; renders `<LogoMarquee logos={TRUSTED_BY_LOGOS} />` on desktop (`hidden sm:block`); 2-column static grid of logos on mobile (`grid grid-cols-2 sm:grid-cols-3 sm:hidden`)
- [ ] T023 Wire `TrustedBySection` into `LandingPage.tsx`

---

## Phase 7 — AI Decision Visualisation Section

> Product differentiator — animated AI workspace demo panel.

- [ ] T024 Create `apps/web/src/features/landing/components/ai-viz/AiDecisionPanel.tsx` — `aria-hidden="true"` glass card (`bg-[var(--landing-glass-bg)] border border-[var(--landing-glass-border)] backdrop-blur-sm rounded-xl`); inner layout: query input row (static text "Compare: Enterprise SaaS vendors Q3"), pulsing ring CSS animation (processing indicator), 3 scored result rows (vendor name + animated score bar); `useGSAP` with `ScrollTrigger` from `gsap-config.ts`: when panel enters viewport at `start: 'top 75%'`, animate bars `width: 0% → final%` and reveal rows with stagger; if `useReducedMotion()` is true, `gsap.set()` all elements to completed output state immediately
- [ ] T025 Create `apps/web/src/features/landing/components/ai-viz/AiDecisionSection.tsx` — `<section id="ai-decision">`; two-column on ≥ 1024 px (copy left, panel right), single column on mobile; copy column: eyebrow "Powered by AI", `<h2>` "Your AI Analyst, Always On.", supporting paragraph, "Try the AI Workspace" `<Button asChild><Link to="/auth">`; panel column: `<AiDecisionPanel />`; `<p className="sr-only">` describing: "Interactive demonstration: ShopWise AI analyses a vendor comparison query and returns scored recommendations in real time."; GSAP scroll-reveal on copy column (`opacity 0→1`, `y 40→0`) guarded by `useReducedMotion()`
- [ ] T026 Wire `AiDecisionSection` into `LandingPage.tsx`

---

## Phase 8 — Features Section

> Scannable product capability card grid.

- [ ] T027 Create `apps/web/src/features/landing/components/features/FeatureCard.tsx` — renders as `<li>`; uses `@shopwise/ui` `Card`, `CardContent`; glass surface: `bg-[var(--landing-glass-bg)] border-[var(--landing-glass-border)]`; props: `item: FeatureItem`; renders `<item.icon className="w-6 h-6 mb-4" aria-hidden="true" />`, `<h3>` title, `<p>` description; hover: CSS `transition-shadow transition-colors` for border glow (`hover:border-[var(--landing-gradient-accent)] hover:shadow-[0_0_20px_var(--landing-glow)]`); no GSAP on hover
- [ ] T028 Create `apps/web/src/features/landing/components/features/FeaturesSection.tsx` — `<section id="features">`; centred `<h2>` "Everything You Need to Buy Smarter" + optional subheadline; `<ul>` with `grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6`; maps `FEATURES` to `<FeatureCard item={f} key={f.id} />`; GSAP ScrollTrigger stagger on `<ul>` children (`opacity 0→1`, `y 30→0`, `stagger: 0.08`, `start: 'top 80%'`) guarded by `useReducedMotion()`
- [ ] T029 Wire `FeaturesSection` into `LandingPage.tsx`

---

## Phase 9 — Workflow Section

> Sequential step timeline with animated connector line.

- [ ] T030 Create `apps/web/src/features/landing/components/workflow/WorkflowTimeline.tsx` — props: `steps: WorkflowStep[]`; desktop ≥ 1024 px: horizontal `<ol>` with steps in a row; CSS `::before` pseudo-element connector line between steps (except last); mobile: vertical `<ol>` with vertical left-edge connector; each step: numbered `<span>` badge (step number in a circle), `<h3>` title, `<p>` description; GSAP ScrollTrigger: connector line draws (`scaleX 0→1` on desktop, `scaleY 0→1` on mobile) then steps stagger in (`opacity 0→1`); if `useReducedMotion()` is true, line fully rendered and all steps visible immediately
- [ ] T031 Create `apps/web/src/features/landing/components/workflow/WorkflowSection.tsx` — `<section id="workflow">`; `<h2>` "From Question to Decision in Four Steps" + supporting subheadline; renders `<WorkflowTimeline steps={WORKFLOW_STEPS} />`
- [ ] T032 Wire `WorkflowSection` into `LandingPage.tsx`

---

## Phase 10 — Testimonials Section

> Auto-advancing social proof carousel.

- [ ] T033 Create `apps/web/src/features/landing/components/testimonials/TestimonialsSection.tsx` — `<section id="testimonials">`; `<h2>` "What Our Customers Say"; wrapping `<div role="region" aria-label="Customer testimonials">`; uses `@shopwise/ui` `Carousel` with `opts={{ align: 'start', loop: true }}` and `setApi` to capture carousel API; each `CarouselItem` uses `basis-full md:basis-1/2` (1 visible mobile, 2 desktop); card content: large opening quote glyph `"`, `<blockquote>` with quote text, `<Avatar>` (from `@shopwise/ui`, with `AvatarImage` + `AvatarFallback` initials), author name, role, company; `CarouselPrevious` and `CarouselNext` with explicit `aria-label`; auto-advance: `useEffect` with `setInterval(5000)` calling `api.scrollNext()`; clear interval on `onMouseEnter`/`onFocus`; restart on `onMouseLeave`/`onBlur`; no auto-advance when `useReducedMotion()` is true; visible slide region has `aria-live="polite"`
- [ ] T034 Wire `TestimonialsSection` into `LandingPage.tsx`

---

## Phase 11 — Pricing Section

> Self-service plan comparison with billing toggle.

- [ ] T035 Create `apps/web/src/features/landing/components/pricing/BillingToggle.tsx` — props: `value: 'monthly' | 'annual'`, `onChange: (v: 'monthly' | 'annual') => void`; renders two label buttons ("Monthly" and "Annual") with a toggle track; Annual label shows `<Badge>` "Save 20%" when in annual mode; full ARIA: `role="switch"` on the toggle element, `aria-checked={value === 'annual'}`; CSS transition on toggle indicator position
- [ ] T036 Create `apps/web/src/features/landing/components/pricing/PricingCard.tsx` — props: `tier: PricingTier`, `billingPeriod: 'monthly' | 'annual'`; uses `@shopwise/ui` `Card`, `CardHeader`, `CardContent`, `CardFooter`, `Badge`, `Button`, `Separator`; plan name `<h3>` + optional subtitle; `<Badge>` "Most Popular" when `tier.badge` is set (visible in DOM, not just visual); price display: `tier[${billingPeriod}Price] ?? 'Custom'` with "/mo" suffix; feature `<ul>` with `lucide-react` `Check` icon per feature bullet; CTA `<Button>` (filled for highlighted, outline for others) as `<Link>` or `<a>` per `tier.ctaHref`; when `tier.highlighted`: elevated border `border-[var(--landing-gradient-accent)]` + subtle outer glow; hover CSS transition `translateY(-4px)` + shadow increase
- [ ] T037 Create `apps/web/src/features/landing/components/pricing/PricingSection.tsx` — `<section id="pricing">`; `<h2>` "Simple, Transparent Pricing"; `billingPeriod` state (`'monthly'` default); renders `<BillingToggle value={billingPeriod} onChange={setBillingPeriod} />`; renders `PRICING_TIERS.map(t => <PricingCard tier={t} billingPeriod={billingPeriod} key={t.id} />)` in `grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8`; note below cards: "Need a custom plan? " + "Talk to Sales" link
- [ ] T038 Wire `PricingSection` into `LandingPage.tsx`

---

## Phase 12 — FAQ Section

> Keyboard-accessible single-open accordion.

- [ ] T039 Create `apps/web/src/features/landing/components/faq/FaqItem.tsx` — props: `item: FaqItem`, `isOpen: boolean`, `onToggle: () => void`; uses `@shopwise/ui` `Collapsible`, `CollapsibleTrigger`, `CollapsibleContent`; trigger `<button>` renders question text + `lucide-react` `ChevronDown` (rotates 180° when open via CSS `transition-transform`); content panel: `id={item.id}`, `role="region"`, `aria-labelledby={triggerId}`; trigger has `aria-expanded={isOpen}`, `aria-controls={item.id}`; CSS `max-height` transition on content panel; `@media (prefers-reduced-motion: reduce) { transition: none }` for instant toggle
- [ ] T040 Create `apps/web/src/features/landing/components/faq/FaqSection.tsx` — `<section id="faq">`; `<h2>` "Frequently Asked Questions"; `openId: string | null` state (default `null`); maps `FAQ_ITEMS` to `<FaqItem item={f} isOpen={openId === f.id} onToggle={() => setOpenId(openId === f.id ? null : f.id)} key={f.id} />`; max-width 800 px centred
- [ ] T041 Wire `FaqSection` into `LandingPage.tsx`

---

## Phase 13 — Final CTA Section

> High-intent conversion section; closes the page before the footer.

- [ ] T042 Create `apps/web/src/features/landing/components/final-cta/FinalCtaSection.tsx` — `<section id="cta">`; high-contrast gradient background referencing `--landing-gradient-accent`; centred single-column: eyebrow text "Ready to transform your procurement?", `<h2>` "Start Making Smarter Decisions Today.", `<p>` supporting statement "Join hundreds of enterprises already using ShopWise. No credit card required.", primary filled `<Button asChild><Link to="/auth">` "Start for Free", secondary outline `<Button asChild><a href="mailto:sales@shopwise.io" rel="noopener">` "Talk to Sales"; CSS `@keyframes` pulse on primary button `box-shadow` using `--landing-glow` (paused via `animation-play-state: paused` under `prefers-reduced-motion`); GSAP scroll-entry: `opacity 0→1`, `scale 0.96→1` when section enters viewport, guarded by `useReducedMotion()`
- [ ] T043 Wire `FinalCtaSection` into `LandingPage.tsx`

---

## Phase 14 — Assembly, Polish & Quality Assurance

> Full page wired, visual polish, all quality gates green.

- [ ] T044 Confirm `LandingPage.tsx` composes all 11 sections in the correct order: `LandingNavbar` → `HeroSection` → `TrustedBySection` → `AiDecisionSection` → `FeaturesSection` → `WorkflowSection` → `TestimonialsSection` → `PricingSection` → `FaqSection` → `FinalCtaSection` → `LandingFooter` — verify no import stubs remain
- [ ] T045 Add Lenis + GSAP ScrollTrigger integration fix in `apps/web/src/lib/gsap-config.ts` — after `gsap.registerPlugin(ScrollTrigger)`, import and call `lenis.on('scroll', ScrollTrigger.update)` pattern (see GSAP docs); ensure Lenis is compatible with ScrollTrigger by adding `ScrollTrigger.normalizeScroll(false)`
- [ ] T046 Verify section `id` anchors are present and match nav href targets: `hero`, `ai-decision`, `features`, `workflow`, `testimonials`, `pricing`, `faq`, `cta` — fix any mismatches
- [ ] T047 Verify all GSAP `useGSAP` hooks guard with `useReducedMotion()`: if true, call `gsap.set()` to end-state and return — audit all 6 animation-bearing components: `HeroSection`, `AiDecisionPanel`, `AiDecisionSection`, `FeaturesSection`, `WorkflowTimeline`, `FinalCtaSection`
- [ ] T048 [P] Verify CSS animations (`HeroVisual` float, `LogoMarquee` scroll, `FinalCtaSection` button pulse) all have `@media (prefers-reduced-motion: reduce)` overrides disabling or pausing them
- [ ] T049 [P] Responsive audit — in browser devtools, test the following viewports and fix any overflow, clipping, or layout issues: 375 px (iPhone SE), 768 px (tablet), 1024 px (desktop breakpoint boundary), 1440 px (standard desktop)
- [ ] T050 [P] Accessibility audit — run keyboard-only navigation through the full page; verify: visible focus rings on all interactive elements, logical Tab order (top-to-bottom left-to-right), no focus traps outside mobile menu, ARIA attributes on Navbar (hamburger), Carousel (prev/next labels, `aria-live`), FAQ (`aria-expanded`, `aria-controls`, `role="region"`), Pricing toggle (`role="switch"`, `aria-checked`), Footer social links (`sr-only` labels)
- [ ] T051 [P] Contrast audit — use browser DevTools accessibility checker or axe to verify WCAG AA contrast ratios on all text rendered over gradient backgrounds (hero headline, subtitle, eyebrows, section headings over dark gradient sections)
- [ ] T052 Verify `LandingLayout` is no longer imported anywhere in `apps/web/src/routes/` — confirm only `index.tsx` was the consumer and it now uses `LandingPage` directly
- [ ] T053 [P] Remove any `console.log`, `debugger`, or `alert` statements from all files in `apps/web/src/features/landing/` and `apps/web/src/lib/gsap-config.ts`
- [ ] T054 Verify no `any` types exist across all landing feature files — run `grep -r ": any" apps/web/src/features/landing` and `apps/web/src/lib/gsap-config.ts`; replace with typed alternatives
- [ ] T055 Run `pnpm dlx ultracite fix` to auto-fix lint and formatting issues across `apps/web/src/features/landing/` and `packages/ui/src/styles/globals.css`
- [ ] T056 Run `pnpm dlx ultracite check` — fix any remaining lint errors not auto-resolved by T055
- [ ] T057 Run `pnpm --filter web check-types` — fix all TypeScript errors until exit code is 0
- [ ] T058 Run `pnpm --filter web test` — verify all existing tests still pass (no regressions from LandingLayout removal or route change)
- [ ] T059 Run `pnpm build` — verify the monorepo builds successfully with zero errors
- [ ] T060 Manual walkthrough of all 12 quickstart validation scenarios from `specs/001-marketing-landing-page/quickstart.md` (VS-01 through VS-12) — document any failures as follow-up tasks
- [ ] T061 [P] Update `specs/001-marketing-landing-page/spec.md` status from `Draft` to `Implemented`
- [ ] T062 [P] Update `specs/001-marketing-landing-page/checklists/requirements.md` — mark any items that required implementation-time decisions as resolved with brief notes
- [ ] T063 Commit with message: `feat(web): implement marketing landing page (001-marketing-landing-page)`

---

## Parallelisation Guide

| After | Can run in parallel |
|---|---|
| T001–T004 (Phase 1) complete | T005–T010 (all data files) |
| T011–T013 (Phase 3 shell) complete | T014–T016 (Navbar + Footer, Phase 4) AND T018–T019 (Hero, Phase 5) AND T021–T022 (Trusted By, Phase 6) |
| T014–T017 (Navbar+Footer wired) complete | T024–T025 (AI Viz), T027–T028 (Features), T030–T031 (Workflow), T033 (Testimonials), T035–T037 (Pricing), T039–T040 (FAQ) |
| T042–T043 (Final CTA + assembly) complete | T046–T054 (all audits and polish tasks) |
| T055–T059 (quality gates) complete | T060–T062 (docs and status updates) |

---

## MVP Scope

The minimum viable increment that is visually presentable and navigable:

**Phases 1–5 + Phase 4 (Navbar/Footer) + Phase 14 quality gates**

This delivers: foundation → data → shell → navbar/footer → hero → trusted by — enough to show the page loads, the dark-first gradient aesthetic, the primary CTA, and social proof. Remaining sections (Phases 7–13) add the full conversion funnel on top.

---

## Implementation Notes

- All task file paths are relative to the repo root `c:\D\SHOPWISE`.
- Imports in `apps/web` use the `@/` alias (mapped to `apps/web/src/`).
- All `@shopwise/ui` imports follow the deep-import pattern: `@shopwise/ui/components/<component>`.
- GSAP must always be imported from `apps/web/src/lib/gsap-config.ts`, never directly from `gsap`.
- `useReducedMotion()` must be called at the top level of every component that uses GSAP or CSS animation.
- Tailwind classes must follow canonical order enforced by `prettier-plugin-tailwindcss` (run `pnpm dlx ultracite fix` after each milestone).

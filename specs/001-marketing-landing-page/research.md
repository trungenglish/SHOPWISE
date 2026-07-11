# Research: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page **Completed**: 2026-07-11 **Status**: All v2.1.0 Decisions Logged

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
- "Launch Decision Workspace" → `/auth` (registration flow, same auth route with a mode param or redirect)
- "AI Workspace" → `/auth` (post-login redirect; no dedicated workspace route exists yet — link to `/auth` with a `redirect` search param intention)
- "Talk to Sales" / "Request Enterprise Demo" → external `mailto:sales@shopwise.io` (placeholder) or a `#contact` anchor.  
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
| --- | --- |
| `Button` | Navbar CTA, Hero CTAs, final CTA |
| `Card` | Feature cards (Meet the Workspace), Decision Confidence visual panel |
| `Badge` | Eyebrow badges on sections |
| `Sheet` | Mobile nav drawer |
| `Separator` | Footer divider |
| `collapsible.tsx` | _Removed from FAQ (no longer used)_ |
| `NavigationMenu` | Desktop nav links |

---

### D-06: GSAP usage pattern

**Decision**: Use `useGSAP` from `@gsap/react` for all GSAP animations. Register `ScrollTrigger` plugin once in a `gsap-config.ts` utility module at `apps/web/src/lib/gsap-config.ts`. Components import from there to guarantee single registration.  
**Rationale**: GSAP plugin registration must be idempotent. Centralising it avoids double-registration bugs across components.  
**Reduced-motion guard**: All `useGSAP` hooks check `window.matchMedia('(prefers-reduced-motion: reduce)').matches` before creating timelines. If true, elements are set to their end-state immediately via `gsap.set()`.

---

### D-07: Typed static data files

**Decision**: All repeated content (features list, workflow steps, trusted-by logos) lives in typed `.ts` data files under `apps/web/src/features/landing/data/`.  
**Rationale**: Separates content from structure, makes copy updates trivial, and ensures TypeScript enforces shape consistency.  
**File structure**:

```
apps/web/src/features/landing/data/
  features.ts
  workflow-steps.ts
  trusted-by.ts
```

---

### D-08: Component file structure

**Decision**: Scoped strictly to `apps/web/src/features/landing/` to co-locate features as per project Constitution.

---

### D-09: LandingLayout disposition

**Decision**: The existing `LandingLayout` in `apps/web/src/components/layouts/landing-layout.tsx` will be **replaced** by the new `LandingNavbar` and `LandingFooter` components composed directly inside `LandingPage.tsx`.

---

### D-10: CSS theme — current palette mismatch

**Decision**: The current `--primary` token (green/teal) is NOT used on the landing page gradient surfaces. Scoped landing variables are used instead.

---

### D-11: Target Audience Hierarchy (Clarified 2026-07-11)

**Decision**: Speaks directly to procurement specialists first (Procurement Team Leaders, Category Managers, Purchasing Specialists). Hero copy, Search Crisis, and Workspace sections use procurement-specific pain points (comparing suppliers across tabs, document fragmentation, technical spec validation). Executive value (ROI, audit trails, governance) is presented later in the page (Decision Confidence, final CTA).  
**Rationale**: Focuses the messaging on the users who champion and adopt the platform first, avoiding a generic, diluted C-suite message at the top.

---

### D-12: "ShopWise OS" Public Branding Removal (Clarified 2026-07-11)

**Decision**: Drop "OS" from all public-facing text, navbar, and headings. Keep "ShopWise" as the official name. Use "Decision Operating System" or "decision environment" in body copy/diagrams only to explain the conceptual framing. Standalone section renamed to "Beyond Search. Beyond Chat." and merged into Search Crisis.  
**Rationale**: Avoids branding confusion for enterprise buyers who associate "OS" with complex IT software or operating systems.

---

### D-13: Decorative-Only AI Workspace Simulations (Clarified 2026-07-11)

**Decision**: Both AI Decision Workspace and Decision Confidence visualisations are decorative GSAP simulations with auto-play on scroll. No click-driven state changes, no mock workflows requiring state management. Under prefers-reduced-motion, they snap immediately to their completed/final static state. Hover effects are restricted to CSS visual-only feedback (no content changes).  
**Rationale**: Minimizes implementation complexity, avoids maintaining redundant state machines on a marketing page, and guarantees consistent presentation.

---

### D-14: Conditional Trusted By Section (Clarified 2026-07-11)

**Decision**: The Trusted By section is optional and hidden by default if the `TRUSTED_BY_LOGOS` array is empty or disabled in content configuration. No generic placeholders or fictional logos will ever be rendered in production.  
**Rationale**: Prevents visual inauthenticity on an enterprise-facing site while maintaining structural flexibility for future logo addition.

---

### D-15: Narrative Restructuring & Search Crisis Merger (Clarified 2026-07-11)

**Decision**: Remove the standalone differentiation section ("Why ShopWise OS") and merge its chatbot/search differentiation message directly into "The Search Crisis" problem/solution comparison. The narrative flows from Problem Statement + Differentiation directly to Workspace Demonstration.  
**Rationale**: Corrects the visitor's mental model early before showing them the workspace, preventing them from classifying ShopWise as a search tool or chatbot.

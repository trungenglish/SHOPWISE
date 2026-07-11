# Quickstart Validation Guide: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page **Created**: 2026-07-10 **Last Updated**: 2026-07-11

> This document describes how to run and validate the landing page after implementation. It is not an implementation guide — see [data-model.md](../data-model.md) and [contracts/ui-contracts.md](../contracts/ui-contracts.md) for design details.

---

## Prerequisites

1. All workspace dependencies installed:
   ```bash
   pnpm install
   ```
2. No TypeScript errors in `apps/web`:
   ```bash
   pnpm --filter web check-types
   ```
3. Ultracite lint passes:
   ```bash
   pnpm dlx ultracite check
   ```

---

## Start the development server

```bash
pnpm dev:web
```

Open: **http://localhost:3001**

---

## Validation Scenarios

### VS-01: Full page load and visual check

**Steps**:

1. Open the app in a Chromium-based browser.
2. Open DevTools → Rendering tab → Enable "Dark prefers-color-scheme" (should already match app default dark theme).
3. Verify the page loads without layout shift (CLS = 0 on first paint).
4. Scroll through the page in order.

**Expected outcomes**:

- Navbar is visible at the top; "Launch Decision Workspace" CTA is distinct from nav links.
- Hero section occupies full viewport height with a blue/violet gradient background.
- All 8 sections are present in order: Hero, Trusted By (if enabled), Search Crisis, AI Decision Workspace, Meet the Workspace, Intelligence Layer, Decision Confidence, Final CTA.
- No horizontal scroll bar at 1440 px viewport width.
- Spacing and typography are visually consistent. No pricing, FAQ, or customer testimonial sections are rendered.

---

### VS-02: Navbar scroll behaviour

**Steps**:

1. Load the page.
2. Observe the navbar at scroll position 0.
3. Scroll down 150 px.

**Expected outcomes**:

- At scroll 0: navbar background is transparent or near-transparent.
- At scroll > 100 px: navbar background transitions to a dark surface with backdrop blur.
- Navbar remains sticky at the top of the viewport throughout scrolling.

---

### VS-03: Mobile layout (375 px viewport)

**Steps**:

1. Open DevTools → set device to iPhone SE (375 × 667 px) or manually set viewport to 375 px.
2. Load the page and scroll through all sections.

**Expected outcomes**:

- Zero horizontal overflow at any scroll position.
- Navbar shows logotype + hamburger icon only (nav links hidden).
- All feature cards, problem comparison blocks, and workflow steps are single-column.
- All tap targets (buttons, menu toggle) have a touch area ≥ 44 × 44 px.
- Hero visual is hidden or deprioritised; headline and CTAs are visible above the fold.

---

### VS-04: Mobile nav drawer

**Steps**:

1. At 375 px viewport, click the hamburger icon.

**Expected outcomes**:

- A drawer slides in from the right (Sheet component).
- All nav links (Problem, Workspace, Intelligence, Confidence, Contact) are visible inside the drawer.
- Pressing Escape closes the drawer.
- Clicking any nav link closes the drawer and scrolls to the relevant section.
- Focus is trapped inside the open drawer.

---

### VS-05: Smooth scroll to section anchors

**Steps**:

1. At desktop width, click each nav link (Problem, Workspace, Intelligence, Confidence, Contact).

**Expected outcomes**:

- Page scrolls smoothly to the corresponding section.
- Scroll motion is fluid (Lenis smooth scroll active).
- Active nav link is highlighted after scroll settles.
- "Contact" link scrolls to the Final CTA section.

---

### VS-06: The Search Crisis (Problem & Differentiation Column Check)

**Steps**:

1. Navigate to the Search Crisis section.
2. Observe the layout on desktop.

**Expected outcomes**:

- Left column (Legacy approach) uses muted, degraded styling with at least 4 pain points.
- Right column (ShopWise approach) uses accent/glowing styling with at least 4 corresponding outcomes.
- The solution column or a block below contains the clear differentiation statement: "not another search engine, chatbot, or marketplace — but an AI Decision Intelligence Platform..."
- The visual presentation clearly establishes this distinction.

---

### VS-07: Decision Confidence Signal Check

**Steps**:

1. Scroll to the Decision Confidence section.
2. Observe the decorative visual mockup.

**Expected outcomes**:

- Mockup displays four signal indicators: Evidence Status, Confidence Indicator, Risk Status, Auditability Signal.
- Each signal has a descriptive textual label — no information is conveyed solely by color.
- All values represent realistic, static enterprise data. No clickable state changes occur.
- Copy successfully targets executive value (governance, auditability, justified ROI).

---

### VS-08: prefers-reduced-motion compliance

**Steps**:

1. In DevTools → Rendering tab → Enable "Emulate CSS media feature prefers-reduced-motion: reduce".
2. Reload the page and scroll through all sections.

**Expected outcomes**:

- No GSAP timeline animations play.
- No CSS `@keyframes` animations play (Trusted By marquee is static if enabled, hero elements appear statically).
- All content is visible in its final/end state (nothing is hidden behind un-triggered opacity 0 or clip-path).
- Visualisation panels in AI Workspace and Decision Confidence render in their final static state immediately.

---

### VS-09: Keyboard navigation

**Steps**:

1. Reload the page with keyboard focus at the browser address bar.
2. Press Tab repeatedly to navigate through all interactive elements.

**Expected outcomes**:

- Every interactive element (nav links, CTA buttons, footer links) receives visible focus (focus ring).
- Tab order follows visual reading order (top-to-bottom, left-to-right).
- No focus is lost or trapped outside the mobile menu context.
- Shift+Tab navigates backwards through all the same elements.

---

### VS-10: TypeScript type-check

```bash
pnpm --filter web check-types
```

**Expected outcome**: Zero TypeScript errors.

---

### VS-11: Lint check

```bash
pnpm dlx ultracite check
```

**Expected outcome**: Zero lint errors (auto-fixable issues resolved first with `pnpm dlx ultracite fix`).

---

## References

- Spec: [spec.md](../spec.md)
- Data model: [data-model.md](../data-model.md)
- UI contracts: [contracts/ui-contracts.md](../contracts/ui-contracts.md)
- Constitution: [.specify/memory/constitution.md](../../../.specify/memory/constitution.md)

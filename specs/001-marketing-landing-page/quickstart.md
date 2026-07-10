# Quickstart Validation Guide: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page
**Created**: 2026-07-10

> This document describes how to run and validate the landing page after implementation.
> It is not an implementation guide — see [data-model.md](../data-model.md) and
> [contracts/ui-contracts.md](../contracts/ui-contracts.md) for design details.

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
4. Scroll through all 11 sections in order.

**Expected outcomes**:
- Navbar is visible at the top; CTA button is distinct from nav links.
- Hero section occupies full viewport height with a blue/violet gradient background.
- All 11 sections are present and content is legible (high contrast against background).
- No horizontal scroll bar at 1440 px viewport width.

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
- All feature cards, pricing cards, and workflow steps are single-column.
- All tap targets (buttons, FAQ toggles) have a touch area ≥ 44 × 44 px.
- Hero visual is hidden or deprioritised; headline and CTAs are visible above the fold.

---

### VS-04: Mobile nav drawer

**Steps**:
1. At 375 px viewport, click the hamburger icon.

**Expected outcomes**:
- A drawer slides in from the right (Sheet component).
- All nav links are visible inside the drawer.
- Pressing Escape closes the drawer.
- Clicking any nav link closes the drawer and scrolls to the relevant section.
- Focus is trapped inside the open drawer.

---

### VS-05: Smooth scroll to section anchors

**Steps**:
1. At desktop width, click each nav link (Features, Workflow, Pricing, FAQ).

**Expected outcomes**:
- Page scrolls smoothly to the corresponding section.
- Scroll motion is fluid (Lenis smooth scroll active).
- Active nav link is highlighted after scroll settles.

---

### VS-06: FAQ accordion

**Steps**:
1. Navigate to the FAQ section.
2. Click the first question.
3. Click a second question.
4. Using keyboard only: Tab to the third question, press Enter.

**Expected outcomes**:
- Clicking the first question expands its answer (height transition).
- Clicking the second question expands it and collapses the first (single-open).
- Keyboard Enter on the third question expands it.
- Each trigger button shows `aria-expanded="true"` when open (verify in DevTools Accessibility panel).

---

### VS-07: Pricing billing toggle

**Steps**:
1. Navigate to the Pricing section.
2. Observe displayed prices (Monthly active by default).
3. Click/activate the Annual toggle.

**Expected outcomes**:
- All plan prices update to annual values (lower amounts with a discount indicator).
- Toggle state is reflected visually (active indicator moves to "Annual").
- Toggle's `aria-checked` attribute updates in the DOM.

---

### VS-08: Testimonials carousel

**Steps**:
1. Navigate to the Testimonials section.
2. Wait 5 seconds.
3. Click the "Next" arrow.
4. Click the second dot indicator.

**Expected outcomes**:
- After 5 seconds, carousel auto-advances to the next testimonial (cross-fade).
- Clicking "Next" advances manually.
- Clicking a dot jumps to that testimonial.
- `aria-live="polite"` region announces the slide change (verify with screen reader or DevTools).

---

### VS-09: prefers-reduced-motion compliance

**Steps**:
1. In DevTools → Rendering tab → Enable "Emulate CSS media feature prefers-reduced-motion: reduce".
2. Reload the page and scroll through all sections.

**Expected outcomes**:
- No GSAP timeline animations play.
- No CSS `@keyframes` animations play (marquee stops, hero elements appear statically).
- All content is visible in its final/end state (nothing is hidden behind un-triggered opacity 0).
- FAQ accordion opens/closes without height animation (instant show/hide).
- Testimonials carousel does not auto-advance.

---

### VS-10: Keyboard navigation

**Steps**:
1. Reload the page with keyboard focus at the browser address bar.
2. Press Tab repeatedly to navigate through all interactive elements.

**Expected outcomes**:
- Every interactive element (nav links, CTA buttons, FAQ toggles, carousel arrows, dot indicators, pricing CTAs, footer links) receives visible focus (focus ring).
- Tab order follows visual reading order (top-to-bottom, left-to-right).
- No focus is lost or trapped outside the mobile menu context.
- Shift+Tab navigates backwards through all the same elements.

---

### VS-11: TypeScript type-check

```bash
pnpm --filter web check-types
```

**Expected outcome**: Zero TypeScript errors.

---

### VS-12: Lint check

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

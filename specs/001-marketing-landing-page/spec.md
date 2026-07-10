# Marketing Landing Page

**Version**: 1.0.0
**Status**: Draft
**Feature Directory**: specs/001-marketing-landing-page
**Created**: 2026-07-10
**Last Updated**: 2026-07-10

---

## Overview

The ShopWise marketing landing page is the primary public-facing surface that
introduces enterprise buyers, procurement teams, and retail decision makers to
ShopWise as an AI-powered shopping decision platform. Its singular objective is to
create an immediate, high-trust, premium first impression that persuades visitors to
explore the platform and take the primary conversion action: launching the AI workspace.

The page delivers a dark-first, visually excellent, enterprise-grade experience aligned
with ShopWise's UI philosophy (Premium · Minimal · Enterprise · Modern · Dark-First)
and must meet WCAG 2.1 AA accessibility criteria while performing to Core Web Vitals
thresholds on all screen sizes.

---

## Problem Statement

ShopWise currently lacks a dedicated marketing surface to introduce the product to
prospective enterprise customers. Without it, the first impression visitors form is
dictated by the application UI itself — an environment designed for active users, not
for evaluators. This mismatch leads to low conversion, poor brand perception, and an
inability to communicate value to decision makers who do not yet have accounts.

Enterprise buyers need to understand the platform's value proposition, see credibility
signals, grasp the AI capabilities, understand pricing options, and find answers to
common objections — all before they invest time in a trial or demo.

---

## Goals

- Communicate ShopWise's value proposition to enterprise audiences within the first
  10 seconds of a visit.
- Drive visitors to the primary conversion action (launch AI workspace / start trial)
  with measurable, positive intent.
- Establish brand credibility through trust signals, social proof, and premium visual
  execution.
- Answer common pre-purchase objections (capability, pricing, workflow fit, security)
  without a sales call.
- Meet WCAG 2.1 AA accessibility standards so the page is usable by all visitors.
- Achieve Core Web Vitals thresholds (LCP < 2.5 s, CLS < 0.1, INP < 200 ms) on
  first load.

## Non-Goals

- This page is not a product documentation hub; it does not replace help docs or
  API references.
- It does not implement authentication, sign-up forms, or user account creation flows
  — CTAs link to the appropriate app routes.
- It does not serve as a blog or content marketing channel.
- Dynamic personalisation based on visitor identity or firmographics is out of scope
  for this version.
- A/B testing infrastructure is out of scope; the page should be built to be
  A/B-test-ready (clean section boundaries, clear CTA targets) but no testing tooling
  is required.

---

## Target Users

| User Type | Description | Primary Need |
|---|---|---|
| Enterprise Buyer | Senior decision maker at a large retail or procurement organisation; evaluates ROI and vendor credibility | Understand the platform's value and trust signals quickly; find pricing |
| Procurement Team Member | Operational buyer or category manager who would use ShopWise daily | Understand how the AI workspace fits their existing workflow |
| Retail Business Owner | SMB or mid-market owner considering ShopWise for their team | Understand pricing tiers and whether the product fits their scale |
| Technical Evaluator | An IT or engineering stakeholder validating integration fit | Confirm the platform is enterprise-grade, secure, and maintainable |

---

## User Scenarios & Testing

### Scenario 1: Enterprise buyer forms a first impression

**Given**: A prospective enterprise buyer clicks a link from a sales email or search ad
and lands on the page for the first time on a desktop browser.
**When**: They read the hero headline, subtitle, and scan the page for credibility signals.
**Then**: They understand what ShopWise does, for whom it is designed, and they see
recognisable trusted-by logos within the first viewport.
**Acceptance**: A first-time visitor can correctly describe ShopWise's core value
proposition after viewing only the hero section — validated via user testing with ≥ 70%
accuracy.

### Scenario 2: Visitor explores AI capabilities

**Given**: A visitor scrolls past the Trusted By section and reaches the AI Decision
Visualisation section.
**When**: The visualisation animates (or appears statically if motion is reduced) showing
the AI processing a shopping decision.
**Then**: The visitor understands the AI capability differentiator without needing to read
technical documentation.
**Acceptance**: The AI visualisation section is self-explanatory to a non-technical visitor
and does not require supplementary tooltips for comprehension.

### Scenario 3: Visitor evaluates pricing and initiates conversion

**Given**: A visitor has scrolled through the page and is in the Pricing section.
**When**: They compare plan tiers and click the primary CTA for their tier.
**Then**: They are routed to the appropriate app destination (trial launch, AI workspace,
or sales contact) without any friction.
**Acceptance**: Every Pricing CTA navigates to a valid, reachable route. The Pricing
section is self-contained and answers: what do I get, what does it cost, and how do I
start?

### Scenario 4: Visitor with reduced-motion preference

**Given**: A visitor has enabled `prefers-reduced-motion` at the operating-system level
and arrives on the page.
**When**: They scroll through all sections.
**Then**: No GSAP timelines or CSS transitions play. All content is statically visible
and fully readable. No content is hidden behind animation-gated reveal states.
**Acceptance**: Every section passes a visual audit with `prefers-reduced-motion: reduce`
active in browser devtools. Zero elements are invisible or clipped due to an unanimated
initial state.

### Scenario 5: Mobile visitor

**Given**: A visitor opens the page on a 375 px-wide mobile device.
**When**: They scroll through the entire page.
**Then**: All sections reflow to single-column layouts, all text is legible, all
interactive elements have tap targets ≥ 44 × 44 px, and no horizontal overflow exists.
**Acceptance**: The page passes a manual review on a 375 px viewport with zero horizontal
scroll and all tap targets meeting the size requirement.

### Scenario 6: Keyboard-only visitor navigates the page

**Given**: A visitor navigates the page using keyboard only (Tab, Enter, Space, arrow keys).
**When**: They tab through all interactive elements (nav links, CTAs, FAQ toggles).
**Then**: Focus is always visible, focus order is logical (top-to-bottom, left-to-right),
and no focus trap exists except in intentionally modal contexts.
**Acceptance**: A keyboard-only navigation audit finds no focus loss, no invisible focus
rings, and no unreachable interactive elements.

---

## Functional Requirements

### Navigation (Navbar)

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-NAV-01 | The navbar displays the ShopWise logotype on the left and primary navigation links on the right. | Must Have | Logotype and nav links are visible on all viewport widths ≥ 375 px. |
| FR-NAV-02 | The navbar includes a primary CTA button ("Get Started" or equivalent) distinct from nav links. | Must Have | CTA button is visually distinct (filled/outlined vs. text links) and navigates to the correct destination. |
| FR-NAV-03 | On scroll, the navbar becomes sticky and applies a subtle backdrop blur / elevated surface effect. | Must Have | Navbar remains at top of viewport after scrolling 100 px. Background opacity transitions smoothly. |
| FR-NAV-04 | On mobile (< 768 px), navigation links collapse into a hamburger/menu toggle. | Must Have | Menu toggle is visible on mobile. Activating it reveals all nav links. Focus is trapped within the open menu. Menu closes on link click or pressing Escape. |
| FR-NAV-05 | Nav links scroll smoothly to their respective page sections. | Should Have | Clicking a nav link scrolls to the correct section anchor. Active link is visually highlighted based on scroll position. |
| FR-NAV-06 | The navbar provides a "Sign In" link for returning users. | Should Have | "Sign In" link is present and navigates to the correct app authentication route. |

### Hero Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-HERO-01 | The hero presents a bold, large-format headline that communicates the core value proposition in ≤ 10 words. | Must Have | Headline is the largest typographic element on the page. It is immediately visible on load without scrolling. |
| FR-HERO-02 | A supporting subtitle expands on the headline in ≤ 30 words. | Must Have | Subtitle is present below the headline with a clear typographic hierarchy (smaller weight/size). |
| FR-HERO-03 | Two CTAs are present: a primary ("Start for Free" or equivalent) and a secondary ("Watch Demo" or equivalent). | Must Have | Both CTAs are keyboard-operable, have accessible labels, and navigate to valid routes. |
| FR-HERO-04 | A hero visual (AI visualisation graphic, abstract 3D element, or premium illustration) supports the headline. | Must Have | Visual is present on desktop (≥ 1024 px). On mobile, the visual is deprioritised or hidden to maintain content hierarchy. |
| FR-HERO-05 | The hero applies a soft blue/purple gradient background that establishes the dark-first visual language. | Must Have | Background gradient is visible. Text contrast against the gradient meets WCAG AA (≥ 4.5:1 for body, ≥ 3:1 for large text). |
| FR-HERO-06 | Entry animations (headline, subtitle, CTAs staggered reveal) play on first load. | Should Have | Animations are absent or replaced by static end-state when `prefers-reduced-motion` is active. |

### Trusted By Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-TRUST-01 | The section displays ≥ 5 and ≤ 10 representative customer / partner logos. | Must Have | Logos are present and legible at all viewport widths. |
| FR-TRUST-02 | Logos are displayed in a single horizontal row (desktop) or a 2-column grid (mobile), with no overflow. | Must Have | No horizontal scroll exists. Logos wrap or reduce gracefully on small screens. |
| FR-TRUST-03 | Each logo includes an accessible text alternative (visually hidden or `alt` attribute). | Must Have | Screen readers announce each logo's company name. |
| FR-TRUST-04 | A subtle auto-scrolling marquee is applied on desktop to animate logo rows. | Should Have | Marquee pauses on hover. Marquee is absent when `prefers-reduced-motion` is active. |
| FR-TRUST-05 | An introductory label (e.g., "Trusted by leading enterprises") precedes the logos. | Should Have | Label is present and uses a muted, secondary text style. |

### AI Decision Visualisation

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-AIVIZ-01 | The section presents an interactive or animated visualisation of the AI analysing a shopping decision (e.g., comparing products, scoring vendors, synthesising signals). | Must Have | Visualisation is present and comprehensible without tooltips on first view. |
| FR-AIVIZ-02 | A headline and supporting copy explain the AI capability in plain language. | Must Have | Copy is present. No technical jargon is used without plain-language explanation. |
| FR-AIVIZ-03 | The visualisation responds to scroll position (scroll-triggered animation) or user pointer (hover micro-interaction). | Should Have | Interaction is absent or replaced by a static representation when `prefers-reduced-motion` is active. |
| FR-AIVIZ-04 | The section uses a two-column layout on desktop (copy left, visualisation right) and single-column on mobile. | Must Have | Columns are side-by-side on ≥ 1024 px viewports. Single column on < 1024 px. |
| FR-AIVIZ-05 | A CTA linking to the AI workspace is present within or immediately below this section. | Should Have | CTA is keyboard-operable and navigates to the correct route. |

### Features Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-FEAT-01 | The section presents ≥ 4 and ≤ 8 product features. Each feature has an icon, a short title (≤ 5 words), and a description (≤ 25 words). | Must Have | All features are present. Icon, title, and description are visually distinct. |
| FR-FEAT-02 | Features are arranged in a responsive grid: 3-column on desktop, 2-column on tablet, 1-column on mobile. | Must Have | Grid reflows correctly at breakpoints. No overflow or clipped content. |
| FR-FEAT-03 | Icons use `lucide-react` exclusively. | Must Have | No other icon library is used. Icons have accessible labels (aria-hidden or aria-label as appropriate). |
| FR-FEAT-04 | Hovering a feature card applies a subtle glass-effect highlight or border glow. | Should Have | Hover state is visible. CSS transition is used (not GSAP). Hover is absent on touch devices. |
| FR-FEAT-05 | A section headline and optional subheadline introduce the feature grid. | Must Have | Section headline is an `<h2>` element. |

### Workflow Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-WORK-01 | The section illustrates the end-to-end user workflow in ≥ 3 and ≤ 5 numbered steps. Each step has a title and a short description (≤ 20 words). | Must Have | All steps are present. Step number, title, and description are present and legible. |
| FR-WORK-02 | Steps are connected visually (e.g., a vertical line, horizontal connector, or dashed path) to convey sequence. | Should Have | Connector is present on desktop. Connector is simplified or hidden on mobile without losing meaning. |
| FR-WORK-03 | The section uses a timeline or step-progress layout: horizontal on desktop, vertical on mobile. | Must Have | Layout transitions correctly at breakpoints. |
| FR-WORK-04 | A short supporting headline and optional introductory copy precede the steps. | Must Have | Section headline is present as an `<h2>`. |

### Testimonials Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-TEST-01 | The section presents ≥ 3 and ≤ 6 testimonials. Each includes a quote (≤ 40 words), the author's name, role, and company name. | Must Have | All testimonials are present and fully rendered. |
| FR-TEST-02 | Testimonials are displayed in a carousel (desktop: 2–3 visible; mobile: 1 visible) with navigation controls (prev/next arrows and dot indicators). | Must Have | Carousel controls are keyboard-operable. Current slide is announced to screen readers via `aria-live`. |
| FR-TEST-03 | The carousel auto-advances every 5 seconds. Auto-advance pauses on hover or focus and is disabled when `prefers-reduced-motion` is active. | Should Have | Auto-advance is verifiable by timing. Pause on hover and focus is verifiable. |
| FR-TEST-04 | An optional avatar or company logo accompanies each testimonial. | Should Have | Avatars/logos have `alt` text or are aria-hidden if purely decorative. |
| FR-TEST-05 | A section headline (e.g., "What our customers say") introduces the testimonials. | Must Have | Section headline is present as an `<h2>`. |

### Pricing Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-PRICE-01 | The section presents ≥ 2 and ≤ 4 pricing tiers. Each tier includes a plan name, price (or "Contact Sales"), a list of ≥ 3 included features, and a CTA button. | Must Have | All tiers are fully rendered. |
| FR-PRICE-02 | One tier is visually highlighted as recommended (e.g., a border accent, badge, or elevated card). | Must Have | Recommended tier is visually distinct. Distinction is communicated to screen readers (e.g., "Recommended" label). |
| FR-PRICE-03 | A billing-period toggle (Monthly / Annual with discount indicator) is present. Switching the toggle updates all displayed prices. | Must Have | Prices update on toggle without a page reload. Toggle is keyboard-operable. |
| FR-PRICE-04 | Each plan's CTA navigates to the correct destination (trial, checkout, or sales contact). | Must Have | All CTA links are valid and reachable. |
| FR-PRICE-05 | A brief note below the tiers addresses enterprise / custom pricing (e.g., "Need a custom plan? Talk to sales."). | Should Have | Note and its link are present and navigable. |
| FR-PRICE-06 | Pricing tiers display in a responsive horizontal row (desktop) and vertical stack (mobile). | Must Have | Layout reflows without horizontal overflow at all breakpoints. |

### FAQ Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-FAQ-01 | The section presents ≥ 5 and ≤ 10 frequently asked questions as an accordion (one expanded at a time or multiple). | Must Have | All FAQs are present. Default state: all collapsed. |
| FR-FAQ-02 | Each FAQ item is a keyboard-operable toggle: pressing Enter or Space on the question expands or collapses the answer. | Must Have | Keyboard operation is fully functional. |
| FR-FAQ-03 | Expanded/collapsed state is communicated to screen readers via `aria-expanded` on the trigger and `aria-controls` pointing to the answer panel. | Must Have | Screen reader audit confirms correct ARIA state changes on toggle. |
| FR-FAQ-04 | Expand/collapse applies a smooth height transition. Transition is removed when `prefers-reduced-motion` is active. | Should Have | Transition is present by default. Static show/hide is applied under reduced-motion. |
| FR-FAQ-05 | A section headline (e.g., "Frequently Asked Questions") introduces the accordion. | Must Have | Section headline is present as an `<h2>`. |
| FR-FAQ-06 | FAQs address the most common objections: pricing, security, integration, AI accuracy, and support SLA. | Must Have | At minimum one question per objection category is present. |

### Final CTA Section

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-CTA-01 | The section presents a single, bold call-to-action with a headline (≤ 8 words), a supporting statement (≤ 20 words), and the primary CTA button. | Must Have | Headline, supporting statement, and CTA button are all present. |
| FR-CTA-02 | The primary CTA button matches the hero's primary CTA in label and destination. | Must Have | CTA label is identical to hero primary CTA. Destination route is identical. |
| FR-CTA-03 | An optional secondary CTA ("Talk to Sales" or "Book a Demo") is present below the primary button. | Should Have | Secondary CTA is present and navigates to a valid destination. |
| FR-CTA-04 | The section uses a high-contrast background (gradient or solid dark) that clearly differentiates it from the preceding FAQ section. | Must Have | Visual separation between FAQ and Final CTA sections is evident at a glance. |

### Footer

| ID | Requirement | Priority | Acceptance Criteria |
|---|---|---|---|
| FR-FOOT-01 | The footer includes the ShopWise logotype, a brief descriptor (≤ 10 words), and grouped navigation links (Product, Company, Legal). | Must Have | All three link groups are present. |
| FR-FOOT-02 | Each link group has an accessible heading (visually rendered as a label above the links). | Must Have | Headings are present and use the correct heading level or `role="group"` with `aria-labelledby`. |
| FR-FOOT-03 | Legal links (Privacy Policy, Terms of Service) are present and navigate to valid routes. | Must Have | Both links are present and reachable. |
| FR-FOOT-04 | Social media links are present (at minimum LinkedIn and GitHub). Each link opens in a new tab with `rel="noopener"` and has a visually hidden accessible label. | Should Have | Links are present. `target="_blank"` is paired with `rel="noopener"`. Screen readers announce link purpose. |
| FR-FOOT-05 | A copyright notice with the current year is present at the bottom of the footer. | Must Have | Copyright text is present and displays the correct year. |
| FR-FOOT-06 | The footer uses a multi-column layout on desktop (logo + description left, link groups right) and single-column stack on mobile. | Must Have | Layout reflows correctly on all viewports. |

---

## Section Specifications

> This section provides the complete design brief for each of the 11 page sections.
> Implementation teams should treat each sub-section as a standalone design contract.

---

### 1. Navbar

**Purpose**: Persistent wayfinding and primary conversion entry point accessible from any scroll position.

**User Value**: Lets visitors orient themselves, jump to key sections, and take the conversion action without needing to scroll back to the top.

**Layout**:
- Fixed/sticky to top of viewport.
- Full-width container, max-width constrained (e.g., 1280 px) and centred.
- Left: logotype (wordmark).
- Centre (desktop): nav links — Features, Workflow, Pricing, FAQ.
- Right (desktop): "Sign In" text link + "Get Started" filled button.
- Mobile (< 768 px): logotype left, hamburger icon right. Nav links in a slide-down or overlay drawer.

**Key Content**:
- ShopWise logotype.
- Nav links: Features · Workflow · Pricing · FAQ.
- "Sign In" (secondary, text weight).
- "Get Started" (primary CTA, filled, accent colour).

**UX Behaviour**:
- On initial load: transparent or very low-opacity background.
- After scrolling > 100 px: background transitions to a dark surface with `backdrop-filter: blur` and a subtle bottom border.
- Active section highlight: nav link corresponding to the visible section is highlighted (colour or underline).
- Mobile menu: hamburger toggles a full-width drawer. Drawer closes on link tap or Escape key.
- Smooth scroll to anchor on link click.

**Accessibility**:
- `<nav aria-label="Main navigation">` wrapping the nav element.
- Mobile menu button: `aria-expanded`, `aria-controls` pointing to menu panel.
- All links and the CTA button are keyboard-focusable with visible focus rings.
- Focus trapped within open mobile drawer.

**Responsive Behaviour**:
- ≥ 768 px: full desktop nav.
- < 768 px: hamburger menu, drawer overlay.

---

### 2. Hero

**Purpose**: Create an immediate, high-impact first impression that communicates ShopWise's core value proposition and triggers the primary conversion action.

**User Value**: Answers "What is this?" and "Why should I care?" within the first 5 seconds, with a clear path to act.

**Layout**:
- Full-viewport-height (100 dvh) section.
- Two-column on desktop (≥ 1024 px): content left (60%), visual right (40%).
- Single-column centred on mobile: headline, subtitle, CTAs stacked; visual below or hidden.
- Gradient background: soft radial/linear blend from deep navy/charcoal (#0a0a0f base) through indigo to soft violet, applied to the section background.

**Key Content**:
- Eyebrow label (e.g., "AI-Powered Procurement" — small caps, accent colour).
- Headline: ≤ 10 words, e.g., "Make Smarter Buying Decisions, Powered by AI."
- Subtitle: ≤ 30 words expanding on the headline — who it's for and what they gain.
- Primary CTA: "Start for Free" → app trial route.
- Secondary CTA: "See How It Works" → smooth scroll to AI Decision Visualisation.
- Hero visual: abstract 3D or isometric representation of AI data processing, or a premium illustrated mockup of the AI workspace dashboard.

**UX Behaviour**:
- Entry animation (GSAP): headline words stagger in from below (opacity 0 → 1, y 30px → 0), subtitle fades in, CTAs slide in, visual scales in.
- All animations are absent / replaced with static end-state under `prefers-reduced-motion`.
- Secondary CTA scrolls the page smoothly to the AI section anchor.
- Subtle particle or grid overlay (CSS or GSAP) adds depth without distracting.

**Accessibility**:
- Headline is the `<h1>` of the page — exactly one `<h1>` exists.
- All animated elements start in their visible end-state when `prefers-reduced-motion` is active (no content hidden behind un-triggered animations).
- CTAs have descriptive accessible labels (not just "Click here").
- Gradient background must maintain WCAG AA contrast against overlaid text.

**Responsive Behaviour**:
- ≥ 1024 px: two-column layout; large headline (clamp 48–80 px).
- 768–1023 px: single column, headline clamps smaller, visual moves below CTAs.
- < 768 px: single column, visual hidden or reduced to a small thumbnail; headline font-size clamps to 36–48 px.

---

### 3. Trusted By

**Purpose**: Establish immediate credibility and social proof by surfacing recognisable customer or partner brands.

**User Value**: Reassures sceptical enterprise visitors that peers and respected organisations have validated ShopWise.

**Layout**:
- Compact section (low vertical height — not competing with hero).
- Introductory label centred above logos.
- Single horizontal row of logos on desktop, 2-column grid or horizontal scroll on mobile.
- Logos are monochrome (low-opacity white or greyscale) to keep visual noise low.

**Key Content**:
- Section label: "Trusted by leading enterprises worldwide" (muted secondary text).
- 6–8 company logos (placeholder brands for initial implementation; designed to be swapped).
- Optional: a marquee/carousel for > 6 logos.

**UX Behaviour**:
- Auto-scrolling marquee on desktop (CSS animation or GSAP), seamlessly looping.
- Marquee pauses on pointer hover or keyboard focus on any logo.
- No marquee on mobile — static grid instead.
- No marquee under `prefers-reduced-motion` — logos display as a static row/grid.

**Accessibility**:
- Each logo `<img>` has a meaningful `alt` attribute (company name).
- If logos are rendered as SVG inline, they have a `<title>` element and `aria-labelledby`.
- Marquee container has `aria-label="Trusted by — company logos"` and `role="region"`.

**Responsive Behaviour**:
- ≥ 1024 px: horizontal marquee row.
- 768–1023 px: static 3-column grid.
- < 768 px: static 2-column grid or scrollable single row.

---

### 4. AI Decision Visualisation

**Purpose**: Demonstrate the AI capability in a way that is viscerally understood without reading documentation — the product's key differentiator made tangible.

**User Value**: Transforms an abstract concept ("AI helps you buy better") into a concrete, credible visual that builds confidence in the product.

**Layout**:
- Two-column on desktop (≥ 1024 px): copy/context left, visualisation right.
- Single-column on mobile: copy above, visualisation below.
- Dark section background, slightly lighter than the hero to create visual rhythm.

**Key Content**:
- Section eyebrow: "Powered by AI" (accent colour, small caps).
- Headline (`<h2>`): ≤ 8 words — e.g., "Your AI Analyst, Always On."
- Supporting paragraph (≤ 40 words): explains what the AI does in the context of a purchasing decision.
- AI Visualisation: an animated card/panel showing:
  - Input: a product search or vendor comparison query.
  - AI processing indicator (pulsing ring, token stream, or data graph populating).
  - Output: a scored recommendation list or decision summary.
- CTA: "Try the AI Workspace" → workspace route.

**UX Behaviour**:
- Scroll-triggered reveal: visualisation animates in when its container enters the viewport (IntersectionObserver or GSAP ScrollTrigger).
- The AI panel simulates processing: text tokens appear sequentially, scores increment, or a graph draws itself.
- All animations are absent under `prefers-reduced-motion`; the panel shows the completed/output state statically.

**Accessibility**:
- Animated visualisation has `aria-hidden="true"`. A companion `<p class="sr-only">` describes the demonstrated capability for screen reader users.
- Section heading is `<h2>`.
- CTA has a descriptive label.

**Responsive Behaviour**:
- ≥ 1024 px: side-by-side, visualisation is 55% width.
- < 1024 px: single column, visualisation is full width.

---

### 5. Features

**Purpose**: Enumerate the product's core capabilities in a scannable, visually organised grid that empowers visitors to evaluate fit quickly.

**User Value**: Lets evaluators match their specific pain points against ShopWise's features without a sales call.

**Layout**:
- Section headline (`<h2>`) + optional sub-headline centred above the grid.
- Feature card grid: 3 columns (desktop ≥ 1024 px), 2 columns (tablet 768–1023 px), 1 column (mobile < 768 px).
- Each card: icon (top), title, description.
- Cards use a subtle glass surface (semi-transparent dark background, thin border at 10–15% opacity white) to layer over the section background.

**Key Content** (proposed 6 features — final copy TBD by content):
1. **AI Vendor Scoring** — Automatically rank vendors by quality, price, and reliability signals.
2. **Real-Time Market Intelligence** — Live pricing benchmarks and supply-chain alerts.
3. **Procurement Workflow Automation** — Automate approvals, POs, and audit trails.
4. **Spend Analytics** — Visual dashboards for category spend and savings tracking.
5. **Collaborative Decision Workspace** — Team-based review, comments, and final approval flows.
6. **Enterprise Integrations** — Native connectors for SAP, Oracle, and Salesforce.

**UX Behaviour**:
- Cards reveal with a staggered fade-in + subtle upward translate on scroll entry (GSAP ScrollTrigger or CSS `@keyframes` with `animation-delay` stagger).
- On hover: card border transitions to a soft blue/purple glow (`box-shadow` with accent colour at low opacity); icon colour lightens.
- All hover/animation states are CSS transitions (not GSAP). Motion is removed under `prefers-reduced-motion`.

**Accessibility**:
- Icons are `aria-hidden="true"`; feature titles and descriptions provide full meaning independently.
- Each card is a semantic `<article>` or `<li>` within an appropriate list/grid container.
- Section headline is `<h2>`; card titles are `<h3>`.

**Responsive Behaviour**:
- Grid reflows as described above. No horizontal overflow at any breakpoint.
- Card minimum height is not fixed — content determines height; cards in the same row align to the tallest.

---

### 6. Workflow

**Purpose**: Show how a user moves from problem to outcome using ShopWise, removing the "how does it work?" barrier to conversion.

**User Value**: Eliminates uncertainty about onboarding complexity and time-to-value; builds confidence that getting started is straightforward.

**Layout**:
- Section headline (`<h2>`) + optional sub-headline.
- 4-step horizontal timeline on desktop: steps connected by a progress line.
- Vertical numbered list on mobile: steps stacked, connected by a vertical line.
- Each step: step number badge, title (≤ 5 words), description (≤ 20 words).

**Key Content** (proposed 4 steps):
1. **Connect Your Data** — Link your ERP, procurement system, or supplier database in minutes.
2. **Define Your Goals** — Tell the AI your priorities: cost, quality, speed, or sustainability.
3. **Run AI Analysis** — The AI surfaces ranked recommendations with full reasoning.
4. **Decide with Confidence** — Approve, adjust, or escalate — with a full audit trail.

**UX Behaviour**:
- Progress line draws from left to right (desktop) or top to bottom (mobile) as the section enters the viewport (GSAP ScrollTrigger on desktop; CSS on mobile).
- Step cards fade in sequentially as the line reaches them.
- All animation is absent under `prefers-reduced-motion`; all steps are visible and the line is fully rendered statically.

**Accessibility**:
- Steps are rendered as an ordered list (`<ol>`) semantically.
- Step numbers are visible in the DOM and not conveyed only via CSS counter or pseudo-element.
- Section headline is `<h2>`; step titles are `<h3>`.

**Responsive Behaviour**:
- ≥ 1024 px: horizontal 4-step layout.
- < 1024 px: vertical stacked steps.

---

### 7. Testimonials

**Purpose**: Provide third-party social proof from recognisable roles and organisations to reduce purchase risk.

**User Value**: Validates the platform's effectiveness with real-world outcomes from peers in similar roles.

**Layout**:
- Section headline (`<h2>`) centred above carousel.
- Carousel: 2 cards visible simultaneously on desktop (≥ 1024 px), 1 card on mobile.
- Each card: large quote mark glyph, quote text, author name, role, company name, optional avatar or company logo.
- Navigation: prev/next arrows flanking the carousel, dot indicators below.

**Key Content** (proposed 4 testimonials — final copy TBD):
1. VP Procurement, Fortune 500 retailer — quantified cost saving.
2. Head of Category Management, global logistics firm — time-to-decision reduction.
3. CTO, mid-market e-commerce brand — integration ease.
4. Director of Finance, manufacturing company — compliance improvement.

**UX Behaviour**:
- Auto-advance every 5 s.
- Pauses on hover (pointer enter) or keyboard focus on any card element.
- Pauses when `prefers-reduced-motion` is active (no auto-advance).
- Slide transition: opacity cross-fade (not a slide that would trigger motion sensitivity).
- Dot indicators update to reflect current slide.

**Accessibility**:
- Carousel root has `role="region"` and `aria-label="Customer testimonials"`.
- Active slide region has `aria-live="polite"`.
- Prev/next buttons have descriptive `aria-label` ("Previous testimonial" / "Next testimonial").
- Each dot indicator button has `aria-label="Go to testimonial N"` and `aria-current="true"` for the active dot.

**Responsive Behaviour**:
- ≥ 1024 px: 2-card visible layout.
- < 1024 px: 1-card layout; arrows and dots remain.

---

### 8. Pricing

**Purpose**: Communicate plan options, prices, and value clearly enough for a visitor to self-qualify and take the next step without sales intervention.

**User Value**: Removes pricing uncertainty — the most common conversion blocker — and gives each audience segment a clear, relevant path forward.

**Layout**:
- Section headline (`<h2>`) + billing toggle above pricing cards.
- Pricing cards in a horizontal row (desktop): 3 tiers.
- Cards stack vertically on mobile.
- Recommended tier: elevated (slightly larger, or accent border, or "Most Popular" badge).

**Key Content** (proposed 3 tiers):

| Tier | Price (Monthly) | Price (Annual) | Target |
|---|---|---|---|
| Starter | $49/mo | $39/mo (–20%) | Small teams, SMB |
| Pro | $149/mo | $119/mo (–20%) | Growing teams, mid-market |
| Enterprise | Contact Sales | Contact Sales | Large orgs, custom needs |

Each card includes:
- Plan name + optional subtitle.
- Price (or "Custom").
- 4–6 plan-specific feature bullets.
- CTA button (Starter/Pro: "Start Free Trial"; Enterprise: "Talk to Sales").
- Optional: "14-day free trial, no credit card required" note under Starter/Pro.

**UX Behaviour**:
- Billing toggle: clicking "Annual" updates all prices with a smooth number transition (or simple swap).
- Hovering a non-highlighted card subtly elevates it (transform: translateY(-4px), shadow increase).
- Hover is CSS transition.

**Accessibility**:
- Billing toggle: `role="switch"`, `aria-checked` reflects current state.
- Recommended badge: "Most Popular" label is present in the DOM, not just visually applied.
- Each pricing card is a `<section>` or `<article>` with a heading for the plan name.
- Feature lists are `<ul>` with `<li>` items.

**Responsive Behaviour**:
- ≥ 1024 px: 3-column horizontal layout.
- 768–1023 px: 2-column layout (Enterprise below).
- < 768 px: single column vertical stack.

---

### 9. FAQ

**Purpose**: Proactively address the most common purchase objections to reduce friction and build confidence.

**User Value**: Gives visitors answers without requiring a sales call, accelerating the decision-making process.

**Layout**:
- Section headline (`<h2>`) centred above accordion.
- Accordion list, max-width 800 px, centred.
- Each item: question as a toggle trigger, answer as collapsible panel.

**Key Content** (proposed questions):
1. How accurate is the AI's purchasing recommendations?
2. How does ShopWise integrate with existing ERP or procurement systems?
3. Is our procurement data secure and compliant with GDPR / SOC 2?
4. How long does onboarding take?
5. What support SLA do you offer?
6. Can we customise the AI's scoring criteria to match our supplier policy?
7. Is there a free trial?

**UX Behaviour**:
- Only one panel open at a time (single-open accordion behaviour).
- Expand: panel height animates from 0 to auto using a CSS transition on `max-height` or `grid-template-rows`.
- Transition removed under `prefers-reduced-motion`.
- Chevron icon on the right of each trigger rotates 180° on open; 0° on close.

**Accessibility**:
- Each trigger is a `<button>` with `aria-expanded` and `aria-controls` pointing to the panel.
- Each panel has a unique `id` and `role="region"` with `aria-labelledby` pointing to the trigger.
- Section headline is `<h2>`; question triggers convey their purpose without a heading role.

**Responsive Behaviour**:
- Accordion is single-column at all breakpoints. Width constrained to 800 px max on desktop.

---

### 10. Final CTA

**Purpose**: Provide a high-intent conversion moment for visitors who have consumed the full page and are ready to act.

**User Value**: Captures motivated visitors with a clear, low-friction path to start — without requiring them to scroll back to the hero.

**Layout**:
- Full-width section, high visual contrast from surrounding sections.
- Centred content: eyebrow label, headline, supporting statement, primary CTA, optional secondary CTA.
- Background: gradient (accent indigo/violet), or a bold dark surface with a glow/bloom effect behind the CTA area.

**Key Content**:
- Eyebrow: "Ready to transform your procurement?" (accent colour).
- Headline: ≤ 8 words — e.g., "Start Making Smarter Decisions Today."
- Supporting statement: ≤ 20 words — e.g., "Join hundreds of enterprises already using ShopWise. No credit card required."
- Primary CTA: "Start for Free" → same destination as hero primary CTA.
- Secondary CTA: "Talk to Sales" → sales contact route.

**UX Behaviour**:
- Section fades and scales in on scroll entry (GSAP or CSS). Static under `prefers-reduced-motion`.
- Primary CTA has a subtle pulse/glow animation drawing attention (paused under `prefers-reduced-motion`).

**Accessibility**:
- Section headline is `<h2>`.
- Primary and secondary CTAs are `<a>` elements or `<button>` elements with descriptive labels.
- No content is hidden only by animation-triggered opacity.

**Responsive Behaviour**:
- Centred single-column at all breakpoints. Padding increases on larger viewports.

---

### 11. Footer

**Purpose**: Provide complete navigational closure — links to all major destinations, legal compliance, and brand reinforcement.

**User Value**: Gives visitors who have scrolled to the bottom a clear map of where to go next without feeling stranded.

**Layout**:
- Multi-column on desktop: ShopWise logotype + descriptor (col 1), Product links (col 2), Company links (col 3), Legal links (col 4).
- Single-column stack on mobile.
- Copyright bar below main footer area: centred copyright notice and optional social links.

**Key Content**:

*Link groups:*
- **Product**: Features · Pricing · AI Workspace · Integrations · Changelog.
- **Company**: About · Blog · Careers · Press.
- **Legal**: Privacy Policy · Terms of Service · Cookie Policy.
- **Social**: LinkedIn · GitHub (icon links, new tab).

*Brand:*
- ShopWise logotype.
- Descriptor (≤ 10 words): "AI-powered procurement intelligence for modern enterprises."

*Copyright:*
- © [current year] ShopWise. All rights reserved.

**UX Behaviour**:
- Static — no animations in footer.
- Social icon links open in a new tab with `rel="noopener noreferrer"`.

**Accessibility**:
- Footer wrapped in `<footer>` landmark.
- Each link group has a visible heading rendered as `<h3>` (or `<p>` with `role="group"` + `aria-labelledby` if not in a heading hierarchy).
- Social links have visually hidden labels (e.g., "ShopWise on LinkedIn") via `<span class="sr-only">`.
- All links are keyboard-reachable.

**Responsive Behaviour**:
- ≥ 1024 px: 4-column layout.
- 768–1023 px: 2-column layout (brand + legal left; product + company right).
- < 768 px: single-column stack.

---

## Success Criteria

| Criterion | Metric | Target |
|---|---|---|
| Core Web Vitals — LCP | Time for largest contentful paint on first load | < 2.5 seconds |
| Core Web Vitals — CLS | Cumulative Layout Shift score | < 0.1 |
| Core Web Vitals — INP | Interaction to Next Paint | < 200 ms |
| Accessibility compliance | WCAG 2.1 AA audit findings | Zero critical or serious violations |
| Mobile usability | Horizontal overflow on 375 px viewport | Zero overflow at any scroll position |
| Reduced-motion compliance | Elements visible under `prefers-reduced-motion: reduce` | 100% — no content hidden by un-triggered animations |
| Hero conversion clarity | First-time visitors can describe the value proposition after hero alone | ≥ 70% accuracy in user testing |
| Pricing section self-sufficiency | Visitors can identify their recommended plan without external help | ≥ 80% task-completion rate in usability test |
| Keyboard navigability | All interactive elements reachable via keyboard Tab | 100% — no unreachable interactive elements |
| CTA reachability | All CTA buttons navigate to valid, non-404 destinations | 100% |

---

## Assumptions

- ShopWise has an existing app authentication route (for "Sign In") and an AI workspace route (for "Launch AI Workspace" / "Start for Free") that this page will link to. Exact route paths to be confirmed with the routing configuration in `apps/web/src/routes/`.
- Lenis smooth scroll is already configured at the app root in `apps/web`; this page can leverage it for smooth anchor scrolling. If not configured, smooth scroll will be handled via native CSS `scroll-behavior: smooth` on the `<html>` element.
- Company/customer logos for the Trusted By section are placeholder assets during initial implementation; final brand-approved logos will be supplied by the content team.
- Testimonial quotes are placeholder copy during initial implementation; final copy will be supplied by the marketing team.
- Pricing figures are illustrative; final pricing strategy is to be confirmed by the business team before launch.
- The page is a standalone TanStack Router route (e.g., `/` or `/landing`) within `apps/web`, rendered on the client side. SSR is not required for this version.
- The "AI Decision Visualisation" is a premium UI component (animated card/panel) — it does not call a live AI API during the page render; it is a designed demonstration.
- `@shopwise/ui` provides base primitives (Button, Card, Badge, etc.) that this page will consume. Any page-specific compound components will be built in `apps/web/src/features/landing/components/`.

## Dependencies

- `@shopwise/ui` — shared component primitives (Button, Card, Badge, Accordion, Tabs, etc.).
- `lucide-react` — icon library.
- `gsap` + `@gsap/react` — complex animations (Hero reveal, AI visualisation, Workflow timeline).
- `lenis` — smooth scroll (if already configured at app root).
- TanStack Router — routing and anchor navigation.
- Tailwind CSS v4 — all styling.
- Company/partner logos from content team (Trusted By section).
- Marketing copy from content team (Testimonials, final Pricing figures, FAQ answers).

## Risks

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| Lenis not configured at app root | Medium | Low | Fall back to CSS `scroll-behavior: smooth`; add a note to implementation tasks to verify Lenis setup first. |
| Final content (logos, testimonials, pricing) not available at implementation time | High | Medium | Use high-quality placeholder content; design components to accept content via props so swapping is trivial. |
| GSAP animations cause layout shift or CLS regression | Medium | High | Measure CLS after each animation phase; ensure GSAP animations do not modify layout-affecting properties (width, height, margin). |
| Glass effects cause performance regression on low-end devices | Low | Medium | Conditionally remove `backdrop-filter` on devices that report `prefers-reduced-transparency` or on detected low-end GPU. |
| Pricing section pricing strategy changes post-implementation | High | Low | Keep pricing values as configuration constants, not hard-coded strings. |

---

## Open Questions

- What are the exact TanStack Router route paths for: Sign In, AI Workspace, Free Trial sign-up, and Sales Contact? These determine all CTA href values.
- Is Lenis currently configured at the `apps/web` app root, and if so, at which route level is `<ReactLenis>` mounted?

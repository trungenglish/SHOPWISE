# ShopWise AI Decision Intelligence — Marketing Landing Page

**Version**: 2.1.0 **Status**: Draft **Feature Directory**: specs/001-marketing-landing-page **Created**: 2026-07-10 **Last Updated**: 2026-07-11

---

## Overview

The ShopWise marketing landing page is the primary public-facing surface that positions ShopWise as an **enterprise AI Decision Intelligence Platform** for procurement teams, category managers, purchasing teams, and business decision makers. Its singular objective is to create a high-trust, premium first impression that articulates the platform's value — transforming fragmented purchasing decisions into a single unified, explainable workspace — and drives visitors to launch a decision workspace or request an enterprise demonstration.

The page delivers a dark-first, visually excellent, enterprise-grade experience aligned with ShopWise's UI philosophy (Premium · Minimal · Enterprise · Modern · Dark-First) and must meet WCAG 2.1 AA accessibility criteria while performing to Core Web Vitals thresholds on all screen sizes.

---

## Clarifications

### Session 2026-07-11

- Q: Who is the primary audience the hero and opening sections should speak to directly? → A: Procurement specialists first — Procurement Team Leaders, Category Managers, and Purchasing Specialists. Hero copy speaks directly to procurement pain (supplier comparison, product evaluation, technical specifications, confident purchasing decisions). Executive value (cost reduction, governance, auditability, ROI) is deferred to Decision Confidence and Final CTA sections.
- Q: Should "ShopWise OS" appear as a visible, public-facing label on the landing page? → A: Drop "OS" from all public-facing text and headings. The public product name remains "ShopWise". "Operating System" is a positioning concept only. The section uses natural language headings (e.g., "Beyond Search. Beyond Chat.", "Why ShopWise Works Differently."). The term "Decision Operating System" may appear in supporting body copy, diagrams, or illustrations but not as the primary product name or navbar branding.
- Q: Are the AI visualisations on this marketing page decorative simulations or genuinely interactive? → A: Decorative simulation only. Scroll-triggered GSAP animations that auto-play once on viewport entry. Use realistic but static enterprise data. No click-driven state changes. No mock workflows requiring state management. Hover effects allowed only for subtle visual feedback — they must not change underlying content. Interactive product exploration belongs in the authenticated application or a dedicated product demo, not on the marketing landing page.
- Q: Is the Trusted By section a hard launch requirement or conditional on real enterprise logos? → A: Optional — must not block the landing page launch. Must not use fictional companies or placeholder logos implying commercial relationships. The section should be omittable if no verified logos are available. Credibility comes from the product narrative, workspace, and explainable decision process — not from unverified logos. Implementation must make the section easy to enable later through content configuration.
- Q: Should the differentiation section ("Why ShopWise Works Differently") appear before or after the workspace demonstration sections, or be merged? → A: Merge differentiation into The Search Crisis section. Within the solution column, explain that ShopWise is not another search engine, chatbot, or marketplace — but an AI Decision Intelligence Platform that maintains context, evidence, reasoning, workflows, and decision history. Remove "Why ShopWise OS" / "Why ShopWise Works Differently" as a standalone section. Section order after Search Crisis: AI Decision Workspace → Meet the Workspace → Intelligence Layer → Decision Confidence.

---

## Problem Statement

Enterprise purchasing decisions are fragmented across multiple browser tabs, spreadsheets, product catalogues, supplier quotations, technical documents, review sites, email threads, and internal approval workflows. This fragmentation creates:

- **Conflicting information**: Contradictory claims from multiple sources with no single source of truth.
- **Slow comparison**: Manually cross-referencing dozens of data points across disparate tools.
- **Low decision confidence**: Recommendations that are difficult to justify or trace back to evidence.
- **Weak audit trails**: No structured record of why a purchasing decision was made.
- **Difficult stakeholder collaboration**: No shared environment for teams to evaluate and approve together.

ShopWise currently lacks a dedicated marketing surface that communicates this problem clearly and positions the platform as the solution. Without it, prospective enterprise customers cannot understand ShopWise's value from a cold introduction, and conversion to workspace adoption or sales engagement is low.

---

## Goals

- Position ShopWise unambiguously as an enterprise AI Decision Intelligence Platform — not a consumer shopping app, marketplace, or generic SaaS product.
- Communicate the enterprise purchasing problem in a way that resonates immediately with procurement professionals and business decision makers.
- Demonstrate how ShopWise creates a unified decision workspace that produces confident, explainable purchasing decisions.
- Present the Intelligence Layer — the AI pipeline from business goal to human-approved recommendation — in plain, credible language.
- Drive visitors to the primary conversion action: launching a decision workspace or requesting an enterprise demonstration.
- Meet WCAG 2.1 AA accessibility standards.
- Achieve Core Web Vitals thresholds (LCP < 2.5 s, CLS < 0.1, INP < 200 ms).

## Non-Goals

- This page does not display pricing plans or free-trial offers.
- It does not present fictional customer testimonials or unverifiable performance claims.
- It does not implement authentication, sign-up forms, or account-creation flows — all CTAs link to app routes.
- It does not serve as a documentation hub, blog, or content marketing channel.
- Dynamic personalisation based on visitor firmographics is out of scope for this version.
- A/B testing infrastructure is out of scope; the page must be structured to be A/B-test-ready (clean section boundaries, clear CTA targets) without requiring testing tooling.

---

## Target Users

**Primary audience** — the hero and early sections speak directly to these roles:

| User Type | Description | Primary Need |
| --- | --- | --- |
| Procurement Team Leader | VP or Director of Procurement at a mid-to-large enterprise; accountable for purchasing outcomes and compliance | Replace fragmented comparison tools with a defensible, auditable decision environment for supplier evaluation |
| Category Manager | Operational buyer responsible for comparing suppliers and managing category spend | Accelerate supplier and product comparison; surface ranked, evidenced recommendations quickly |
| Purchasing Specialist | Day-to-day buyer who evaluates products, validates technical specifications, and manages supplier quotations | Compare options against requirements; validate technical claims and quotation terms in one place |

**Secondary audience** — addressed later in the page (Decision Confidence, Final CTA):

| User Type | Description | Primary Need |
| --- | --- | --- |
| Business Decision Maker | C-suite or department head who approves significant purchasing decisions | Gain confidence that ShopWise produces traceable, justifiable, auditable decisions their organisation can stand behind |
| Technical Evaluator | IT or engineering stakeholder validating enterprise-grade integration, security, and maintainability | Confirm the platform is robust, auditable, and integrates with existing enterprise systems |
| Executive Approver | Senior leader who reviews and signs off on procurement decisions; concerned with governance and ROI | Understand how ShopWise produces auditable decisions that reduce governance risk and justify cost |

---

## User Scenarios & Testing

### Scenario 1: Procurement leader forms a first impression

**Given**: A VP of Procurement clicks a link from a sales email or LinkedIn ad and lands on the page for the first time on a desktop browser. **When**: They read the hero headline, subtitle, and scan the page layout. **Then**: They immediately understand that ShopWise is an AI Decision Intelligence Platform for procurement — not a consumer app or a search engine — and see a clear CTA to explore the workspace. **Acceptance**: A first-time visitor from a procurement background can correctly describe ShopWise's core value proposition and target audience after viewing only the hero section — validated via user testing with >= 70% accuracy.

### Scenario 2: Visitor recognises the procurement fragmentation problem

**Given**: A category manager scrolls to the Search Crisis section. **When**: They read the problem-versus-solution comparison. **Then**: They recognise their own fragmented workflow in the problem column and understand that ShopWise resolves it with a unified workspace. **Acceptance**: The section is self-explanatory without supplementary tooltips. A procurement professional can map at least two of their own pain points to the listed legacy problems after a single read.

### Scenario 3: Visitor understands the AI Intelligence Layer

**Given**: A visitor scrolls to the Intelligence Layer section. **When**: They read the decision pipeline stages. **Then**: They understand that ShopWise processes a business goal through structured AI stages — planning, data gathering, specialist analysis, comparison, and recommendation — resulting in a human-approachable output. **Acceptance**: A non-technical visitor can describe the overall flow of the AI pipeline in plain terms after reading the section once.

### Scenario 4: Visitor navigates to conversion action

**Given**: A visitor has scrolled through the full page and reached the Final CTA section. **When**: They click "Launch Decision Workspace" or "Request Enterprise Demo". **Then**: They are routed to the correct app destination without friction. **Acceptance**: Both CTA buttons navigate to valid, non-404 routes. The final CTA section does not introduce any new obstacles (no pricing gates, no form walls before routing).

### Scenario 5: Visitor with reduced-motion preference

**Given**: A visitor has enabled prefers-reduced-motion at the operating-system level. **When**: They scroll through all page sections. **Then**: No GSAP timelines or CSS transitions play. All content is statically visible and fully readable. No content is hidden behind animation-gated reveal states. **Acceptance**: Every section passes a visual audit with prefers-reduced-motion: reduce active in browser devtools. Zero elements are invisible or clipped due to an unanimated initial state.

### Scenario 6: Mobile visitor

**Given**: A visitor opens the page on a 375 px-wide mobile device. **When**: They scroll through the entire page. **Then**: All sections reflow to single-column layouts, all text is legible, all interactive elements have tap targets >= 44 x 44 px, and no horizontal overflow exists. **Acceptance**: The page passes a manual review on a 375 px viewport with zero horizontal scroll and all tap targets meeting the size requirement.

### Scenario 7: Keyboard-only visitor navigates the page

**Given**: A visitor navigates the page using keyboard only (Tab, Enter, Space, arrow keys). **When**: They tab through all interactive elements (nav links, CTAs, accordion toggles). **Then**: Focus is always visible, focus order is logical (top-to-bottom, left-to-right), and no focus trap exists except in the intentionally modal mobile menu. **Acceptance**: A keyboard-only navigation audit finds no focus loss, no invisible focus rings, and no unreachable interactive elements.

---

## Functional Requirements

### Navigation (Navbar)

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-NAV-01 | The navbar displays the ShopWise logotype on the left and primary navigation links on the right. | Must Have | Logotype and nav links are visible on all viewport widths >= 375 px. |
| FR-NAV-02 | The navbar includes a primary CTA button ("Launch Decision Workspace") distinct from nav links. | Must Have | CTA button is visually distinct (filled vs. text links) and navigates to the correct destination. |
| FR-NAV-03 | On scroll, the navbar becomes sticky and applies a subtle backdrop blur / elevated surface effect. | Must Have | Navbar remains at top of viewport after scrolling 100 px. Background opacity transitions smoothly. |
| FR-NAV-04 | On mobile (< 768 px), navigation links collapse into a hamburger/menu toggle. | Must Have | Menu toggle is visible on mobile. Activating it reveals all nav links. Focus is trapped within the open menu. Menu closes on link click or pressing Escape. |
| FR-NAV-05 | Nav links scroll smoothly to their respective page sections. | Should Have | Clicking a nav link scrolls to the correct section anchor. Active link is visually highlighted based on scroll position. |
| FR-NAV-06 | The navbar provides a "Sign In" link for returning users. | Should Have | "Sign In" link is present and navigates to the correct app authentication route. |

**Navigation link targets**: Problem · Workspace · Intelligence · Confidence · Contact

### Hero Section

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-HERO-01 | The hero presents a bold, large-format headline that communicates the enterprise AI Decision Intelligence positioning in <= 10 words. | Must Have | Headline is the largest typographic element on the page. Immediately visible on load. |
| FR-HERO-02 | A supporting subtitle explains how ShopWise moves teams from fragmented comparison to confident purchasing decisions in <= 30 words. | Must Have | Subtitle is present below the headline with a clear typographic hierarchy. |
| FR-HERO-03 | Two CTAs are present: primary ("Launch Decision Workspace") and secondary ("See How ShopWise Reasons"). | Must Have | Both CTAs are keyboard-operable, have accessible labels, and navigate to valid routes. The secondary CTA scrolls to the Intelligence Layer section. |
| FR-HERO-04 | A hero visual — abstract representation of the AI decision workspace or intelligence pipeline — supports the headline. | Must Have | Visual is present on desktop (>= 1024 px). On mobile, the visual is deprioritised or hidden. |
| FR-HERO-05 | The hero applies a dark gradient background (deep navy/charcoal through indigo accents) that establishes the enterprise visual language. | Must Have | Gradient is visible. Text contrast meets WCAG AA (>= 4.5:1 for body, >= 3:1 for large text). |
| FR-HERO-06 | Entry animations (headline, subtitle, CTAs staggered reveal) play on first load. | Should Have | Animations are absent or replaced by the static end-state when prefers-reduced-motion is active. |

### Trusted By Section

> **This section is optional and must not block the landing page launch.** It should only appear in production when verified, brand-approved enterprise logos are available. If no approved logos are supplied, the section is omitted entirely from the rendered page.

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-TRUST-01 | The section is implemented with a content-configuration flag; when no verified logos are available, the section is hidden and the page renders correctly without it. | Must Have | The page layout is correct and no gap or broken region exists when the section is disabled. |
| FR-TRUST-02 | When enabled, the section displays >= 5 and <= 10 verified, brand-approved enterprise logos. Fictional companies, generic placeholders, or logos implying unverified commercial relationships must not appear in production. | Must Have | All visible logos are approved by the content team before the section is enabled in production. |
| FR-TRUST-03 | Each logo includes an accessible text alternative (alt attribute or visually hidden label). | Must Have | Screen readers announce each logo's organisation name. |
| FR-TRUST-04 | Logos are displayed in a single horizontal row on desktop, a 2-column grid on mobile, with no overflow. | Must Have | No horizontal scroll. Logos wrap or reduce gracefully on small screens. |
| FR-TRUST-05 | A subtle auto-scrolling marquee is applied on desktop when >= 6 logos are present. | Should Have | Marquee pauses on hover. Marquee absent when prefers-reduced-motion is active. |
| FR-TRUST-06 | An introductory label (e.g., "Trusted by leading enterprises") precedes the logos. | Should Have | Label uses a muted, secondary text style. |

### The Search Crisis (Problem Statement + Differentiation Section)

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-PROB-01 | The section presents a clear problem/solution comparison contrasting the legacy fragmented procurement approach with the ShopWise unified workspace approach. | Must Have | Both the problem state and the ShopWise outcome are clearly labelled and visually distinct. |
| FR-PROB-02 | The legacy problem column presents >= 4 concrete pain points expressed in procurement-specific language: comparing suppliers across dozens of tabs; conflicting technical claims with no verification; spreadsheet comparisons that become stale; quotation and specification documents in separate tools; purchasing decisions with no audit trail; low confidence and slow approval cycles. | Must Have | Pain points use language recognisable to Procurement Team Leaders, Category Managers, and Purchasing Specialists. |
| FR-PROB-03 | The ShopWise outcome column presents >= 4 corresponding outcomes: one decision workspace for all comparison activity; verified evidence and structured supplier data; AI-powered comparison of suppliers, products, and specifications; full decision audit trail traceable to source evidence; confident recommendations with explainable reasoning. | Must Have | Outcomes directly correspond to the problem items and are expressed as concrete procurement benefits. |
| FR-PROB-04 | The ShopWise outcome column includes a brief differentiation statement explaining that ShopWise is not another search engine, chatbot, or marketplace — it is an AI Decision Intelligence Platform that maintains context, evidence, reasoning, workflows, and decision history across the full procurement lifecycle. | Must Have | The differentiation statement is present within the solution column. It does not use the label "ShopWise OS". |
| FR-PROB-05 | Section headline (h2) frames the fragmentation crisis in plain language. | Must Have | Headline is present as h2. |
| FR-PROB-06 | The section does not present performance statistics that cannot be verified at this stage. | Must Have | No quantified claims (e.g., "50% faster") appear unless clearly marked as illustrative. |

### AI Decision Workspace Section

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-WORK-01 | The section presents a decorative simulation visualisation demonstrating how the platform compares suppliers or products within a decision workspace. The visualisation is scroll-triggered and auto-plays once on viewport entry — it is not interactive. No click-driven state changes occur. | Must Have | Visualisation is present, self-playing, comprehensible without tooltips, and does not respond to user clicks. |
| FR-WORK-02 | The visualisation uses realistic but static enterprise data. It illustrates >= 3 workspace capabilities: requirement matching, vendor scoring, risk evaluation, evidence verification, recommendation output. Each capability is labelled within the panel. | Must Have | Labels are present. Data is realistic but clearly illustrative (e.g., fictional supplier names). No live API calls. |
| FR-WORK-03 | A headline (h2) and supporting copy explain the workspace concept in plain language (<= 40 words for body copy). | Must Have | Headline and body copy are present. No unexplained technical jargon. |
| FR-WORK-04 | A two-column layout is used on desktop (>= 1024 px): copy left, visualisation right. Single-column on mobile. | Must Have | Columns are side-by-side on >= 1024 px. Single column on < 1024 px. |
| FR-WORK-05 | The visualisation is GSAP scroll-triggered: it animates once when entering the viewport. Under prefers-reduced-motion, all elements are immediately visible in their final static state. Hover effects may provide subtle visual feedback (e.g., a row highlight) but must not change the underlying content or data. | Must Have | Animation plays once on scroll entry. Static final state shown under prefers-reduced-motion. Hover changes only visual appearance, not content. |
| FR-WORK-06 | A CTA ("Launch Decision Workspace") is present within or immediately below this section. | Should Have | CTA is keyboard-operable and navigates to the correct route. |

### Meet the Workspace Section

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-MEET-01 | The section presents >= 4 and <= 6 core workspace capabilities, each with an icon, a short title (<= 5 words), and a description (<= 25 words). | Must Have | All capabilities are present. Icon, title, and description are visually distinct per item. |
| FR-MEET-02 | Capabilities are arranged in a responsive grid: 3-column on desktop, 2-column on tablet, 1-column on mobile. | Must Have | Grid reflows correctly at breakpoints. No overflow or clipped content. |
| FR-MEET-03 | Icons use lucide-react exclusively. | Must Have | No other icon library is used. Icons are aria-hidden with meaning conveyed by text. |
| FR-MEET-04 | Hovering a capability card applies a subtle glass-effect highlight or border glow. | Should Have | Hover state is CSS transition only (not GSAP). Hover absent on touch devices. |
| FR-MEET-05 | Section headline (h2) introduces the capability grid. Card titles are h3. | Must Have | Section headline is h2. |

<!-- FR-WHY section removed. Differentiation messaging merged into The Search Crisis section (FR-PROB-04). See Clarifications > Session 2026-07-11, Q5. -->

### Intelligence Layer Section

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-INTEL-01 | The section presents the AI decision pipeline as 7 sequential labelled stages: Business Goal, Planning, Data and Evidence, Specialist Analysis, Comparison Logic, Recommendation, Human Approval. | Must Have | All 7 pipeline stages are present and labelled. |
| FR-INTEL-02 | Each stage has a short label (<= 4 words) and a brief descriptor (<= 15 words). | Must Have | Labels and descriptors are present for all stages. |
| FR-INTEL-03 | The pipeline is rendered as a horizontal flow on desktop and a vertical sequence on mobile. | Must Have | Layout transitions correctly at breakpoints without horizontal overflow. |
| FR-INTEL-04 | A connecting visual (line, arrow, or path) links the stages to convey sequence. | Should Have | Connector is present on desktop. Simplified or absent on mobile without loss of meaning. |
| FR-INTEL-05 | A headline (h2) introduces the section with framing copy (<= 30 words). | Must Have | Headline is h2. |
| FR-INTEL-06 | Pipeline stages animate in sequentially as the section enters the viewport (scroll-triggered). | Should Have | Sequential animation plays on scroll. All stages immediately visible in final state under prefers-reduced-motion. |

### Decision Confidence Section

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-CONF-01 | The section shows the final transition from AI analysis to an actionable purchasing decision, featuring: evidence status, confidence indicator, risk status, and auditability signal. | Must Have | All four signal types are present and labelled. |
| FR-CONF-02 | The section does not claim unsupported real-world performance statistics. All output indicators are presented as illustrative examples within the platform UI. | Must Have | No quantified performance claims appear (e.g., no "reduces costs by X%"). |
| FR-CONF-03 | A headline (h2) and supporting copy (<= 30 words) frame the decision confidence concept. The copy targets executive approvers and business decision makers — it emphasises governance, auditability, and justified ROI rather than operational procurement detail. | Must Have | Headline is h2. Copy speaks to executive value without repeating procurement-specialist language from earlier sections. |
| FR-CONF-04 | A visual output card or panel illustrates a confident, explainable recommendation using a decorative simulation. The panel uses realistic but static data. It does not respond to user clicks and requires no state management. Hover effects may be used for subtle visual feedback only. | Must Have | Visual is present, self-contained, and comprehensible without supplementary explanation. No interactivity beyond hover feedback. |
| FR-CONF-05 | A CTA ("Launch Decision Workspace") is present within or immediately below this section. | Should Have | CTA is keyboard-operable and navigates to the correct route. |

### Final CTA Section

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-CTA-01 | The section presents a single, bold call-to-action with a headline (<= 8 words), a supporting statement (<= 20 words), and a primary CTA button. | Must Have | Headline, supporting statement, and primary CTA button are all present. |
| FR-CTA-02 | The primary CTA button label is "Launch Decision Workspace" — matching the hero primary CTA. | Must Have | CTA label is consistent with the hero. Destination route is the same. |
| FR-CTA-03 | A secondary CTA ("Request Enterprise Demo" or "Talk to Sales") is present below the primary button. | Must Have | Secondary CTA is present and navigates to a valid sales contact or demo-request destination. |
| FR-CTA-04 | The section uses a high-contrast background (gradient or solid dark with glow) that differentiates it from the preceding Decision Confidence section. | Must Have | Visual separation between sections is evident at a glance. |
| FR-CTA-05 | The section contains no pricing plans, free-trial claims, or consumer-SaaS conversion patterns. | Must Have | No pricing copy, "no credit card" language, or trial-related messaging appears in this section. |

### Footer

| ID | Requirement | Priority | Acceptance Criteria |
| --- | --- | --- | --- |
| FR-FOOT-01 | The footer includes the ShopWise logotype, a brief descriptor (<= 10 words), and grouped navigation links (Platform, Company, Legal). | Must Have | All three link groups are present. |
| FR-FOOT-02 | Each link group has an accessible heading rendered as the correct heading level or a labelled group. | Must Have | Headings are present and use the correct semantic element. |
| FR-FOOT-03 | Legal links (Privacy Policy, Terms of Service) are present and navigate to valid routes. | Must Have | Both links are present and reachable. |
| FR-FOOT-04 | Social media links are present (at minimum LinkedIn and GitHub). Each opens in a new tab with rel="noopener noreferrer" and has a visually hidden accessible label. | Should Have | Links are present. Screen readers announce link purpose. |
| FR-FOOT-05 | A copyright notice with the current year is present at the bottom of the footer. | Must Have | Copyright text is present and displays the correct year. |
| FR-FOOT-06 | The footer uses a multi-column layout on desktop and a single-column stack on mobile. | Must Have | Layout reflows correctly on all viewports. |
| FR-FOOT-07 | The footer descriptor reflects the enterprise AI Decision Intelligence positioning. | Must Have | Footer descriptor does not use consumer-SaaS language. |

---

## Section Specifications

> This section provides the complete design brief for each of the 9 content sections (plus Navbar and Footer). The standalone "Why ShopWise OS" section has been removed; its differentiation content is merged into Section 3 (The Search Crisis). Implementation teams should treat each sub-section as a standalone design contract.

---

### 1. Navbar

**Purpose**: Persistent wayfinding and primary conversion entry point accessible from any scroll position.

**User Value**: Lets visitors orient themselves, jump to key sections, and take the conversion action without scrolling back to the top.

**Layout**:

- Fixed/sticky to top of viewport.
- Full-width container, max-width constrained (1280 px) and centred.
- Left: ShopWise logotype (wordmark).
- Centre (desktop): nav links — Problem · Workspace · Intelligence · Confidence · Contact.
- Right (desktop): "Sign In" text link + "Launch Decision Workspace" filled button.
- Mobile (< 768 px): logotype left, hamburger icon right. Nav links in a slide-down or overlay drawer.

**Key Content**:

- ShopWise logotype.
- Nav links: Problem · Workspace · Intelligence · Confidence · Contact.
- "Sign In" (secondary, text weight).
- "Launch Decision Workspace" (primary CTA, filled, accent colour).

**UX Behaviour**:

- On initial load: transparent or very low-opacity background.
- After scrolling > 100 px: background transitions to a dark surface with backdrop-filter: blur and a subtle bottom border.
- Active section highlight: nav link corresponding to the visible section is highlighted.
- Mobile menu: hamburger toggles a full-width drawer. Drawer closes on link tap or Escape key.
- Smooth scroll to anchor on link click.

**Accessibility**:

- nav element with aria-label="Main navigation".
- Mobile menu button: aria-expanded, aria-controls pointing to the menu panel.
- All links and the CTA button are keyboard-focusable with visible focus rings.
- Focus trapped within the open mobile drawer.

---

### 2. Hero

**Purpose**: Create an immediate, high-impact first impression that communicates ShopWise's enterprise AI Decision Intelligence positioning and triggers the primary conversion action.

**User Value**: Answers "What is this?", "Who is it for?", and "Why should I care?" within the first 5 seconds, with a clear path to act.

**Layout**:

- Full-viewport-height (100 dvh) section.
- Two-column on desktop (>= 1024 px): content left (60%), visual right (40%).
- Single-column centred on mobile: headline, subtitle, CTAs stacked; visual below or hidden.
- Gradient background: deep navy/charcoal (#0a0a0f base) through indigo to soft blue/violet, with subtle grid or glow overlay.

**Key Content**:

- Eyebrow label (e.g., "AI Decision Intelligence for Enterprise Procurement" — small caps, accent colour).
- Headline: <= 10 words — e.g., "From Fragmented Research to Confident Purchasing Decisions."
- Subtitle: <= 30 words — explains that ShopWise unifies comparison, evidence, and AI reasoning into one decision workspace for procurement teams.
- Primary CTA: "Launch Decision Workspace" — workspace route.
- Secondary CTA: "See How ShopWise Reasons" — smooth scroll to Intelligence Layer section.
- Hero visual: abstract representation of the AI decision workspace — e.g., an animated intelligence pipeline, a decision panel with confidence score and evidence status, or an isometric workspace illustration.

**UX Behaviour**:

- Entry animation (GSAP): headline staggered word reveal (opacity 0->1, y 30px->0), subtitle fades in, CTAs slide in, visual scales in.
- All animations absent / replaced with static end-state under prefers-reduced-motion.
- Secondary CTA scrolls smoothly to the Intelligence Layer section anchor.
- Subtle particle, grid, or noise overlay adds depth without distracting from the headline.

**Accessibility**:

- Headline is the h1 of the page — exactly one h1 exists.
- All animated elements start in their visible end-state when prefers-reduced-motion is active.
- CTAs have descriptive accessible labels.
- Gradient background must maintain WCAG AA contrast against overlaid text.

**Responsive Behaviour**:

- > = 1024 px: two-column layout; large headline (clamp 48–80 px).
- 768–1023 px: single column, headline clamps smaller, visual moves below CTAs.
- < 768 px: single column, visual hidden or reduced; headline font-size clamps to 36–48 px.

---

### 3. Trusted By (Optional)

> **This section is optional.** It must not block the landing page launch. Render this section only when at least one verified, brand-approved enterprise logo is available. If no approved logos are supplied, omit the section entirely and ensure the page layout remains intact without it. Implement via a content-configuration flag (e.g., an empty logos array causes the section to return null).

**Purpose**: Establish enterprise credibility by surfacing verified partner or customer organisations.

**User Value**: Reassures enterprise procurement visitors that ShopWise is trusted by peers, reducing evaluation risk — but only when logos are genuine.

**Layout**:

- Compact section (low vertical height).
- Introductory label centred above logos.
- Single horizontal row of logos on desktop, 2-column grid on mobile.
- Logos are monochrome (low-opacity white or greyscale).

**Key Content**:

- Section label: "Trusted by leading enterprises" (muted secondary text).
- > = 5 verified, brand-approved enterprise logos. No fictional companies, generic shields, or placeholders implying commercial relationships.

**UX Behaviour**:

- Auto-scrolling marquee on desktop when >= 6 logos are present, seamlessly looping.
- Marquee pauses on pointer hover or keyboard focus.
- No marquee on mobile — static grid.
- No marquee under prefers-reduced-motion — logos display as a static row/grid.

**Accessibility**:

- Each logo img has a meaningful alt attribute (organisation name).
- Marquee container has aria-label="Trusted by — enterprise logos" and role="region".

---

### 4. The Search Crisis

**Purpose**: Articulate the fragmented enterprise procurement problem in concrete, procurement-specific terms; position ShopWise as the structural solution; and correct the visitor's mental model about what kind of platform ShopWise is — before the workspace demonstration sections.

**User Value**: Validates the pain points that Procurement Team Leaders, Category Managers, and Purchasing Specialists live with daily, then shows how ShopWise resolves them — and establishes that ShopWise is fundamentally different from search engines, chatbots, or marketplaces.

**Layout**:

- Section headline (h2) + optional introductory copy (1–2 sentences) centred above.
- Two-column comparison layout on desktop: Legacy approach (left) vs. ShopWise approach (right).
- Single-column on mobile (legacy column above, ShopWise column below).
- Visual differentiation between the two columns: muted/degraded styling for legacy vs. accent/glowing treatment for ShopWise.

**Key Content**:

- Headline: <= 8 words — e.g., "Enterprise Procurement is Broken."
- Legacy column heading: e.g., "How teams buy today."
- Legacy problem items (>= 4, expressed in procurement language):
  - Comparing suppliers across dozens of browser tabs with no single source of truth.
  - Conflicting technical claims and quotation terms with no verification layer.
  - Spreadsheet comparisons that become stale as data changes.
  - Product catalogues, technical documents, and supplier quotes in separate disconnected tools.
  - Purchasing decisions with no structured evidence trail or audit record.
  - Low confidence, slow approval cycles, and hard-to-justify recommendations.
- ShopWise column heading: e.g., "How ShopWise works."
- ShopWise outcome items (>= 4, directly paired with problem items):
  - One decision workspace with all suppliers, products, and evidence in one place.
  - Verified evidence and structured supplier data — claims checked against source documents.
  - AI-powered comparison of suppliers, products, and technical specifications that updates as data changes.
  - All documents, quotations, and specifications managed within the workspace.
  - Full decision audit trail — every step, source, and approval recorded and traceable.
  - Confident, explainable recommendations the whole organisation can stand behind.
- **Differentiation statement** (within the ShopWise column or as a closing note below):
  - ShopWise is not another search engine, a chatbot, or a marketplace.
  - It is an AI Decision Intelligence Platform — a persistent environment that maintains context, evidence, reasoning, workflows, collaboration, and decision history across the full procurement process.
  - The term "Decision Operating System" may appear here in supporting copy or an illustrative diagram to reinforce this concept.

**Accessibility**:

- Both columns are semantically grouped (e.g., two section or div role="group" elements with aria-labelledby).
- List items use ul/li semantics.
- Section headline is h2.
- Differentiation statement is part of the normal content flow — not hidden or aria-hidden.

---

### 5. AI Decision Workspace

**Purpose**: Demonstrate the platform's core capability — comparing suppliers or products within a unified workspace — viscerally and without requiring documentation.

**User Value**: Transforms the abstract concept of "AI helps you decide" into a concrete, credible visual that builds confidence in the product.

**Layout**:

- Two-column on desktop (>= 1024 px): copy left, visualisation right.
- Single-column on mobile: copy above, visualisation below.
- Dark section background, slightly lighter than the hero.

**Key Content**:

- Section eyebrow: "AI Decision Workspace" (accent colour, small caps).
- Headline (h2): <= 8 words — e.g., "One Workspace. Every Signal. One Decision."
- Supporting paragraph (<= 40 words): explains how the workspace aggregates requirements, evidence, and AI analysis into a single environment.
- Workspace visualisation: a decorative simulation panel — scroll-triggered, auto-playing GSAP animation — showing requirement matching matrix, vendor/supplier score comparison, risk and evidence status indicators, and recommendation output. Uses realistic but static fictional enterprise data. No interactivity beyond subtle hover feedback on individual rows.
- CTA: "Launch Decision Workspace" — workspace route.

**UX Behaviour**:

- The visualisation panel animates once when entering the viewport (GSAP ScrollTrigger). Fields appear, scores populate, and status indicators update sequentially.
- Under prefers-reduced-motion: the panel is immediately visible in its completed final state.
- Hover may apply a subtle row highlight (CSS transition only). Hover must not change the underlying data or content.
- No click-driven state changes. No mock workflow state management.

**Accessibility**:

- Animated visualisation has aria-hidden="true". A companion sr-only paragraph describes the demonstrated capability for screen reader users.
- Section heading is h2. CTA has a descriptive label.

---

### 6. Meet the Workspace

**Purpose**: Enumerate the platform's core workspace capabilities in a scannable, visually organised grid.

**User Value**: Lets procurement professionals and technical evaluators quickly assess whether ShopWise covers their required use cases without a sales call.

**Layout**:

- Section headline (h2) + optional sub-headline centred above the grid.
- Capability card grid: 3 columns (desktop >= 1024 px), 2 columns (tablet 768–1023 px), 1 column (mobile < 768 px).
- Each card: icon (top), title, description.
- Cards use a subtle glass surface (semi-transparent dark background, thin border at 10–15% opacity white).

**Key Content** (6 capabilities):

1. Product and Supplier Canvas — A unified view of all options, structured for direct comparison.
2. Requirement Matching — Automatically map business requirements to each supplier or product.
3. Trust and Evidence Centre — Verify claims, documents, and supplier credentials in one place.
4. Audit Trail — A full, traceable record of every decision step and approval action.
5. Team Collaboration — Assign reviewers, gather comments, and reach consensus without leaving the workspace.
6. Enterprise Integrations — Connect to ERP, procurement, and supplier management systems.

**Accessibility**:

- Icons are aria-hidden; titles and descriptions provide full meaning independently.
- Each card is a semantic article or li within an appropriate list/grid container.
- Section headline is h2; card titles are h3.

---

<!-- Section 7 (Why ShopWise OS) removed. Differentiation content merged into Section 4 (The Search Crisis). See Clarifications > Session 2026-07-11, Q5. -->

### 7. Intelligence Layer

**Purpose**: Present the AI decision pipeline in plain language — showing how ShopWise transforms a business goal into a human-approved, explainable recommendation.

**User Value**: Builds trust in the AI's approach by showing it is structured, transparent, and human-gated — not a black box.

**Layout**:

- Section headline (h2) + introductory copy (<= 30 words).
- Horizontal pipeline/timeline on desktop (>= 1024 px): stages left to right, connected by a line or arrow path.
- Vertical sequence on mobile: stages stacked top to bottom.
- Each stage: numbered or icon badge, short label, brief descriptor.

**Key Content** (7 pipeline stages):

1. Business Goal — The team defines what they need to buy and why.
2. Planning — ShopWise structures the decision: criteria, constraints, and stakeholders.
3. Data and Evidence — The platform gathers and verifies supplier data, documents, and market signals.
4. Specialist Analysis — AI agents evaluate each option against the defined criteria.
5. Comparison Logic — Options are ranked with full reasoning and evidence links.
6. Recommendation — A confident, explainable recommendation is surfaced to the team.
7. Human Approval — The team reviews, adjusts, and approves — creating the audit record.

**UX Behaviour**:

- Scroll-triggered: stages animate in sequentially as the pipeline enters the viewport (GSAP ScrollTrigger).
- Connector line draws from left to right (desktop) or top to bottom (mobile) as stages appear.
- Under prefers-reduced-motion: all stages visible immediately in their final state; connector fully rendered statically.

**Accessibility**:

- Stages rendered as an ordered list (ol) semantically.
- Stage numbers visible in the DOM (not conveyed only via CSS pseudo-elements).
- Section headline is h2; stage labels are h3.
- Connector graphics are aria-hidden="true".

---

### 8. Decision Confidence

**Purpose**: Show the final output of the ShopWise decision process — a confident, explainable purchasing decision that the team can act on and justify to stakeholders and executive approvers.

**User Value**: Demonstrates the tangible outcome of using ShopWise: a recommendation that is evidenced, risk-assessed, and audit-ready. This section speaks primarily to the **secondary audience** — Business Decision Makers, Executive Approvers, and Technical Evaluators — who need to understand the governance, auditability, and ROI justification that ShopWise provides.

**Layout**:

- Section headline (h2) + supporting copy.
- A large, premium decision output card or panel as the centrepiece — decorative simulation only.
- CTA below the card.

**Key Content**:

- Headline: <= 8 words — e.g., "Every Decision. Evidenced. Explainable. Approved."
- Supporting copy (<= 30 words): targets executive approvers — ShopWise produces a decision the whole organisation can stand behind, traceable to evidence, assessed for risk, and approved by the right people.
- Decision output panel (illustrative decorative simulation — static, fictional data, auto-plays on scroll entry):
  - Evidence status: e.g., "12 sources verified — 3 flagged for review."
  - Confidence indicator: e.g., a visual score or signal (High / Medium / Low — labelled, not just colour-coded).
  - Risk status: e.g., "2 risk factors identified — 1 mitigated."
  - Auditability signal: e.g., "Full audit trail available — 7 decision steps recorded."
- CTA: "Launch Decision Workspace" — workspace route.

**UX Behaviour**:

- Panel animates once on scroll entry (GSAP ScrollTrigger): signal indicators appear sequentially.
- Under prefers-reduced-motion: panel is immediately visible in its completed final state.
- Hover may apply subtle visual feedback (e.g., a field highlight). No click-driven changes to content or data.

**Accessibility**:

- All four decision signal types are labelled in text — no information conveyed by colour alone.
- Section headline is h2. Panel content is readable by screen readers (not aria-hidden).
- CTA has a descriptive label.

---

### 9. Final CTA

**Purpose**: Provide a high-intent conversion moment for visitors who have consumed the full page and are ready to engage.

**User Value**: Captures motivated visitors with a clear, enterprise-appropriate path to launch a workspace or speak to sales.

**Layout**:

- Full-width section, high visual contrast from the Decision Confidence section.
- Centred content: eyebrow label, headline, supporting statement, primary CTA, secondary CTA.
- Background: deep gradient (accent indigo/blue) or a bold dark surface with a glow/bloom effect.

**Key Content**:

- Eyebrow: "Ready to transform your procurement?" (accent colour).
- Headline: <= 8 words — e.g., "Start Making Confident Purchasing Decisions."
- Supporting statement: <= 20 words — e.g., "Join enterprise teams using ShopWise to unify comparison, evidence, and AI reasoning."
- Primary CTA: "Launch Decision Workspace" — workspace route.
- Secondary CTA: "Request Enterprise Demo" or "Talk to Sales" — sales contact or demo-request route.

**UX Behaviour**:

- Section fades and scales in on scroll entry (GSAP or CSS). Static under prefers-reduced-motion.
- Primary CTA has a subtle pulse/glow animation (paused under prefers-reduced-motion).

**Accessibility**:

- Section headline is h2. Primary and secondary CTAs are keyboard-operable with descriptive labels.
- No content hidden only by animation-triggered opacity.

---

### 10. Footer

**Purpose**: Provide complete navigational closure — links to all major destinations, legal compliance, and brand reinforcement.

**Layout**:

- Multi-column on desktop: ShopWise logotype + descriptor (col 1), Platform links (col 2), Company links (col 3), Legal links (col 4).
- Single-column stack on mobile.
- Copyright bar below main footer area.

**Key Content**:

- Platform links: AI Workspace · Decision Intelligence · Enterprise Integrations · Security.
- Company links: About · Careers · Press · Contact.
- Legal links: Privacy Policy · Terms of Service · Cookie Policy.
- Social links: LinkedIn · GitHub (icon links, new tab).
- Descriptor (<= 10 words): "AI Decision Intelligence for Enterprise Procurement."
- Copyright: © [current year] ShopWise. All rights reserved.

**Accessibility**:

- Footer wrapped in footer landmark.
- Each link group has a visible heading rendered as h3.
- Social links have visually hidden labels via sr-only span.
- All links are keyboard-reachable.

**Responsive Behaviour**:

- > = 1024 px: 4-column layout.
- 768–1023 px: 2-column layout.
- < 768 px: single-column stack.

---

## Success Criteria

| Criterion | Metric | Target |
| --- | --- | --- |
| Core Web Vitals — LCP | Time for largest contentful paint on first load | < 2.5 seconds |
| Core Web Vitals — CLS | Cumulative Layout Shift score | < 0.1 |
| Core Web Vitals — INP | Interaction to Next Paint | < 200 ms |
| Accessibility compliance | WCAG 2.1 AA audit findings | Zero critical or serious violations |
| Mobile usability | Horizontal overflow on 375 px viewport | Zero overflow at any scroll position |
| Reduced-motion compliance | Elements visible under prefers-reduced-motion: reduce | 100% — no content hidden by un-triggered animations |
| Enterprise positioning clarity | First-time visitors from procurement can correctly describe ShopWise as an AI Decision Intelligence Platform after hero section alone | >= 70% accuracy in user testing |
| Problem resonance | Procurement professionals can map >= 2 pain points to the Search Crisis section after a single read | >= 65% in usability test |
| Intelligence Layer comprehension | Non-technical visitors can describe the AI pipeline flow in plain terms after reading the section | >= 60% accuracy in usability test |
| CTA reachability | All CTA buttons navigate to valid, non-404 destinations | 100% |
| Keyboard navigability | All interactive elements reachable via keyboard Tab | 100% — no unreachable interactive elements |

---

## Assumptions

- ShopWise has an existing app authentication route (for "Sign In") and an AI workspace route (for "Launch Decision Workspace"). Exact route paths to be confirmed with apps/web/src/routes/.
- A sales contact or demo-request route exists or will exist for the secondary CTA destinations.
- Lenis smooth scroll is already configured at the app root in apps/web; this page can leverage it for smooth anchor scrolling.
- The Trusted By section is conditionally rendered. It will not appear in production until verified, brand-approved enterprise logos are supplied by the content team. The implementation uses a content-configuration flag (empty logos array → section returns null). No placeholder or fictional logos appear in production.
- The page is a standalone TanStack Router route (e.g., / or /landing) within apps/web, rendered on the client side. SSR is not required for this version.
- The AI Decision Workspace visualisation and Decision Confidence panel are premium UI components — they do not call a live AI API during the page render; they are designed demonstrations of the platform's output.
- @shopwise/ui provides base primitives (Button, Card, Badge, etc.) that this page will consume. Page-specific compound components will be built in apps/web/src/features/landing/components/.
- All performance statistics and outcome claims are illustrative and described as platform capabilities, not verified customer results, until real customer data is approved by the business team.

## Dependencies

- @shopwise/ui — shared component primitives (Button, Card, Badge, etc.).
- lucide-react — icon library.
- gsap + @gsap/react — complex animations (Hero reveal, AI Workspace visualisation, Intelligence Layer pipeline).
- lenis — smooth scroll (if configured at app root).
- TanStack Router — routing and anchor navigation.
- Tailwind CSS v4 — all styling.
- Enterprise partner/customer logos from content team (Trusted By section).
- Sales contact or demo-request route from routing team (Final CTA secondary CTA destination).

## Risks

| Risk | Likelihood | Impact | Mitigation |
| --- | --- | --- | --- |
| Workspace and Intelligence Layer visualisations require bespoke UI components not in @shopwise/ui | High | Medium | Build visualisation components locally in apps/web/src/features/landing/components/; migrate to packages/ui only if reuse is confirmed. |
| Lenis not configured at app root | Medium | Low | Fall back to CSS scroll-behavior: smooth; verify Lenis setup in implementation tasks. |
| Decision Confidence panel complexity increases implementation time | Medium | Medium | Start with a static card layout; add animation in a second pass. |
| Enterprise logos not available at implementation time | High | Low | Use clearly labelled placeholder assets; design components to accept content via props so swapping is a content change only. |
| GSAP animations cause layout shift or CLS regression | Medium | High | Measure CLS after each animation phase; ensure GSAP does not modify layout-affecting properties (width, height, margin). |
| Glass/blur effects cause performance regression on low-end devices | Low | Medium | Conditionally remove backdrop-filter on devices that report prefers-reduced-transparency or detected low-end GPU. |

---

## Open Questions

- What are the exact TanStack Router route paths for: Sign In, AI Workspace, and Sales Contact / Demo Request? These determine all CTA href values.
- Is Lenis currently configured at the apps/web app root, and if so, at which route level is ReactLenis mounted?
- Will a dedicated sales contact or demo-request route be available at the time of implementation, or should the secondary CTA initially link to an email mailto: address as a placeholder?

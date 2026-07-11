# Data Model: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page **Created**: 2026-07-10 **Last Updated**: 2026-07-11

> This page is entirely static/content-driven. There are no server-fetched entities or user state mutations. All entities below describe the TypeScript data structures that hold static content for each landing section.

---

## Entity: NavLink

**File**: `apps/web/src/features/landing/components/navbar/LandingNavbar.tsx` (inline in component)

```ts
type NavLink = {
  label: string; // e.g. "Problem", "Workspace"
  id: string; // e.g. "problem", "workspace"
  href: string; // e.g. "#problem", "#workspace"
};
```

**Valid values**:

- Target list: Problem (`#problem`), Workspace (`#workspace`), Intelligence (`#intelligence`), Confidence (`#confidence`), Contact (`#contact`).

---

## Entity: TrustedByLogo

**File**: `apps/web/src/features/landing/data/trusted-by.ts`

```ts
type TrustedByLogo = {
  id: string; // Unique slug, e.g. "acme-corp"
  name: string; // Company name for alt text
  src: string; // Path to logo SVG/PNG asset
  width: number; // Intrinsic width (px) for <img> sizing
  height: number; // Intrinsic height (px)
};
```

**Validation rules**:

- `name` must be non-empty (used as `alt` text — accessibility requirement).
- `src` must resolve to a valid asset path.
- In production, this array defaults to empty (`[]`), causing the section to return `null`. It is populated only when verified, brand-approved logos are supplied.

---

## Entity: FeatureItem (Workspace Capabilities)

**File**: `apps/web/src/features/landing/data/features.ts`

```ts
type FeatureItem = {
  id: string; // Unique slug, e.g. "canvas"
  icon: LucideIcon; // Lucide icon component reference
  title: string; // ≤ 5 words
  description: string; // ≤ 25 words
};
```

**Validation rules**:

- `title.split(' ').length <= 5`
- `description` word count ≤ 25
- Exactly 6 items representing the workspace capabilities defined in spec section 6:
  1. Product and Supplier Canvas
  2. Requirement Matching
  3. Trust and Evidence Centre
  4. Audit Trail
  5. Team Collaboration
  6. Enterprise Integrations

---

## Entity: WorkflowStep (Intelligence Layer Stages)

**File**: `apps/web/src/features/landing/data/workflow-steps.ts`

```ts
type WorkflowStep = {
  step: number; // 1-based step number (1 to 7)
  title: string; // ≤ 4 words
  description: string; // ≤ 15 words
};
```

**Validation rules**:

- `step` values are consecutive integers starting at 1.
- Exactly 7 items representing the intelligence pipeline stages defined in spec section 8:
  1. Business Goal
  2. Planning
  3. Data and Evidence
  4. Specialist Analysis
  5. Comparison Logic
  6. Recommendation
  7. Human Approval

---

## Entity: DecisionConfidenceSignal (Decision Confidence Panel)

**File**: `apps/web/src/features/landing/components/confidence/DecisionConfidenceSection.tsx` (or inline in visual component)

```ts
type DecisionConfidenceSignal = {
  label: string; // e.g. "Evidence status", "Confidence indicator"
  value: string; // e.g. "12 sources verified — 3 flagged for review"
  icon: LucideIcon; // Visual decorator icon
  status?: "success" | "warning" | "info" | "error"; // Signal severity/color styling
};
```

---

## Local UI State (not persisted)

| Component | State | Type | Description |
| --- | --- | --- | --- |
| `LandingNavbar` | `isScrolled` | `boolean` | Triggers sticky blur effect after 100 px |
| `LandingNavbar` | `isMobileMenuOpen` | `boolean` | Controls Sheet open state |

---

## Removed Entities

The following entities from v1.0.0 are **removed** and no longer exist in the codebase:

- `Testimonial` (formerly in `testimonials.ts`)
- `PricingTier` (formerly in `pricing.ts`)
- `FaqItem` (formerly in `faq.ts`)

---

## No Server Entities

This feature has no API calls, no database models, and no authentication-gated dynamic content. All data is statically imported at build time.

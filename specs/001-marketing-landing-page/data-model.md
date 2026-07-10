# Data Model: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page
**Created**: 2026-07-10

> This page is entirely static/content-driven. There are no server-fetched entities
> or user state mutations. All entities below describe the TypeScript data structures
> that hold static content for each landing section.

---

## Entity: NavLink

**File**: `apps/web/src/features/landing/data/` (inline in `LandingNavbar`)

```ts
type NavLink = {
  label: string;       // Display text, e.g. "Features"
  href: string;        // Anchor href, e.g. "#features"
};
```

---

## Entity: TrustedByLogo

**File**: `apps/web/src/features/landing/data/trusted-by.ts`

```ts
type TrustedByLogo = {
  id: string;           // Unique slug, e.g. "acme-corp"
  name: string;         // Company name for alt text
  src: string;          // Path to logo SVG/PNG asset
  width: number;        // Intrinsic width (px) for <img> sizing
  height: number;       // Intrinsic height (px)
};
```

**Validation rules**:
- `name` must be non-empty (used as `alt` text — accessibility requirement).
- `src` must resolve to a valid asset path.
- 6–8 items for the initial implementation.

---

## Entity: FeatureItem

**File**: `apps/web/src/features/landing/data/features.ts`

```ts
type FeatureItem = {
  id: string;           // Unique slug, e.g. "ai-vendor-scoring"
  icon: LucideIcon;     // Lucide icon component reference
  title: string;        // ≤ 5 words
  description: string;  // ≤ 25 words
};
```

**Validation rules**:
- `title.split(' ').length <= 5`
- `description` word count ≤ 25
- 4–8 items

---

## Entity: WorkflowStep

**File**: `apps/web/src/features/landing/data/workflow-steps.ts`

```ts
type WorkflowStep = {
  step: number;         // 1-based step number (rendered in badge)
  title: string;        // ≤ 5 words
  description: string;  // ≤ 20 words
};
```

**Validation rules**:
- `step` values are consecutive integers starting at 1.
- 3–5 items.

---

## Entity: Testimonial

**File**: `apps/web/src/features/landing/data/testimonials.ts`

```ts
type Testimonial = {
  id: string;           // Unique slug
  quote: string;        // ≤ 40 words
  author: string;       // Full name
  role: string;         // Job title
  company: string;      // Company name
  avatarSrc?: string;   // Optional avatar image path
  avatarAlt?: string;   // Required if avatarSrc is present
};
```

**Validation rules**:
- `quote` word count ≤ 40.
- If `avatarSrc` is present, `avatarAlt` must also be present.
- 3–6 items.

---

## Entity: PricingTier

**File**: `apps/web/src/features/landing/data/pricing.ts`

```ts
type PricingTier = {
  id: string;                    // Unique slug, e.g. "pro"
  name: string;                  // Plan name
  subtitle?: string;             // Optional tagline
  monthlyPrice: number | null;   // null = contact sales
  annualPrice: number | null;    // null = contact sales
  features: string[];            // 4–6 feature bullet strings
  cta: string;                   // CTA button label
  ctaHref: string;               // Navigation target
  highlighted: boolean;          // true = recommended tier
  badge?: string;                // e.g. "Most Popular"
};
```

**Validation rules**:
- Exactly one tier has `highlighted: true`.
- `features.length` is between 4 and 6.
- If `monthlyPrice` is null, `annualPrice` must also be null.
- 2–4 tiers total.

**State transitions**:
The `BillingToggle` component maintains `billingPeriod: 'monthly' | 'annual'` in local React state. Components display `monthlyPrice` or `annualPrice` accordingly. No server interaction.

---

## Entity: FaqItem

**File**: `apps/web/src/features/landing/data/faq.ts`

```ts
type FaqItem = {
  id: string;           // Unique slug (used as DOM id for aria-controls)
  question: string;
  answer: string;       // Plain text; may contain simple markdown-style formatting
};
```

**Validation rules**:
- `id` must be URL-safe (no spaces, lowercase).
- 5–10 items.

---

## Entity: FooterLinkGroup

**Inline in `LandingFooter`** (not a data file — small enough to be inline):

```ts
type FooterLink = {
  label: string;
  href: string;
  external?: boolean;  // true → target="_blank" + rel="noopener noreferrer"
};

type FooterLinkGroup = {
  heading: string;
  links: FooterLink[];
};
```

---

## Local UI State (not persisted)

| Component | State | Type | Description |
|---|---|---|---|
| `BillingToggle` | `billingPeriod` | `'monthly' \| 'annual'` | Controls which price column is shown |
| `FaqItem` | `isOpen` | `boolean` | Accordion open/close (via Collapsible) |
| `LandingNavbar` | `isScrolled` | `boolean` | Triggers sticky blur effect after 100 px |
| `LandingNavbar` | `isMobileMenuOpen` | `boolean` | Controls Sheet open state |
| `TestimonialsSection` | `activeIndex` | `number` | Carousel current slide (managed by @shopwise/ui Carousel) |

---

## No Server Entities

This feature has no API calls, no database models, and no authentication-gated content.
All data is statically imported at build time from the `data/` files listed above.

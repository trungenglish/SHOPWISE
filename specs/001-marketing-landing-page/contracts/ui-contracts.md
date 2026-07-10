# UI Contracts: Marketing Landing Page

**Feature**: specs/001-marketing-landing-page
**Created**: 2026-07-10

> These contracts define the public interface (props) of each custom component built
> for this feature. They are the implementation contract between the page assembly
> (`LandingPage.tsx`) and the individual section components. All components live in
> `apps/web/src/features/landing/components/`.

---

## `LandingPage`

**File**: `apps/web/src/features/landing/LandingPage.tsx`
**Role**: Page root — composes all sections in order. Mounted by the index route.

```ts
// No props — self-contained page root
type LandingPageProps = Record<string, never>;
```

**Composition order**:
1. `<LandingNavbar />`
2. `<HeroSection />`
3. `<TrustedBySection />`
4. `<AiDecisionSection />`
5. `<FeaturesSection />`
6. `<WorkflowSection />`
7. `<TestimonialsSection />`
8. `<PricingSection />`
9. `<FaqSection />`
10. `<FinalCtaSection />`
11. `<LandingFooter />`

**Lenis**: `<ReactLenis root>` wraps the entire `LandingPage` output.

---

## `LandingNavbar`

**File**: `components/navbar/LandingNavbar.tsx`

```ts
// No external props — reads scroll position internally
type LandingNavbarProps = Record<string, never>;
```

**Internal state**: `isScrolled: boolean`, `isMobileMenuOpen: boolean`
**Children rendered**: `MobileNavDrawer` (when mobile menu is open)
**@shopwise/ui used**: `Button`, `Sheet` (for mobile drawer)

---

## `MobileNavDrawer`

**File**: `components/navbar/MobileNavDrawer.tsx`

```ts
type MobileNavDrawerProps = {
  isOpen: boolean;
  onClose: () => void;
};
```

**@shopwise/ui used**: `Sheet`, `SheetContent`, `Button`

---

## `HeroSection`

**File**: `components/hero/HeroSection.tsx`

```ts
// No external props — all content is internal/static
type HeroSectionProps = Record<string, never>;
```

**Children rendered**: `HeroVisual`
**@shopwise/ui used**: `Button`, `Badge`
**Animation**: GSAP timeline via `useGSAP` (guarded by `useReducedMotion`)

---

## `HeroVisual`

**File**: `components/hero/HeroVisual.tsx`

```ts
// No props — purely presentational animated graphic
type HeroVisualProps = Record<string, never>;
```

**Animation**: GSAP or CSS keyframe animation (floating/pulsing abstract graphic)
**Accessibility**: `aria-hidden="true"` — purely decorative

---

## `TrustedBySection`

**File**: `components/trusted-by/TrustedBySection.tsx`

```ts
// No external props — consumes data from trusted-by.ts
type TrustedBySectionProps = Record<string, never>;
```

**Children rendered**: `LogoMarquee`
**Animation**: CSS marquee (desktop), static grid (mobile / reduced-motion)

---

## `LogoMarquee`

**File**: `components/trusted-by/LogoMarquee.tsx`

```ts
type LogoMarqueeProps = {
  logos: TrustedByLogo[];
  pauseOnHover?: boolean;   // default: true
};
```

---

## `AiDecisionSection`

**File**: `components/ai-viz/AiDecisionSection.tsx`

```ts
// No external props
type AiDecisionSectionProps = Record<string, never>;
```

**Children rendered**: `AiDecisionPanel`
**@shopwise/ui used**: `Button`
**Animation**: GSAP ScrollTrigger on panel reveal

---

## `AiDecisionPanel`

**File**: `components/ai-viz/AiDecisionPanel.tsx`

```ts
// No external props — self-animating demo panel
type AiDecisionPanelProps = Record<string, never>;
```

**Accessibility**: `aria-hidden="true"` + companion `<p class="sr-only">` in parent section
**Animation**: sequential token reveal / score animation via `useGSAP`

---

## `FeaturesSection`

**File**: `components/features/FeaturesSection.tsx`

```ts
// No external props — consumes data from features.ts
type FeaturesSectionProps = Record<string, never>;
```

**Children rendered**: `FeatureCard[]` (mapped from data)

---

## `FeatureCard`

**File**: `components/features/FeatureCard.tsx`

```ts
type FeatureCardProps = {
  item: FeatureItem;
};
```

**@shopwise/ui used**: `Card`, `CardContent`
**Animation**: CSS transition on hover (border glow / icon color shift)
**Scroll reveal**: GSAP stagger or CSS `animation-delay` on `FeaturesSection`

---

## `WorkflowSection`

**File**: `components/workflow/WorkflowSection.tsx`

```ts
// No external props — consumes data from workflow-steps.ts
type WorkflowSectionProps = Record<string, never>;
```

**Children rendered**: `WorkflowTimeline`

---

## `WorkflowTimeline`

**File**: `components/workflow/WorkflowTimeline.tsx`

```ts
type WorkflowTimelineProps = {
  steps: WorkflowStep[];
};
```

**Animation**: GSAP ScrollTrigger draws connector line; steps stagger in
**Layout**: horizontal on ≥ 1024 px, vertical on < 1024 px

---

## `TestimonialsSection`

**File**: `components/testimonials/TestimonialsSection.tsx`

```ts
// No external props — consumes data from testimonials.ts
type TestimonialsSectionProps = Record<string, never>;
```

**@shopwise/ui used**: `Carousel`, `CarouselContent`, `CarouselItem`, `CarouselPrevious`, `CarouselNext`, `Avatar`, `AvatarImage`, `AvatarFallback`
**Accessibility**: `aria-live="polite"` on carousel viewport; prev/next have `aria-label`

---

## `PricingSection`

**File**: `components/pricing/PricingSection.tsx`

```ts
// No external props — manages billingPeriod state internally
type PricingSectionProps = Record<string, never>;
```

**Children rendered**: `BillingToggle`, `PricingCard[]`

---

## `BillingToggle`

**File**: `components/pricing/BillingToggle.tsx`

```ts
type BillingToggleProps = {
  value: 'monthly' | 'annual';
  onChange: (value: 'monthly' | 'annual') => void;
};
```

**@shopwise/ui used**: `Toggle` or `Switch` for the toggle affordance
**Accessibility**: `role="switch"`, `aria-checked`

---

## `PricingCard`

**File**: `components/pricing/PricingCard.tsx`

```ts
type PricingCardProps = {
  tier: PricingTier;
  billingPeriod: 'monthly' | 'annual';
};
```

**@shopwise/ui used**: `Card`, `CardHeader`, `CardContent`, `CardFooter`, `Badge`, `Button`, `Separator`

---

## `FaqSection`

**File**: `components/faq/FaqSection.tsx`

```ts
// No external props — consumes data from faq.ts
type FaqSectionProps = Record<string, never>;
```

**Children rendered**: `FaqItem[]`

---

## `FaqItem`

**File**: `components/faq/FaqItem.tsx`

```ts
type FaqItemProps = {
  item: FaqItem;
};
```

**@shopwise/ui used**: `Collapsible`, `CollapsibleTrigger`, `CollapsibleContent`
**Animation**: CSS `max-height` transition (removed under `prefers-reduced-motion`)
**Accessibility**: `aria-expanded`, `aria-controls`, `role="region"`

---

## `FinalCtaSection`

**File**: `components/final-cta/FinalCtaSection.tsx`

```ts
// No external props
type FinalCtaSectionProps = Record<string, never>;
```

**@shopwise/ui used**: `Button`
**Animation**: GSAP scroll-entry fade + CTA pulse (CSS, paused under reduced-motion)

---

## `LandingFooter`

**File**: `components/footer/LandingFooter.tsx`

```ts
// No external props
type LandingFooterProps = Record<string, never>;
```

**@shopwise/ui used**: `Separator`

---

## Custom Hooks

### `useReducedMotion`

**File**: `apps/web/src/features/landing/hooks/useReducedMotion.ts`

```ts
// Returns true if prefers-reduced-motion: reduce is active
function useReducedMotion(): boolean;
```

**Implementation**: Wraps `window.matchMedia('(prefers-reduced-motion: reduce)')` with
a `useEffect` listener for runtime changes.

---

### `useScrollSpy`

**File**: `apps/web/src/features/landing/hooks/useScrollSpy.ts`

```ts
type UseScrollSpyOptions = {
  sectionIds: string[];
  offset?: number;  // px from top of viewport to trigger active state (default: 100)
};

function useScrollSpy(options: UseScrollSpyOptions): string | null;
// Returns the id of the currently active section, or null if none
```

**Used by**: `LandingNavbar` to highlight the active nav link.

---

## GSAP Configuration

**File**: `apps/web/src/lib/gsap-config.ts`

```ts
// Registers ScrollTrigger plugin once.
// All GSAP-using components import from this file instead of gsap directly.
import { gsap } from 'gsap';
import { ScrollTrigger } from 'gsap/ScrollTrigger';

gsap.registerPlugin(ScrollTrigger);

export { gsap, ScrollTrigger };
```

export type PricingTier = {
  id: string;
  name: string;
  /** Optional one-line tagline below the plan name. */
  subtitle?: string;
  /** Monthly price in USD. null = "Contact Sales". */
  monthlyPrice: number | null;
  /** Annual price in USD (per month, billed annually). null = "Contact Sales". */
  annualPrice: number | null;
  /** 4–6 feature bullets displayed in the card. */
  features: string[];
  /** CTA button label. */
  cta: string;
  /** Navigation target — internal route or external href. */
  ctaHref: string;
  /** When true, this card is visually elevated as the recommended plan. */
  highlighted: boolean;
  /** Optional badge text, e.g. "Most Popular". */
  badge?: string;
};

export const PRICING_TIERS: PricingTier[] = [
  {
    id: "starter",
    name: "Starter",
    subtitle: "For small teams getting started",
    monthlyPrice: 49,
    annualPrice: 39,
    features: [
      "Up to 5 users",
      "AI vendor scoring (50 queries/month)",
      "Basic spend analytics dashboard",
      "2 ERP integrations",
      "Email support",
    ],
    cta: "Start Free Trial",
    ctaHref: "/auth",
    highlighted: false,
  },
  {
    id: "pro",
    name: "Pro",
    subtitle: "For growing procurement teams",
    monthlyPrice: 149,
    annualPrice: 119,
    features: [
      "Up to 25 users",
      "Unlimited AI vendor scoring",
      "Advanced spend analytics & forecasting",
      "Unlimited ERP integrations",
      "Collaborative decision workspace",
      "Priority email & chat support",
    ],
    cta: "Start Free Trial",
    ctaHref: "/auth",
    highlighted: true,
    badge: "Most Popular",
  },
  {
    id: "enterprise",
    name: "Enterprise",
    subtitle: "For large organisations with custom needs",
    monthlyPrice: null,
    annualPrice: null,
    features: [
      "Unlimited users",
      "Custom AI scoring models",
      "Full audit trail & SOC 2 compliance",
      "Dedicated customer success manager",
      "SLA-backed 99.9% uptime",
      "Custom integrations & SSO",
    ],
    cta: "Talk to Sales",
    ctaHref: "mailto:sales@shopwise.io",
    highlighted: false,
  },
];

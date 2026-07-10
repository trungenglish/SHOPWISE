export type Testimonial = {
  id: string;
  /** ≤ 40 words. */
  quote: string;
  author: string;
  role: string;
  company: string;
  /** Optional path to avatar image. If present, avatarAlt is required. */
  avatarSrc?: string;
  /** Alt text for the avatar image — required when avatarSrc is set. */
  avatarAlt?: string;
};

export const TESTIMONIALS: Testimonial[] = [
  {
    id: "sarah-chen",
    quote:
      "ShopWise cut our vendor evaluation cycle from three weeks to two days. The AI scoring is eerily accurate — it consistently surfaces the same vendors our experts would have chosen.",
    author: "Sarah Chen",
    role: "VP Procurement",
    company: "Meridian Retail Group",
  },
  {
    id: "james-okafor",
    quote:
      "We reduced time-to-decision by 60% in the first quarter. The collaborative workspace means our category managers, finance, and legal teams all work from the same source of truth.",
    author: "James Okafor",
    role: "Head of Category Management",
    company: "Atlas Global Logistics",
  },
  {
    id: "priya-nair",
    quote:
      "Integration with our SAP instance took less than a day. The ShopWise team actually cares about enterprise readiness — it's not just a demo product.",
    author: "Priya Nair",
    role: "CTO",
    company: "Horizon Commerce",
  },
  {
    id: "michael-hart",
    quote:
      "Our audit pass rate went from 74% to 98% after we moved procurement approvals into ShopWise. The audit trail is immaculate and the compliance team finally trusts the data.",
    author: "Michael Hart",
    role: "Director of Finance",
    company: "Stratos Manufacturing",
  },
];

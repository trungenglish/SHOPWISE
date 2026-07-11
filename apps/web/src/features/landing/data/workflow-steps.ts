export type WorkflowStep = {
  /** 1-based step number rendered in the timeline badge. */
  step: number;
  /** ≤ 5 words. */
  title: string;
  /** ≤ 20 words. */
  description: string;
};

export const WORKFLOW_STEPS: WorkflowStep[] = [
  {
    step: 1,
    title: "Customer Query",
    description: "Customer asks a question about a product type or model.",
  },
  {
    step: 2,
    title: "Intent Analysis",
    description: "AI understands intent and extracts search criteria.",
  },
  {
    step: 3,
    title: "Product Search",
    description: "Searches the Phong Vu catalogue for matching options.",
  },
  {
    step: 4,
    title: "Inventory Check",
    description: "Retrieves live stock status, pricing, and active promos.",
  },
  {
    step: 5,
    title: "Comparison Logic",
    description: "Compares product specs and matches user requirements.",
  },
  {
    step: 6,
    title: "Recommendation",
    description: "Recommends the best product fit with explainable details.",
  },
  {
    step: 7,
    title: "Guided Checkout",
    description: "Guides the customer directly to purchase completion.",
  },
];

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
    title: "Business Goal",
    description: "The team defines what they need to buy and why.",
  },
  {
    step: 2,
    title: "Planning",
    description:
      "ShopWise structures the decision: criteria, constraints, and stakeholders.",
  },
  {
    step: 3,
    title: "Data and Evidence",
    description:
      "The platform gathers and verifies supplier data, documents, and market signals.",
  },
  {
    step: 4,
    title: "Specialist Analysis",
    description: "AI agents evaluate each option against the defined criteria.",
  },
  {
    step: 5,
    title: "Comparison Logic",
    description: "Options are ranked with full reasoning and evidence links.",
  },
  {
    step: 6,
    title: "Recommendation",
    description:
      "A confident, explainable recommendation is surfaced to the team.",
  },
  {
    step: 7,
    title: "Human Approval",
    description:
      "The team reviews, adjusts, and approves — creating the audit record.",
  },
];

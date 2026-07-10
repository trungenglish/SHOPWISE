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
    title: "Connect Your Data",
    description:
      "Link your ERP, procurement system, or supplier database in minutes with no-code connectors.",
  },
  {
    step: 2,
    title: "Define Your Goals",
    description:
      "Tell the AI your priorities — cost, quality, speed, or sustainability — and set scoring weights.",
  },
  {
    step: 3,
    title: "Run AI Analysis",
    description:
      "The AI surfaces ranked recommendations with full reasoning, sourced from your live data.",
  },
  {
    step: 4,
    title: "Decide with Confidence",
    description:
      "Approve, adjust, or escalate decisions with a complete audit trail for compliance.",
  },
];

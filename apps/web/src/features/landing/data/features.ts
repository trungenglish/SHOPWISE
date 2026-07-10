import type { LucideIcon } from "lucide-react";
import { Bot, TrendingUp, Zap, BarChart3, Users, Plug } from "lucide-react";

export type FeatureItem = {
  id: string;
  icon: LucideIcon;
  /** ≤ 5 words. */
  title: string;
  /** ≤ 25 words. */
  description: string;
};

export const FEATURES: FeatureItem[] = [
  {
    id: "ai-vendor-scoring",
    icon: Bot,
    title: "AI Vendor Scoring",
    description:
      "Automatically rank vendors by quality, price, and reliability signals using real-time AI analysis.",
  },
  {
    id: "market-intelligence",
    icon: TrendingUp,
    title: "Real-Time Market Intelligence",
    description:
      "Access live pricing benchmarks and supply-chain alerts to stay ahead of market shifts.",
  },
  {
    id: "workflow-automation",
    icon: Zap,
    title: "Procurement Workflow Automation",
    description:
      "Automate approvals, purchase orders, and audit trails — removing manual bottlenecks.",
  },
  {
    id: "spend-analytics",
    icon: BarChart3,
    title: "Spend Analytics",
    description:
      "Visual dashboards that surface category spend, savings opportunities, and budget tracking.",
  },
  {
    id: "decision-workspace",
    icon: Users,
    title: "Collaborative Decision Workspace",
    description:
      "Team-based review, threaded comments, and structured final-approval flows in one place.",
  },
  {
    id: "enterprise-integrations",
    icon: Plug,
    title: "Enterprise Integrations",
    description:
      "Native connectors for SAP, Oracle, and Salesforce keep ShopWise in sync with your existing stack.",
  },
];

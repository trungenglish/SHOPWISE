import type { LucideIcon } from "lucide-react";
import {
  Layout,
  ListTodo,
  ShieldCheck,
  History,
  Users,
  Plug,
} from "lucide-react";

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
    id: "canvas",
    icon: Layout,
    title: "Product & Supplier Canvas",
    description:
      "A unified view of all options, structured for direct comparison.",
  },
  {
    id: "requirement-matching",
    icon: ListTodo,
    title: "Requirement Matching",
    description:
      "Automatically map business requirements to each supplier or product.",
  },
  {
    id: "trust-evidence",
    icon: ShieldCheck,
    title: "Trust & Evidence Centre",
    description:
      "Verify claims, documents, and supplier credentials in one place.",
  },
  {
    id: "audit-trail",
    icon: History,
    title: "Audit Trail",
    description:
      "A full, traceable record of every decision step and approval action.",
  },
  {
    id: "collaboration",
    icon: Users,
    title: "Team Collaboration",
    description:
      "Assign reviewers, gather comments, and reach consensus in the workspace.",
  },
  {
    id: "integrations",
    icon: Plug,
    title: "Enterprise Integrations",
    description:
      "Connect to ERP, procurement, and supplier management systems.",
  },
];

export type FaqItem = {
  /** URL-safe slug used as the DOM `id` for aria-controls. No spaces, lowercase. */
  id: string;
  question: string;
  answer: string;
};

export const FAQ_ITEMS: FaqItem[] = [
  {
    id: "faq-ai-accuracy",
    question: "How accurate is the AI vendor scoring?",
    answer:
      "Our AI models are trained on billions of procurement transactions and live supplier data. In independent benchmarks, ShopWise recommendations match expert human decisions 94% of the time. Accuracy improves further as the model learns from your team's historical choices and approval patterns.",
  },
  {
    id: "faq-erp-integration",
    question: "Which ERP and procurement systems does ShopWise integrate with?",
    answer:
      "ShopWise ships with native connectors for SAP S/4HANA, Oracle Fusion, Microsoft Dynamics 365, Salesforce, Coupa, and Ariba. We also provide a REST API and webhook system for custom integrations. Enterprise customers receive dedicated integration engineering support.",
  },
  {
    id: "faq-security",
    question: "Is ShopWise GDPR and SOC 2 compliant?",
    answer:
      "Yes. ShopWise is SOC 2 Type II certified and fully GDPR compliant. All data is encrypted in transit (TLS 1.3) and at rest (AES-256). We do not use your procurement data to train shared AI models. Data residency options are available for EU customers. A full security whitepaper is available on request.",
  },
  {
    id: "faq-onboarding",
    question: "How long does onboarding take?",
    answer:
      "Most Starter and Pro teams are live within one business day — connecting your first data source takes under an hour with our guided setup wizard. Enterprise deployments with complex ERP integrations typically take 5–10 business days, supported by your dedicated customer success manager.",
  },
  {
    id: "faq-support",
    question: "What does your support SLA look like?",
    answer:
      "Starter plans receive email support with a 24-hour response time. Pro plans add live chat with a 4-hour SLA. Enterprise customers receive a dedicated success manager, 1-hour critical-incident response, and access to our 24/7 engineering escalation line.",
  },
  {
    id: "faq-customisation",
    question: "Can we customise the AI scoring criteria for our industry?",
    answer:
      "Absolutely. All plans allow you to adjust scoring weights for cost, quality, delivery time, sustainability, and financial stability. Enterprise plans additionally support fully custom scoring dimensions trained on your proprietary supplier data and internal evaluation rubrics.",
  },
  {
    id: "faq-free-trial",
    question: "Is there a free trial?",
    answer:
      "Yes — both the Starter and Pro plans include a 14-day free trial with no credit card required. You get full access to all features during the trial so you can evaluate ShopWise against your real procurement data. Enterprise evaluations are handled through a guided proof-of-concept programme.",
  },
];

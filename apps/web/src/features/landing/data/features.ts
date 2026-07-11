import { Search, Sparkles, Layout, Database, Tag, ShoppingBag } from "lucide-react";
import type { LucideIcon } from "lucide-react";

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
    id: "product-search",
    icon: Search,
    title: "AI Product Search",
    description: "Find products on Phong Vu using natural language and semantic queries.",
  },
  {
    id: "product-recommendation",
    icon: Sparkles,
    title: "Product Recommendation",
    description: "Recommend products customized to buyer preferences and budget limits.",
  },
  {
    id: "product-comparison",
    icon: Layout,
    title: "Product Comparison",
    description: "Compare technical specs, key features, and price tags side-by-side.",
  },
  {
    id: "inventory-lookup",
    icon: Database,
    title: "Live Inventory Lookup",
    description: "Retrieve real-time pricing and stock status directly from warehouses.",
  },
  {
    id: "promotion-finder",
    icon: Tag,
    title: "Promotion Finder",
    description: "Locate active discount vouchers, bank coupons, and bundle deals.",
  },
  {
    id: "checkout-guidance",
    icon: ShoppingBag,
    title: "Checkout Guidance",
    description: "Guide customers directly to the shopping cart for simplified checkout.",
  },
];

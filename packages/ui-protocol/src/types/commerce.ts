// AI Components

export interface ChatMessageProps {
  text: string;
  sender: "user" | "ai" | "system";
}

export interface ThinkingIndicatorProps {
  statusText?: string;
}

export interface RecommendationExplanationProps {
  text?: string;
  highlights?: string[];
}

export interface DecisionReasoningProps {
  reasoning?: string;
}

export interface ConfidenceIndicatorProps {
  score?: number; // 0-100
  label?: string;
}

// Commerce Components

export interface ProductCardProps {
  productId: string;
  name: string;
  priceVND: number;
  imageUrl?: string;
  rating?: number;
}

export interface ProductCarouselProps {
  title?: string;
}

export interface ProductComparisonProps {
  productIds?: string[];
  features?: string[];
}

export interface SpecificationTableProps {
  specs?: Record<string, string>;
}

export interface PromotionBannerProps {
  text?: string;
  discountCode?: string;
}

export interface WarrantyInformationProps {
  months?: number;
  details?: string;
}

export interface InventoryStatusProps {
  inStock?: boolean;
  quantity?: number;
  storeId?: string;
}

export interface CheckoutSummaryItem {
  productId?: string;
  quantity?: number;
  priceVND?: number;
}

export interface CheckoutSummaryProps {
  items?: CheckoutSummaryItem[];
  totalPriceVND?: number;
}

// Commerce Actions

export interface AddToComparisonPayload {
  productId: string;
}

export interface RemoveProductPayload {
  productId: string;
}

export interface SaveDecisionPayload {
  decisionId: string;
  productId: string;
}

export interface ResumeConversationPayload {
  conversationId: string;
}

export interface CheckoutReadinessPayload {
  cartId?: string;
  ready: boolean;
}

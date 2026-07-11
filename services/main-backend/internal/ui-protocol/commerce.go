package uiprotocol

// AI Components

type ChatMessageProps struct {
	Text   string `json:"text"`
	Sender string `json:"sender"` // "user", "ai", "system"
}

type ThinkingIndicatorProps struct {
	StatusText *string `json:"statusText,omitempty"`
}

type RecommendationExplanationProps struct {
	Text       *string  `json:"text,omitempty"`
	Highlights []string `json:"highlights,omitempty"`
}

type DecisionReasoningProps struct {
	Reasoning *string `json:"reasoning,omitempty"`
}

type ConfidenceIndicatorProps struct {
	Score *float64 `json:"score,omitempty"` // 0-100
	Label *string  `json:"label,omitempty"`
}

// Commerce Components

type ProductCardProps struct {
	ProductID string  `json:"productId"`
	Name      string  `json:"name"`
	PriceVND  int     `json:"priceVND"`
	ImageURL  *string `json:"imageUrl,omitempty"`
	Rating    *float64 `json:"rating,omitempty"`
}

type ProductCarouselProps struct {
	Title *string `json:"title,omitempty"`
}

type ProductComparisonProps struct {
	ProductIDs []string `json:"productIds,omitempty"`
	Features   []string `json:"features,omitempty"`
}

type SpecificationTableProps struct {
	Specs map[string]string `json:"specs,omitempty"`
}

type PromotionBannerProps struct {
	Text         *string `json:"text,omitempty"`
	DiscountCode *string `json:"discountCode,omitempty"`
}

type WarrantyInformationProps struct {
	Months  *int    `json:"months,omitempty"`
	Details *string `json:"details,omitempty"`
}

type InventoryStatusProps struct {
	InStock  *bool   `json:"inStock,omitempty"`
	Quantity *int    `json:"quantity,omitempty"`
	StoreID  *string `json:"storeId,omitempty"`
}

type CheckoutSummaryItem struct {
	ProductID *string `json:"productId,omitempty"`
	Quantity  *int    `json:"quantity,omitempty"`
	PriceVND  *int    `json:"priceVND,omitempty"`
}

type CheckoutSummaryProps struct {
	Items         []CheckoutSummaryItem `json:"items,omitempty"`
	TotalPriceVND *int                  `json:"totalPriceVND,omitempty"`
}

// Commerce Actions

type AddToComparisonPayload struct {
	ProductID string `json:"productId"`
}

type RemoveProductPayload struct {
	ProductID string `json:"productId"`
}

type SaveDecisionPayload struct {
	DecisionID string `json:"decisionId"`
	ProductID  string `json:"productId"`
}

type ResumeConversationPayload struct {
	ConversationID string `json:"conversationId"`
}

type CheckoutReadinessPayload struct {
	CartID *string `json:"cartId,omitempty"`
	Ready  bool    `json:"ready"`
}

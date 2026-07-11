package uiprotocol

type ErrorStateProps struct {
	Title    string  `json:"title"`
	Message  string  `json:"message"`
	Code     *string `json:"code,omitempty"`
	Severity *string `json:"severity,omitempty"` // "info", "warning", "error", "critical"
}

type RetryActionProps struct {
	Label    *string `json:"label,omitempty"`
	ActionID *string `json:"actionId,omitempty"`
}

type ContinueCachedDataProps struct {
	Label *string `json:"label,omitempty"`
}

type ModifySearchProps struct {
	Label         *string `json:"label,omitempty"`
	OriginalQuery *string `json:"originalQuery,omitempty"`
}

type AlternativeRecommendationProps struct {
	Message    *string  `json:"message,omitempty"`
	ProductIDs []string `json:"productIds,omitempty"`
}

package domain

import (
	"time"
)

// WorkspaceState represents the state of a ComparisonWorkspace
type WorkspaceState string

const (
	WorkspaceStateInit               WorkspaceState = "INIT"
	WorkspaceStateAddingProduct      WorkspaceState = "ADDING_PRODUCT"
	WorkspaceStateMaxCapacity        WorkspaceState = "MAX_CAPACITY"
	WorkspaceStatePendingReplacement WorkspaceState = "PENDING_REPLACEMENT"
	WorkspaceStateCompleted          WorkspaceState = "COMPLETED"
	WorkspaceStateActive             WorkspaceState = "ACTIVE"
)

// ComparisonWorkspace represents a user's active comparison session.
type ComparisonWorkspace struct {
	WorkspaceID string         `json:"workspace_id"`
	UserID      string         `json:"user_id"`
	CategoryID  string         `json:"category_id"`
	ProductIDs  []string       `json:"product_ids"`
	State       WorkspaceState `json:"state"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// Availability represents product stock status
type Availability string

const (
	AvailabilityInStock    Availability = "IN_STOCK"
	AvailabilityOutOfStock Availability = "OUT_OF_STOCK"
	AvailabilityPreorder   Availability = "PREORDER"
)

// HighlightType represents visual differences for a specification
type HighlightType string

const (
	HighlightBetter            HighlightType = "BETTER"
	HighlightWorse             HighlightType = "WORSE"
	HighlightEqual             HighlightType = "EQUAL"
	HighlightMissing           HighlightType = "MISSING"
	HighlightCategoryAdvantage HighlightType = "CATEGORY_ADVANTAGE"
)

// ComparisonDifferenceHighlight represents the visual differences for a specific specification
type ComparisonDifferenceHighlight struct {
	SpecificationKey string                   `json:"specification_key"`
	Highlights       map[string]HighlightType `json:"highlights"` // mapping product_id -> HighlightType
}

// ComparisonProductDetail represents the runtime data for a product in the comparison view
type ComparisonProductDetail struct {
	ProductID      string            `json:"product_id"`
	Name           string            `json:"name"`
	Brand          string            `json:"brand"`
	Category       string            `json:"category"`
	PriceVND       int64             `json:"price_vnd"`
	Promotion      any       `json:"promotion,omitempty"`
	Availability   Availability      `json:"availability"`
	Warranty       string            `json:"warranty"`
	Images         []string          `json:"images"`
	Specifications map[string]string `json:"specifications"`
}

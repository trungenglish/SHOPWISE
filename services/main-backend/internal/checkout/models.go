package checkout

import "time"

type CheckoutReadinessStatus string

const (
	StatusReady              CheckoutReadinessStatus = "READY"
	StatusMissingInformation CheckoutReadinessStatus = "MISSING_INFORMATION"
	StatusInventoryChanged   CheckoutReadinessStatus = "INVENTORY_CHANGED"
	StatusPromotionChanged   CheckoutReadinessStatus = "PROMOTION_CHANGED"
	StatusOutOfStock         CheckoutReadinessStatus = "OUT_OF_STOCK"
	StatusStoreUnavailable   CheckoutReadinessStatus = "STORE_UNAVAILABLE"
	StatusRequiresUserAction CheckoutReadinessStatus = "REQUIRES_USER_ACTION"
)

type ValidationType string

const (
	ValidationInventory         ValidationType = "INVENTORY"
	ValidationPrice             ValidationType = "PRICE"
	ValidationPromotion         ValidationType = "PROMOTION"
	ValidationCustomerInfo      ValidationType = "CUSTOMER_INFO"
	ValidationStoreAvailability ValidationType = "STORE_AVAILABILITY"
)

type DeliveryPreference string

const (
	DeliveryHome  DeliveryPreference = "HOME_DELIVERY"
	DeliveryStore DeliveryPreference = "STORE_PICKUP"
	DeliveryUnset DeliveryPreference = "UNSELECTED"
)

type ValidationItem struct {
	Type    ValidationType `json:"type"`
	Passed  bool           `json:"passed"`
	Message string         `json:"message"`
}

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	District string `json:"district"`
	Ward     string `json:"ward"`
}

type Store struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CustomerProfile struct {
	Name               *string            `json:"name,omitempty"`
	PhoneNumber        *string            `json:"phoneNumber,omitempty"`
	DeliveryPreference DeliveryPreference `json:"deliveryPreference"`
	DeliveryAddress    *Address           `json:"deliveryAddress,omitempty"`
	SelectedStore      *Store             `json:"selectedStore,omitempty"`
}

type ProductSnapshot struct {
	ProductID string `json:"productId"`
	Name      string `json:"name"`
	PriceVND  int    `json:"priceVnd"`
}

type PromotionContext struct {
	PromotionID       string     `json:"promotionId"`
	Description       string     `json:"description"`
	DiscountAmountVND int        `json:"discountAmountVnd"`
	ExpiresAt         *time.Time `json:"expiresAt,omitempty"`
}

type CheckoutReadinessState struct {
	SessionID           string                  `json:"sessionId"`
	CartVersion         string                  `json:"cartVersion"`
	Status              CheckoutReadinessStatus `json:"status"`
	LastValidatedAt     time.Time               `json:"lastValidatedAt"`
	ValidationChecklist []ValidationItem        `json:"validationChecklist"`
	ProductSnapshot     *ProductSnapshot        `json:"productSnapshot,omitempty"`
	CustomerProfile     *CustomerProfile        `json:"customerProfile,omitempty"`
	PromotionContext    *PromotionContext       `json:"promotionContext,omitempty"`
}

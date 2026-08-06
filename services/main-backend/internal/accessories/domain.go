package accessories

import (
	"time"

	"github.com/google/uuid"
)

type Kind string
type VerificationStatus string
type OfferState string

const (
	KindExternal Kind = "external"
	KindUpgrade  Kind = "upgrade"

	VerificationCategory VerificationStatus = "category"
	VerificationModel    VerificationStatus = "verified_model"

	OfferFresh       OfferState = "fresh"
	OfferStale       OfferState = "stale"
	OfferUnavailable OfferState = "unavailable"
)

type Accessory struct {
	ID             uuid.UUID
	Kind           Kind
	Category       string
	Name           string
	Brand          string
	ImageURL       string
	Specifications map[string]string
}

type RetailerOffer struct {
	ID                uuid.UUID
	AccessoryID       uuid.UUID
	RetailerProductID string
	SourceURL         string
	Price             int64
	OriginalPrice     int64
	InStock           bool
	FetchedAt         time.Time
}

type AccessoryCompatibility struct {
	ID                 uuid.UUID
	AccessoryID        uuid.UUID
	LaptopCategory     string
	ProductID          *uuid.UUID
	Reason             string
	VerificationStatus VerificationStatus
}

type LaptopCompatibilityProfile struct {
	ID               uuid.UUID
	ProductID        uuid.UUID
	ModelCode        string
	RAMType          string
	RAMSlots         int
	MaximumRAMGB     int
	StorageInterface string
	FreeStorageSlots int
	Charger          string
	SourceURL        string
	Verified         bool
	VerifiedAt       time.Time
}

type RetailerSyncState struct {
	Retailer      string
	LastRunAt     *time.Time
	LastSuccessAt *time.Time
	LastError     string
}

type LaptopTarget struct {
	ID                     uuid.UUID
	Name                   string
	Category               string
	UpgradeProfileVerified bool
}

type Candidate struct {
	Accessory       Accessory
	Offer           RetailerOffer
	Compatibilities []AccessoryCompatibility
}

type RecommendationSnapshot struct {
	Products   []LaptopTarget
	Candidates []Candidate
}

type Recommendation struct {
	Accessory            Accessory
	Offer                RetailerOffer
	CompatibleProductIDs []uuid.UUID
	CompatibilityReason  string
	CompatibilityStatus  VerificationStatus
	OfferState           OfferState
	CheckoutAvailable    bool
}

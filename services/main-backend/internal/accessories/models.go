package accessories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AccessoryModel struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Kind           string         `gorm:"size:16;not null;index"`
	Category       string         `gorm:"size:64;not null;index"`
	Name           string         `gorm:"size:255;not null"`
	Brand          string         `gorm:"size:100"`
	ImageURL       string         `gorm:"type:text"`
	Specifications datatypes.JSON `gorm:"type:jsonb"`
}

func (AccessoryModel) TableName() string { return "accessories" }

type RetailerOfferModel struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	AccessoryID       uuid.UUID `gorm:"type:uuid;not null;index"`
	Retailer          string    `gorm:"size:32;not null;uniqueIndex:retailer_product"`
	RetailerProductID string    `gorm:"size:128;not null;uniqueIndex:retailer_product"`
	SourceURL         string    `gorm:"type:text;not null"`
	Price             int64     `gorm:"type:bigint;not null;check:price >= 0"`
	OriginalPrice     int64     `gorm:"type:bigint;not null;default:0;check:original_price >= 0"`
	InStock           bool      `gorm:"not null;default:false"`
	FetchedAt         time.Time `gorm:"type:timestamptz;not null;index"`
}

func (RetailerOfferModel) TableName() string { return "retailer_offers" }

type AccessoryCompatibilityModel struct {
	ID                 uuid.UUID  `gorm:"type:uuid;primaryKey"`
	AccessoryID        uuid.UUID  `gorm:"type:uuid;not null;index"`
	LaptopCategory     string     `gorm:"size:64;index"`
	ProductID          *uuid.UUID `gorm:"type:uuid;index"`
	Reason             string     `gorm:"type:text;not null"`
	VerificationStatus string     `gorm:"size:32;not null"`
}

func (AccessoryCompatibilityModel) TableName() string { return "accessory_compatibilities" }

type LaptopCompatibilityProfileModel struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID        uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	ModelCode        string    `gorm:"size:128;not null;uniqueIndex"`
	RAMType          string    `gorm:"size:64"`
	RAMSlots         int
	MaximumRAMGB     int
	StorageInterface string `gorm:"size:128"`
	FreeStorageSlots int
	Charger          string    `gorm:"size:255"`
	SourceURL        string    `gorm:"type:text;not null"`
	Verified         bool      `gorm:"not null;default:false"`
	VerifiedAt       time.Time `gorm:"type:timestamptz"`
}

func (LaptopCompatibilityProfileModel) TableName() string { return "laptop_compatibility_profiles" }

type RetailerSyncStateModel struct {
	Retailer      string     `gorm:"size:32;primaryKey"`
	LastRunAt     *time.Time `gorm:"type:timestamptz"`
	LastSuccessAt *time.Time `gorm:"type:timestamptz"`
	LastError     string     `gorm:"type:text;not null;default:''"`
}

func (RetailerSyncStateModel) TableName() string { return "retailer_sync_states" }

package promotions

import (
	"time"

	"github.com/google/uuid"
)

type PromotionStatus string

const (
	StatusActive                    PromotionStatus = "ACTIVE"
	StatusExpired                   PromotionStatus = "EXPIRED"
	StatusCompletedPendingVoucher   PromotionStatus = "COMPLETED_PENDING_VOUCHER"
	StatusVoucherIssued             PromotionStatus = "VOUCHER_ISSUED"
	StatusErrorRecoverable          PromotionStatus = "ERROR_RECOVERABLE"
	StatusIssueRetryRequired        PromotionStatus = "ISSUE_RETRY_REQUIRED"
)

type VoucherStatus string

const (
	VoucherStatusPending VoucherStatus = "PENDING"
	VoucherStatusIssued  VoucherStatus = "ISSUED"
	VoucherStatusFailed  VoucherStatus = "FAILED"
)

type TriggerType string

const (
	TriggerSavedProduct   TriggerType = "saved_product"
	TriggerStartedCheckout TriggerType = "started_checkout"
)

type CheckoutPromotion struct {
	PromotionID   uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SessionID     string          `gorm:"type:varchar(255);not null;uniqueIndex"`
	UserID        *uuid.UUID      `gorm:"type:uuid;index"`
	DeviceID      *string         `gorm:"type:varchar(255);index"`
	TriggerType   TriggerType     `gorm:"type:varchar(50);not null"`
	Status        PromotionStatus `gorm:"type:varchar(50);not null"`
	StartedAt     time.Time       `gorm:"type:timestamptz;not null"`
	ExpiresAt     time.Time       `gorm:"type:timestamptz;not null"`
	CompletedAt   *time.Time      `gorm:"type:timestamptz"`
	VoucherID     *uuid.UUID      `gorm:"type:uuid"`
	VoucherStatus *VoucherStatus  `gorm:"type:varchar(50)"`

	CampaignCode string `gorm:"type:varchar(255);not null;default:'ACCESSORY_NEXT_PURCHASE_15'"`
	RewardType   string `gorm:"type:varchar(50);not null;default:'percentage'"`
	RewardValue  int    `gorm:"not null;default:15"`
	RewardScope  string `gorm:"type:varchar(255);not null;default:'accessories_next_purchase'"`

	CooldownUntil *time.Time `gorm:"type:timestamptz"`

	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
}

package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Promotion struct {
	ID                    uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Campaign              string         `gorm:"type:varchar(255);not null"`
	CouponCode            string         `gorm:"type:varchar(50);unique"`
	Discount              datatypes.JSON `gorm:"type:jsonb;not null"`
	DiscountType          string         `gorm:"type:varchar(32);not null;default:PERCENTAGE"`
	DiscountValue         int64          `gorm:"type:bigint;not null;default:0"`
	MinimumSubtotal       int64          `gorm:"type:bigint;not null;default:0"`
	MaximumDiscountAmount *int64         `gorm:"type:bigint"`
	Active                bool           `gorm:"not null;default:true"`
	StartsAt              *time.Time     `gorm:"type:timestamptz"`
	EndsAt                *time.Time     `gorm:"type:timestamptz"`
}

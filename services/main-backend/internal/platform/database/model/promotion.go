package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Promotion struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Campaign   string         `gorm:"type:varchar(255);not null"`
	CouponCode string         `gorm:"type:varchar(50);unique"`
	Discount   datatypes.JSON `gorm:"type:jsonb;not null"`
}

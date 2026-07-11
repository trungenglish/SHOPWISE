package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Order struct {
	ID       uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID   uuid.UUID      `gorm:"type:uuid;index;not null"`
	Status   string         `gorm:"type:varchar(50);not null"`
	Payment  datatypes.JSON `gorm:"type:jsonb"`
	Delivery datatypes.JSON `gorm:"type:jsonb"`
}

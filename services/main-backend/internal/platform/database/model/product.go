package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Product struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SKU            string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name           string         `gorm:"type:varchar(255);not null"`
	Brand          string         `gorm:"type:varchar(100);index"`
	Category       string         `gorm:"type:varchar(100);index"`
	Price          int64          `gorm:"type:bigint;index;not null;check:price >= 0"`
	Specifications datatypes.JSON `gorm:"type:jsonb"`
	Metadata       datatypes.JSON `gorm:"type:jsonb"`

	Inventory *Inventory `gorm:"foreignKey:ProductID"`
}

type Inventory struct {
	ProductID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	Stock        int       `gorm:"default:0"`
	Warehouse    string    `gorm:"type:varchar(100)"`
	Availability bool      `gorm:"index"`
}

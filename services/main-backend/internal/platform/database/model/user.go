package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Identity  string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`

	Preferences UserPreference `gorm:"foreignKey:UserID"`
}

type UserPreference struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID      uuid.UUID      `gorm:"type:uuid;index;not null"`
	Preferences datatypes.JSON `gorm:"type:jsonb;not null"`
	UpdatedAt   time.Time      `gorm:"default:now()"`
}

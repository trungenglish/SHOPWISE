package postgres

import (
	"gorm.io/gorm"
)

// Migrate automates the migration of decision memory tables
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&DecisionSession{},
		&SessionMessage{},
		&ExtractedPreference{},
	)
}

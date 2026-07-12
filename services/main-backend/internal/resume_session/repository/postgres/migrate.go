package postgres

import (
	"gorm.io/gorm"
)

// Migrate automates the migration of resume session tables
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&ResumeToken{},
		&NotificationLog{},
	)
}

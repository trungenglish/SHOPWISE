package database

import (
	"fmt"

	"gorm.io/gorm"
)

const migrateProductPriceToWholeVND = `
DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_schema = 'public'
          AND table_name = 'products'
          AND column_name = 'price'
          AND data_type <> 'bigint'
    ) THEN
        ALTER TABLE products
        ALTER COLUMN price TYPE BIGINT
        USING ROUND(price)::BIGINT;
    END IF;
END $$;
`

func migrateCommerceMoney(database *gorm.DB) error {
	if err := database.Exec(migrateProductPriceToWholeVND).Error; err != nil {
		return fmt.Errorf("migrate product price to whole VND: %w", err)
	}
	return nil
}

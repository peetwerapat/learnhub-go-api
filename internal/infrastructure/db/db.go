package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfgHost, cfgUser, cfgPassword, cfgName, cfgPort string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfgHost, cfgUser, cfgPassword, cfgName, cfgPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}

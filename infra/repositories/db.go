package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	connectionUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&Status{},
		&Kind{},
		&PaymentType{},
		&Property{},
	); err != nil {
		return nil, err
	}

	return db, nil
}

package database

import (
	"log"
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
		log.Fatal("Unable to connect to database")
	}

	autoMigrateError := db.AutoMigrate(
		&Status{},
		&Kind{},
		&PaymentType{},
		&Property{},
	)

	if autoMigrateError != nil {
		log.Fatal("Unable to run AutoMigrate")
	}

	return db, nil
}

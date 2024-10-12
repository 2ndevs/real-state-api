package database

import (
	"errors"
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
		return nil, errors.New("Unable to create database connection")
	}

	autoMigrateError := db.AutoMigrate(
		&Status{},
		&Kind{},
		&PaymentType{},
		&Property{},
		&User{},
		&Role{},
	)
	if autoMigrateError != nil {
		return nil, errors.New("Unable to auto migrate schemas")
	}

	return db, nil
}

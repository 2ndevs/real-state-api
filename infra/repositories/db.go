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
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		return nil, errors.New("Unable to create database connection")
	}

	autoMigrateError := db.AutoMigrate(
		&Status{},
		&Property{},
		&Kind{},
		&PaymentType{},
		&User{},
		&Role{},
		&NegotiationType{},
		&UnitOfMeasurement{},
		&InterestedUser{},
		&Visit{},
	)
	if autoMigrateError != nil {
		return nil, errors.New("Unable to auto migrate schemas")
	}

	return db, nil
}

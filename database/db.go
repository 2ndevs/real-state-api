package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	connectionUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	db.AutoMigrate(&Status{}, &Kind{}, &PaymentType{}, &Property{})
}

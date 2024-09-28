package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

// STATUS FIELD
//  - ACTIVE
//  - INACTIVE/DISABLED

type Status struct {
	gorm.Model

	Name string `gorm:"index"`
}

type Kind struct {
	gorm.Model

	Name   string `gorm:"index"`
	Status Status
}

type PaymentType struct {
	gorm.Model

	Name   string `gorm:"index"`
	Status Status
}

type Property struct {
	gorm.Model

	Kind         *Kind `gorm:"index"`
	Price        float64
	Size         uint
	Rooms        uint
	Kitchens     uint
	Bathrooms    uint
	Address      string
	Summary      *string
	Details      []*string
	Coords       []uint
	PaymentTypes []PaymentType

	Status Status
}

func Schema() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env variables")
	}

	connectionUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	instance, err := db.DB()
	if err != nil {
		log.Fatal("Unable to create generic database interface")
	}

	defer instance.Close()

	db.AutoMigrate(&Status{}, &Kind{}, &PaymentType{}, &Property{})
}

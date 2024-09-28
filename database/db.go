package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model

	Name string `gorm:"index"`
}

type Kind struct {
	gorm.Model

	Name   string `gorm:"index"`

	Status Status
	StatusID uint
}

type PaymentType struct {
	gorm.Model

	Name string `gorm:"index"`

	Status   Status
	StatusID uint
}

type Property struct {
	gorm.Model

	Size      uint
	Rooms     uint
	Kitchens  uint
	Bathrooms uint
	Address   string
	Summary   string
	Details   string
	latitude  float64
	longitude float64
	Price     float64

	Kind          Kind
	KindID        uint `gorm:"index"`
	Status        Status
	StatusID      uint
	PaymentType   PaymentType
	PaymentTypeID uint
}

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

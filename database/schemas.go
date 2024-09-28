package database

import "gorm.io/gorm"

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
	Latitude  float64
	Longitude float64
	Price     float64

	Kind          Kind
	KindID        uint `gorm:"index"`
	Status        Status
	StatusID      uint
	PaymentType   PaymentType
	PaymentTypeID uint
}

package entities

import "gorm.io/gorm"

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

	KindID        uint `gorm:"index"`
	StatusID      uint
	PaymentTypeID uint
}

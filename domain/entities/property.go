package entities

import "gorm.io/gorm"

type Property struct {
	gorm.Model

  Size      uint `validate:"required,min=1"`
  Rooms     uint `validate:"required,min=0"`
  Kitchens  uint `validate:"required,min=0"`
  Bathrooms uint `validate:"required,min=0"`
  Address   string `validate:"required"`
  Summary   string `validate:"required"`
  Details   string `validate:"required"`
  Latitude  float64 `validate:"required,min=0"`
  Longitude float64 `validate:"required,min=0"`
  Price     float64  `validate:"required,min=1"`

	KindID        uint `gorm:"index"`
  StatusID      uint `validate:"required,min=1"`
  PaymentTypeID uint `validate:"required,min=1"`
}

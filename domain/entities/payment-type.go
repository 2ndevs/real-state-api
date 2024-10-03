package entities

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model

  Name string `gorm:"index" validate:"required"`

  StatusID uint `validate:"required,min=1"`
}

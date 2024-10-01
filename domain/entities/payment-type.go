package entities

import "gorm.io/gorm"

type PaymentType struct {
	gorm.Model

	Name string `gorm:"index"`

	StatusID uint
}

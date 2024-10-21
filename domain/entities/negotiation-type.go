package entities

import "gorm.io/gorm"

type NegotiationType struct {
	gorm.Model

	Name string `gorm:"index" validate:"required"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
}

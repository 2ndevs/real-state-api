package entities

import "main/core"

type NegotiationType struct {
	core.Model

	Name string `json:"name" gorm:"index" validate:"required"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
}

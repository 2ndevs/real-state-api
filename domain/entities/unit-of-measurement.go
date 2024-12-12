package entities

import (
	"main/core"
)

type UnitOfMeasurement struct {
	core.Model

	Name         string `json:"name" gorm:"index,unique" validate:"required,gte=3,lte=100"`
	Abbreviation string `json:"abbreviation" validate:"required,gte=2"`

	Status   *Status `json:"status"`
	StatusID uint    `json:"status_id" validate:"required,min=1"`
}

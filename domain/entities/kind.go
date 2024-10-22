package entities

import "main/core"

type Kind struct {
	core.Model

	Name string `json:"name" gorm:"index" validate:"required,gte=3,lte=100"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
}

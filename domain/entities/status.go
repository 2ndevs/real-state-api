package entities

import "main/core"

type Status struct {
	core.Model

	Name string `json:"name" gorm:"index" validate:"required"`
}

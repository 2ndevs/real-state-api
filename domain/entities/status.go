package entities

import "gorm.io/gorm"

type Status struct {
	gorm.Model

  Name string `gorm:"index" validate:"required"`
}

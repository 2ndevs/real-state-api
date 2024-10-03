package entities

import (
	"gorm.io/gorm"
)

type Kind struct {
	gorm.Model

  Name string `gorm:"index" validate:"required,gte=3,lte=100"`

  StatusID uint `validate:"required,min=1"`
}

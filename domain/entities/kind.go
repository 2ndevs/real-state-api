package entities

import (
	"gorm.io/gorm"
)

type Kind struct {
	gorm.Model

	Name     string `gorm:"index"`
	StatusID uint
}

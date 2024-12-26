package entities

import (
	"main/core"
)

type Visit struct {
	core.Model

	PropertyID uint   `gorm:"index;not null" validate:"required"`
	UserID     string `gorm:"index;not null" validate:"required"`

	Property Property `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-"`
}

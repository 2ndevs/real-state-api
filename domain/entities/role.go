package entities

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name        string         `json:"name" validate:"required,gte=3"`
	Permissions pq.StringArray `gorm:"type:text[]" json:"permissions" validate:"required,min=1"`

	StatusID uint `json:"status_id" validate:"required,gte=1"`
}

package entities

import (
	"main/core"

	"github.com/lib/pq"
)

type Role struct {
	core.Model

	Name        string         `json:"name" validate:"required,gte=3"`
	Permissions pq.StringArray `gorm:"type:text[]" json:"permissions" validate:"required,min=1"`

	StatusID uint `json:"status_id" validate:"required,gte=1"`
}

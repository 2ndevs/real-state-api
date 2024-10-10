package entities

import "gorm.io/gorm"

type Role struct {
  gorm.Model

  Name string `json:"name" validate:"required,gte=3"`

  StatusID uint `json:"status_id" validate:"required,gte=1"`
}

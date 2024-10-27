package entities

import "main/core"

type User struct {
	core.Model

	Name         string `json:"name" validate:"required,gte=5,lte=25"`
	Email        string `gorm:"embedded" json:"email" validate:"required,email"`
	PasswordHash string `json:"password_hash" validate:"required"`
	RefreshToken string `json:"refresh_token"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
	RoleID   uint `json:"role_id" validate:"required,min=1"`
}

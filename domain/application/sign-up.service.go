package application

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SignUpService struct {
	Database  *gorm.DB
	Validator *validator.Validate
}

type SignUpRequest struct {
	Name         string `json:"name" validate:"required,gte=5,lte=25"`
	Email        string `gorm:"embedded" json:"email" validate:"required,email"`
	Password string `json:"password_hash" validate:"required,gte=6,lte=36"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
	RoleID   uint `json:"role_id" validate:"required,min=1"`
}

type SignUpResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`

	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`

	RoleID   uint `json:"role_id"`
}

func (self SignUpService) Execute(user SignUpRequest) (*SignUpResponse, error) {
	var response SignUpResponse 

	return &response, nil 
}

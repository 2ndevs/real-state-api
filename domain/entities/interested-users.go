package entities

import "main/core"

type InterestedUser struct {
	core.Model

	FirstName string  `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string  `json:"last_name" validate:"required,min=2,max=50"`
	Phone     string  `json:"phone" validate:"required,min=11,max=11"`
	Email     *string `json:"email"`
	Answered  *bool   `json:"answered" gorm:"default:false"`

	StatusID *uint `json:"status_id" validate:"required,min=1"`
	Status   *Status

	PropertyID uint `json:"property_id" validate:"required,min=1"`
	Property   *Property
}

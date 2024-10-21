package application

import (
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreatePropertyService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

type CreatePropertyServiceRequest struct {
	entities.Property
	PreviewImages []byte `validated:"required,min=1"`
}

func (self *CreatePropertyService) Execute(property CreatePropertyServiceRequest) (*entities.Property, error) {
	validationErr := self.Validated.Struct(property)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	createPropertyTransaction := self.Database.Create(&property)
	if createPropertyTransaction.Error != nil {
		return nil, createPropertyTransaction.Error
	}

	return &property.Property, nil
}

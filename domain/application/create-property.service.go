package application

import (
	"errors"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreatePropertyService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *CreatePropertyService) Execute(property entities.Property) (*entities.Property, error) {
	validationErr := self.Validated.Struct(property)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	createPropertyTransaction := self.Database.Create(&property)
	if createPropertyTransaction.Error != nil {
		return nil, createPropertyTransaction.Error
	}

	return &property, nil
}

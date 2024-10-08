package application

import (
	"errors"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"

	"gorm.io/gorm"
)

type CreatePropertyService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (propertyService *CreatePropertyService) Execute(property entities.Property) (*entities.Property, error) {
	validate, ctxErr := middlewares.GetValidator(propertyService.Request)
	if ctxErr != nil {
		return nil, ctxErr
	}

	validationErr := validate.Struct(property)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	createPropertyTransaction := propertyService.Database.Create(&property)
	if createPropertyTransaction.Error != nil {
		return nil, createPropertyTransaction.Error
	}

	return &property, nil
}

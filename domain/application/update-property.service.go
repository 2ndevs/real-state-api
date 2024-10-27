package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdatePropertyService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdatePropertyService) Execute(property entities.Property, propertyID uint64) (*entities.Property, error) {
	validationErr := self.Validated.Struct(property)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingProperty *entities.Property
	query := self.Database.Model(&entities.Property{}).Where("id = ? and deleted_at IS NULL and is_sold != true", propertyID)

	existingPropertyDatabaseResponse := query.First(&existingProperty)
	if errors.Is(existingPropertyDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingPropertyDatabaseResponse.Error != nil {
		return nil, existingPropertyDatabaseResponse.Error
	}

	property.ID = existingProperty.ID

	updatePropertyTransaction := self.Database.Save(&property)
	if updatePropertyTransaction.Error != nil {
		return nil, updatePropertyTransaction.Error
	}

	return &property, nil
}

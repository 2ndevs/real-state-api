package application

import (
	"errors"
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
		return nil, errors.Join(errors.New("validation errors: "), validationErr)
	}

	var existingProperty *entities.Property
	query := self.Database.Model(&entities.Property{}).Where("id = ?", propertyID)

	existingPropertyDatabaseResponse := query.First(&existingProperty)
	if errors.Is(existingPropertyDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("property not found")
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

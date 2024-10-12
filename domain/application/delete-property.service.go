package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeletePropertyService struct {
	Database *gorm.DB
}

func (self *DeletePropertyService) Execute(propertyID uint64) (*entities.Property, error) {
	var existingProperty *entities.Property
	query := self.Database.Model(&entities.Property{}).Where("id = ?", propertyID)

	existingPropertyDatabaseResponse := query.First(&existingProperty)
	if errors.Is(existingPropertyDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingPropertyDatabaseResponse.Error != nil {
		return nil, existingPropertyDatabaseResponse.Error
	}

	deletePropertyTransaction := self.Database.Delete(existingProperty)
	if deletePropertyTransaction.Error != nil {
		return nil, deletePropertyTransaction.Error
	}

	return existingProperty, nil
}

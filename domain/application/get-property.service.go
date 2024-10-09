package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetPropertyService struct {
	PropertyID uint64
	Database   *gorm.DB
}

func (self *GetPropertyService) Execute() (*entities.Property, error) {
	property := entities.Property{}

	getPropertyTransaction := self.Database.Find(&property, self.PropertyID).Where("deleted_at IS NULL").First(&property)
	if getPropertyTransaction.Error != nil {
		return nil, getPropertyTransaction.Error
	}

	return &property, nil
}

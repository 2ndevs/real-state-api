package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetPropertyService struct {
	Database *gorm.DB
}

func (self *GetPropertyService) Execute(propertyID uint64, userIdentity *string) (*entities.Property, error) {
	property := entities.Property{}

	getPropertyTransaction := self.Database.Find(&property, propertyID).Where("deleted_at IS NULL").First(&property)
	if getPropertyTransaction.Error != nil {
		return nil, getPropertyTransaction.Error
	}

	if userIdentity == nil {
		return &property, nil
	}

	property.VisitedBy = append(property.VisitedBy, *userIdentity)
	self.Database.Save(&property)

	return &property, nil
}

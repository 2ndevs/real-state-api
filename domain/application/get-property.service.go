package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetPropertyService struct {
	Database *gorm.DB
}

func (self *GetPropertyService) Execute(propertyID uint64, userIdentity *string) (*entities.Property, error) {
	property := entities.Property{}

	if userIdentity != nil {
		visit := entities.Visit{
			PropertyID: uint(propertyID),
			UserID:     *userIdentity,
		}

		visitTransaction := self.Database.Save(&visit)
		if visitTransaction.Error != nil {
			return nil, visitTransaction.Error
		}
	}

	getPropertyTransaction := self.Database.Preload(clause.Associations).Preload("Visits").Find(&property, propertyID).Where("deleted_at IS NULL and sold_at IS NULL").First(&property)
	if getPropertyTransaction.Error != nil {
		return nil, getPropertyTransaction.Error
	}

	return &property, nil
}

package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetPropertyService struct {
	Request    *http.Request
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

package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetHighlightedPropertiesService struct {
	Database *gorm.DB
}

func (self *GetHighlightedPropertiesService) Execute() (*[]entities.Property, error) {
	var properties []entities.Property
	query := self.Database.Model(&entities.Property{})

	query = query.Where("deleted_at IS NULL AND is_highlight = true AND is_sold != true")
	query = query.Order("updated_at DESC")
	query = query.Limit(10)

	getPropertiesTransaction := query.Preload(clause.Associations).Find(&properties)

	if getPropertiesTransaction.Error != nil {
		return nil, getPropertiesTransaction.Error
	}

	return &properties, nil
}

package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetManyPropertiesService struct {
	SearchFilter *string
	Database     *gorm.DB
}

func (self *GetManyPropertiesService) Execute() (*[]entities.Property, error) {
	var properties []entities.Property
	query := self.Database.Model(&entities.Property{})

	if *self.SearchFilter != "" {
		query = query.Where("details ILIKE ?", "%"+*self.SearchFilter+"%").
			Or("address ILIKE ?", "%"+*self.SearchFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")

	getPropertiesTransaction := query.Find(&properties)

	if getPropertiesTransaction.Error != nil {
		return nil, getPropertiesTransaction.Error
	}

	return &properties, nil
}

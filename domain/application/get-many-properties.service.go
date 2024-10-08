package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyPropertiesService struct {
	Request      *http.Request
	SearchFilter *string
	Database     *gorm.DB
}

func (propertiesService *GetManyPropertiesService) Execute() (*[]entities.Property, error) {
	var properties []entities.Property
	query := propertiesService.Database.Model(&entities.Property{})

	if *propertiesService.SearchFilter != "" {
		query = query.Where("details ILIKE ?", "%"+*propertiesService.SearchFilter+"%").
			Or("address ILIKE ?", "%"+*propertiesService.SearchFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")

	getPropertiesTransaction := query.Find(&properties)

	if getPropertiesTransaction.Error != nil {
		return nil, getPropertiesTransaction.Error
	}

	return &properties, nil
}

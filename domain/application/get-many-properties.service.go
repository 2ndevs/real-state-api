package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyPropertiesService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (propertiesService *GetManyPropertiesService) Execute() (*[]entities.Property, error) {
	nameFilter := propertiesService.Request.URL.Query().Get("search")

	var properties []entities.Property
	query := propertiesService.Database.Model(&entities.Property{})

	if nameFilter != "" {
		query = query.Where("details ILIKE ?", "%"+nameFilter+"%").
			Or("address ILIKE ?", "%"+nameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")

	getPropertiesTransaction := query.Find(&properties)

	if getPropertiesTransaction.Error != nil {
		return nil, getPropertiesTransaction.Error
	}

	return &properties, nil
}

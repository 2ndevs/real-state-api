package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyKindsService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (kindService *GetManyKindsService) Execute() (*[]entities.Kind, error) {
	nameFilter := kindService.Request.URL.Query().Get("name")

	var kinds []entities.Kind
	query := kindService.Database.Model(&entities.Kind{})

	if nameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+nameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getKindsTransaction := query.Find(&kinds)

	if getKindsTransaction.Error != nil {
		return nil, getKindsTransaction.Error
	}

	return &kinds, nil
}

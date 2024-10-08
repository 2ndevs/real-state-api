package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyKindsService struct {
	Request    *http.Request
	NameFilter *string
	Database   *gorm.DB
}

func (kindService *GetManyKindsService) Execute() (*[]entities.Kind, error) {
	var kinds []entities.Kind
	query := kindService.Database.Model(&entities.Kind{})

	if *kindService.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*kindService.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getKindsTransaction := query.Find(&kinds)

	if getKindsTransaction.Error != nil {
		return nil, getKindsTransaction.Error
	}

	return &kinds, nil
}

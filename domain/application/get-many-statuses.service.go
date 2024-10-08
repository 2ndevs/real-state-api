package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyStatusesService struct {
	Request    *http.Request
	NameFilter *string
	Database   *gorm.DB
}

func (statusesService *GetManyStatusesService) Execute() (*[]entities.Status, error) {
	var statuses []entities.Status
	query := statusesService.Database.Model(&entities.Status{})

	if *statusesService.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*statusesService.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getStatusesTransaction := query.Find(&statuses)

	if getStatusesTransaction.Error != nil {
		return nil, getStatusesTransaction.Error
	}

	return &statuses, nil
}

package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyStatusesService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (statusesService *GetManyStatusesService) Execute() (*[]entities.Status, error) {
	nameFilter := statusesService.Request.URL.Query().Get("name")

	var statuses []entities.Status
	query := statusesService.Database.Model(&entities.Status{})

	if nameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+nameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getStatusesTransaction := query.Find(&statuses)

	if getStatusesTransaction.Error != nil {
		return nil, getStatusesTransaction.Error
	}

	return &statuses, nil
}

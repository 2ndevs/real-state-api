package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetStatusService struct {
	Request  *http.Request
	StatusID uint64
	Database *gorm.DB
}

func (statusService *GetStatusService) Execute() (*entities.Status, error) {

	status := entities.Status{}
	getStatusTransaction := statusService.Database.Find(&status, statusService.StatusID).Where("deleted_at IS NULL").First(&status)
	if getStatusTransaction.Error != nil {
		return nil, getStatusTransaction.Error
	}

	return &status, nil
}

package application

import (
	"errors"
	"main/domain/entities"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GetStatusService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (statusService *GetStatusService) Execute() (*entities.Status, error) {
	idParam := statusService.Request.URL.Path[len("/statuses/"):]

	statusId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		return nil, errors.Join(errors.New("invalid id"), validationErr)
	}

	status := entities.Status{}
	getStatusTransaction := statusService.Database.Find(&status, statusId).Where("deleted_at IS NULL").First(&status)
	if getStatusTransaction.Error != nil {
		return nil, getStatusTransaction.Error
	}

	return &status, nil
}

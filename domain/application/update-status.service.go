package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateStatusService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdateStatusService) Execute(status entities.Status, statusID uint64) (*entities.Status, error) {
	validationErr := self.Validated.Struct(status)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingStatus *entities.Status
	query := self.Database.Model(&entities.Status{}).Where("id = ?", statusID)

	existingStatusDatabaseResponse := query.First(&existingStatus)
	if errors.Is(existingStatusDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingStatusDatabaseResponse.Error != nil {
		return nil, existingStatusDatabaseResponse.Error
	}

	status.ID = existingStatus.ID

	updateStatusTransaction := self.Database.Save(&status)
	if updateStatusTransaction.Error != nil {
		return nil, updateStatusTransaction.Error
	}

	return &status, nil
}

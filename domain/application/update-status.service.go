package application

import (
	"errors"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateStatusService struct {
	Validated *validator.Validate
	StatusID  uint64
	Database  *gorm.DB
}

func (self *UpdateStatusService) Execute(status entities.Status) (*entities.Status, error) {
	validationErr := self.Validated.Struct(status)
	if validationErr != nil {
		return nil, errors.Join(errors.New("Erros de validação: "), validationErr)
	}

	var existingStatus *entities.Status
	query := self.Database.Model(&entities.Status{}).Where("id = ?", self.StatusID)

	existingStatusDatabaseResponse := query.First(&existingStatus)
	if errors.Is(existingStatusDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Status não encontrado")
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

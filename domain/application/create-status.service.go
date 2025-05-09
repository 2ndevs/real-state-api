package application

import (
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateStatusService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *CreateStatusService) Execute(status entities.Status) (*entities.Status, error) {
	validationErr := self.Validated.Struct(status)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingStatus *entities.Status

	query := self.Database.Model(&entities.Status{}).Where("name = ?", status.Name)
	response := query.First(&existingStatus)

	if response.Error == nil {
		return nil, core.EntityAlreadyExistsError
	}

	createStatusTransaction := self.Database.Create(&status)
	if createStatusTransaction.Error != nil {
		return nil, createStatusTransaction.Error
	}

	return &status, nil
}

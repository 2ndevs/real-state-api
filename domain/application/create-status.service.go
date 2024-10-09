package application

import (
	"errors"
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
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	createStatusTransaction := self.Database.Create(&status)
	if createStatusTransaction.Error != nil {
		return nil, createStatusTransaction.Error
	}

	return &status, nil
}

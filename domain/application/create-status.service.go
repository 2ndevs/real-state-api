package application

import (
	"errors"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"

	"gorm.io/gorm"
)

type CreateStatusService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (self *CreateStatusService) Execute(status entities.Status) (*entities.Status, error) {
	validate, ctxErr := middlewares.GetValidator(self.Request)
	if ctxErr != nil {
		return nil, ctxErr
	}

	validationErr := validate.Struct(status)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	createStatusTransaction := self.Database.Create(&status)
	if createStatusTransaction.Error != nil {
		return nil, createStatusTransaction.Error
	}

	return &status, nil
}

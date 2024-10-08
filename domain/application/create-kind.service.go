package application

import (
	"errors"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"

	"gorm.io/gorm"
)

type CreateKindService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (kindService *CreateKindService) Execute(kind entities.Kind) (*entities.Kind, error) {
	validate, ctxErr := middlewares.GetValidator(kindService.Request)
	if ctxErr != nil {
		return nil, ctxErr
	}

	validationErr := validate.Struct(kind)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	createKindTransaction := kindService.Database.Create(&kind)
	if createKindTransaction.Error != nil {
		return nil, createKindTransaction.Error
	}

	return &kind, nil
}

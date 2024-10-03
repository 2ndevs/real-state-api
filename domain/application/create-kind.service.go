package application

import (
	"errors"
	"log"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"

	"gorm.io/gorm"
)

type CreateKindService struct {
	Request *http.Request
	Database      *gorm.DB
}

func (self *CreateKindService) Execute(kind entities.Kind) (*entities.Kind, error) {
	validate, ctxErr := middlewares.GetValidator(self.Request)
	if ctxErr != nil {
		return nil, ctxErr
	}

	validationErr := validate.Struct(kind)
	if validationErr != nil {
    return nil, errors.Join(errors.New("Validation error: "), validationErr)
	}

  log.Printf("%v", validationErr)

	createKindTransaction := self.Database.Create(&kind)
	if createKindTransaction.Error != nil {
		return nil, createKindTransaction.Error
	}

	return &kind, nil
}

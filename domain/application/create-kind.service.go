package application

import (
	"errors"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateKindService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *CreateKindService) Execute(kind entities.Kind) (*entities.Kind, error) {
	validationErr := self.Validated.Struct(kind)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation errors: "), validationErr)
	}

	createKindTransaction := self.Database.Create(&kind)
	if createKindTransaction.Error != nil {
		return nil, createKindTransaction.Error
	}

	return &kind, nil
}

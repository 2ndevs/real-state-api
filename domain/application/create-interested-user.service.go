package application

import (
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateInterestedUserService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *CreateInterestedUserService) Execute(interestedUser entities.InterestedUser) (*entities.InterestedUser, error) {
	validationErr := self.Validated.Struct(interestedUser)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	createInterestedUserTransaction := self.Database.Create(&interestedUser)
	if createInterestedUserTransaction.Error != nil {
		return nil, createInterestedUserTransaction.Error
	}

	return &interestedUser, nil
}

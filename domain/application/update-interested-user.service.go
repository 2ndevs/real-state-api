package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateInterestedUserService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdateInterestedUserService) Execute(interestedUser entities.InterestedUser, interestedUserID uint64) (*entities.InterestedUser, error) {
	validationErr := self.Validated.Struct(interestedUser)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingInterestedUser *entities.InterestedUser
	query := self.Database.Model(&entities.InterestedUser{}).Where("id = ?", interestedUserID)

	existingInterestedUserDatabaseResponse := query.First(&existingInterestedUser)
	if errors.Is(existingInterestedUserDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingInterestedUserDatabaseResponse.Error != nil {
		return nil, existingInterestedUserDatabaseResponse.Error
	}

	interestedUser.ID = existingInterestedUser.ID

	updateInterestedUserTransaction := self.Database.Save(&interestedUser)
	if updateInterestedUserTransaction.Error != nil {
		return nil, updateInterestedUserTransaction.Error
	}

	return &interestedUser, nil
}

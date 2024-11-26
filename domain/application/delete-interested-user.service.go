package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteInterestedUserService struct {
	Database *gorm.DB
}

func (self *DeleteInterestedUserService) Execute(interestedUserID uint64) (*entities.InterestedUser, error) {
	var existingInterestedUser *entities.InterestedUser
	query := self.Database.Model(&entities.InterestedUser{}).Where("id = ?", interestedUserID)

	existingInterestedUserDatabaseResponse := query.First(&existingInterestedUser)
	if errors.Is(existingInterestedUserDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingInterestedUserDatabaseResponse.Error != nil {
		return nil, existingInterestedUserDatabaseResponse.Error
	}

	deleteInterestedUserTransaction := self.Database.Delete(existingInterestedUser)
	if deleteInterestedUserTransaction.Error != nil {
		return nil, deleteInterestedUserTransaction.Error
	}

	return existingInterestedUser, nil
}

package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteUserService struct {
	Database *gorm.DB
}

func (self DeleteUserService) Execute(id string) error {
	user := &entities.User {}

	existingUserTransaction := self.Database.Model(user).Where("id = ?", id).First(user)
	if errors.Is(existingUserTransaction.Error, gorm.ErrRecordNotFound) {
		return core.NotFoundError
	}

	if existingUserTransaction.Error != nil {
		return existingUserTransaction.Error
	}

	deleteUserTransaction := self.Database.Delete(user)
	if deleteUserTransaction.Error != nil {
		return deleteUserTransaction.Error
	}
	
	return nil
}

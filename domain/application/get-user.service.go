package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetUserService struct {
	UserID   uint64
	Database *gorm.DB
}

func (self *GetUserService) Execute() (*entities.User, error) {
	user := entities.User{}

	getUserTransaction := self.Database.Model(&user).Preload(clause.Associations).Find(&user, self.UserID).Where("deleted_at IS NULL").First(&user)
	if getUserTransaction.Error != nil {
		return nil, getUserTransaction.Error
	}

	return &user, nil
}

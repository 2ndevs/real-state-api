package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetManyUsersService struct {
	Database *gorm.DB
}

func (self GetManyUsersService) Execute() ([]entities.User, error) {
	var users []entities.User
	query := self.Database.Model(&entities.User{}).Where("deleted_at IS NULL").Preload(clause.Associations)

	transaction := query.Find(&users)
	if transaction.Error != nil {
		return nil, transaction.Error
	}

	return users, nil
}

package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetManyInterestedUsersService struct {
	Database *gorm.DB
}

func (self *GetManyInterestedUsersService) Execute() (*[]entities.InterestedUser, error) {
	var interestedUsers []entities.InterestedUser
	query := self.Database.Model(&entities.InterestedUser{}).Order("created_at DESC")

	query = query.Where("deleted_at IS NULL")

	getInterestedUsersTransaction := query.Preload(clause.Associations).Find(&interestedUsers)

	if getInterestedUsersTransaction.Error != nil {
		return nil, getInterestedUsersTransaction.Error
	}

	return &interestedUsers, nil
}

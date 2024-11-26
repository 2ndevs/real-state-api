package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetInterestedUserService struct {
	Database *gorm.DB
}

func (self *GetInterestedUserService) Execute(ID uint) (*entities.InterestedUser, error) {
	interestedUser := entities.InterestedUser{}

	getInterestedUserTransaction := self.Database.Model(&interestedUser).Find(&interestedUser, ID).Where("deleted_at IS NULL").First(&interestedUser)
	if getInterestedUserTransaction.Error != nil {
		return nil, getInterestedUserTransaction.Error
	}

	return &interestedUser, nil
}

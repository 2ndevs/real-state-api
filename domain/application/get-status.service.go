package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetStatusService struct {
	StatusID uint64
	Database *gorm.DB
}

func (self *GetStatusService) Execute() (*entities.Status, error) {

	status := entities.Status{}
	getStatusTransaction := self.Database.Find(&status, self.StatusID).Where("deleted_at IS NULL").First(&status)
	if getStatusTransaction.Error != nil {
		return nil, getStatusTransaction.Error
	}

	return &status, nil
}

package application

import (
	"errors"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteStatusService struct {
	Database *gorm.DB
}

func (self *DeleteStatusService) Execute(statusID uint64) (*entities.Status, error) {
	var existingStatus *entities.Status
	query := self.Database.Model(&entities.Status{}).Where("id = ?", statusID)

	existingStatusDatabaseResponse := query.First(&existingStatus)
	if errors.Is(existingStatusDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("status not found")
	}

	if existingStatusDatabaseResponse.Error != nil {
		return nil, existingStatusDatabaseResponse.Error
	}

	deleteStatusTransaction := self.Database.Delete(existingStatus)
	if deleteStatusTransaction.Error != nil {
		return nil, deleteStatusTransaction.Error
	}

	return existingStatus, nil
}

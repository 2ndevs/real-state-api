package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteKindService struct {
	Database *gorm.DB
}

func (self *DeleteKindService) Execute(kindID uint64) (*entities.Kind, error) {
	var existingKind *entities.Kind
	query := self.Database.Model(&entities.Kind{}).Where("id = ?", kindID)

	existingKindDatabaseResponse := query.First(&existingKind)
	if errors.Is(existingKindDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingKindDatabaseResponse.Error != nil {
		return nil, existingKindDatabaseResponse.Error
	}

	deleteKindTransaction := self.Database.Delete(existingKind)
	if deleteKindTransaction.Error != nil {
		return nil, deleteKindTransaction.Error
	}

	return existingKind, nil
}

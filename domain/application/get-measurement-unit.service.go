package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetMeasurementUnitService struct {
	Database *gorm.DB
}

func (self GetMeasurementUnitService) Execute(id uint) (*entities.UnitOfMeasurement, error) {
	entity := entities.UnitOfMeasurement {}
	
	transaction := self.Database.Find(&entity, id).Where("deleted_at IS NULL").First(&entity)
	if errors.Is(transaction.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if transaction.Error != nil {
		return nil, transaction.Error
	}

	return &entity, nil
}

package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteMeasurementUnitService struct {
	Database *gorm.DB
}

func (self DeleteMeasurementUnitService) Execute(id uint) (*entities.UnitOfMeasurement, error){
	var entity *entities.UnitOfMeasurement
	
	existingEntityTransaction := self.Database.Model(&entities.UnitOfMeasurement{}).Where("id = ?", id).First(&entity)
	if errors.Is(existingEntityTransaction.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingEntityTransaction.Error != nil {
		return nil, existingEntityTransaction.Error
	}

	deleteTransaction := self.Database.Delete(entity)
	if deleteTransaction.Error != nil {
		return nil, deleteTransaction.Error
	}

	return entity, nil
}

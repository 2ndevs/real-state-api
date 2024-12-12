package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetManyMeasurementUnitService struct {
	Database *gorm.DB
}

func (self GetManyMeasurementUnitService) Execute() ([]entities.UnitOfMeasurement, error) {
	var response []entities.UnitOfMeasurement
	
	model := self.Database.Model(&entities.UnitOfMeasurement{})
	query := model.Preload(clause.Associations).Where("deleted_at IS NULL")

	transaction := query.Find(&response)
	if transaction.Error != nil {
		return nil, transaction.Error
	}

	return response, nil
}

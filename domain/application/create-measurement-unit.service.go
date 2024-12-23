package application

import (
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateMeasurementUnitService struct {
	Validator *validator.Validate
	Database  *gorm.DB
}

func (self CreateMeasurementUnitService) Execute(data entities.UnitOfMeasurement) (*entities.UnitOfMeasurement, error) {
	validationErr := self.Validator.Struct(data)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	transaction := self.Database.Create(&data)
	if transaction.Error != nil {
		return nil, transaction.Error
	}

	return &data, nil
}

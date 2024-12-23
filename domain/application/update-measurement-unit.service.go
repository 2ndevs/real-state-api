package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateMeasurementUnitService struct {
	Database  *gorm.DB
	Validator *validator.Validate
}

func (self UpdateMeasurementUnitService) Execute(id uint, payload entities.UnitOfMeasurement) (*entities.UnitOfMeasurement, error) {
	entity := entities.UnitOfMeasurement{}

	validationErr := self.Validator.Struct(&payload)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	existingEntityTransaction := self.Database.Find(&entity, id).Where("deleted_at IS NULL").First(&entity)
	if errors.Is(existingEntityTransaction.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingEntityTransaction.Error != nil {
		return nil, existingEntityTransaction.Error
	}

	entity.Name = payload.Name
	entity.StatusID = payload.StatusID
	entity.Abbreviation = payload.Abbreviation

	updateEntityTransaction := self.Database.Save(entity)
	if updateEntityTransaction.Error != nil {
		return nil, updateEntityTransaction.Error
	}

	return &entity, nil
}

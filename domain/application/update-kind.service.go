package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateKindService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdateKindService) Execute(kind entities.Kind, kindID uint64) (*entities.Kind, error) {
	validationErr := self.Validated.Struct(kind)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingKind *entities.Kind
	query := self.Database.Model(&entities.Kind{}).Where("id = ?", kindID)

	existingKindDatabaseResponse := query.First(&existingKind)
	if errors.Is(existingKindDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingKindDatabaseResponse.Error != nil {
		return nil, existingKindDatabaseResponse.Error
	}

	kind.ID = existingKind.ID

	updateKindTransaction := self.Database.Save(&kind)
	if updateKindTransaction.Error != nil {
		return nil, updateKindTransaction.Error
	}

	return &kind, nil
}

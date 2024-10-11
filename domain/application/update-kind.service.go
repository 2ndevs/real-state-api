package application

import (
	"errors"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateKindService struct {
	Validated *validator.Validate
	KindID    uint64
	Database  *gorm.DB
}

func (self *UpdateKindService) Execute(kind entities.Kind) (*entities.Kind, error) {
	validationErr := self.Validated.Struct(kind)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation errors: "), validationErr)
	}

	var existingKind *entities.Kind
	query := self.Database.Model(&entities.Kind{}).Where("id = ?", self.KindID)

	existingKindDatabaseResponse := query.First(&existingKind)
	if errors.Is(existingKindDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("kind not found")
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

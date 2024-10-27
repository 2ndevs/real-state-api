package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateNegotiationTypeService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdateNegotiationTypeService) Execute(negotiationType entities.NegotiationType, negotiationTypeID uint64) (*entities.NegotiationType, error) {
	validationErr := self.Validated.Struct(negotiationType)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingNegotiationType *entities.NegotiationType
	query := self.Database.Model(&entities.NegotiationType{}).Where("id = ?", negotiationTypeID)

	existingNegotiationTypeDatabaseResponse := query.First(&existingNegotiationType)
	if errors.Is(existingNegotiationTypeDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingNegotiationTypeDatabaseResponse.Error != nil {
		return nil, existingNegotiationTypeDatabaseResponse.Error
	}

	var sameNegotiationType *entities.NegotiationType

	findSameQuery := self.Database.Model(&entities.NegotiationType{}).Where("name = ? AND id != ?", negotiationType.Name, existingNegotiationType.ID)
	response := findSameQuery.First(&sameNegotiationType)

	if response.Error == nil {
		return nil, core.EntityAlreadyExistsError
	}

	negotiationType.ID = existingNegotiationType.ID

	updateNegotiationTypeTransaction := self.Database.Save(&negotiationType)
	if updateNegotiationTypeTransaction.Error != nil {
		return nil, updateNegotiationTypeTransaction.Error
	}

	return &negotiationType, nil
}

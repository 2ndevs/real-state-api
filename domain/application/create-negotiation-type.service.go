package application

import (
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateNegotiationTypeService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *CreateNegotiationTypeService) Execute(negotiationType entities.NegotiationType) (*entities.NegotiationType, error) {
	validationErr := self.Validated.Struct(negotiationType)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingNegotiationType *entities.NegotiationType

	query := self.Database.Model(&entities.NegotiationType{}).Where("name = ?", negotiationType.Name)
	response := query.First(&existingNegotiationType)

	if response.Error == nil {
		return nil, core.EntityAlreadyExistsError
	}

	createNegotiationTypeTransaction := self.Database.Create(&negotiationType)
	if createNegotiationTypeTransaction.Error != nil {
		return nil, createNegotiationTypeTransaction.Error
	}

	return &negotiationType, nil
}

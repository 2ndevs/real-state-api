package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteNegotiationTypeService struct {
	Database *gorm.DB
}

func (self *DeleteNegotiationTypeService) Execute(negotiationTypeID uint64) (*entities.NegotiationType, error) {
	var existingNegotiationType *entities.NegotiationType
	query := self.Database.Model(&entities.NegotiationType{}).Where("id = ?", negotiationTypeID)

	existingNegotiationTypeDatabaseResponse := query.First(&existingNegotiationType)
	if errors.Is(existingNegotiationTypeDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingNegotiationTypeDatabaseResponse.Error != nil {
		return nil, existingNegotiationTypeDatabaseResponse.Error
	}

	deleteNegotiationTypeTransaction := self.Database.Delete(existingNegotiationType)
	if deleteNegotiationTypeTransaction.Error != nil {
		return nil, deleteNegotiationTypeTransaction.Error
	}

	return existingNegotiationType, nil
}

package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetNegotiationTypeService struct {
	NegotiationTypeID uint64
	Database          *gorm.DB
}

func (self *GetNegotiationTypeService) Execute() (*entities.NegotiationType, error) {
	negotiationType := entities.NegotiationType{}

	getNegotiationTypeTransaction := self.Database.Find(&negotiationType, self.NegotiationTypeID).Where("deleted_at IS NULL").First(&negotiationType)
	if getNegotiationTypeTransaction.Error != nil {
		return nil, getNegotiationTypeTransaction.Error
	}

	return &negotiationType, nil
}

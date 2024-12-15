package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetManyNegotiationTypesService struct {
	NameFilter *string
	Database   *gorm.DB
}

func (self *GetManyNegotiationTypesService) Execute() (*[]entities.NegotiationType, error) {
	var negotiationTypes []entities.NegotiationType
	query := self.Database.Model(&entities.NegotiationType{})

	if *self.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*self.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getNegotiationTypesTransaction := query.Preload(clause.Associations).Find(&negotiationTypes)

	if getNegotiationTypesTransaction.Error != nil {
		return nil, getNegotiationTypesTransaction.Error
	}

	return &negotiationTypes, nil
}

package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetManyKindsService struct {
	NameFilter *string
	Database   *gorm.DB
}

func (self *GetManyKindsService) Execute() (*[]entities.Kind, error) {
	var kinds []entities.Kind
	query := self.Database.Model(&entities.Kind{})

	if *self.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*self.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getKindsTransaction := query.Find(&kinds)

	if getKindsTransaction.Error != nil {
		return nil, getKindsTransaction.Error
	}

	return &kinds, nil
}

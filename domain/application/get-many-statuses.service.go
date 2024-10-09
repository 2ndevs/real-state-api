package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetManyStatusesService struct {
	NameFilter *string
	Database   *gorm.DB
}

func (self *GetManyStatusesService) Execute() (*[]entities.Status, error) {
	var statuses []entities.Status
	query := self.Database.Model(&entities.Status{})

	if *self.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*self.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getStatusesTransaction := query.Find(&statuses)

	if getStatusesTransaction.Error != nil {
		return nil, getStatusesTransaction.Error
	}

	return &statuses, nil
}

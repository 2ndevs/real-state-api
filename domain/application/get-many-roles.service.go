package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GetManyRolesService struct {
	NameFilter *string
	Database   *gorm.DB
}

func (self *GetManyRolesService) Execute() (*[]entities.Role, error) {
	var roles []entities.Role
	query := self.Database.Model(&entities.Role{})

	if *self.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*self.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getRolesTransaction := query.Preload(clause.Associations).Find(&roles)

	if getRolesTransaction.Error != nil {
		return nil, getRolesTransaction.Error
	}

	return &roles, nil
}

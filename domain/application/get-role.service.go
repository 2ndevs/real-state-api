package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetRoleService struct {
	RoleID   uint64
	Database *gorm.DB
}

func (self *GetRoleService) Execute() (*entities.Role, error) {
	role := entities.Role{}

	getRoleTransaction := self.Database.Model(&role).Find(&role, self.RoleID).Where("deleted_at IS NULL").First(&role)
	if getRoleTransaction.Error != nil {
		return nil, getRoleTransaction.Error
	}

	return &role, nil
}

package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeleteRoleService struct {
	Database *gorm.DB
}

func (self *DeleteRoleService) Execute(roleID uint64) (*entities.Role, error) {
	var existingRole *entities.Role
	query := self.Database.Model(&entities.Role{}).Where("id = ?", roleID)

	existingRoleDatabaseResponse := query.First(&existingRole)
	if errors.Is(existingRoleDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingRoleDatabaseResponse.Error != nil {
		return nil, existingRoleDatabaseResponse.Error
	}

	deleteRoleTransaction := self.Database.Delete(existingRole)
	if deleteRoleTransaction.Error != nil {
		return nil, deleteRoleTransaction.Error
	}

	return existingRole, nil
}

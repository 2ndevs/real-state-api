package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdateRoleService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdateRoleService) Execute(role entities.Role, roleID uint64) (*entities.Role, error) {
	validationErr := self.Validated.Struct(role)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingRole *entities.Role
	query := self.Database.Model(&entities.Role{}).Where("id = ?", roleID)

	existingRoleDatabaseResponse := query.First(&existingRole)
	if errors.Is(existingRoleDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingRoleDatabaseResponse.Error != nil {
		return nil, existingRoleDatabaseResponse.Error
	}

	var sameRole *entities.Role

	findSameQuery := self.Database.Model(&entities.Role{}).Where("name = ? AND id != ?", role.Name, existingRole.ID)
	response := findSameQuery.First(&sameRole)

	if response.Error == nil {
		return nil, core.EntityAlreadyExistsError
	}

	role.ID = existingRole.ID

	updateRoleTransaction := self.Database.Save(&role)
	if updateRoleTransaction.Error != nil {
		return nil, updateRoleTransaction.Error
	}

	return &role, nil
}

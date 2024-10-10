package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateRoleService struct {
	Validate *validator.Validate
	Database *gorm.DB
}

func (self CreateRoleService) Execute(role entities.Role) (*entities.Role, error) {
	validationErr := self.Validate.Struct(role)
	if validationErr != nil {
		return nil, errors.Join(core.InvalidParametersError, validationErr)
	}

	var existingRole *entities.Role
	query := self.Database.Model(&entities.Role{}).Where("name = ?", role.Name)

	existingRoleDatabaseResponse := query.First(&existingRole)
	if existingRoleDatabaseResponse.Error != nil {
		return nil, existingRoleDatabaseResponse.Error
	}

	if existingRole != nil {
		return nil, core.EntityAlreadyExistsError
	}

	databaseResponse := self.Database.Create(&role)
	if databaseResponse.Error != nil {
		return nil, databaseResponse.Error
	}

	return &role, nil
}

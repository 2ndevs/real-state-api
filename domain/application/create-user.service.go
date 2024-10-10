package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateUserService struct {
	Validate *validator.Validate
	Database *gorm.DB
}

func (self CreateUserService) Execute(user entities.User) (*entities.User, error) {
	validationErr := self.Validate.Struct(user)
	if validationErr != nil {
		return nil, errors.Join(core.InvalidParametersError, validationErr)
	}

	var existingUser *entities.User
	query := self.Database.Model(&entities.User{}).Where("email = ?", user.Email)

	existingUserDatabaseResponse := query.First(&existingUser)
	if existingUserDatabaseResponse.Error != nil {
		return nil, existingUserDatabaseResponse.Error
	}

	if existingUser != nil {
		return nil, core.EntityAlreadyExistsError
	}

	databaseResponse := self.Database.Create(&user)
	if databaseResponse.Error != nil {
		return nil, databaseResponse.Error
	}

	return &user, nil
}

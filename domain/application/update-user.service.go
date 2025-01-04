package application

import (
	"errors"
	"main/core"
	"main/domain/entities"
	"main/utils/libs"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UpdateUserService struct {
	Database  *gorm.DB
	Validator *validator.Validate
	Hasher libs.Hashing
}

type UpdateUserServiceRequest struct {
	Name         string `json:"name" validate:"required,gte=3,lte=25"`
	Email        string `gorm:"embedded" json:"email" validate:"required,email"`
	PasswordHash *string `json:"password_hash"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
}

func (self *UpdateUserService) Execute(id string, payload UpdateUserServiceRequest) (*entities.User, error) {
	response := entities.User{}

	err := self.Validator.Struct(payload)
	if err != nil {
		return nil, err
	}

	existingUserTransaction := self.Database.Model(&response).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(&response)
	if errors.Is(existingUserTransaction.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingUserTransaction.Error != nil {
		return nil, existingUserTransaction.Error
	}

	response.Name = payload.Name
	response.Email = payload.Email
	response.StatusID = payload.StatusID
	response.Status.ID = payload.StatusID

	if payload.PasswordHash != nil {
		hashedPassword, err := self.Hasher.EncryptPassword(*payload.PasswordHash)
		if err != nil {
			return nil, err
		}

		response.PasswordHash = *hashedPassword
	}

	updateUserTransaction := self.Database.Save(&response)
	if updateUserTransaction.Error != nil {
		return nil, updateUserTransaction.Error
	}
	
	return &response, nil
}

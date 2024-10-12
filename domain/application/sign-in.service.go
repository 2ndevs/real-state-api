package application

import (
	"errors"
	"main/core"
	"main/domain/entities"
	"main/utils/libs"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignInService struct {
	Database     *gorm.DB
	Validate     *validator.Validate
	Serilization libs.JWT
	Hashing      libs.Hashing
}

func (self SignInService) Execute(params SignInRequest) (*SignInResponse, error) {
	validationErr := self.Validate.Struct(params)
	if validationErr != nil {
		return nil, errors.Join(core.InvalidParametersError, validationErr)
	}

	var user entities.User
	query := self.Database.Model(&entities.User{}).Where("email = ?", params.Email)

	databaseResponse := query.First(&user)
	if databaseResponse.Error != nil {
		return nil, errors.Join(core.InvalidEmailError, databaseResponse.Error)
	}

	passwordErr := self.Hashing.IsValidPassword(params.Password, user.PasswordHash)
	if passwordErr != nil {
		return nil, core.InvalidPasswordError
	}

	token, tokenErr := self.Serilization.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
	})
	if tokenErr != nil {
		return nil, tokenErr
	}

	refreshToken, refreshErr := self.Serilization.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
	})
	if refreshErr != nil {
		return nil, refreshErr
	}

  updateTokenDatabaseResponse := self.Database.Model(&entities.User{}).Where("id = ?", user.ID).Update("refresh_token", refreshToken)
  if updateTokenDatabaseResponse.Error != nil {
    return nil, core.UnableToPersistToken
  }

	response := SignInResponse{
		Token:        *token,
		RefreshToken: *refreshToken,
	}

	return &response, nil
}

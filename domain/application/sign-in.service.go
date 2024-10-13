package application

import (
	"main/core"
	"main/domain/entities"
	"main/utils/libs"
	"time"

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
	Database *gorm.DB
	Validate *validator.Validate
	Parser   libs.JWT
	Hasher   libs.Hashing
}

func (self SignInService) Execute(params SignInRequest) (*SignInResponse, error) {
	validationErr := self.Validate.Struct(params)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var user entities.User
	query := self.Database.Model(&entities.User{}).Where("email = ?", params.Email)

	databaseResponse := query.First(&user)
	if databaseResponse.Error != nil {
		return nil, core.InvalidEmailError
	}

	passwordErr := self.Hasher.IsValidPassword(params.Password, user.PasswordHash)
	if passwordErr != nil {
		return nil, core.InvalidPasswordError
	}

	token, err := self.Parser.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
		Time: time.Now().Add(time.Hour * 2).Unix(),
		Data: map[string]any{
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
	if err != nil {
		return nil, err
	}

	refreshToken, err := self.Parser.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
		Time: time.Now().Add(time.Hour * 24).Unix(),
	})
	if err != nil {
		return nil, err
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

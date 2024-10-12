package application

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"main/core"
	"main/domain/entities"
	"main/utils/libs"
)

type SignUpService struct {
	Database  *gorm.DB
	Validator *validator.Validate
	Parser    libs.JWT
	Hasher    libs.Hashing
}

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,gte=5,lte=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password_hash" validate:"required,gte=6,lte=36"`

	StatusID uint `json:"status_id" validate:"required,min=1"`
	RoleID   uint `json:"role_id" validate:"required,min=1"`
}

type SignUpResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`

	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`

	RoleID uint `json:"role_id"`
}

func (self SignUpService) Execute(request SignUpRequest) (*SignUpResponse, error) {
	var response SignUpResponse

	err := self.Validator.Struct(request)
	if err != nil {
		return nil, errors.Join(core.InvalidParametersError, err)
	}

	var existingUser *entities.User

	query := self.Database.Find(&entities.User{}).Where("email = ?", request.Email).First(&existingUser)
	if query.Error != nil && !errors.Is(query.Error, gorm.ErrRecordNotFound) {
		return nil, errors.Join(core.EntityAlreadyExistsError, query.Error)
	}

	hashedPassword, err := self.Hasher.EncryptPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := entities.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: *hashedPassword,

		RoleID:   1,
		StatusID: 1,
	}

	userTransaction := self.Database.Create(&user)
	if userTransaction.Error != nil {
		return nil, userTransaction.Error
	}

	token, err := self.Parser.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
		Data: map[string]any{
			"email": user.Email,
			"name":  user.Name,
		},
	})
	if err != nil {
		return nil, errors.Join(core.UnableToPersistTokenButEntityCreated, err)
	}

	refreshToken, err := self.Parser.Generate(libs.CreateJWTParams{
		Sub:  user.ID,
		Role: user.RoleID,
	})
	if err != nil {
		return nil, errors.Join(core.UnableToPersistTokenButEntityCreated, err)
	}

	response.Token = *token
	response.RefreshToken = *refreshToken

	return &response, nil
}

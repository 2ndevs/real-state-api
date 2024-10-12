package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type SignUpPresenter struct{}

type SignUpFromHTTP struct {
	Name     string `json:"name" validate:"required,gte=3,lte=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=36"`
}

type SignUpToHTTP struct {
	ID uint `json:"id"`

	Name         string `json:"name"`
	Email        string `json:"email"`
	RefreshToken string `json:"refresh_token"`

	RoleID   uint `json:"role_id"`
	StatusID uint `json:"status_id"`
}

func (SignUpPresenter) FromHTTP(request *http.Request) (*SignUpFromHTTP, error) {
	var signUpRequest SignUpFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&signUpRequest)
	if err != nil {
		return nil, err
	}

	return &signUpRequest, nil
}

func (SignUpPresenter) ToHTTP(user *entities.User) SignUpToHTTP {
	return SignUpToHTTP{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		RefreshToken: user.RefreshToken,
		StatusID:     user.StatusID,
		RoleID:       user.RoleID,
	}
}

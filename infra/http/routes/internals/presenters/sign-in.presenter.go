package presenters

import (
	"encoding/json"
	"net/http"
)

type SignInPresenter struct{}

type SignInFromHTTP struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6,lte=36"`
}

type SignInToHTTP struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (SignInPresenter) FromHTTP(request *http.Request) (*SignInFromHTTP, error) {
	var signInRequest SignInFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&signInRequest)
	if err != nil {
		return nil, err
	}

	return &signInRequest, nil
}

func (SignInPresenter) ToHTTP(token, refreshToken string) SignInToHTTP {
	return SignInToHTTP{
		Token:        token,
		RefreshToken: refreshToken,
	}
}

package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type UserPresenter struct{}

type UserFromHTTP struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     string  `json:"name" validate:"required,gte=3,lte=25"`
	Password *string `json:"password" validate:"gte=6,lte=36"`

	RoleId   uint `json:"role_id" validate:"required,gte=1"`
	StatusId uint `json:"status_id" validate:"required,gte=1"`
}

type UserToHTTP struct {
	ID uint `json:"id"`

	Name  string `json:"name"`
	Email string `json:"email"`

	Role   *RoleToHTTP   `json:"role"`
	Status *StatusToHTTP `json:"status"`

	RoleId   uint `json:"role_id"`
	StatusId uint `json:"status_id"`
}

func (UserPresenter) FromHTTP(request *http.Request) (*UserFromHTTP, error) {
	var userRequest UserFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&userRequest)
	if err != nil {
		return nil, err
	}

	return &userRequest, nil
}

func (UserPresenter) ToHTTP(user *entities.User) UserToHTTP {
	return UserToHTTP{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role: &RoleToHTTP{
			ID:          user.Role.ID,
			Name:        user.Role.Name,
			Permissions: user.Role.Permissions,
			StatusID:    user.Role.StatusID,
		},
		Status: &StatusToHTTP{
			ID:   user.Status.ID,
			Name: user.Status.Name,
		},
		RoleId:   user.RoleID,
		StatusId: user.StatusID,
	}
}

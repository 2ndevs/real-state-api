package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type RolePresenter struct{}

type RoleFromHTTP struct {
	Name        string   `json:"name" validate:"required,gte=3,lte=20"`
	Permissions []string `json:"permissions" validate:"required"`
}

type RoleToHTTP struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	StatusID    uint     `json:"status_id"`
}

func (RolePresenter) FromHTTP(request *http.Request) (*RoleFromHTTP, error) {
	var roleRequest RoleFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&roleRequest)
	if err != nil {
		return nil, err
	}

	return &roleRequest, nil
}

func (RolePresenter) ToHTTP(role entities.Role) RoleToHTTP {
	return RoleToHTTP{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: role.Permissions,
		StatusID:    role.StatusID,
	}
}

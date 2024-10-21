package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type StatusPresenter struct{}

type StatusFromHTTP struct {
	Name string `json:"name" validate:"required,gte=3,lte=100"`
}

type StatusToHTTP struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
}

func (StatusPresenter) FromHTTP(request *http.Request) (*StatusFromHTTP, error) {
	var statusRequest StatusFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&statusRequest)
	if err != nil {
		return nil, err
	}

	return &statusRequest, nil
}

func (StatusPresenter) ToHTTP(status entities.Status) StatusToHTTP {
	return StatusToHTTP{
		ID:       status.ID,
		Name:     status.Name,
	}
}

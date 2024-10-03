package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type KindPresenter struct{}

type kindFromHTTP struct {
	Name string `json:"name"`
}

type kindToHTTP struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func (KindPresenter) FromHTTP(request *http.Request) (*kindFromHTTP, error) {
	var kindRequest kindFromHTTP

	err := json.NewDecoder(request.Body).Decode(&kindRequest)
	if err != nil {
		return nil, err
	}

	return &kindRequest, nil
}

func (KindPresenter) ToHTTP(kind entities.Kind) kindToHTTP {
	return kindToHTTP{
		ID:       kind.ID,
		Name:     kind.Name,
		StatusID: kind.StatusID,
	}
}

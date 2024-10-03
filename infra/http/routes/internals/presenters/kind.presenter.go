package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type KindPresenter struct{}

type KindFromHTTP struct {
	Name string `json:"name"`
}

type KindToHTTP struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func (KindPresenter) FromHTTP(request *http.Request) (*KindFromHTTP, error) {
	var kindRequest KindFromHTTP

	err := json.NewDecoder(request.Body).Decode(&kindRequest)
	if err != nil {
		return nil, err
	}

	return &kindRequest, nil
}

func (KindPresenter) ToHTTP(kind entities.Kind) KindToHTTP {
	return KindToHTTP{
		ID:       kind.ID,
		Name:     kind.Name,
		StatusID: kind.StatusID,
	}
}

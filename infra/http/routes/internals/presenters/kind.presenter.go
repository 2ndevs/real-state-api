package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type KindPresenter struct{}

type KindFromHTTP struct {
	Name     string `json:"name" validate:"required,gte=3,lte=100"`
	StatusID *uint  `json:"status_id" validate:"gte=1,lte=2"`
}

type KindToHTTP struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	StatusID uint            `json:"status_id"`
	Status   entities.Status `json:"status"`
}

func (KindPresenter) FromHTTP(request *http.Request) (*KindFromHTTP, error) {
	var kindRequest KindFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&kindRequest)
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
		Status:   kind.Status,
	}
}

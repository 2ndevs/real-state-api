package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type NegotiationTypePresenter struct{}

type NegotiationTypeFromHTTP struct {
	Name     string `json:"name" validate:"required,gte=3,lte=100"`
	StatusID *uint  `json:"status_id" validate:"gte=1,lte=2"`
}

type NegotiationTypeToHTTP struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	StatusID uint            `json:"status_id"`
	Status   entities.Status `json:"status"`
}

func (NegotiationTypePresenter) FromHTTP(request *http.Request) (*NegotiationTypeFromHTTP, error) {
	var negotiationTypeRequest NegotiationTypeFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&negotiationTypeRequest)
	if err != nil {
		return nil, err
	}

	return &negotiationTypeRequest, nil
}

func (NegotiationTypePresenter) ToHTTP(negotiationType entities.NegotiationType) NegotiationTypeToHTTP {
	return NegotiationTypeToHTTP{
		ID:       negotiationType.ID,
		Name:     negotiationType.Name,
		Status:   negotiationType.Status,
		StatusID: negotiationType.StatusID,
	}
}

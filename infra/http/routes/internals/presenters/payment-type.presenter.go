package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type PaymentTypePresenter struct{}

type PaymentTypeFromHTTP struct {
	Name string `json:"name" validate:"required,gte=3,lte=100"`
}

type PaymentTypeToHTTP struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func (PaymentTypePresenter) FromHTTP(request *http.Request) (*PaymentTypeFromHTTP, error) {
	var paymentTypeRequest PaymentTypeFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&paymentTypeRequest)
	if err != nil {
		return nil, err
	}

	return &paymentTypeRequest, nil
}

func (PaymentTypePresenter) ToHTTP(paymentType entities.PaymentType) PaymentTypeToHTTP {
	return PaymentTypeToHTTP{
		ID:       paymentType.ID,
		Name:     paymentType.Name,
		StatusID: paymentType.StatusID,
	}
}

package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type MeasurementUnitPresenter struct{}

type MeasurementUnitFromHTTP struct {
	Name         string `json:"name" validate:"required,gte=3,lte=100"`
	Abbreviation string `json:"abbreviation" validate:"required,gte=2"`
}

type MeasurementUnitToHTTP struct {
	ID uint `json:"id"`

	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`

	StatusID uint             `json:"status_id"`
	Status   *entities.Status `json:"status"`
}

func (self MeasurementUnitPresenter) FromHTTP(request *http.Request) (*MeasurementUnitFromHTTP, error) {
	var measurementUnit MeasurementUnitFromHTTP

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&measurementUnit)
	if err != nil {
		return nil, err
	}

	return &measurementUnit, nil
}

func (self MeasurementUnitPresenter) ToHTTP(item entities.UnitOfMeasurement) MeasurementUnitToHTTP {
	return MeasurementUnitToHTTP{
		ID: item.ID,

		Name:         item.Name,
		Abbreviation: item.Abbreviation,

		Status:   item.Status,
		StatusID: item.StatusID,
	}
}

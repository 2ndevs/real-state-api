package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateMeasurementUnit(write http.ResponseWriter, request *http.Request) {
	presenter := presenters.MeasurementUnitPresenter{}
	rawId := chi.URLParam(request, "id")

	id, err := strconv.ParseUint(rawId, 10, 32)
	if err != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	data, err := presenter.FromHTTP(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	validator, err := middlewares.GetValidator(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	service := application.UpdateMeasurementUnitService {
		Database: database,
		Validator: validator, 
	}

	payload := entities.UnitOfMeasurement {
		Name: data.Name,
		Abbreviation: data.Abbreviation,
		StatusID: *data.StatusID,
	}

	response, err := service.Execute(uint(id), payload)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	entity := presenter.ToHTTP(*response)

	encoder := json.NewEncoder(write)
	err = encoder.Encode(entity)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	write.WriteHeader(http.StatusOK)
}

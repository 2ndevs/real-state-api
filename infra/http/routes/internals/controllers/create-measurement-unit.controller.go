package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func CreateMeasurementUnit(write http.ResponseWriter, request *http.Request) {
	presenter := presenters.MeasurementUnitPresenter{}

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

	body, err := presenter.FromHTTP(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	service := application.CreateMeasurementUnitService{
		Validator: validator,
		Database:  database,
	}

	entity := entities.UnitOfMeasurement{
		Name:         body.Name,
		Abbreviation: body.Abbreviation,
		StatusID:     *body.StatusID,
	}

	if body.StatusID == nil {
		entity.StatusID = 1
	}

	data, err := service.Execute(entity)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	response := presenter.ToHTTP(*data)

	encoder := json.NewEncoder(write)
	err = encoder.Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}
}

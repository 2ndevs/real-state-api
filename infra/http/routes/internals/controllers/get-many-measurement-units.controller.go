package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyMeasurementUnits(write http.ResponseWriter, request *http.Request) {
	presenter := presenters.MeasurementUnitPresenter{}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	service := application.GetManyMeasurementUnitService{
		Database: database,
	}

	data, err := service.Execute()
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	var response []presenters.MeasurementUnitToHTTP

	for _, item := range data {
		response = append(response, presenter.ToHTTP(item))
	}

	encoder := json.NewEncoder(write)
	err = encoder.Encode(response)
	
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}
}

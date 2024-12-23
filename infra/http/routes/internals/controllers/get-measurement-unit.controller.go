package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetMeasurementUnit(write http.ResponseWriter, request *http.Request) {
	presenter := presenters.MeasurementUnitPresenter {}
	rawId := chi.URLParam(request, "id")

	id, err := strconv.ParseUint(rawId, 10, 32)
	if err != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}
	
	service := application.GetMeasurementUnitService {
		Database: database,
	}

	data, err := service.Execute(uint(id))
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

	write.WriteHeader(http.StatusOK)
}

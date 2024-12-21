package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteMeasurementUnit(write http.ResponseWriter, request *http.Request) {
	rawId := chi.URLParam(request, "id")
	if len(rawId) == 0 {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	service := application.DeleteMeasurementUnitService{
		Database: database,
	}

	id, err := strconv.ParseUint(rawId, 10, 32)
	if err != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	response, err := service.Execute(uint(id))
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	encoder := json.NewEncoder(write)
	err = encoder.Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	write.WriteHeader(http.StatusOK)
}

package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"
)

func VisitsByFeature(write http.ResponseWriter, request *http.Request) {
	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	feature := request.URL.Query().Get("feature")
	value, err := strconv.Atoi(request.URL.Query().Get("value"))
	if err != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	service := application.VisitsByFeatureService{Database: database, Feature: feature, Value: value}
	result, err := service.Execute()
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	presenter := presenters.VisitsByFeaturePresenter{}
	response := presenter.ToHTTP(result)

	write.WriteHeader(http.StatusOK)
	error := json.NewEncoder(write).Encode(response)

	if error != nil {
		core.HandleHTTPStatus(write, err)
	}
}

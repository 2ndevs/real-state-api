package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyProperties(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	propertiesService := application.GetManyPropertiesService{Database: database}

	filters := httpPresenter.GetSearchParams(request)

	properties, getPropertiesErr := propertiesService.Execute(filters)
	if getPropertiesErr != nil {
		core.HandleHTTPStatus(write, getPropertiesErr)
		return
	}

	var response []presenters.PropertyToHTTP

	for _, property := range *properties {
		response = append(response, httpPresenter.ToHTTP(property))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyProperties(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	searchFilter := request.URL.Query().Get("search")
	propertiesService := application.GetManyPropertiesService{SearchFilter: &searchFilter, Database: database}

	properties, getPropertiesErr := propertiesService.Execute()
	if getPropertiesErr != nil {
		http.Error(write, getPropertiesErr.Error(), http.StatusInternalServerError)
		return
	}

	var response []presenters.PropertyToHTTP

	for _, property := range *properties {
		response = append(response, httpPresenter.ToHTTP(property))
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Erro no servidor", http.StatusInternalServerError)
	}
}

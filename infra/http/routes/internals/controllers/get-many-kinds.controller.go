package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyKinds(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	nameFilter := request.URL.Query().Get("name")
	kindService := application.GetManyKindsService{NameFilter: &nameFilter, Database: database}

	kinds, getKindsErr := kindService.Execute()
	if getKindsErr != nil {
		core.HandleHTTPStatus(write, getKindsErr)
		return
	}

	var response []presenters.KindToHTTP

	for _, kind := range *kinds {
		response = append(response, httpPresenter.ToHTTP(kind))
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

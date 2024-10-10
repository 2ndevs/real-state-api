package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyKinds(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	nameFilter := request.URL.Query().Get("name")
	kindService := application.GetManyKindsService{NameFilter: &nameFilter, Database: database}

	kinds, getKindsErr := kindService.Execute()
	if getKindsErr != nil {
		http.Error(write, getKindsErr.Error(), http.StatusInternalServerError)
		return
	}

	var response []presenters.KindToHTTP

	for _, kind := range *kinds {
		response = append(response, httpPresenter.ToHTTP(kind))
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Erro no servidor", http.StatusInternalServerError)
	}
}

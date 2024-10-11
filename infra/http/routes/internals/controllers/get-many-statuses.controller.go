package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyStatuses(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.StatusPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}
	nameFilter := request.URL.Query().Get("name")
	statusesService := application.GetManyStatusesService{NameFilter: &nameFilter, Database: database}

	statuses, getStatusesErr := statusesService.Execute()
	if getStatusesErr != nil {
		http.Error(write, getStatusesErr.Error(), http.StatusInternalServerError)
		return
	}

	var response []presenters.StatusToHTTP

	for _, status := range *statuses {
		response = append(response, httpPresenter.ToHTTP(status))
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetStatus(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.StatusPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	statusService := application.GetStatusService{Request: request, Database: database}

	status, getStatusErr := statusService.Execute()
	if getStatusErr != nil {
		http.Error(write, getStatusErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*status)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

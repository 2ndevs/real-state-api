package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func CreateStatus(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.StatusPresenter{}

	statusRequest, parseError := httpPresenter.FromHTTP(request)
	if parseError != nil {
		http.Error(write, parseError.Error(), http.StatusBadRequest)
		return
	}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	validated, ctxErr := middlewares.GetValidator(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusBadRequest)
		return
	}

	statusService := application.CreateStatusService{Validated: validated, Database: database}
	statusPayload := entities.Status{
		Name: statusRequest.Name,
	}

	status, createStatusErr := statusService.Execute(statusPayload)
	if createStatusErr != nil {
		http.Error(write, createStatusErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*status)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

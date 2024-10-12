package controllers

import (
	"encoding/json"
	"main/core"
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
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	validated, ctxErr := middlewares.GetValidator(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	statusService := application.CreateStatusService{Validated: validated, Database: database}
	statusPayload := entities.Status{
		Name: statusRequest.Name,
	}

	status, createStatusErr := statusService.Execute(statusPayload)
	if createStatusErr != nil {
		core.HandleHTTPStatus(write, createStatusErr)
		return
	}

	response := httpPresenter.ToHTTP(*status)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

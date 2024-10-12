package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateStatus(write http.ResponseWriter, request *http.Request) {
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
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	statusId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	statusService := application.UpdateStatusService{Validated: validated, Database: database}
	statusPayload := entities.Status{
		Name: statusRequest.Name,
	}

	status, updateStatusErr := statusService.Execute(statusPayload, statusId)
	if updateStatusErr != nil {
		core.HandleHTTPStatus(write, updateStatusErr)
		return
	}

	response := httpPresenter.ToHTTP(*status)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

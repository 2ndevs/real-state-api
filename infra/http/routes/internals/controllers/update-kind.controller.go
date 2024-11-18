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

func UpdateKind(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}

	kindRequest, parseError := httpPresenter.FromHTTP(request)
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
	kindId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	if kindRequest.StatusID == nil {
		defaultStatusValue := uint(1)
		kindRequest.StatusID = &defaultStatusValue
	}

	kindService := application.UpdateKindService{Validated: validated, Database: database}
	kindPayload := entities.Kind{
		Name:     kindRequest.Name,
		StatusID: *kindRequest.StatusID,
	}

	kind, updateKindErr := kindService.Execute(kindPayload, kindId)
	if updateKindErr != nil {
		core.HandleHTTPStatus(write, updateKindErr)
		return
	}

	response := httpPresenter.ToHTTP(*kind)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteNegotiationType(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.NegotiationTypePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	negotiationTypeId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	negotiationTypeService := application.DeleteNegotiationTypeService{Database: database}

	negotiationType, deleteNegotiationTypeErr := negotiationTypeService.Execute(negotiationTypeId)
	if deleteNegotiationTypeErr != nil {
		core.HandleHTTPStatus(write, deleteNegotiationTypeErr)
		return
	}

	response := httpPresenter.ToHTTP(*negotiationType)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

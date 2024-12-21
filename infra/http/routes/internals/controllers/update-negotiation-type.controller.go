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

func UpdateNegotiationType(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.NegotiationTypePresenter{}

	negotiationTypeRequest, parseError := httpPresenter.FromHTTP(request)
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
	negotiationTypeId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	if negotiationTypeRequest.StatusID == nil {
		defaultStatusValue := uint(1)
		negotiationTypeRequest.StatusID = &defaultStatusValue
	}

	negotiationTypeService := application.UpdateNegotiationTypeService{Validated: validated, Database: database}
	negotiationtypePayload := entities.NegotiationType{
		Name:     negotiationTypeRequest.Name,
		StatusID: *negotiationTypeRequest.StatusID,
	}

	negotiationType, updateNegotiationTypeErr := negotiationTypeService.Execute(negotiationtypePayload, negotiationTypeId)
	if updateNegotiationTypeErr != nil {
		core.HandleHTTPStatus(write, updateNegotiationTypeErr)
		return
	}

	response := httpPresenter.ToHTTP(*negotiationType)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

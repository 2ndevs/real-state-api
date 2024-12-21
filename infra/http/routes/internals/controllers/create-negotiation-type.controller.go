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

func CreateNegotiationType(write http.ResponseWriter, request *http.Request) {
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
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	if negotiationTypeRequest.StatusID == nil {
		defaultStatusValue := uint(1)
		negotiationTypeRequest.StatusID = &defaultStatusValue
	}

	negotiationtypeService := application.CreateNegotiationTypeService{Validated: validated, Database: database}
	negotiationtypePayload := entities.NegotiationType{
		Name:     negotiationTypeRequest.Name,
		StatusID: *negotiationTypeRequest.StatusID,
	}

	negotiationtype, createNegotiationTypeErr := negotiationtypeService.Execute(negotiationtypePayload)
	if createNegotiationTypeErr != nil {
		core.HandleHTTPStatus(write, createNegotiationTypeErr)
		return
	}

	response := httpPresenter.ToHTTP(*negotiationtype)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

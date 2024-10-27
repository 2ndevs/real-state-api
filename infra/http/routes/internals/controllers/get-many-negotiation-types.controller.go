package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyNegotiationTypes(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.NegotiationTypePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	nameFilter := request.URL.Query().Get("name")
	negotiationTypeService := application.GetManyNegotiationTypesService{NameFilter: &nameFilter, Database: database}

	negotiationtypes, getNegotiationTypesErr := negotiationTypeService.Execute()
	if getNegotiationTypesErr != nil {
		core.HandleHTTPStatus(write, getNegotiationTypesErr)
		return
	}

	var response []presenters.NegotiationTypeToHTTP

	for _, negotiationtype := range *negotiationtypes {
		response = append(response, httpPresenter.ToHTTP(negotiationtype))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

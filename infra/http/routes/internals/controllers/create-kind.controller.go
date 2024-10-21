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

func CreateKind(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}
	//
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
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	kindService := application.CreateKindService{Validated: validated, Database: database}
	kindPayload := entities.Kind{
		Name:     kindRequest.Name,
		StatusID: 1,
	}

	kind, createKindErr := kindService.Execute(kindPayload)
	if createKindErr != nil {
		core.HandleHTTPStatus(write, createKindErr)
		return
	}

	response := httpPresenter.ToHTTP(*kind)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

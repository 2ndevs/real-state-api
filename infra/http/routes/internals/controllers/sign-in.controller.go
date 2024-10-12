package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"main/utils/libs"
	"net/http"
)

func SignIn(writer http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.SignInPresenter {}

	body, err := httpPresenter.FromHTTP(request)
	if err != nil {
		core.HandleHTTPStatus(writer, core.InvalidParametersError)
		return
	}

	validator, err := middlewares.GetValidator(request)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return 
	}

	service := application.SignInService {
		Database: database,
		Validate: validator,
		Parser: libs.JWT{},
		Hasher: libs.Hashing{},	
	}
	payload := application.SignInRequest {
		Email: body.Email,
		Password: body.Password,
	}

	response, err := service.Execute(payload)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	encondeErr := json.NewEncoder(writer).Encode(response)
	if encondeErr != nil {
		core.HandleHTTPStatus(writer, encondeErr)
	}
}

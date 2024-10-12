package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func SignUp(writer http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.SignUpPresenter{}

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

	service := application.SignUpService {
		Validator: validator,
		Database: database,
	}
	payload := application.SignUpRequest {
		Email: body.Email,
		Name: body.Name,
		Password: body.Password,
		StatusID: 1,
		RoleID: 1,
	}

	response, err := service.Execute(payload)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}
	
	writer.WriteHeader(http.StatusCreated)

	encondeErr := json.NewEncoder(writer).Encode(response)
	if encondeErr != nil {
		core.HandleHTTPStatus(writer, encondeErr)
	} 
}

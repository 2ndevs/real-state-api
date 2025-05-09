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

func CreateRole(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.RolePresenter{}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	parser, err := middlewares.GetValidator(request)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	body, err := httpPresenter.FromHTTP(request)
	if err != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	createRoleService := application.CreateRoleService{
		Validate: parser,
		Database: database,
	}

	if body.StatusID == nil {
		defaultStatusValue := uint(1)
		body.StatusID = &defaultStatusValue
	}

	payload := entities.Role{
		Name:        body.Name,
		Permissions: body.Permissions,
		StatusID:    *body.StatusID,
	}

	response, err := createRoleService.Execute(payload)
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	write.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyRoles(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.RolePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	nameFilter := request.URL.Query().Get("name")
	roleService := application.GetManyRolesService{NameFilter: &nameFilter, Database: database}

	roles, getRolesErr := roleService.Execute()
	if getRolesErr != nil {
		core.HandleHTTPStatus(write, getRolesErr)
		return
	}

	var response []presenters.RoleToHTTP

	for _, role := range *roles {
		response = append(response, httpPresenter.ToHTTP(role))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

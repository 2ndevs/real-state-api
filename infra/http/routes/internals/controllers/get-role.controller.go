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

func GetRole(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.RolePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	roleId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	roleService := application.GetRoleService{RoleID: roleId, Database: database}

	role, getRoleErr := roleService.Execute()
	if getRoleErr != nil {
		core.HandleHTTPStatus(write, getRoleErr)
		return
	}

	response := httpPresenter.ToHTTP(*role)

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

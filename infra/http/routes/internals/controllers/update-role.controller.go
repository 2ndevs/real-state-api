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

func UpdateRole(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.RolePresenter{}

	roleRequest, parseError := httpPresenter.FromHTTP(request)
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
	roleId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	if roleRequest.StatusID == nil {
		defaultStatusValue := uint(1)
		roleRequest.StatusID = &defaultStatusValue
	}

	roleService := application.UpdateRoleService{Validated: validated, Database: database}
	rolePayload := entities.Role{
		Name:        roleRequest.Name,
		Permissions: roleRequest.Permissions,
		StatusID:    *roleRequest.StatusID,
	}

	role, updateRoleErr := roleService.Execute(rolePayload, roleId)
	if updateRoleErr != nil {
		core.HandleHTTPStatus(write, updateRoleErr)
		return
	}

	response := httpPresenter.ToHTTP(*role)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

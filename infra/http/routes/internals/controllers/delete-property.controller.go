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

func DeleteProperty(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)

		return
	}

	idParam := chi.URLParam(request, "id")
	propertyId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	propertyService := application.DeletePropertyService{Database: database}

	property, deletePropertyErr := propertyService.Execute(propertyId)
	if deletePropertyErr != nil {
		core.HandleHTTPStatus(write, deletePropertyErr)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

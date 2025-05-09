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

func GetProperty(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	propertyId, validationErr := strconv.Atoi(idParam)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	identity, _ := httpPresenter.GetIdentity(request)
	propertyService := application.GetPropertyService{Database: database}

	property, getPropertyErr := propertyService.Execute(uint64(propertyId), identity)
	if getPropertyErr != nil {
		core.HandleHTTPStatus(write, getPropertyErr)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

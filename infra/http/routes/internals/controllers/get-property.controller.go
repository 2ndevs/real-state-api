package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"
)

func GetProperty(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	idParam := request.URL.Path[len("/properties/"):]

	propertyId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "invalid id", http.StatusBadRequest)
		return
	}

	propertyService := application.GetPropertyService{PropertyID: propertyId, Database: database}

	property, getPropertyErr := propertyService.Execute()
	if getPropertyErr != nil {
		http.Error(write, getPropertyErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

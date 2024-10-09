package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func CreateProperty(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	propertyRequest, parseError := httpPresenter.FromHTTP(request)
	if parseError != nil {
		http.Error(write, parseError.Error(), http.StatusBadRequest)
		return
	}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	validated, ctxErr := middlewares.GetValidator(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusBadRequest)
		return
	}

	propertyService := application.CreatePropertyService{Validated: validated, Database: database}
	propertyPayload := entities.Property{
		Size:      propertyRequest.Size,
		Rooms:     propertyRequest.Rooms,
		Kitchens:  propertyRequest.Kitchens,
		Bathrooms: propertyRequest.Bathrooms,
		Address:   propertyRequest.Address,
		Summary:   propertyRequest.Summary,
		Details:   propertyRequest.Details,
		Latitude:  propertyRequest.Latitude,
		Longitude: propertyRequest.Longitude,
		Price:     propertyRequest.Price,

		KindID:        propertyRequest.KindID,
		PaymentTypeID: propertyRequest.PaymentTypeID,
		StatusID:      1,
	}

	property, createPropertyErr := propertyService.Execute(propertyPayload)
	if createPropertyErr != nil {
		http.Error(write, createPropertyErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

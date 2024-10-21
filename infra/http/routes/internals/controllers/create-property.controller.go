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

func CreateProperty(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PropertyPresenter{}

	propertyRequest, parseError := httpPresenter.FromHTTP(request)
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
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	propertyService := application.CreatePropertyService{Validated: validated, Database: database}
	propertyPayload := entities.Property{
		Size:             propertyRequest.Size,
		Rooms:            propertyRequest.Rooms,
		Kitchens:         propertyRequest.Kitchens,
		Bathrooms:        propertyRequest.Bathrooms,
		Address:          propertyRequest.Address,
		Summary:          propertyRequest.Summary,
		Details:          propertyRequest.Details,
		Latitude:         propertyRequest.Latitude,
		Longitude:        propertyRequest.Longitude,
		Price:            propertyRequest.Price,
		IsHighlight:      propertyRequest.IsHighlight,
		Discount:         propertyRequest.Discount,
		ConstructionYear: propertyRequest.ConstructionYear,
		IsSold:           propertyRequest.IsSold,

		KindID:            propertyRequest.KindID,
		PaymentTypeID:     propertyRequest.PaymentTypeID,
		NegotiationTypeId: propertyRequest.NegotiationTypeID,
		StatusID:          1,
	}

	property, createPropertyErr := propertyService.Execute(propertyPayload)
	if createPropertyErr != nil {
		core.HandleHTTPStatus(write, createPropertyErr)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

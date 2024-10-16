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

func UpdateProperty(write http.ResponseWriter, request *http.Request) {
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
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	propertyId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	propertyService := application.UpdatePropertyService{Validated: validated, Database: database}
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

		KindID:        propertyRequest.KindID,
		PaymentTypeID: propertyRequest.PaymentTypeID,
		StatusID:      1,
	}

	property, updatePropertyErr := propertyService.Execute(propertyPayload, propertyId)
	if updatePropertyErr != nil {
		core.HandleHTTPStatus(write, updatePropertyErr)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

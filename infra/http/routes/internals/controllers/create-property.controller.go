package controllers

import (
	"encoding/json"
	"log"
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
		log.Println(parseError)
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

	bucket, ctxErr := middlewares.GetBucketContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	propertyService := application.CreatePropertyService{Validated: validated, Database: database, Bucket: bucket}
	propertyPayload := application.CreatePropertyServiceRequest{
		Property: entities.Property{
			BuiltArea:        propertyRequest.BuiltArea,
			TotalArea:        propertyRequest.TotalArea,
			Rooms:            propertyRequest.Rooms,
			Suites:           propertyRequest.Suites,
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
			ContactNumber:    propertyRequest.ContactNumber,

			KindID:              propertyRequest.KindID,
			PaymentTypeID:       propertyRequest.PaymentTypeID,
			UnitOfMeasurementID: propertyRequest.UnitOfMeasurementID,
			StatusID:            1,
		},

		PreviewImages: propertyRequest.PreviewImages,
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

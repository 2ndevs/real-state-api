package controllers

import (
	"encoding/json"
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

	idParam := chi.URLParam(request, "id")
	propertyId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	propertyService := application.UpdatePropertyService{Validated: validated, PropertyID: propertyId, Database: database}
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

	property, updatePropertyErr := propertyService.Execute(propertyPayload)
	if updatePropertyErr != nil {
		http.Error(write, updatePropertyErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*property)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Erro no servidor", http.StatusInternalServerError)
	}
}

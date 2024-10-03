package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"
)

type CreatePropertyRequest struct {
	Size      uint    `json:"size"`
	Rooms     uint    `json:"rooms"`
	Kitchens  uint    `json:"kitchens"`
	Bathrooms uint    `json:"bathrooms"`
	Address   string  `json:"address"`
	Summary   string  `json:"summary"`
	Details   string  `json:"details"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Price     float64 `json:"price"`

	KindID        uint `json:"kind_id"`
	StatusID      uint `json:"status_id"`
	PaymentTypeID uint `json:"payment_type_id"`
}

type CreatePropertyResponse struct {
	ID uint `json:"id"`

	Size      uint    `json:"size"`
	Rooms     uint    `json:"rooms"`
	Kitchens  uint    `json:"kitchens"`
	Bathrooms uint    `json:"bathrooms"`
	Address   string  `json:"address"`
	Summary   string  `json:"summary"`
	Details   string  `json:"details"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Price     float64 `json:"price"`

	KindID        uint `json:"kind_id"`
	StatusID      uint `json:"status_id"`
	PaymentTypeID uint `json:"payment_type_id"`
}

func CreateProperty(writer http.ResponseWriter, request *http.Request) {
	var propertyRequest CreatePropertyRequest

	parseError := json.NewDecoder(request.Body).Decode(&propertyRequest)

	if parseError != nil {
		http.Error(writer, parseError.Error(), http.StatusBadRequest)
		return
	}

	db, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(writer, "Database connection not found", http.StatusInternalServerError)
		return
	}

	Property := entities.Property{
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
		StatusID:      propertyRequest.StatusID,
		PaymentTypeID: propertyRequest.PaymentTypeID,
	}
	createPropertyError := db.Create(&Property).Error

	if createPropertyError != nil {
		http.Error(writer, "Unable to create Property", http.StatusInternalServerError)
		return
	}

	response := CreatePropertyResponse{
		ID: Property.ID,

		Size:      Property.Size,
		Rooms:     Property.Rooms,
		Kitchens:  Property.Kitchens,
		Bathrooms: Property.Bathrooms,
		Address:   Property.Address,
		Summary:   Property.Summary,
		Details:   Property.Details,
		Latitude:  Property.Latitude,
		Longitude: Property.Longitude,
		Price:     Property.Price,

		KindID:        Property.KindID,
		StatusID:      Property.StatusID,
		PaymentTypeID: Property.PaymentTypeID,
	}

	writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(writer).Encode(response)

	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
	}
}

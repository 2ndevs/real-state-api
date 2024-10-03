package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"
)

type CreatePaymentTypeRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	StatusID uint   `json:"status_id" binding:"required"`
}

type CreatePaymentTypeResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func CreatePaymentType(writer http.ResponseWriter, request *http.Request) {
	var paymentTypeRequest CreatePaymentTypeRequest

	parseError := json.NewDecoder(request.Body).Decode(&paymentTypeRequest)

	if parseError != nil {
		http.Error(writer, parseError.Error(), http.StatusBadRequest)
		return
	}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(writer, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	PaymentType := entities.PaymentType{Name: paymentTypeRequest.Name, StatusID: paymentTypeRequest.StatusID}
	createPaymentTypeError := database.Create(&PaymentType).Error

	if createPaymentTypeError != nil {
		http.Error(writer, "Unable to create PaymentType", http.StatusInternalServerError)
		return
	}

	response := CreatePaymentTypeResponse{
		ID:       PaymentType.ID,
		Name:     PaymentType.Name,
		StatusID: PaymentType.StatusID,
	}

	writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(writer).Encode(response)

	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
	}
}

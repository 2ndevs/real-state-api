package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"
)

func GetPaymentType(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PaymentTypePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	idParam := request.URL.Path[len("/payment-types/"):]

	paymentTypeId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "invalid id", http.StatusBadRequest)
		return
	}

	paymentTypeService := application.GetPaymentTypeService{PaymentTypeID: paymentTypeId, Database: database}

	paymentType, getPaymentTypeErr := paymentTypeService.Execute()
	if getPaymentTypeErr != nil {
		http.Error(write, getPaymentTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*paymentType)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

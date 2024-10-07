package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyPaymentTypes(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PaymentTypePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	paymentTypeService := application.GetManyPaymentTypesService{Request: request, Database: database}

	paymenttypes, getPaymentTypesErr := paymentTypeService.Execute()
	if getPaymentTypesErr != nil {
		http.Error(write, getPaymentTypesErr.Error(), http.StatusInternalServerError)
		return
	}

	var response []presenters.PaymentTypeToHTTP

	for _, paymenttype := range *paymenttypes {
		response = append(response, httpPresenter.ToHTTP(paymenttype))
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

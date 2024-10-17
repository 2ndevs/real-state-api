package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyPaymentTypes(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PaymentTypePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	nameFilter := request.URL.Query().Get("name")
	paymentTypeService := application.GetManyPaymentTypesService{NameFilter: &nameFilter, Database: database}

	paymenttypes, getPaymentTypesErr := paymentTypeService.Execute()
	if getPaymentTypesErr != nil {
		core.HandleHTTPStatus(write, getPaymentTypesErr)
		return
	}

	var response []presenters.PaymentTypeToHTTP

	for _, paymenttype := range *paymenttypes {
		response = append(response, httpPresenter.ToHTTP(paymenttype))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

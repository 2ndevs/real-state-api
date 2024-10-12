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

func CreatePaymentType(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PaymentTypePresenter{}

	paymentTypeRequest, parseError := httpPresenter.FromHTTP(request)
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

	paymenttypeService := application.CreatePaymentTypeService{Validated: validated, Database: database}
	paymenttypePayload := entities.PaymentType{
		Name:     paymentTypeRequest.Name,
		StatusID: 1,
	}

	paymenttype, createPaymentTypeErr := paymenttypeService.Execute(paymenttypePayload)
	if createPaymentTypeErr != nil {
		core.HandleHTTPStatus(write, createPaymentTypeErr)
		return
	}

	response := httpPresenter.ToHTTP(*paymenttype)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

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

func UpdatePaymentType(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PaymentTypePresenter{}

	paymentTypeRequest, parseError := httpPresenter.FromHTTP(request)
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
	paymentTypeId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "ID inválido", http.StatusBadRequest)
		return
	}

	paymentTypeService := application.UpdatePaymentTypeService{Validated: validated, PaymentTypeID: paymentTypeId, Database: database}
	paymenttypePayload := entities.PaymentType{
		Name:     paymentTypeRequest.Name,
		StatusID: 1,
	}

	paymentType, updatePaymentTypeErr := paymentTypeService.Execute(paymenttypePayload)
	if updatePaymentTypeErr != nil {
		http.Error(write, updatePaymentTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*paymentType)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Erro no servidor", http.StatusInternalServerError)
	}
}

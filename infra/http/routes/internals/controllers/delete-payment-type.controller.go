package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeletePaymentType(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.PaymentTypePresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	idParam := chi.URLParam(request, "id")
	paymentTypeId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "invalid ID", http.StatusBadRequest)
		return
	}

	paymentTypeService := application.DeletePaymentTypeService{Database: database}

	paymentType, deletePaymentTypeErr := paymentTypeService.Execute(paymentTypeId)
	if deletePaymentTypeErr != nil {
		http.Error(write, deletePaymentTypeErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*paymentType)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

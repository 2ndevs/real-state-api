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

func DeleteStatus(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.StatusPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	idParam := chi.URLParam(request, "id")
	statusId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "invalid ID", http.StatusBadRequest)
		return
	}

	statusService := application.DeleteStatusService{Database: database}

	status, deleteStatusErr := statusService.Execute(statusId)
	if deleteStatusErr != nil {
		http.Error(write, deleteStatusErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*status)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

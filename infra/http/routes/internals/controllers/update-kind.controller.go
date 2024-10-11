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

func UpdateKind(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}

	kindRequest, parseError := httpPresenter.FromHTTP(request)
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
	kindId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "invalid ID", http.StatusBadRequest)
		return
	}

	kindService := application.UpdateKindService{Validated: validated, KindID: kindId, Database: database}
	kindPayload := entities.Kind{
		Name:     kindRequest.Name,
		StatusID: 1,
	}

	kind, updateKindErr := kindService.Execute(kindPayload)
	if updateKindErr != nil {
		http.Error(write, updateKindErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*kind)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

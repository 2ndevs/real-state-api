package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteKind(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	kindId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	kindService := application.DeleteKindService{Database: database}

	kind, deleteKindErr := kindService.Execute(kindId)
	if deleteKindErr != nil {
		core.HandleHTTPStatus(write, deleteKindErr)
		return
	}

	response := httpPresenter.ToHTTP(*kind)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

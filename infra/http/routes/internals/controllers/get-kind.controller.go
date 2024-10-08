package controllers

import (
	"encoding/json"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"
)

func GetKind(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.KindPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		http.Error(write, ctxErr.Error(), http.StatusInternalServerError)
		return
	}

	idParam := request.URL.Path[len("/kinds/"):]
	kindId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		http.Error(write, "invalid id", http.StatusBadRequest)
		return
	}

	kindService := application.GetKindService{Request: request, KindID: kindId, Database: database}

	kind, getKindErr := kindService.Execute()
	if getKindErr != nil {
		http.Error(write, getKindErr.Error(), http.StatusInternalServerError)
		return
	}

	response := httpPresenter.ToHTTP(*kind)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

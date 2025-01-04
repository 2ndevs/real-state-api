package controllers

import (
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if len(id) == 0 {
		core.HandleHTTPStatus(writer, core.InvalidParametersError)
		return
	}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	service := application.DeleteUserService{
		Database: database,
	}

	err = service.Execute(id)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

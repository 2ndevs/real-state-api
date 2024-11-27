package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"

	"net/http"
)

func GetManyUsers(writer http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.UserPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(writer, ctxErr)
		return
	}

	service := application.GetManyUsersService{Database: database}

	users, serviceErr := service.Execute()
	if serviceErr != nil {
		core.HandleHTTPStatus(writer, serviceErr)
		return
	}

	var response []presenters.UserToHTTP

	for _, user := range users {
		response = append(
			response,
			httpPresenter.ToHTTP(&user),
		)
	}

	writer.WriteHeader(http.StatusOK)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
	}
}

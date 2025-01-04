package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"main/utils/libs"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	presenter := presenters.UserPresenter{}

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

	validator, err := middlewares.GetValidator(request)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	hasher := libs.Hashing {}

	data, err := presenter.FromHTTP(request)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	payload := application.UpdateUserServiceRequest{
		StatusID: data.StatusId,
		Email:    data.Email,
		Name:     data.Name,
		PasswordHash: nil,
	}

	if len(*data.Password) > 0 {
		payload.PasswordHash = data.Password
	}

	service := application.UpdateUserService{
		Validator: validator,
		Database:  database,
		Hasher: hasher,
	}

	responseData, err := service.Execute(id, payload)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	response := presenter.ToHTTP(responseData)
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

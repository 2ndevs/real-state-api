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

func GetInterestedUser(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.InterestedUsersPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	interestedUserId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	interestedUserService := application.GetInterestedUserService{Database: database}

	interestedUser, getInterestedUserErr := interestedUserService.Execute(uint(interestedUserId))
	if getInterestedUserErr != nil {
		core.HandleHTTPStatus(write, getInterestedUserErr)
		return
	}

	response := httpPresenter.ToHTTP(*interestedUser)

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

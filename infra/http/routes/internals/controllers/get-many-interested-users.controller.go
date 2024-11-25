package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyInterestedUsers(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.InterestedUsersPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	interestedUserService := application.GetManyInterestedUsersService{Database: database}

	interestedUsers, getInterestedUsersErr := interestedUserService.Execute()
	if getInterestedUsersErr != nil {
		core.HandleHTTPStatus(write, getInterestedUsersErr)
		return
	}

	var response []presenters.InterestedUsersToHTTP

	for _, interestedUser := range *interestedUsers {
		response = append(response, httpPresenter.ToHTTP(interestedUser))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

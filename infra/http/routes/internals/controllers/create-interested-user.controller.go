package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func CreateInterestedUser(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.InterestedUsersPresenter{}

	interestedUserRequest, parseError := httpPresenter.FromHTTP(request)
	if parseError != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	validated, ctxErr := middlewares.GetValidator(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	if interestedUserRequest.StatusID == nil {
		defaultStatusValue := uint(1)
		interestedUserRequest.StatusID = &defaultStatusValue
	}

	interestedUserService := application.CreateInterestedUserService{Validated: validated, Database: database}
	interestedUserPayload := entities.InterestedUser{
		FirstName: interestedUserRequest.FirstName,
		LastName:  interestedUserRequest.LastName,
		Phone:     interestedUserRequest.Phone,
		Email:     interestedUserRequest.Email,
		Answered:  interestedUserRequest.Answered,

		StatusID:   interestedUserRequest.StatusID,
		PropertyID: interestedUserRequest.PropertyID,
	}

	interestedUser, createInterestedUserErr := interestedUserService.Execute(interestedUserPayload)
	if createInterestedUserErr != nil {
		core.HandleHTTPStatus(write, createInterestedUserErr)
		return
	}

	response := httpPresenter.ToHTTP(*interestedUser)

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

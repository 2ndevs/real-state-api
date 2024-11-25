package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateInterestedUser(write http.ResponseWriter, request *http.Request) {
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
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	idParam := chi.URLParam(request, "id")
	interestedUserId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		core.HandleHTTPStatus(write, core.InvalidParametersError)
		return
	}

	if interestedUserRequest.StatusID == nil {
		defaultStatusValue := uint(1)
		interestedUserRequest.StatusID = &defaultStatusValue
	}

	interestedUserService := application.UpdateInterestedUserService{Validated: validated, Database: database}
	interestedUserPayload := entities.InterestedUser{
		FirstName: interestedUserRequest.FirstName,
		LastName:  interestedUserRequest.LastName,
		Phone:     interestedUserRequest.Phone,
		Email:     interestedUserRequest.Email,
		Answered:  interestedUserRequest.Answered,

		StatusID:   interestedUserRequest.StatusID,
		PropertyID: interestedUserRequest.PropertyID,
	}

	interestedUser, updateInterestedUserErr := interestedUserService.Execute(interestedUserPayload, interestedUserId)
	if updateInterestedUserErr != nil {
		core.HandleHTTPStatus(write, updateInterestedUserErr)
		return
	}

	response := httpPresenter.ToHTTP(*interestedUser)

	write.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

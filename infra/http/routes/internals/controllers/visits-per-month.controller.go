package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func VisitsPerMonth(write http.ResponseWriter, request *http.Request) {
	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	service := application.VisitsPerMonthService{Database: database}
	result, err := service.Execute()
	if err != nil {
		core.HandleHTTPStatus(write, err)
		return
	}

	presenter := presenters.VisitsPerMonthPresenter{}
	response := presenter.ToHTTP(result)

	write.WriteHeader(http.StatusOK)
	error := json.NewEncoder(write).Encode(response)

	if error != nil {
		core.HandleHTTPStatus(write, err)
	}
}

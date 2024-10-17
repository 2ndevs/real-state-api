package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetManyStatuses(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.StatusPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}
	nameFilter := request.URL.Query().Get("name")
	statusesService := application.GetManyStatusesService{NameFilter: &nameFilter, Database: database}

	statuses, getStatusesErr := statusesService.Execute()
	if getStatusesErr != nil {
		core.HandleHTTPStatus(write, getStatusesErr)
		return
	}

	var response []presenters.StatusToHTTP

	for _, status := range *statuses {
		response = append(response, httpPresenter.ToHTTP(status))
	}

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

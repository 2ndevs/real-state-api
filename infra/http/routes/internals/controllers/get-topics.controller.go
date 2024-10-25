package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func GetTopics(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.TopicPresenter{}

	database, ctxErr := middlewares.GetDatabaseFromContext(request)
	if ctxErr != nil {
		core.HandleHTTPStatus(write, ctxErr)
		return
	}

	TopicsService := application.GetTopicsService{Database: database}

	topics, getTopicsErr := TopicsService.Execute()
	if getTopicsErr != nil {
		core.HandleHTTPStatus(write, getTopicsErr)
		return
	}

	response := httpPresenter.ToHTTP(*topics)

	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(write, err)
	}
}

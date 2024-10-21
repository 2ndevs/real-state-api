package controllers

import (
	"encoding/json"
	"main/core"
	"main/domain/application"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"main/utils/libs"
	"net/http"
)

func RefreshToken(writer http.ResponseWriter, request *http.Request) {
	presenter := presenters.RefreshTokenPresenter{}
	token := request.Header.Get("X-Refresh-Token")

	if len(token) <= 0 {
		core.HandleHTTPStatus(writer, core.MissingRefreshTokenError)
		return
	}

	parser := libs.JWT{}
	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
    return
	}

	service := application.RefreshTokenService {
		Database: database,
		Parser:   parser,
	}
  
  tokens, err := service.Execute(token)
  if err != nil {
    core.HandleHTTPStatus(writer, err)
    return
  }

	response := presenter.ToHTTP(*tokens.Token, *tokens.RefreshToken)

	writer.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		core.HandleHTTPStatus(writer, err)
	}
}

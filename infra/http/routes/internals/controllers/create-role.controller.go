package controllers

import (
	"main/domain/application"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"main/infra/http/routes/internals/presenters"
	"net/http"
)

func CreateRole(write http.ResponseWriter, request *http.Request) {
	httpPresenter := presenters.RolePresenter{}

	database, err := middlewares.GetDatabaseFromContext(request)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

	parser, err := middlewares.GetValidator(request)
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
	}

	body, err := httpPresenter.FromHTTP(request)
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
	}
	
	createRoleService := application.CreateRoleService{
		Validate: parser,
		Database: database,
	}

	payload := entities.Role{
		Name:        body.Name,
		Permissions: body.Permissions,
		StatusID:    1,
	}

	_, err = createRoleService.Execute(payload)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// write.WriteHeader(http.StatusCreated)
	// err = json.NewEncoder(write).Encode(response)

	// if err != nil {
	// 	http.Error(write, err.Error(), http.StatusInternalServerError)
	// }
}

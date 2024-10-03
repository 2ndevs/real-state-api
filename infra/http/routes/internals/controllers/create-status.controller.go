package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"
)

type CreateStatusRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

type CreateStatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateStatus(writer http.ResponseWriter, request *http.Request) {
	var statusRequest CreateStatusRequest

	parseError := json.NewDecoder(request.Body).Decode(&statusRequest)

	if parseError != nil {
		http.Error(writer, parseError.Error(), http.StatusBadRequest)
		return
	}

	db := middlewares.GetDBFromContext(request.Context())
	if db == nil {
		http.Error(writer, "Database connection not found", http.StatusInternalServerError)
		return
	}

	Status := entities.Status{Name: statusRequest.Name}
	createStatusError := db.Create(&Status).Error

	if createStatusError != nil {
		http.Error(writer, "Unable to create Status", http.StatusInternalServerError)
		return
	}

	response := CreateStatusResponse{
		ID:   Status.ID,
		Name: Status.Name,
	}

	writer.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(writer).Encode(response)

	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
	}
}

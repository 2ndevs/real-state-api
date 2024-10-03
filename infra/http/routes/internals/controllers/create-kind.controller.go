package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"
)

type CreateKindRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	StatusID uint   `json:"status_id" binding:"required"`
}

type CreateKindResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func CreateKind(write http.ResponseWriter, request *http.Request) {
	var kindRequest CreateKindRequest

	parseError := json.NewDecoder(request.Body).Decode(&kindRequest)

	if parseError != nil {
		http.Error(write, parseError.Error(), http.StatusBadRequest)
		return
	}

	db := middlewares.GetDBFromContext(request.Context())
	if db == nil {
		http.Error(write, "Database connection not found", http.StatusInternalServerError)
		return
	}

	kind := entities.Kind{Name: kindRequest.Name, StatusID: kindRequest.StatusID}
	createKindError := db.Create(&kind).Error

	if createKindError != nil {
		http.Error(write, "Unable to create kind", http.StatusInternalServerError)
		return
	}

	response := CreateKindResponse{
		ID:       kind.ID,
		Name:     kind.Name,
		StatusID: kind.StatusID,
	}

	write.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

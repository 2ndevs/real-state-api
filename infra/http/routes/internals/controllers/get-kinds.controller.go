package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"
)

type GetKindsResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func GetKinds(write http.ResponseWriter, request *http.Request) {
	nameFilter := request.URL.Query().Get("name")

	db := middlewares.GetDBFromContext(request.Context())

	if db == nil {
		http.Error(write, "Database connection not found", http.StatusInternalServerError)
		return
	}

	var kinds []entities.Kind
	query := db.Model(&entities.Kind{})

	if nameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+nameFilter+"%")
	}

	findError := query.Find(&kinds).Error

	if findError != nil {
		http.Error(write, "Unable to retrieve kinds", http.StatusInternalServerError)
		return
	}

	var response []GetKindsResponse

	for _, kind := range kinds {
		response = append(response, GetKindsResponse{
			ID:       kind.ID,
			Name:     kind.Name,
			StatusID: kind.StatusID,
		})
	}

	write.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(write).Encode(response); err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

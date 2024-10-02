package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GetKindResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	StatusID uint   `json:"status_id"`
}

func GetKind(write http.ResponseWriter, request *http.Request) {
	idParam := request.URL.Path[len("/kinds/"):]

	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		http.Error(write, "Invalid ID", http.StatusBadRequest)
		return
	}

	db := middlewares.GetDBFromContext(request.Context())

	if db == nil {
		http.Error(write, "Database connection not found", http.StatusInternalServerError)
		return
	}

	var kind entities.Kind

	findError := db.Model(&entities.Kind{}).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&kind).Error

	if findError == gorm.ErrRecordNotFound {
		http.Error(write, "Kind not found", http.StatusNotFound)
		return
	}

	if findError != nil {
		http.Error(write, "Unable to retrieve kind", http.StatusInternalServerError)
		return
	}

	response := GetKindResponse{
		ID:       kind.ID,
		Name:     kind.Name,
		StatusID: kind.StatusID,
	}

	write.WriteHeader(http.StatusOK)
	err = json.NewEncoder(write).Encode(response)

	if err != nil {
		http.Error(write, "Server error", http.StatusInternalServerError)
	}
}

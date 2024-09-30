package controllers

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type CreateKindRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

type CreateKindResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type KindController struct {
	DB *gorm.DB
}

func (kc *KindController) CreateKind(w http.ResponseWriter, r *http.Request) {
	var request CreateKindRequest

	// Validate body request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kind := entities.Kind{Name: request.Name}

	// Create new kind inside db
	if err := kc.DB.Create(&kind).Error; err != nil {
		http.Error(w, "Unable to create kind", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := CreateKindResponse{
		ID:   kind.ID,
		Name: kind.Name,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

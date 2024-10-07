package application

import (
	"errors"
	"main/domain/entities"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GetPropertyService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (propertyService *GetPropertyService) Execute() (*entities.Property, error) {
	idParam := propertyService.Request.URL.Path[len("/properties/"):]

	propertyId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		return nil, errors.Join(errors.New("invalid id"), validationErr)
	}

	property := entities.Property{}
	getPropertyTransaction := propertyService.Database.Find(&property, propertyId).Where("deleted_at IS NULL").First(&property)
	if getPropertyTransaction.Error != nil {
		return nil, getPropertyTransaction.Error
	}

	return &property, nil
}

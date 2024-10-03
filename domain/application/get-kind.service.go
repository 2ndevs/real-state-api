package application

import (
	"errors"
	"main/domain/entities"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GetKindService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (kindService *GetKindService) Execute() (*entities.Kind, error) {
	idParam := kindService.Request.URL.Path[len("/kinds/"):]

	kindId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		return nil, errors.Join(errors.New("invalid id"), validationErr)
	}

	kind := entities.Kind{}
	getKindTransaction := kindService.Database.Find(&kind, kindId).Where("deleted_at IS NULL").First(&kind)
	if getKindTransaction.Error != nil {
		return nil, getKindTransaction.Error
	}

	return &kind, nil
}

package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetKindService struct {
	Request  *http.Request
	KindID   uint64
	Database *gorm.DB
}

func (kindService *GetKindService) Execute() (*entities.Kind, error) {
	kind := entities.Kind{}
	
	getKindTransaction := kindService.Database.Find(&kind, kindService.KindID).Where("deleted_at IS NULL").First(&kind)
	if getKindTransaction.Error != nil {
		return nil, getKindTransaction.Error
	}

	return &kind, nil
}

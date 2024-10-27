package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetKindService struct {
	KindID   uint64
	Database *gorm.DB
}

func (self *GetKindService) Execute() (*entities.Kind, error) {
	kind := entities.Kind{}

	getKindTransaction := self.Database.Model(&kind).Find(&kind, self.KindID).Where("deleted_at IS NULL").First(&kind)
	if getKindTransaction.Error != nil {
		return nil, getKindTransaction.Error
	}

	return &kind, nil
}

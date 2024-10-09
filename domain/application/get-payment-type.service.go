package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetPaymentTypeService struct {
	PaymentTypeID uint64
	Database      *gorm.DB
}

func (self *GetPaymentTypeService) Execute() (*entities.PaymentType, error) {
	paymentType := entities.PaymentType{}

	getPaymentTypeTransaction := self.Database.Find(&paymentType, self.PaymentTypeID).Where("deleted_at IS NULL").First(&paymentType)
	if getPaymentTypeTransaction.Error != nil {
		return nil, getPaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetPaymentTypeService struct {
	Request       *http.Request
	PaymentTypeID uint64
	Database      *gorm.DB
}

func (paymentTypeService *GetPaymentTypeService) Execute() (*entities.PaymentType, error) {
	paymentType := entities.PaymentType{}

	getPaymentTypeTransaction := paymentTypeService.Database.Find(&paymentType, paymentTypeService.PaymentTypeID).Where("deleted_at IS NULL").First(&paymentType)
	if getPaymentTypeTransaction.Error != nil {
		return nil, getPaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

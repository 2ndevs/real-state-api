package application

import (
	"errors"
	"main/domain/entities"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GetPaymentTypeService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (paymentTypeService *GetPaymentTypeService) Execute() (*entities.PaymentType, error) {
	idParam := paymentTypeService.Request.URL.Path[len("/payment-types/"):]

	paymentTypeId, validationErr := strconv.ParseUint(idParam, 10, 32)
	if validationErr != nil {
		return nil, errors.Join(errors.New("invalid id"), validationErr)
	}

	paymentType := entities.PaymentType{}
	getPaymentTypeTransaction := paymentTypeService.Database.Find(&paymentType, paymentTypeId).Where("deleted_at IS NULL").First(&paymentType)
	if getPaymentTypeTransaction.Error != nil {
		return nil, getPaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

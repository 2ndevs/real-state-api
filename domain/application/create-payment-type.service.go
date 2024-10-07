package application

import (
	"errors"
	"log"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"

	"gorm.io/gorm"
)

type CreatePaymentTypeService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (paymentTypeService *CreatePaymentTypeService) Execute(paymentType entities.PaymentType) (*entities.PaymentType, error) {
	validate, ctxErr := middlewares.GetValidator(paymentTypeService.Request)
	if ctxErr != nil {
		return nil, ctxErr
	}

	validationErr := validate.Struct(paymentType)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	log.Printf("%v", validationErr)

	createPaymentTypeTransaction := paymentTypeService.Database.Create(&paymentType)
	if createPaymentTypeTransaction.Error != nil {
		return nil, createPaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

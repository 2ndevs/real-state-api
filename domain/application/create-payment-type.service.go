package application

import (
	"errors"
	"main/domain/entities"
	"main/infra/http/middlewares"
	"net/http"

	"gorm.io/gorm"
)

type CreatePaymentTypeService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (self *CreatePaymentTypeService) Execute(paymentType entities.PaymentType) (*entities.PaymentType, error) {
	validate, ctxErr := middlewares.GetValidator(self.Request)
	if ctxErr != nil {
		return nil, ctxErr
	}

	validationErr := validate.Struct(paymentType)
	if validationErr != nil {
		return nil, errors.Join(errors.New("validation error: "), validationErr)
	}

	createPaymentTypeTransaction := self.Database.Create(&paymentType)
	if createPaymentTypeTransaction.Error != nil {
		return nil, createPaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

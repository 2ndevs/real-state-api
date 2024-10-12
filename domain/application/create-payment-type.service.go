package application

import (
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreatePaymentTypeService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *CreatePaymentTypeService) Execute(paymentType entities.PaymentType) (*entities.PaymentType, error) {
	validationErr := self.Validated.Struct(paymentType)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	createPaymentTypeTransaction := self.Database.Create(&paymentType)
	if createPaymentTypeTransaction.Error != nil {
		return nil, createPaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

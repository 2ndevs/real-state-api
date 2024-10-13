package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdatePaymentTypeService struct {
	Validated *validator.Validate
	Database  *gorm.DB
}

func (self *UpdatePaymentTypeService) Execute(paymentType entities.PaymentType, paymentTypeID uint64) (*entities.PaymentType, error) {
	validationErr := self.Validated.Struct(paymentType)
	if validationErr != nil {
		return nil, core.InvalidParametersError
	}

	var existingPaymentType *entities.PaymentType
	query := self.Database.Model(&entities.PaymentType{}).Where("id = ?", paymentTypeID)

	existingPaymentTypeDatabaseResponse := query.First(&existingPaymentType)
	if errors.Is(existingPaymentTypeDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingPaymentTypeDatabaseResponse.Error != nil {
		return nil, existingPaymentTypeDatabaseResponse.Error
	}

	var samePaymentType *entities.PaymentType

	findSameQuery := self.Database.Model(&entities.PaymentType{}).Where("name = ? AND id != ?", paymentType.Name, existingPaymentType.ID)
	response := findSameQuery.First(&samePaymentType)

	if response.Error == nil {
		return nil, core.EntityAlreadyExistsError
	}

	paymentType.ID = existingPaymentType.ID

	updatePaymentTypeTransaction := self.Database.Save(&paymentType)
	if updatePaymentTypeTransaction.Error != nil {
		return nil, updatePaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

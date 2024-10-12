package application

import (
	"errors"
	"main/core"
	"main/domain/entities"

	"gorm.io/gorm"
)

type DeletePaymentTypeService struct {
	Database *gorm.DB
}

func (self *DeletePaymentTypeService) Execute(paymentTypeID uint64) (*entities.PaymentType, error) {
	var existingPaymentType *entities.PaymentType
	query := self.Database.Model(&entities.PaymentType{}).Where("id = ?", paymentTypeID)

	existingPaymentTypeDatabaseResponse := query.First(&existingPaymentType)
	if errors.Is(existingPaymentTypeDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, core.NotFoundError
	}

	if existingPaymentTypeDatabaseResponse.Error != nil {
		return nil, existingPaymentTypeDatabaseResponse.Error
	}

	deletePaymentTypeTransaction := self.Database.Delete(existingPaymentType)
	if deletePaymentTypeTransaction.Error != nil {
		return nil, deletePaymentTypeTransaction.Error
	}

	return existingPaymentType, nil
}

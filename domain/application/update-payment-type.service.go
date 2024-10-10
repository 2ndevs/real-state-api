package application

import (
	"errors"
	"main/domain/entities"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdatePaymentTypeService struct {
	Validated     *validator.Validate
	PaymentTypeID uint64
	Database      *gorm.DB
}

func (self *UpdatePaymentTypeService) Execute(paymentType entities.PaymentType) (*entities.PaymentType, error) {
	validationErr := self.Validated.Struct(paymentType)
	if validationErr != nil {
		return nil, errors.Join(errors.New("Erros de validação: "), validationErr)
	}

	var existingPaymentType *entities.PaymentType
	query := self.Database.Model(&entities.PaymentType{}).Where("id = ?", self.PaymentTypeID)

	existingPaymentTypeDatabaseResponse := query.First(&existingPaymentType)
	if errors.Is(existingPaymentTypeDatabaseResponse.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Tipo de pagamento não encontrado")
	}

	if existingPaymentTypeDatabaseResponse.Error != nil {
		return nil, existingPaymentTypeDatabaseResponse.Error
	}

	paymentType.ID = existingPaymentType.ID

	updatePaymentTypeTransaction := self.Database.Save(&paymentType)
	if updatePaymentTypeTransaction.Error != nil {
		return nil, updatePaymentTypeTransaction.Error
	}

	return &paymentType, nil
}

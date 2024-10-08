package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyPaymentTypesService struct {
	Request  *http.Request
	NameFilter *string
	Database *gorm.DB
}

func (paymentTypeService *GetManyPaymentTypesService) Execute() (*[]entities.PaymentType, error) {
	var paymentTypes []entities.PaymentType
	query := paymentTypeService.Database.Model(&entities.PaymentType{})

	if *paymentTypeService.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*paymentTypeService.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getPaymentTypesTransaction := query.Find(&paymentTypes)

	if getPaymentTypesTransaction.Error != nil {
		return nil, getPaymentTypesTransaction.Error
	}

	return &paymentTypes, nil
}

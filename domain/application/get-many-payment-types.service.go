package application

import (
	"main/domain/entities"
	"net/http"

	"gorm.io/gorm"
)

type GetManyPaymentTypesService struct {
	Request  *http.Request
	Database *gorm.DB
}

func (paymentTypeService *GetManyPaymentTypesService) Execute() (*[]entities.PaymentType, error) {
	nameFilter := paymentTypeService.Request.URL.Query().Get("name")

	var paymentTypes []entities.PaymentType
	query := paymentTypeService.Database.Model(&entities.PaymentType{})

	if nameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+nameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL")
	query = query.Order("name ASC")

	getPaymentTypesTransaction := query.Find(&paymentTypes)

	if getPaymentTypesTransaction.Error != nil {
		return nil, getPaymentTypesTransaction.Error
	}

	return &paymentTypes, nil
}

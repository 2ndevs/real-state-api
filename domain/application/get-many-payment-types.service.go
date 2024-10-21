package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetManyPaymentTypesService struct {
	NameFilter *string
	Database   *gorm.DB
}

func (self *GetManyPaymentTypesService) Execute() (*[]entities.PaymentType, error) {
	var paymentTypes []entities.PaymentType
	query := self.Database.Model(&entities.PaymentType{})

	if *self.NameFilter != "" {
		query = query.Where("name ILIKE ?", "%"+*self.NameFilter+"%")
	}

	query = query.Where("deleted_at IS NULL AND is_sold != true")
	query = query.Order("name ASC")

	getPaymentTypesTransaction := query.Find(&paymentTypes)

	if getPaymentTypesTransaction.Error != nil {
		return nil, getPaymentTypesTransaction.Error
	}

	return &paymentTypes, nil
}

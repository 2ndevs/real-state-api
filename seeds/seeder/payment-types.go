package seeder

import (
	"fmt"
	"main/domain/entities"

	"gorm.io/gorm"
)

func SeedPaymentTypes(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE payment_types RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table payment_types: %w", err)
	}

	paymentTypes := []entities.PaymentType{
		{Name: "Cartão de Crédito", StatusID: 1},
		{Name: "Cartão de Débito", StatusID: 1},
		{Name: "Transferência Bancária", StatusID: 1},
		{Name: "Dinheiro", StatusID: 1},
		{Name: "Financiamento", StatusID: 1},
		{Name: "Boleto Bancário", StatusID: 1},
		{Name: "PIX", StatusID: 1},
		{Name: "Cheque", StatusID: 1},
		{Name: "Depósito", StatusID: 1},
		{Name: "Consórcio", StatusID: 1},
	}

	for _, paymentType := range paymentTypes {
		db.Create(&paymentType)
	}

	fmt.Println("All payment types have been created.")
	fmt.Println("----------------------------")
}

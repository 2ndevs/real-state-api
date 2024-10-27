package seeder

import (
	"fmt"
	"main/domain/entities"

	"gorm.io/gorm"
)

func SeedNegotiationTypes(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE negotiation_types RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table negotiation_types: %w", err)
	}

	negotiarionTypes := []entities.NegotiationType{
		{Name: "Venda", StatusID: 1},
		{Name: "Financiamento", StatusID: 1},
		{Name: "Aluguel", StatusID: 1},
		{Name: "Permuta", StatusID: 1},
		{Name: "Leasing", StatusID: 1},
		{Name: "Consórcio", StatusID: 1},
		{Name: "Arrendamento", StatusID: 1},
		{Name: "Cessão de Direitos", StatusID: 1},
		{Name: "Parceria", StatusID: 1},
	}

	for _, negotiarionType := range negotiarionTypes {
		db.Create(&negotiarionType)
	}

	fmt.Println("All negotiation types have been created.")
	fmt.Println("----------------------------")
}

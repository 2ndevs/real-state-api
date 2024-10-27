package seeder

import (
	"fmt"
	"main/domain/entities"

	"gorm.io/gorm"
)

func SeedStatuses(db *gorm.DB) {
	statuses := []entities.Status{
		{Name: "Ativo"},
		{Name: "Inativo"},
	}

	err := db.Exec("TRUNCATE TABLE statuses RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table statuses: %w", err)
	}

	for _, status := range statuses {
		db.Create(&status)
	}

	fmt.Println("All status have been created.")
	fmt.Println("----------------------------")
}

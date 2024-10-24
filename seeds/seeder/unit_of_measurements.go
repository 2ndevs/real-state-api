package seeder

import (
	"fmt"
	"main/domain/entities"

	"gorm.io/gorm"
)

func SeedUnitOfMeasurements(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE unit_of_measurements RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table unit_of_measurements: %w", err)
	}

	unitsOfMeasurement := []entities.UnitOfMeasurement{
		{Name: "Metro Quadrado", Abbreviation: "m²", StatusID: 1},
		{Name: "Hectare", Abbreviation: "ha", StatusID: 1},
		{Name: "Acre", Abbreviation: "ac", StatusID: 1},
		{Name: "Alqueire", Abbreviation: "alq", StatusID: 1},
		{Name: "Quilômetro Quadrado", Abbreviation: "km²", StatusID: 1},
	}

	for _, unitOfMeasurement := range unitsOfMeasurement {
		db.Create(&unitOfMeasurement)
	}

	fmt.Println("All units of measurement have been created.")
	fmt.Println("----------------------------")
}

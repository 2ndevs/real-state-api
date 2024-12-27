package seeder

import (
	"fmt"
	"main/domain/entities"
	"math/rand"

	"gorm.io/gorm"
)

func SeedProperties(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE properties RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table properties: %v\n", err)
	}

	properties := []entities.Property{
		{
			TotalArea:           uint(rand.Float64()*500 + 50),
			BuiltArea:           uint(rand.Float64()*400 + 30),
			Rooms:               uint(rand.Float64()*6 + 1),
			Kitchens:            uint(rand.Float64()*3 + 1),
			Bathrooms:           uint(rand.Float64()*4 + 1),
			Suites:              uint(rand.Float64() * 3),
			Address:             "Avenida Paraná, 123, Umuarama, PR",
			Summary:             "Ótima localização comercial, próximo ao centro.",
			Details:             "Imóvel bem conservado com fácil acesso a mercados, farmácias e bancos.",
			Latitude:            -23.766667,
			Longitude:           -53.320556,
			Price:               rand.Float64()*1000000 + 50000,
			Discount:            rand.Float64() * 50000,
			IsHighlight:         rand.Float64() >= 0.8,
			ConstructionYear:    uint(rand.Float64()*50 + 1970),
			KindID:              uint(rand.Float64()*5 + 1),
			StatusID:            1,
			PaymentTypeID:       uint(rand.Float64()*5 + 1),
			UnitOfMeasurementID: uint(rand.Float64()*5 + 1),
			PreviewImages:       []string{"fallback.jpg"},
		},
		{
			TotalArea:           uint(rand.Float64()*500 + 50),
			BuiltArea:           uint(rand.Float64()*400 + 30),
			Rooms:               uint(rand.Float64()*6 + 1),
			Kitchens:            uint(rand.Float64()*3 + 1),
			Bathrooms:           uint(rand.Float64()*4 + 1),
			Suites:              uint(rand.Float64() * 3),
			Address:             "Rua Maringá, 456, Umuarama, PR",
			Summary:             "Próximo ao shopping e ao terminal de ônibus.",
			Details:             "Ideal para quem busca praticidade e conforto.",
			Latitude:            -23.763611,
			Longitude:           -53.324167,
			Price:               rand.Float64()*1000000 + 50000,
			Discount:            rand.Float64() * 50000,
			IsHighlight:         rand.Float64() >= 0.8,
			ConstructionYear:    uint(rand.Float64()*50 + 1970),
			KindID:              uint(rand.Float64()*5 + 1),
			StatusID:            1,
			PaymentTypeID:       uint(rand.Float64()*5 + 1),
			UnitOfMeasurementID: uint(rand.Float64()*5 + 1),
			PreviewImages:       []string{"fallback.jpg"},
		},
	}

	for _, property := range properties {
		db.Create(&property)
	}

	fmt.Println("All properties have been seeded.")
	fmt.Println("----------------------------")
}

func generateRandomMD5() string {
	return fmt.Sprintf("%x", rand.Int63())
}

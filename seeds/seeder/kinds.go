package seeder

import (
	"fmt"
	"main/domain/entities"

	"gorm.io/gorm"
)

func SeedKinds(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE kinds RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table kinds: %w", err)
	}

	kinds := []entities.Kind{
		{Name: "Apartamento", StatusID: 1},
		{Name: "Casa", StatusID: 1},
		{Name: "Terreno", StatusID: 1},
		{Name: "Sobrado", StatusID: 1},
		{Name: "Cobertura", StatusID: 1},
		{Name: "Flat", StatusID: 1},
		{Name: "Loft", StatusID: 1},
		{Name: "Kitnet", StatusID: 1},
		{Name: "Pousada", StatusID: 1},
		{Name: "Fazenda", StatusID: 1},
		{Name: "Sítio", StatusID: 1},
		{Name: "Chácara", StatusID: 1},
		{Name: "Duplex", StatusID: 1},
		{Name: "Triplex", StatusID: 1},
		{Name: "Studio", StatusID: 1},
		{Name: "Lote", StatusID: 1},
		{Name: "Galpão", StatusID: 1},
		{Name: "Armazém", StatusID: 1},
		{Name: "Barracão", StatusID: 1},
		{Name: "Condomínio", StatusID: 1},
		{Name: "Prédio Comercial", StatusID: 1},
		{Name: "Sala Comercial", StatusID: 1},
		{Name: "Loja", StatusID: 1},
		{Name: "Ponto Comercial", StatusID: 1},
		{Name: "Casa de Praia", StatusID: 1},
		{Name: "Casa de Campo", StatusID: 1},
		{Name: "Mansão", StatusID: 1},
		{Name: "Vila", StatusID: 1},
		{Name: "Cobertura Duplex", StatusID: 1},
		{Name: "Cobertura Triplex", StatusID: 1},
		{Name: "Edícula", StatusID: 1},
		{Name: "Pavilhão", StatusID: 1},
		{Name: "Hangar", StatusID: 1},
		{Name: "Área Rural", StatusID: 1},
		{Name: "Colônia", StatusID: 1},
		{Name: "Vila Rural", StatusID: 1},
		{Name: "Residencial", StatusID: 1},
		{Name: "Industrial", StatusID: 1},
		{Name: "Comercial", StatusID: 1},
		{Name: "Hotel", StatusID: 1},
	}

	for _, kind := range kinds {
		db.Create(&kind)
	}

	fmt.Println("All kinds have been created.")
	fmt.Println("----------------------------")
}

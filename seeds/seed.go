package main

import (
	"bufio"
	"fmt"
	"log"
	"main/seeds/seeder"
	"os"
	"strings"

	database "main/infra/repositories"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func askForSeed(seedName string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Do you want to seed '%s'? Type 'y' to confirm: ", seedName)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	return strings.TrimSpace(input) == "y"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env variables")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	seeds := []struct {
		name string
		fn   func(db *gorm.DB)
	}{
		{"statuses", seeder.SeedStatuses},
		{"kinds", seeder.SeedKinds},
		{"payment types", seeder.SeedPaymentTypes},
		{"unit of measurements", seeder.SeedUnitOfMeasurements},
		{"negotiation_types", seeder.SeedNegotiationTypes},
		{"properties", seeder.SeedProperties},
		{"roles", seeder.SeedRoles},
		{"users", seeder.SeedUsers},
	}

	for _, seed := range seeds {
		if askForSeed(seed.name) {
			fmt.Printf("Seeding %s...\n", seed.name)
			seed.fn(db)
		}
	}
}

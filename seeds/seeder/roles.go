package seeder

import (
	"fmt"
	"main/domain/entities"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE roles RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate table roles: %w", err)
	}

	// Definindo os papéis e suas permissões
	roles := []entities.Role{
		{
			Name:        "Admin",
			Permissions: pq.StringArray{"create", "read", "update", "delete"},
			StatusID:    1,
		},
		{
			Name:        "Editor",
			Permissions: pq.StringArray{"create", "read", "update"},
			StatusID:    1,
		},
		{
			Name:        "Visualizador",
			Permissions: pq.StringArray{"read"},
			StatusID:    1,
		},
		{
			Name:        "Convidado",
			Permissions: pq.StringArray{"read"},
			StatusID:    1,
		},
		{
			Name:        "SuperAdmin",
			Permissions: pq.StringArray{"create", "read", "update", "delete", "manage_users"},
			StatusID:    1,
		},
	}

	// Inserindo os registros de roles

	for _, role := range roles {
		db.Create(&role)
	}

	fmt.Println("All roles have been created.")
	fmt.Println("----------------------------")
}

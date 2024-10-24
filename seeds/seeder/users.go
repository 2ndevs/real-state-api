package seeder

import (
	"fmt"
	"main/domain/entities"
	"main/utils/libs"
	"math/rand"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	if err != nil {
		fmt.Printf("failed to truncate users roles: %w", err)
	}

	hasher := libs.Hashing{}
	hashedPassword, _ := hasher.EncryptPassword("123456")

	users := []entities.User{
		{
			Name:         "Alice Johnson",
			Email:        "alice@example.com",
			PasswordHash: *hashedPassword,
			StatusID:     1,
			RoleID:       uint(rand.Intn(5) + 1),
		},
		{
			Name:         "Bob Smith",
			Email:        "bob@example.com",
			PasswordHash: *hashedPassword,
			StatusID:     1,
			RoleID:       uint(rand.Intn(5) + 1),
		},
		{
			Name:         "Charlie Brown",
			Email:        "charlie@example.com",
			PasswordHash: *hashedPassword,
			StatusID:     1,
			RoleID:       uint(rand.Intn(5) + 1),
		},
		{
			Name:         "David Wilson",
			Email:        "david@example.com",
			PasswordHash: *hashedPassword,
			StatusID:     1,
			RoleID:       uint(rand.Intn(5) + 1),
		},
		{
			Name:         "Eva Green",
			Email:        "eva@example.com",
			PasswordHash: *hashedPassword,
			StatusID:     1,
			RoleID:       uint(rand.Intn(5) + 1),
		},
	}

	for _, user := range users {
		db.Create(&user)
	}

	fmt.Println("All users have been created.")
	fmt.Println("----------------------------")
}

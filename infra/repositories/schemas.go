package database

import (
	"gorm.io/gorm"
	"main/domain/entities"
)

type Status struct {
	gorm.Model
	entities.Status
}

type Kind struct {
	gorm.Model
	entities.Kind

	Status Status
}

type PaymentType struct {
	gorm.Model
	entities.PaymentType

	Status Status
}

type Property struct {
	gorm.Model
	entities.Property

	Kind        Kind
	Status      Status
	PaymentType PaymentType
}

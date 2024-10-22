package database

import (
	"main/domain/entities"
)

type User struct {
	entities.User
}

type Role struct {
	entities.Role

	Users []User
}

type Status struct {
	entities.Status

	Kinds            []Kind `gorm:"foreignKey:StatusID"`
	NegotiationTypes []NegotiationType
	PaymentTypes     []PaymentType
	Properties       []Property
	Roles            []Role
	Users            []User
}

type Kind struct {
	entities.Kind

	Status     Status
	Properties []Property `gorm:"foreignKey:KindID"`
}

type PaymentType struct {
	entities.PaymentType

	Status     Status
	Properties []Property
}

type NegotiationType struct {
	entities.NegotiationType

	Status     Status
	Properties []Property
}

type Property struct {
	entities.Property
}

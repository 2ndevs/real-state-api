package database

import (
	"main/domain/entities"
)

type User struct {
	entities.User
}

type Role struct {
	entities.Role
}

type Status struct {
	entities.Status
}

type Kind struct {
	entities.Kind

	Status Status
}

type PaymentType struct {
	entities.PaymentType

	Status Status
}

type NegotiationType struct {
	entities.NegotiationType

	Status Status
}

type Property struct {
	entities.Property

	Kind            Kind
	Status          Status
	PaymentType     PaymentType
	NegotiationType NegotiationType
}

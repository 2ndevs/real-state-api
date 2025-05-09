package database

import (
	"main/domain/entities"
)

type User struct {
	entities.User

	Role Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role struct {
	entities.Role

	Users []User `gorm:"foreignKey:RoleID"`
}

type Status struct {
	entities.Status

	Kinds              []Kind `gorm:"foreignKey:StatusID"`
	NegotiationTypes   []NegotiationType
	PaymentTypes       []PaymentType
	Properties         []Property
	Roles              []Role
	Users              []User
	UnitsOfMeasurement []UnitOfMeasurement
}

type Kind struct {
	entities.Kind

	Properties []Property `gorm:"foreignKey:KindID"`
}

type PaymentType struct {
	entities.PaymentType

	Status     Status
	Properties []Property
}

type NegotiationType struct {
	entities.NegotiationType

	Status Status
}

type Property struct {
	entities.Property

	Visits []Visit
}

type Visit struct {
	entities.Visit
}

type UnitOfMeasurement struct {
	entities.UnitOfMeasurement

	Status Status
}

type InterestedUser struct {
	entities.InterestedUser
}

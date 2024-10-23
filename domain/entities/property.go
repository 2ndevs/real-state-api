package entities

import (
	"main/core"

	"github.com/lib/pq"
)

type Property struct {
	core.Model

	Size             uint           `validate:"required,min=1"`
	Rooms            uint           `validate:"required,min=0"`
	Kitchens         uint           `validate:"required,min=0"`
	Bathrooms        uint           `validate:"required,min=0"`
	Address          string         `validate:"required"`
	Summary          string         `validate:"required"`
	Details          string         `validate:"required"`
	Latitude         float64        `validate:"required,gte=-90,lte=90"`
	Longitude        float64        `validate:"required,gte=-180,lte=180"`
	Price            float64        `validate:"required,min=1"`
	Discount         float64        `validate:"min=0"`
	IsSold           bool           `gorm:"default:false"`
	IsHighlight      bool           `gorm:"default:false"`
	ConstructionYear uint           `validate:"required,min=1945"`
	VisitedBy        pq.StringArray `gorm:"type:text[]"`
	PreviewImages    pq.StringArray `gorm:"type:text[]"`

	KindID              uint `validate:"required,min=1"`
	StatusID            uint `validate:"required,min=1"`
	PaymentTypeID       uint `validate:"required,min=1"`
	NegotiationTypeID   uint `validate:"required,min=1"`
	UnitOfMeasurementID uint `validate:"required,min=1"`

	Kind              Kind              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	Status            Status            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	PaymentType       PaymentType       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	NegotiationType   NegotiationType   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	UnitOfMeasurement UnitOfMeasurement `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
}

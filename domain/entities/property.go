package entities

import (
	"main/core"
	"time"

	"github.com/lib/pq"
)

type Property struct {
	core.Model

	TotalArea        uint           `validate:"required,min=1"`
	BuiltArea        uint           `validate:"required,min=1"`
	Rooms            uint           `validate:"min=0"`
	Suites           uint           `validate:"min=0"`
	Kitchens         uint           `validate:"min=0"`
	Bathrooms        uint           `validate:"min=0"`
	Address          string         `validate:"required"`
	Summary          string         `validate:"required"`
	Details          string         `validate:"required"`
	Latitude         float64        `validate:"gte=-90,lte=90"`
	Longitude        float64        `validate:"gte=-180,lte=180"`
	Price            float64        `validate:"required,min=1"`
	Discount         float64        `validate:"min=0"`
	SoldAt           *time.Time     `gorm:"default:null"`
	IsHighlight      bool           `gorm:"default:false"`
	ConstructionYear uint           `validate:"required,min=1945"`
	PreviewImages    pq.StringArray `gorm:"type:text[]"`
	ContactNumber    string         `validate:"required,min=13,max=13"`

	KindID              uint `gorm:"index" validate:"required,min=1"`
	StatusID            uint `gorm:"index" validate:"required,min=1"`
	PaymentTypeID       uint `gorm:"index" validate:"required,min=1"`
	UnitOfMeasurementID uint `gorm:"index" validate:"required,min=1"`

	Kind              Kind              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	Status            Status            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	PaymentType       PaymentType       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	UnitOfMeasurement UnitOfMeasurement `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
	Visits            []Visit           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" validate:"-"`
}

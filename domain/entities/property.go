package entities

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model

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

	KindID        uint `gorm:"index"`
	StatusID      uint `validate:"required,min=1"`
	PaymentTypeID uint `gorm:"index" validate:"required,min=1"`
}

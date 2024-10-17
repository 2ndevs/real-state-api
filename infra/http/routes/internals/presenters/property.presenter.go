package presenters

import (
	"encoding/json"
	"main/domain/entities"
	"net/http"
)

type PropertyPresenter struct{}

type PropertyFromHTTP struct {
	Size             uint    `json:"size" validate:"required,min=1"`
	Rooms            uint    `json:"rooms" validate:"required,min=0"`
	Kitchens         uint    `json:"kitchens" validate:"required,min=0"`
	Bathrooms        uint    `json:"bathrooms" validate:"required,min=0"`
	Address          string  `json:"address" validate:"required"`
	Summary          string  `json:"summary" validate:"required"`
	Details          string  `json:"details" validate:"required"`
	Latitude         float64 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude        float64 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Price            float64 `json:"price" validate:"required,min=1"`
	IsHighlight      bool    `json:"is_highlight" validate:"required"`
	Discount         float64 `json:"discount" validate:"min=0"`
	IsSold           bool    `json:"is_sold"`
	ConstructionYear uint    `json:"construction_year" validate:"required,min=1945"`
	VisitedBy        string  `json:"visited_by"`

	KindID        uint `json:"kind_id" validate:"required,min=1"`
	PaymentTypeID uint `json:"payment_type_id" validate:"required,min=1"`
}

type PropertyToHTTP struct {
	ID uint `json:"id"`

	Size             uint    `json:"size"`
	Rooms            uint    `json:"rooms"`
	Kitchens         uint    `json:"kitchens"`
	Bathrooms        uint    `json:"bathrooms"`
	Address          string  `json:"address"`
	Summary          string  `json:"summary"`
	Details          string  `json:"details"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Price            float64 `json:"price"`
	IsHighlight      bool    `json:"is_highlight"`
	Discount         float64 `json:"discount"`
	IsSold           bool    `json:"is_sold"`
	ConstructionYear uint    `json:"construction_year"`

	KindID        uint `json:"kind_id"`
	StatusID      uint `json:"status_id"`
	PaymentTypeID uint `json:"payment_type_id"`
}

func (PropertyPresenter) FromHTTP(request *http.Request) (*PropertyFromHTTP, error) {
	var propertyRequest PropertyFromHTTP
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&propertyRequest)
	if err != nil {
		return nil, err
	}

	return &propertyRequest, nil
}

func (PropertyPresenter) ToHTTP(property entities.Property) PropertyToHTTP {
	return PropertyToHTTP{
		ID: property.ID,

		Size:             property.Size,
		Rooms:            property.Rooms,
		Kitchens:         property.Kitchens,
		Bathrooms:        property.Bathrooms,
		Address:          property.Address,
		Summary:          property.Summary,
		Details:          property.Details,
		Latitude:         property.Latitude,
		Longitude:        property.Longitude,
		Price:            property.Price,
		IsHighlight:      property.IsHighlight,
		Discount:         property.Discount,
		IsSold:           property.IsSold,
		ConstructionYear: property.ConstructionYear,

		KindID:        property.KindID,
		StatusID:      property.StatusID,
		PaymentTypeID: property.PaymentTypeID,
	}
}

func (PropertyPresenter) GetIdentity(request *http.Request) (*string, error) {
	cookie, err := request.Cookie("identity")
	if err != nil {
		return nil, err
	}

	return &cookie.Value, nil
}

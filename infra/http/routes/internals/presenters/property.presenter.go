package presenters

import (
	"io"
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/utils/libs"
	"mime/multipart"
	"net/http"
	"strconv"
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
	PreviewImages    []byte  `json:"preview_images" validate:"required,min=1"`

	KindID            uint `json:"kind_id" validate:"required,min=1"`
	PaymentTypeID     uint `json:"payment_type_id" validate:"required,min=1"`
	NegotiationTypeID uint `json:"negotiation_type_id" validate:"required,min=1"`
}

type PropertyToHTTP struct {
	ID uint `json:"id"`

	Size             uint     `json:"size"`
	Rooms            uint     `json:"rooms"`
	Kitchens         uint     `json:"kitchens"`
	Bathrooms        uint     `json:"bathrooms"`
	Address          string   `json:"address"`
	Summary          string   `json:"summary"`
	Details          string   `json:"details"`
	Latitude         float64  `json:"latitude"`
	Longitude        float64  `json:"longitude"`
	Price            float64  `json:"price"`
	IsHighlight      bool     `json:"is_highlight"`
	Discount         float64  `json:"discount"`
	IsSold           bool     `json:"is_sold"`
	ConstructionYear uint     `json:"construction_year"`
	PreviewImages    []string `json:"preview_images"`

	KindID            uint `json:"kind_id"`
	StatusID          uint `json:"status_id"`
	PaymentTypeID     uint `json:"payment_type_id"`
	NegotiationTypeID uint `json:"negotiation_type_id"`
}

func (PropertyPresenter) FromHTTP(request *http.Request) (*PropertyFromHTTP, error) {
	request.ParseMultipartForm(1024 * 1024 * 15)
	files := request.MultipartForm.File

	var previewImages []byte

	for _, file := range files["preview_images"] {
		fileAsBytes := func(file *multipart.FileHeader) []byte {
			content, err := file.Open()
			defer content.Close()
			if err != nil {
				return nil
			}

			fileByte, err := io.ReadAll(content)
			if err != nil {
				return nil
			}

			return fileByte
		}(file)

		if fileAsBytes == nil {
			break
		}

		previewImages = append(previewImages, fileAsBytes...)
	}

	rooms, err := strconv.ParseUint(request.FormValue("rooms"), 32, 10)
	if err != nil {
		return nil, err
	}

	size, err := strconv.ParseUint(request.FormValue("size"), 32, 64)
	if err != nil {
		return nil, err
	}

	kitchens, err := strconv.ParseUint(request.FormValue("kitchens"), 32, 10)
	if err != nil {
		return nil, err
	}

	bathrooms, err := strconv.ParseUint(request.FormValue("bathrooms"), 32, 10)
	if err != nil {
		return nil, err
	}

	address := request.FormValue("address")
	if len(address) == 0 {
		return nil, core.InvalidParametersError
	}

	summary := request.FormValue("summary")
	if len(summary) == 0 {
		return nil, core.InvalidParametersError
	}

	details := request.FormValue("details")
	if len(details) == 0 {
		return nil, core.InvalidParametersError
	}

	latitude, err := strconv.ParseFloat(request.FormValue("latitude"), 64)
	if err != nil {
		return nil, err
	}

	longitude, err := strconv.ParseFloat(request.FormValue("longitude"), 64)
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(request.FormValue("price"), 64)
	if err != nil {
		return nil, err
	}

	isHighlight, err := strconv.ParseBool(request.FormValue("is_highlight"))
	if err != nil && request.FormValue("is_highlight") != "" {
		return nil, err
	}
	if request.FormValue("is_highlight") == "" {
		isHighlight = false
	}

	discount, err := strconv.ParseFloat(request.FormValue("discount"), 64)
	if err != nil && request.FormValue("discount") != "" {
		return nil, err
	}
	if request.FormValue("discount") == "" {
		discount = 0
	}

	isSold, err := strconv.ParseBool(request.FormValue("is_sold"))
	if err != nil && request.FormValue("is_sold") != "" {
		return nil, err
	}
	if request.FormValue("is_sold") == "" {
		isSold = false
	}

	constructionYear, err := strconv.ParseUint(request.FormValue("construction_year"), 32, 24)
	if err != nil {
		return nil, err
	}

	paymentTypeId, err := strconv.ParseUint(request.FormValue("payment_type_id"), 32, 24)
	if err != nil {
		return nil, err
	}

	negotiationTypeId, err := strconv.ParseUint(request.FormValue("negotiation_type_id"), 32, 24)
	if err != nil {
		return nil, err
	}

	kindId, err := strconv.ParseUint(request.FormValue("kind_id"), 32, 24)
	if err != nil {
		return nil, err
	}

	propertyRequest := PropertyFromHTTP{
		Rooms:            uint(rooms),
		Size:             uint(size),
		Kitchens:         uint(kitchens),
		Bathrooms:        uint(bathrooms),
		Address:          address,
		Summary:          summary,
		Details:          details,
		Latitude:         latitude,
		Longitude:        longitude,
		Price:            price,
		IsHighlight:      isHighlight,
		Discount:         discount,
		IsSold:           isSold,
		ConstructionYear: uint(constructionYear),
		PreviewImages:    previewImages,

		KindID:            uint(kindId),
		PaymentTypeID:     uint(paymentTypeId),
		NegotiationTypeID: uint(negotiationTypeId),
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
		PreviewImages:    property.PreviewImages,

		KindID:            property.KindID,
		StatusID:          property.StatusID,
		PaymentTypeID:     property.PaymentTypeID,
		NegotiationTypeID: property.NegotiationTypeId,
	}
}

func (PropertyPresenter) GetSearchParams(request *http.Request) application.GetManyPropertiesFilters {
	filters := application.GetManyPropertiesFilters{}

	searchFilter := request.URL.Query().Get("search")
	latitudeFilter := request.URL.Query().Get("latitude")
	longitudeFilter := request.URL.Query().Get("longitude")
	isNewFilter := request.URL.Query().Get("is_new")
	withDiscountFilter := request.URL.Query().Get("with_discount")
	recentlySoldFilter := request.URL.Query().Get("recently_sold")
	recentlyBuiltFilter := request.URL.Query().Get("recently_built")
	isSpecialFilter := request.URL.Query().Get("is_special")
	isApartmentFilter := request.URL.Query().Get("is_apartment")
	allowFinancingFilter := request.URL.Query().Get("allow_financing")
	mostVisitedFilter := request.URL.Query().Get("most_visited")

	if searchFilter != "" {
		filters.Search = &searchFilter
	}

	latitude := libs.ValidateAndConvertCoordinate(latitudeFilter, -90, 90)
	longitude := libs.ValidateAndConvertCoordinate(longitudeFilter, -180, 180)
	if latitude != nil && longitude != nil {
		filters.Latitude = latitude
		filters.Longitude = longitude
	}

	isNew, err := strconv.ParseBool(isNewFilter)

	if isNew && err == nil {
		filters.IsNew = &isNew
	}

	withDiscount, err := strconv.ParseBool(withDiscountFilter)
	if withDiscount && err == nil {
		filters.WithDiscount = &withDiscount
	}

	recentlySold, err := strconv.ParseBool(recentlySoldFilter)
	if recentlySold && err == nil {
		filters.RecentlySold = &recentlySold
	}

	recentlyBuilt, err := strconv.ParseBool(recentlyBuiltFilter)
	if recentlyBuilt && err == nil {
		filters.RecentlyBuilt = &recentlyBuilt
	}

	isSpecial, err := strconv.ParseBool(isSpecialFilter)
	if isSpecial && err == nil {
		filters.IsSpecial = &isSpecial
	}

	isApartment, err := strconv.ParseBool(isApartmentFilter)
	if isApartment && err == nil {
		filters.IsApartment = &isApartment
	}

	allowFinancing, err := strconv.ParseBool(allowFinancingFilter)
	if allowFinancing && err == nil {
		filters.AllowFinancing = &allowFinancing
	}

	mostVisited, err := strconv.ParseBool(mostVisitedFilter)
	if mostVisited && err == nil {
		filters.MostVisited = &mostVisited
	}

	return filters
}

func (PropertyPresenter) GetIdentity(request *http.Request) (*string, error) {
	cookie, err := request.Cookie("identity")
	if err != nil {
		return nil, err
	}

	return &cookie.Value, nil
}

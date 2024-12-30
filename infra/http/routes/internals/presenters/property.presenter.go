package presenters

import (
	"main/core"
	"main/domain/application"
	"main/domain/entities"
	"main/utils"
	"main/utils/libs"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

type PropertyPresenter struct{}

type PropertyFromHTTP struct {
	BuiltArea        uint                    `json:"built_area" validate:"required,min=1"`
	TotalArea        uint                    `json:"total_area" validate:"required,min=1"`
	Rooms            uint                    `json:"rooms" validate:"required,min=0"`
	Suites           uint                    `json:"suites" validate:"required,min=0"`
	Kitchens         uint                    `json:"kitchens" validate:"required,min=0"`
	Bathrooms        uint                    `json:"bathrooms" validate:"required,min=0"`
	Address          string                  `json:"address" validate:"required"`
	Summary          string                  `json:"summary" validate:"required"`
	Details          string                  `json:"details" validate:"required"`
	Latitude         float64                 `json:"latitude" validate:"required,gte=-90,lte=90"`
	Longitude        float64                 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Price            float64                 `json:"price" validate:"required,min=1"`
	IsHighlight      bool                    `json:"is_highlight" validate:"required"`
	Discount         float64                 `json:"discount" validate:"min=0"`
	SoldAt           *time.Time              `json:"sold_at"`
	ConstructionYear uint                    `json:"construction_year" validate:"required,min=1945"`
	PreviewImages    []*multipart.FileHeader `json:"preview_images" validate:"required,min=1"`
	ContactNumber    string                  `json:"contact_number" validate:"required,min=13,max=13"`

	KindID              uint  `json:"kind_id" validate:"required,min=1"`
	StatusID            *uint `json:"status_id"`
	PaymentTypeID       uint  `json:"payment_type_id" validate:"required,min=1"`
	UnitOfMeasurementID uint  `json:"unit_of_measurement_id" validate:"required,min=1"`
}

type VisitToHTTP struct {
	ID        uint      `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"create_at"`
}

type PropertyToHTTP struct {
	ID uint `json:"id"`

	BuiltArea        uint       `json:"built_area"`
	TotalArea        uint       `json:"total_area"`
	Rooms            uint       `json:"rooms"`
	Suites           uint       `json:"suites"`
	Kitchens         uint       `json:"kitchens"`
	Bathrooms        uint       `json:"bathrooms"`
	Address          string     `json:"address"`
	Summary          string     `json:"summary"`
	Details          string     `json:"details"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
	Price            float64    `json:"price"`
	IsHighlight      bool       `json:"is_highlight"`
	Discount         float64    `json:"discount"`
	SoldAt           *time.Time `json:"sold_at"`
	ConstructionYear uint       `json:"construction_year"`
	PreviewImages    []string   `json:"preview_images"`
	ContactNumber    string     `json:"contact_number"`

	KindID              *uint `json:"kind_id"`
	StatusID            *uint `json:"status_id"`
	PaymentTypeID       *uint `json:"payment_type_id"`
	UnitOfMeasurementID *uint `json:"unit_of_measurement_id"`

	Status            *entities.Status            `json:"status"`
	Kind              *entities.Kind              `json:"kind"`
	PaymentType       *entities.PaymentType       `json:"payment_type"`
	UnitOfMeasurement *entities.UnitOfMeasurement `json:"unit_of_measurement"`
	Visits            *[]VisitToHTTP              `json:"visits"`
}

func (PropertyPresenter) FromHTTP(request *http.Request) (*PropertyFromHTTP, error) {
	request.ParseMultipartForm(1024 * 1024 * 15)
	files := request.MultipartForm.File

	var previewImages []*multipart.FileHeader

	for _, file := range files["preview_images"] {
		previewImages = append(previewImages, file)
	}

	rooms, err := strconv.ParseUint(request.FormValue("rooms"), 10, 32)
	if err != nil {
		return nil, err
	}

	suites, err := strconv.ParseUint(request.FormValue("suites"), 10, 32)
	if err != nil {
		return nil, err
	}

	builtArea, err := strconv.ParseUint(request.FormValue("built_area"), 10, 32)
	if err != nil {
		return nil, err
	}

	totalArea, err := strconv.ParseUint(request.FormValue("total_area"), 10, 32)
	if err != nil {
		return nil, err
	}

	kitchens, err := strconv.ParseUint(request.FormValue("kitchens"), 10, 32)
	if err != nil {
		return nil, err
	}

	bathrooms, err := strconv.ParseUint(request.FormValue("bathrooms"), 10, 64)
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

	var soldAt *time.Time
	soldAtStr := request.FormValue("sold_at")
	if soldAtStr == "" {
		soldAt = nil
	}
	parsedTime, err := time.Parse(time.RFC3339, soldAtStr)
	if err == nil {
		soldAt = &parsedTime
	}

	constructionYear, err := strconv.ParseUint(request.FormValue("construction_year"), 10, 32)
	if err != nil {
		return nil, err
	}

	paymentTypeId, err := strconv.ParseUint(request.FormValue("payment_type_id"), 10, 32)
	if err != nil {
		return nil, err
	}

	kindId, err := strconv.ParseUint(request.FormValue("kind_id"), 10, 32)
	if err != nil {
		return nil, err
	}

	UnitOfMeasurementId, err := strconv.ParseUint(request.FormValue("unit_of_measurement_id"), 10, 32)
	if err != nil {
		return nil, err
	}

	Status, err := strconv.Atoi(request.FormValue("status_id"))
	if err != nil {
		return nil, err
	}

	StatusId := uint(Status)

	contactNumber := request.FormValue("contact_number")
	if len(contactNumber) == 0 {
		return nil, core.InvalidParametersError
	}

	propertyRequest := PropertyFromHTTP{
		Rooms:            uint(rooms),
		Suites:           uint(suites),
		BuiltArea:        uint(builtArea),
		TotalArea:        uint(totalArea),
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
		SoldAt:           soldAt,
		ConstructionYear: uint(constructionYear),
		PreviewImages:    previewImages,
		ContactNumber:    contactNumber,

		KindID:              uint(kindId),
		PaymentTypeID:       uint(paymentTypeId),
		UnitOfMeasurementID: uint(UnitOfMeasurementId),
		StatusID:            &StatusId,
	}

	return &propertyRequest, nil
}

func (PropertyPresenter) ToHTTP(property entities.Property) PropertyToHTTP {
	hasVisits := len(property.Visits) > 0

	entity := PropertyToHTTP{
		ID: property.ID,

		BuiltArea:        property.BuiltArea,
		TotalArea:        property.TotalArea,
		Rooms:            property.Rooms,
		Suites:           property.Suites,
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
		SoldAt:           property.SoldAt,
		ConstructionYear: property.ConstructionYear,
		PreviewImages:    property.PreviewImages,
		ContactNumber:    property.ContactNumber,

		KindID:              &property.KindID,
		StatusID:            &property.StatusID,
		PaymentTypeID:       &property.PaymentTypeID,
		UnitOfMeasurementID: &property.UnitOfMeasurementID,

		Status:            &property.Status,
		Kind:              &property.Kind,
		PaymentType:       &property.PaymentType,
		UnitOfMeasurement: &property.UnitOfMeasurement,
	}

	if hasVisits {
		visits := make([]VisitToHTTP, len(property.Visits))

		for i, visit := range property.Visits {
			visits[i] = VisitToHTTP{
				ID:        visit.ID,
				UserID:    visit.UserID,
				CreatedAt: visit.CreatedAt,
			}
		}

		entity.Visits = &visits
	}

	return entity
}

func (PropertyPresenter) GetSearchParams(request *http.Request) application.GetManyPropertiesFilters {
	filters := application.GetManyPropertiesFilters{}

	searchFilter := request.URL.Query().Get("search")
	latitudeFilter := request.URL.Query().Get("latitude")
	longitudeFilter := request.URL.Query().Get("longitude")
	isNewFilter := request.URL.Query().Get("is-new")
	withDiscountFilter := request.URL.Query().Get("with-discount")
	recentlySoldFilter := request.URL.Query().Get("recently-sold")
	recentlyBuiltFilter := request.URL.Query().Get("recently-built")
	isSpecialFilter := request.URL.Query().Get("is-special")
	isApartmentFilter := request.URL.Query().Get("is-apartment")
	allowFinancingFilter := request.URL.Query().Get("allow-financing")
	mostVisitedFilter := request.URL.Query().Get("most-visited")
	minValueFilter := request.URL.Query().Get("min-value")
	maxValueFilter := request.URL.Query().Get("max-value")
	negotiationTypesFilter := request.URL.Query().Get("negotiation-types")
	kindsFilter := request.URL.Query().Get("kinds")
	pageFilter := request.URL.Query().Get("page")
	perPageFilter := request.URL.Query().Get("per-page")
	minBedroomsFilter := request.URL.Query().Get("min-bedrooms")
	minBathroomsFilter := request.URL.Query().Get("min-bathrooms")
	maxBedroomsFilter := request.URL.Query().Get("max-bedrooms")
	maxBathroomsFilter := request.URL.Query().Get("max-bathrooms")
	sortByFilter := request.URL.Query().Get("sort-by")

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

	minValue, err := strconv.ParseFloat(minValueFilter, 32)
	if err == nil {
		filters.MinValue = &minValue
	}

	maxValue, err := strconv.ParseFloat(maxValueFilter, 32)
	if err == nil {
		filters.MaxValue = &maxValue
	}

	negotiationTypes := utils.StringToUintArray(negotiationTypesFilter)
	if negotiationTypes != nil {
		filters.NegotiationTypes = &negotiationTypes
	}

	kinds := utils.StringToUintArray(kindsFilter)
	if kinds != nil {
		filters.Kinds = &kinds
	}

	limit, limitErr := strconv.ParseInt(perPageFilter, 10, 32)
	if perPageFilter == "" || limit < 1 || limitErr != nil {
		limit = 15
	}
	filters.Limit = int(limit)

	page, pageErr := strconv.ParseInt(pageFilter, 10, 32)
	if pageFilter == "" || page < 1 || pageErr != nil {
		page = 1
	}
	filters.Offset = int((page - 1) * limit)

	minBedroom, minBedroomErr := strconv.ParseUint(minBedroomsFilter, 10, 32)
	if minBedroomsFilter != "" && minBedroomErr == nil {
		filters.MinBedrooms = &minBedroom
	}

	minBathroom, minBathroomErr := strconv.ParseUint(minBathroomsFilter, 10, 32)
	if minBathroomsFilter != "" && minBathroomErr == nil {
		filters.MinBathrooms = &minBathroom
	}

	maxBedroom, maxBedroomErr := strconv.ParseUint(maxBedroomsFilter, 10, 32)
	if maxBedroomsFilter != "" && maxBedroomErr == nil {
		filters.MaxBedrooms = &maxBedroom
	}

	maxBathroom, maxBathroomErr := strconv.ParseUint(maxBathroomsFilter, 10, 32)
	if maxBathroomsFilter != "" && maxBathroomErr == nil {
		filters.MaxBathrooms = &maxBathroom
	}

	validSortOptions := map[string]application.SortBy{
		"recents":       application.SortByRecents,
		"highest-price": application.SortByHighestPrice,
		"lowest-price":  application.SortByLowestPrice,
		"most-visiteds": application.SortByMostVisiteds,
	}

	sortOption, exists := validSortOptions[sortByFilter]
	if exists {
		filters.SortBy = &sortOption
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

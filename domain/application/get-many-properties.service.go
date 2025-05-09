package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SortBy string

const (
	SortByRecents      SortBy = "recents"
	SortByHighestPrice SortBy = "highest-price"
	SortByLowestPrice  SortBy = "lowest-price"
	SortByMostVisiteds SortBy = "most-visiteds"
)

type GetManyPropertiesFilters struct {
	Search           *string
	Latitude         *float64
	Longitude        *float64
	IsNew            *bool
	WithDiscount     *bool
	IsApartment      *bool
	RecentlySold     *bool
	RecentlyBuilt    *bool
	MostVisited      *bool
	AllowFinancing   *bool
	IsSpecial        *bool
	MinValue         *float64
	MaxValue         *float64
	MinBedrooms      *uint64
	MinBathrooms     *uint64
	MaxBedrooms      *uint64
	MaxBathrooms     *uint64
	NegotiationTypes *[]uint
	Kinds            *[]uint
	Offset           int
	Limit            int
	SortBy           *SortBy
}

type GetManyPropertiesService struct {
	Database *gorm.DB
}

func (self *GetManyPropertiesService) Execute(filters GetManyPropertiesFilters) (*[]entities.Property, error) {
	var properties []entities.Property
	query := self.Database.Model(&entities.Property{})

	if filters.Search != nil {
		query = query.Where("details ILIKE ?", "%"+*filters.Search+"%").
			Or("address ILIKE ?", "%"+*filters.Search+"%")
	}

	if filters.Latitude != nil && filters.Longitude != nil {
		latitude := *filters.Latitude
		longitude := *filters.Longitude

		haversineQuery := `
			6371 * acos(
				cos(radians(?)) * cos(radians(latitude)) * 
				cos(radians(longitude) - radians(?)) + 
				sin(radians(?)) * sin(radians(latitude))
			) <= 5
		`

		query = query.Where(gorm.Expr(haversineQuery, latitude, longitude, latitude))
	}

	if filters.IsNew != nil {
		query = query.Order("created_at DESC")
	}

	if filters.WithDiscount != nil {
		query = query.Where("discount > 0")
	}

	if filters.RecentlySold != nil {
		query = query.Where("sold_at IS NOT NULL").Order("sold_at DESC")
	}

	if filters.RecentlyBuilt != nil {
		query = query.Order("construction_year DESC")
	}

	if filters.IsSpecial != nil {
		query = query.Where("is_highlight = true").Order("updated_at DESC")
	}

	if filters.IsApartment != nil {
		var apartmentKind entities.Kind
		self.Database.Where("name = ?", "apartment").First(&apartmentKind)

		query = query.Where("kind_id = ?", apartmentKind.ID).Order("updated_at DESC")
	}

	if filters.AllowFinancing != nil {
		var financingPaymentType entities.PaymentType
		self.Database.Where("name = ?", "financing").First(&financingPaymentType)

		query = query.Where("payment_type_id = ?", financingPaymentType.ID).Order("updated_at DESC")
	}

	if filters.MostVisited != nil {
		query = query.Order("COALESCE(array_length(visited_by, 1), 0) DESC")
	}

	if filters.MinValue != nil {
		query = query.Where("price >= ?", *filters.MinValue)
	}

	if filters.MaxValue != nil {
		query = query.Where("price <= ?", *filters.MaxValue)
	}

	if filters.NegotiationTypes != nil && len(*filters.NegotiationTypes) > 0 {
		query = query.Where("negotiation_type_id IN ?", *filters.NegotiationTypes)
	}

	if filters.Kinds != nil && len(*filters.Kinds) > 0 {
		query = query.Where("kind_id IN ?", *filters.Kinds)
	}

	if filters.MinBedrooms != nil {
		query = query.Where("rooms >= ?", *filters.MinBedrooms)
	}

	if filters.MaxBedrooms != nil {
		query = query.Where("rooms <= ?", *filters.MaxBedrooms)
	}

	if filters.MinBathrooms != nil {
		query = query.Where("bathrooms >= ?", *filters.MinBathrooms)
	}

	if filters.MaxBathrooms != nil {
		query = query.Where("bathrooms <= ?", *filters.MaxBathrooms)
	}

	if filters.SortBy != nil {
		switch *filters.SortBy {
		case SortByRecents:
			query = query.Order("created_at DESC")
		case SortByHighestPrice:
			query = query.Order("price DESC")
		case SortByLowestPrice:
			query = query.Order("price ASC")
		case SortByMostVisiteds:
			query = query.Order("COALESCE(array_length(visited_by, 1), 0) DESC")
		}
	}

	query = query.Where("deleted_at IS NULL")

	getPropertiesTransaction := query.Preload(clause.Associations).Offset(filters.Offset).Limit(filters.Limit).Find(&properties)

	if getPropertiesTransaction.Error != nil {
		return nil, getPropertiesTransaction.Error
	}

	return &properties, nil
}

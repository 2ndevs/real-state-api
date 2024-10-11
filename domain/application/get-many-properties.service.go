package application

import (
	"main/domain/entities"

	"gorm.io/gorm"
)

type GetManyPropertiesFilters struct {
	Search    *string
	Latitude  *float64
	Longitude *float64
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

	query = query.Where("deleted_at IS NULL")

	getPropertiesTransaction := query.Find(&properties)

	if getPropertiesTransaction.Error != nil {
		return nil, getPropertiesTransaction.Error
	}

	return &properties, nil
}

package application

import (
	"main/domain/entities"
	"time"

	"gorm.io/gorm"
)

type Topics struct {
	NewItems      int64
	WithDiscount  int64
	Apartments    int64
	RecentlySold  int64
	RecentlyBuilt int64
	MostVisited   int64
	Financing     int64
	Special       int64
	TotalItems    int64
}

type GetTopicsService struct {
	Database *gorm.DB
}

func (self *GetTopicsService) Execute() (*Topics, error) {
	var topics Topics
	properties := entities.Property{}
	LAST_MONTH := time.Now().AddDate(0, -1, 0)
	LAST_YEAR := time.Now().AddDate(-1, 0, 0)

	self.Database.Model(properties).Where("created_at >= ? AND sold_at IS NULL AND deleted_at IS NULL", LAST_MONTH).Count(&topics.NewItems)

	self.Database.Model(properties).Where("discount > 0 AND sold_at IS NULL true AND deleted_at IS NULL").Count(&topics.WithDiscount)

	self.Database.Joins("JOIN kinds ON kinds.id = properties.kind_id").Where("kinds.name = ? AND properties.sold_at IS NULL AND properties.deleted_at IS NULL", "Apartamento").Model(properties).Count(&topics.Apartments)

	self.Database.Model(properties).Where("sold_at IS NOT NULL AND updated_at >= ?", LAST_MONTH).Count(&topics.RecentlySold)

	self.Database.Model(properties).Where("construction_year >= ? AND sold_at IS NULL AND deleted_at IS NULL", LAST_YEAR.Year()).Count(&topics.RecentlyBuilt)

	self.Database.Model(properties).Where("COALESCE(array_length(visited_by, 1), 0) > ? AND sold_at IS NULL AND deleted_at IS NULL", 20).Order("COALESCE(array_length(visited_by, 1), 0) DESC").Count(&topics.MostVisited)

	self.Database.Joins("JOIN payment_types ON payment_types.id = properties.payment_type_id").Where("payment_types.name = ? AND properties.sold_at IS NULL true AND properties.deleted_at IS NULL", "Financiamento").Model(properties).Count(&topics.Financing)

	self.Database.Model(properties).Where("is_highlight = true AND sold_at IS NULL AND deleted_at IS NULL").Count(&topics.Special)

	self.Database.Model(properties).Where("sold_at IS NULL AND deleted_at IS NULL").Count(&topics.TotalItems)

	return &topics, nil
}

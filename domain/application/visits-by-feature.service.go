package application

import (
	"fmt"

	"gorm.io/gorm"
)

type VisitsByFeatureService struct {
	Database *gorm.DB
	Feature  string
}

type VisitsByFeatureAndMonthResponse struct {
	Feature string `json:"feature"`
	Month   string `json:"month"`
	Count   int    `json:"count"`
}

func (self *VisitsByFeatureService) Execute() ([]VisitsByFeatureAndMonthResponse, error) {
	var result []VisitsByFeatureAndMonthResponse

	query := self.Database.Table("properties").
		Select("? AS feature, DATE_TRUNC('month', visits.created_at) AT TIME ZONE 'utc' AS month, COUNT(visits.id) AS count", self.Feature).
		Joins("JOIN visits ON visits.property_id = properties.id").
		Group(fmt.Sprintf("%s, month", self.Feature)).
		Order("month, count DESC")

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

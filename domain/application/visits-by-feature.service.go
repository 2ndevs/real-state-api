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
	Feature      string `json:"feature"`
	Month        string `json:"month"`
	FeatureValue int    `json:"feature_value"`
	Count        int    `json:"count"`
	Rank         int    `json:"rank"`
}

func (self *VisitsByFeatureService) Execute() ([]VisitsByFeatureAndMonthResponse, error) {
	var allResult []VisitsByFeatureAndMonthResponse

	err := self.Database.
		Table("properties").
		Select(fmt.Sprintf("? AS feature, DATE_TRUNC('month', visits.created_at) AS month, properties.%s AS feature_value, COUNT(visits.id) AS count", self.Feature), self.Feature).
		Joins("JOIN visits ON visits.property_id = properties.id").
		Group(fmt.Sprintf("DATE_TRUNC('month', visits.created_at), properties.%s", self.Feature)).
		Order("month, count DESC").
		Scan(&allResult).
		Error
	if err != nil {
		return nil, err
	}

	rankedResults := []VisitsByFeatureAndMonthResponse{}
	currentMonth := ""
	rankCount := 0

	for _, result := range allResult {
		if currentMonth != result.Month {
			currentMonth = result.Month
			rankCount = 0
		}

		if rankCount < 2 {
			rankedResults = append(rankedResults, VisitsByFeatureAndMonthResponse{
				Feature:      result.Feature,
				Month:        result.Month,
				FeatureValue: result.FeatureValue,
				Count:        result.Count,
				Rank:         rankCount + 1,
			})
			rankCount++
		}
	}

	return rankedResults, nil
}

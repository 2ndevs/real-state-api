package presenters

import "main/domain/application"

type VisitsByFeaturePresenter struct{}

type VisitsByFeatureToHTTP struct {
	Feature      string `json:"feature"`
	Month        string `json:"month"`
	FeatureValue int    `json:"feature_value"`
	Count        int    `json:"count"`
	Rank         int    `json:"rank"`
}

func (VisitsByFeaturePresenter) ToHTTP(data []application.VisitsByFeatureAndMonthResponse) []VisitsByFeatureToHTTP {
	result := make([]VisitsByFeatureToHTTP, len(data))

	for i, item := range data {
		result[i] = VisitsByFeatureToHTTP{
			Feature:      item.Feature,
			Month:        item.Month,
			FeatureValue: item.FeatureValue,
			Count:        item.Count,
			Rank:         item.Rank,
		}
	}

	return result
}

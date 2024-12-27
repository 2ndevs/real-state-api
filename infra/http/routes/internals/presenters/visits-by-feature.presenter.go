package presenters

import "main/domain/application"

type VisitsByFeaturePresenter struct{}

type VisitsByFeatureToHTTP struct {
	Feature string `json:"feature"`
	Month   string `json:"month"`
	Count   int    `json:"count"`
}

func (VisitsByFeaturePresenter) ToHTTP(data []application.VisitsByFeatureAndMonthResponse) []VisitsByFeatureToHTTP {
	result := make([]VisitsByFeatureToHTTP, len(data))

	for i, item := range data {
		result[i] = VisitsByFeatureToHTTP{
			Feature: item.Feature,
			Month:   item.Month,
			Count:   item.Count,
		}
	}

	return result
}

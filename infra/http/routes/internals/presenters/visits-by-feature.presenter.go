package presenters

import "main/domain/application"

type VisitsByFeaturePresenter struct{}

type VisitsByFeatureToHTTP struct {
	PropertyID uint `json:"property_id"`
	Count      int  `json:"count"`
}

func (VisitsByFeaturePresenter) ToHTTP(data []application.VisitsByFeatureResponse) []VisitsByFeatureToHTTP {
	result := make([]VisitsByFeatureToHTTP, len(data))

	for i, item := range data {
		result[i] = VisitsByFeatureToHTTP{
			PropertyID: item.PropertyID,
			Count:      item.Count,
		}
	}

	return result
}

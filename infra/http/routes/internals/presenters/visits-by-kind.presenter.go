package presenters

import "main/domain/application"

type VisitsByKindPresenter struct{}

type VisitsByKindToHTTP struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
}

func (VisitsByKindPresenter) ToHTTP(data []application.VisitsByKindResponse) []VisitsByKindToHTTP {
	result := make([]VisitsByKindToHTTP, len(data))
	for i, item := range data {
		result[i] = VisitsByKindToHTTP{
			Kind:  item.Kind,
			Count: item.Count,
		}
	}
	return result
}

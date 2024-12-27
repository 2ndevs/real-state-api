package presenters

import (
	"main/domain/application"
)

type VisitsPerMonthPresenter struct{}

type VisitsPerMonthToHTTP struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

func (VisitsPerMonthPresenter) ToHTTP(data []application.VisitsPerMonthResponse) []VisitsPerMonthToHTTP {
	result := make([]VisitsPerMonthToHTTP, len(data))

	for i, item := range data {
		result[i] = VisitsPerMonthToHTTP{
			Month: item.Month,
			Count: item.Count,
		}
	}

	return result
}

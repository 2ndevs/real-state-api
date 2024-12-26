package presenters

import "main/domain/application"

type SoldPerMonthPresenter struct{}

type SoldPerMonthToHTTP struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

func (SoldPerMonthPresenter) ToHTTP(data []application.SoldPerMonthResponse) []SoldPerMonthToHTTP {
	result := make([]SoldPerMonthToHTTP, len(data))
	for i, item := range data {
		result[i] = SoldPerMonthToHTTP{
			Month: item.Month,
			Count: item.Count,
		}
	}
	return result
}

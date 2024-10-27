package presenters

import (
	"main/domain/application"
)

type TopicPresenter struct{}

type TopicsToHTTP struct {
	NewItems      int64 `json:"new_items"`
	WithDiscount  int64 `json:"with_discount"`
	Apartments    int64 `json:"apartments"`
	RecentlySold  int64 `json:"recently_sold"`
	RecentlyBuilt int64 `json:"recently_built"`
	MostVisited   int64 `json:"most_visited"`
	Financing     int64 `json:"financing"`
	Special       int64 `json:"special"`
	TotalItems    int64 `json:"total_items"`
}

func (TopicPresenter) ToHTTP(topic application.Topics) TopicsToHTTP {
	return TopicsToHTTP{
		NewItems:      topic.NewItems,
		WithDiscount:  topic.WithDiscount,
		Apartments:    topic.Apartments,
		RecentlySold:  topic.RecentlySold,
		RecentlyBuilt: topic.RecentlyBuilt,
		MostVisited:   topic.MostVisited,
		Financing:     topic.Financing,
		Special:       topic.Special,
		TotalItems:    topic.TotalItems,
	}
}

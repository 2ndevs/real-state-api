package presenters

import (
	"main/domain/entities"
)

type FiltersPresenter struct{}

type FiltersToHTTP struct {
	Bathrooms uint `json:"bathrooms"`
	Rooms     uint `json:"rooms"`
	Kitchens  uint `json:"itchens"`
}

func (FiltersPresenter) ToHTTP(filters entities.Status) FiltersToHTTP {
	return FiltersToHTTP{
		Bathrooms: filters.ID,
		Rooms:     filters.ID,
		Kitchens:  filters.ID,
	}
}

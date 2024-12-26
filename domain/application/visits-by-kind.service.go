package application

import "gorm.io/gorm"

type VisitsByKindService struct {
	Database *gorm.DB
}

type VisitsByKindResponse struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
}

func (self *VisitsByKindService) Execute() ([]VisitsByKindResponse, error) {
	var result []VisitsByKindResponse
	err := self.Database.
		Table("visits").
		Select("kinds.name AS kind, COUNT(visits.id) AS count").
		Joins("JOIN properties ON visits.property_id = properties.id").
		Joins("JOIN kinds ON properties.kind_id = kinds.id").
		Group("kinds.name").
		Order("count DESC").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

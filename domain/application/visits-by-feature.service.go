package application

import "gorm.io/gorm"

type VisitsByFeatureService struct {
	Database *gorm.DB
	Feature  string
	Value    int
}

type VisitsByFeatureResponse struct {
	PropertyID uint `json:"property_id"`
	Count      int  `json:"count"`
}

func (self *VisitsByFeatureService) Execute() ([]VisitsByFeatureResponse, error) {
	var result []VisitsByFeatureResponse
	query := self.Database.Table("properties").
		Select("properties.id AS property_id, COUNT(visits.id) AS count").
		Joins("JOIN visits ON visits.property_id = properties.id").
		Group("properties.id").
		Order("count DESC")

	switch self.Feature {
	case "bathrooms":
		query = query.Where("bathrooms = ?", self.Value)
	case "rooms":
		query = query.Where("rooms = ?", self.Value)
	case "suites":
		query = query.Where("suites = ?", self.Value)
	case "kitchens":
		query = query.Where("kitchens = ?", self.Value)

	default:
		return nil, nil
	}

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

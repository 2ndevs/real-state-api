package application

import (
	"gorm.io/gorm"
)

type VisitsPerMonthService struct {
	Database *gorm.DB
}

type VisitsPerMonthResponse struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

func (self *VisitsPerMonthService) Execute() ([]VisitsPerMonthResponse, error) {
	var result []VisitsPerMonthResponse
	err := self.Database.
		Table("visits").
		Select("DATE_TRUNC('month', created_at) AS month, COUNT(*) AS count").
		Group("month").
		Order("month").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

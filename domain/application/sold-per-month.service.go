package application

import "gorm.io/gorm"

type SoldPerMonthService struct {
	Database *gorm.DB
}

type SoldPerMonthResponse struct {
	Month string `json:"month"`
	Count int    `json:"count"`
}

func (self *SoldPerMonthService) Execute() ([]SoldPerMonthResponse, error) {
	var result []SoldPerMonthResponse
	err := self.Database.
		Table("properties").
		Select("DATE_TRUNC('month', sold_at) AT TIME ZONE 'utc' AS month, COUNT(*) AS count").
		Where("sold_at IS NOT NULL").
		Group("month").
		Order("month").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

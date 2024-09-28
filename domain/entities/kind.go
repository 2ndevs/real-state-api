package entities

type Kind struct {
	Name   string `gorm:"index"`
	StatusID uint
}

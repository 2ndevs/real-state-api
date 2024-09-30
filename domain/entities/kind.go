package entities

type Kind struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"index"`
	StatusID uint
}

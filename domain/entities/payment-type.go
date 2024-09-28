package entities

type PaymentType struct {
	Name string `gorm:"index"`

	StatusID uint
}

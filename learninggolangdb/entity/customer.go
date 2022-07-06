package entity

type Customer struct {
	ID          string
	Test        string `gorm:"column:name"`
	Address     string
	Description string
}

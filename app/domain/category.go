package domain

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique"`
	Name string `gorm:"not null"`
}

func (p *Category) TableName() string {
	return "categories"
}

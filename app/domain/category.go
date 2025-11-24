package domain

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"unique"`
	Name string `json:"name" gorm:"not null"`
}

func (p *Category) TableName() string {
	return "categories"
}

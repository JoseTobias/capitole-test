package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Code string `json:"code" gorm:"unique"`
	Name string `json:"name" gorm:"not null"`
}

func (p *Category) TableName() string {
	return "categories"
}

type CreateCategoryRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (r CreateCategoryRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Code, validation.Required),
		validation.Field(&r.Name, validation.Required, validation.Length(1, 200)),
	)
}

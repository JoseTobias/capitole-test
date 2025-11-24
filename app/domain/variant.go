package domain

import "github.com/shopspring/decimal"

// Variant represents a product variant in the catalog.
// It includes a unique name, SKU, and an optional price.
// Variants can be used to represent different configurations or options for a product.
type Variant struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	ProductID uint            `json:"productID" gorm:"not null"`
	Name      string          `json:"name" gorm:"not null"`
	SKU       string          `json:"SKU" gorm:"uniqueIndex;not null"`
	Price     decimal.Decimal `json:"price" gorm:"type:decimal(10,2);null"`
}

func (v *Variant) TableName() string {
	return "product_variants"
}

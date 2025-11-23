package domain

import "github.com/shopspring/decimal"

// Product represents a product in the catalog.
// It includes a unique code and a price.
type Product struct {
	ID         uint            `gorm:"primaryKey"`
	Code       string          `gorm:"uniqueIndex;not null"`
	Price      decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	Variants   []Variant       `gorm:"foreignKey:ProductID"`
	CategoryID uint            `gorm:"not null"`
	Category   Category        `gorm:"foreignKey:CategoryID;references:ID"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductResponse struct {
	Code     string   `json:"code"`
	Price    float64  `json:"price"`
	Category Category `json:"category"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
	Paging   Paging    `json:"paging"`
}

type GetProductsRequest struct {
	Limit  int64 `json:"-"`
	Offset int64 `json:"-"`
}

func NewGetProductsRequest(q QueryGetter) *GetProductsRequest {
	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")

	if limitStr == "" {
		limitStr = "10"
	}
	if offsetStr == "" {
		offsetStr = "0"
	}
	return &GetProductsRequest{
		Limit:  min(StringToInt64(limitStr), 100),
		Offset: StringToInt64(offsetStr),
	}
}

type GetProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Paging   Paging            `json:"paging"`
}

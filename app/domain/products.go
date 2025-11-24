package domain

import "github.com/shopspring/decimal"

// Product represents a product in the catalog.
// It includes a unique code and a price.
type Product struct {
	ID         uint            `json:"id" gorm:"primaryKey"`
	Code       string          `json:"code" gorm:"uniqueIndex;not null"`
	Price      decimal.Decimal `json:"price" gorm:"type:decimal(10,2);not null"`
	Variants   []Variant       `json:"variants" gorm:"foreignKey:ProductID"`
	CategoryID uint            `json:"categoryID" gorm:"not null"`
	Category   Category        `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductResponse struct {
	Code     string            `json:"code"`
	Price    float64           `json:"price"`
	Category Category          `json:"category"`
	Variants []VariantResponse `json:"variants,omitempty"`
}

type ProductsResponse struct {
	Products []Product `json:"products"`
	Paging   Paging    `json:"paging"`
}

type GetProductsRequest struct {
	Price      decimal.Decimal `json:"price"`
	CategoryID uint            `json:"category_id"`
	Limit      int             `json:"-"`
	Offset     int             `json:"-"`
}

func NewGetProductsRequest(q QueryGetter) *GetProductsRequest {
	limitStr := q.Get("limit")
	offsetStr := q.Get("offset")
	return &GetProductsRequest{
		Price:      StringToDecimal(q.Get("price")),
		CategoryID: StringToUint(q.Get("category_id")),
		Limit:      min(StringToInt(limitStr, 10), 100),
		Offset:     StringToInt(offsetStr, 0),
	}
}

type GetProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Paging   Paging            `json:"paging"`
}

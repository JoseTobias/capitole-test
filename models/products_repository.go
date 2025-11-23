package models

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) GetAllProducts(req *domain.GetProductsRequest) (*domain.ProductsResponse, error) {
	var (
		products []domain.Product
		total    int64
	)

	baseQuery := r.db.Model(&domain.Product{})

	if req.CategoryID > 0 {
		baseQuery.Where("category_id = ?", req.CategoryID)
	}

	if !req.Price.IsZero() {
		baseQuery.Where("price < ?", req.Price)
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	if err := baseQuery.
		Preload("Variants").
		Preload("Category").
		Limit(req.Limit).
		Offset(req.Offset).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return &domain.ProductsResponse{
		Products: products,
		Paging: domain.Paging{
			Total:  int(total),
			Offset: req.Offset,
			Limit:  req.Limit,
		},
	}, nil
}

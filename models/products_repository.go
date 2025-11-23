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

	if err := r.db.Model(&domain.Product{}).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := r.db.
		Preload("Variants").
		Preload("Category").
		Limit(int(req.Limit)).
		Offset(int(req.Offset)).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return &domain.ProductsResponse{
		Products: products,
		Paging: domain.Paging{
			Total:  total,
			Offset: req.Offset,
			Limit:  req.Limit,
		},
	}, nil
}

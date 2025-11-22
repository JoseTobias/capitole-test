package catalog

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type ProductGetter interface {
	Get() (*domain.GetProductsResponse, error)
}

type ProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
}

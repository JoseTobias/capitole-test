package catalog

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type ProductGetter interface {
	Get(request *domain.GetProductsRequest) (*domain.GetProductsResponse, error)
}

type ProductRepository interface {
	GetAllProducts(request *domain.GetProductsRequest) (*domain.ProductsResponse, error)
}

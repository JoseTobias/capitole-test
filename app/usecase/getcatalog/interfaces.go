package getcatalog

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type ProductRepository interface {
	GetAllProducts(request *domain.GetProductsRequest) (*domain.ProductsResponse, error)
}

package catalogget

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type ProductGetter interface {
	Get(request *domain.GetProductsRequest) (*domain.GetProductsResponse, error)
}

package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
)

type GetCatalog struct {
	repository ProductRepository
}

func NewGetCatalog(r ProductRepository) *GetCatalog {
	return &GetCatalog{
		repository: r,
	}
}

func (s *GetCatalog) Get() (*domain.GetProductsResponse, error) {
	res, err := s.repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	products := make([]domain.ProductResponse, len(res))
	for i, p := range res {
		products[i] = domain.ProductResponse{
			Code:     p.Code,
			Price:    p.Price.InexactFloat64(),
			Category: p.Category,
		}
	}

	re := &domain.GetProductsResponse{
		Products: products,
	}
	return re, err
}

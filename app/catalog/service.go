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

func (s *GetCatalog) Get(req *domain.GetProductsRequest) (*domain.GetProductsResponse, error) {
	prd, err := s.repository.GetAllProducts(req)
	if err != nil {
		return nil, err
	}

	products := make([]domain.ProductResponse, len(prd.Products))
	for i, p := range prd.Products {
		products[i] = domain.ProductResponse{
			Code:     p.Code,
			Price:    p.Price.InexactFloat64(),
			Category: p.Category,
		}
	}

	res := &domain.GetProductsResponse{
		Products: products,
		Paging:   prd.Paging,
	}
	return res, err
}

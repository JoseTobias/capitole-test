package catalogbycode

import (
	"errors"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/repositories/products"
)

type GetCatalogByCode struct {
	repository CatalogRepository
}

func NewGetCatalog(r CatalogRepository) *GetCatalogByCode {
	return &GetCatalogByCode{
		repository: r,
	}
}

func (s *GetCatalogByCode) GetByCode(code string) (*domain.ProductResponse, error) {
	prd, err := s.repository.GetProductByCode(code)
	if err != nil {
		if errors.Is(err, products.ErrProductNotFound) {
			return nil, ErrProductNotFound
		}

		return nil, err
	}

	variants := make([]domain.VariantResponse, len(prd.Variants))
	for i, v := range prd.Variants {
		price := v.Price
		if price.IsZero() {
			price = prd.Price
		}
		variants[i] = domain.VariantResponse{
			ID:        v.ID,
			ProductID: v.ProductID,
			Name:      v.Name,
			SKU:       v.SKU,
			Price:     price.InexactFloat64(),
		}
	}

	res := &domain.ProductResponse{
		Code:     prd.Code,
		Price:    prd.Price.InexactFloat64(),
		Category: prd.Category,
		Variants: variants,
	}

	return res, err
}

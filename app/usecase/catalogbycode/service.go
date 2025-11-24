package catalogbycode

import (
	"errors"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"github.com/mytheresa/go-hiring-challenge/models"
)

type GetCatalogByCode struct {
	repository CatalogRepository
}

func NewGetCatalog(r CatalogRepository) *GetCatalogByCode {
	return &GetCatalogByCode{
		repository: r,
	}
}

func (s *GetCatalogByCode) GetByCode(code string) (*domain.Product, error) {
	prd, err := s.repository.GetProductByCode(code)
	if err != nil {
		if errors.Is(err, models.ErrProductNotFound) {
			return nil, ErrProductNotFound
		}

		return nil, err
	}

	variants := make([]domain.Variant, len(prd.Variants))
	for i, v := range prd.Variants {
		price := v.Price
		if price.IsZero() {
			price = prd.Price
		}
		variants[i] = domain.Variant{
			ID:        v.ID,
			ProductID: v.ProductID,
			Name:      v.Name,
			SKU:       v.SKU,
			Price:     price,
		}
	}

	prd.Variants = variants

	return prd, err
}

package getcatalogbycode

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type GetCatalogByCode interface {
	GetByCode(code string) (*domain.Product, error)
}

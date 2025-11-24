package catalogbyid

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type CatalogRepository interface {
	GetProductByCode(code string) (*domain.Product, error)
}

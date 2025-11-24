package catalogbycode

import "github.com/mytheresa/go-hiring-challenge/app/domain"

//go:generate mockgen -source=./interfaces.go -package=mock -destination=./mock/interfaces.go
type CatalogRepository interface {
	GetProductByCode(code string) (*domain.Product, error)
}

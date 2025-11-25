package createcategory

import "github.com/mytheresa/go-hiring-challenge/app/domain"

//go:generate mockgen -source=./interfaces.go -package=mock -destination=./mock/interfaces.go
type CategoryRepository interface {
	Create(req *domain.Category) error
	GetCategoryByCode(code string) (*domain.Category, error)
}

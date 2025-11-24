package getcategories

import "github.com/mytheresa/go-hiring-challenge/app/domain"

//go:generate mockgen -source=./interfaces.go -package=mock -destination=./mock/interfaces.go
type CategoryRepository interface {
	Get() ([]domain.Category, error)
}

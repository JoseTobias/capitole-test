package getcategories

import "github.com/mytheresa/go-hiring-challenge/app/domain"

type CategoryRepository interface {
	Get() ([]domain.Category, error)
}

package categoriesget

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"net/http"
)

//go:generate mockgen -source=./interfaces.go -package=mock -destination=./mock/interfaces.go
type CategoriesGetter interface {
	Get() ([]domain.Category, error)
}

type Responder interface {
	Ok(w http.ResponseWriter, data any)
	Error(w http.ResponseWriter, status int, message string)
}

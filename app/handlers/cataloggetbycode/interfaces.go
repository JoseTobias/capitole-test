package cataloggetbycode

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"net/http"
)

//go:generate mockgen -source=./interfaces.go -package=mock -destination=./mock/interfaces.go
type GetCatalogByCode interface {
	GetByCode(code string) (*domain.ProductResponse, error)
}

type Responder interface {
	Ok(w http.ResponseWriter, data any)
	Error(w http.ResponseWriter, status int, message string)
}

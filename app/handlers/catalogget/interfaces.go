package catalogget

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"net/http"
)

type ProductGetter interface {
	Get(request *domain.GetProductsRequest) (*domain.GetProductsResponse, error)
}

type Responder interface {
	Ok(w http.ResponseWriter, data any)
	Error(w http.ResponseWriter, status int, message string)
}

package catalogget

import (
	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"net/http"
)

type CatalogHandler struct {
	productGetter ProductGetter
}

func NewCatalogHandler(r ProductGetter) *CatalogHandler {
	return &CatalogHandler{
		productGetter: r,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	req := domain.NewGetProductsRequest(r.URL.Query())

	res, err := h.productGetter.Get(req)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, res)
}

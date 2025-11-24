package catalogget

import (
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"net/http"
)

type CatalogHandler struct {
	productGetter ProductGetter
	responder     Responder
}

func NewCatalogHandler(r ProductGetter, resp Responder) *CatalogHandler {
	return &CatalogHandler{
		productGetter: r,
		responder:     resp,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	req := domain.NewGetProductsRequest(r.URL.Query())

	res, err := h.productGetter.Get(req)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.responder.Ok(w, res)
}

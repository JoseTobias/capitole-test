package catalog

import (
	"encoding/json"
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
	// Return the products as a JSON response
	w.Header().Set("Content-Type", "application/json")

	req := domain.NewGetProductsRequest(r.URL.Query())

	res, err := h.productGetter.Get(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package getcatalogbyid

import (
	"github.com/mytheresa/go-hiring-challenge/app/api"
	"net/http"
)

type Handler struct {
	getter GetCatalogByCode
}

func NewCatalogHandler(r GetCatalogByCode) *Handler {
	return &Handler{
		getter: r,
	}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	res, err := h.getter.GetByCode(code)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, res)
}

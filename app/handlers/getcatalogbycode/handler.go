package getcatalogbycode

import (
	"errors"
	"github.com/mytheresa/go-hiring-challenge/app/api"
	"github.com/mytheresa/go-hiring-challenge/app/usecase/catalogbycode"
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
		if errors.Is(err, catalogbycode.ErrProductNotFound) {
			api.ErrorResponse(w, http.StatusNotFound, err.Error())
		}

		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.OKResponse(w, res)
}

package cataloggetbycode

import (
	"errors"
	"github.com/mytheresa/go-hiring-challenge/app/usecase/catalogbycode"
	"net/http"
)

type Handler struct {
	getter    GetCatalogByCode
	responder Responder
}

func NewCatalogHandler(r GetCatalogByCode, resp Responder) *Handler {
	return &Handler{
		getter:    r,
		responder: resp,
	}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	res, err := h.getter.GetByCode(code)
	if err != nil {
		if errors.Is(err, catalogbycode.ErrProductNotFound) {
			h.responder.Error(w, http.StatusNotFound, err.Error())
			return
		}

		h.responder.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.responder.Ok(w, res)
}

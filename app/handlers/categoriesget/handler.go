package categoriesget

import (
	"net/http"
)

type Handler struct {
	getter    CategoriesGetter
	responder Responder
}

func NewHandler(r CategoriesGetter, resp Responder) *Handler {
	return &Handler{
		getter:    r,
		responder: resp,
	}
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	res, err := h.getter.Get()
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.responder.Ok(w, res)
}

package categorycreate

import (
	"encoding/json"
	"github.com/mytheresa/go-hiring-challenge/app/domain"
	"net/http"
)

type Handler struct {
	creator   CategoriesCreator
	responder Responder
}

func NewHandler(r CategoriesCreator, resp Responder) *Handler {
	return &Handler{
		creator:   r,
		responder: resp,
	}
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var req *domain.CreateCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.responder.Error(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	defer r.Body.Close()

	if err = req.Validate(); err != nil {
		h.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.creator.Create(req)
	if err != nil {
		h.responder.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.responder.Ok(w, res)
}

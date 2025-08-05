package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) GetByIdOrder(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	order, err := h.OrdersService.GetById(id)
	return order, 200, err
}
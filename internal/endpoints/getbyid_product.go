package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) GetByIdProduct(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	product, err := h.ProductsService.GetById(id)
	return product, 200, err
}
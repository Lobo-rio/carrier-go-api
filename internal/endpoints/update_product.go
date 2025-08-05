package endpoints

import (
	contracts "carrierCheck/internal/contracts/products"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	var request contracts.UpdateProduct
	render.DecodeJSON(r.Body, &request)
	err := h.ProductsService.Update(id, request)
	return nil, 204, err
}
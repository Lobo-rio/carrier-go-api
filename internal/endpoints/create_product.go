package endpoints

import (
	contracts "carrierCheck/internal/contracts/products"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contracts.CreateProduct
	render.DecodeJSON(r.Body, &request)

	id, err := h.ProductsService.Create(request)

	return map[string]string{"id": id}, 201, err
}

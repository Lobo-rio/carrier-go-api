package endpoints

import (
	"net/http"
)

func (h *Handler) GetAllProduct(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	products, err := h.ProductsService.GetAll()
	return products, 200, err
}
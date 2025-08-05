package endpoints

import (
	"net/http"
)

func (h *Handler) GetAllOrder(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	orders, err := h.OrdersService.GetAll()
	return orders, 200, err
}
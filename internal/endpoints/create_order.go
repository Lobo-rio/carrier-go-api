package endpoints

import (
	contracts "carrierCheck/internal/contracts/order"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contracts.CreateOrder
	render.DecodeJSON(r.Body, &request)

	id, err := h.OrdersService.Create(request)

	return map[string]string{"id": id}, 201, err
}

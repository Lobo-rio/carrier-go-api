package endpoints

import (
	contracts "carrierCheck/internal/contracts/order"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) UpdateCarrierOrder(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	var request contracts.UpdateCarrierOrder
	render.DecodeJSON(r.Body, &request)
	err := h.OrdersService.UpdateCarrier(id, request)
	return nil, 204, err
}
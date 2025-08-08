package endpoints

import (
	contracts "carrierCheck/internal/contracts/order"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) UpdateStatusInTransit(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	var request contracts.UpdateOrder
	render.DecodeJSON(r.Body, &request)
	err := h.OrdersService.UpdateStatusInTransit(id)
	return nil, 200, err
}
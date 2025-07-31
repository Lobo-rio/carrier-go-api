package endpoints

import (
	contracts "carrierCheck/internal/contracts/carrier"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) UpdateCarrier(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	var request contracts.UpdateCarrier
	render.DecodeJSON(r.Body, &request)
	err := h.CarrierService.Update(id, request)
	return nil, 204, err
}
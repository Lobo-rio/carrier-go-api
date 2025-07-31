package endpoints

import (
	contracts "carrierCheck/internal/contracts/carrier"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateCarrier(w http.ResponseWriter, r *http.Request) (interface{}, int, error){
	var request contracts.CreateCarrier
	render.DecodeJSON(r.Body, &request)

	id, err := h.CarrierService.Create(request)

	return map[string]string{"id": id}, 201, err
}
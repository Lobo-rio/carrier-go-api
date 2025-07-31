package endpoints

import (
	"net/http"
)

func (h *Handler) GetAllCarrier(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	carriers, err := h.CarrierService.GetAll()
	return carriers, 200, err
}
package endpoints

import (
	"net/http"
)

func (h *Handler) CreateGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	carriers, err := h.CarrierService.Repository.GetAll()
	return carriers, 200, err
}
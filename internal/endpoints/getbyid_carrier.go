package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) GetByIdCarrier(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	carrier, err := h.CarrierService.GetById(id)
	return carrier, 200, err
}
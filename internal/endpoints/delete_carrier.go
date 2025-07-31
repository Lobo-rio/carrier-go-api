package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) DeleteCarrier(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CarrierService.Delete(id)
	return nil, 204, err
}
package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) GetByIdClient(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	client, err := h.ClientService.GetById(id)
	return client, 200, err
}
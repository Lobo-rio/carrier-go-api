package endpoints

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) DeleteClient(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.ClientService.Delete(id)
	return nil, 204, err
}
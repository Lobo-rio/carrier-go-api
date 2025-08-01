package endpoints

import (
	contracts "clientCheck/internal/contracts/client"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) UpdateClient(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	var request contracts.UpdateClient
	render.DecodeJSON(r.Body, &request)
	err := h.ClientService.Update(id, request)
	return nil, 204, err
}
package endpoints

import (
	contracts "carrierCheck/internal/contracts/clients"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CreateClient(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contracts.CreateClient
	render.DecodeJSON(r.Body, &request)

	id, err := h.ClientsService.Create(request)

	return map[string]string{"id": id}, 201, err
}

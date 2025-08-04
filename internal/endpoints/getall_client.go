package endpoints

import (
	"net/http"
)

func (h *Handler) GetAllClient(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	clients, err := h.ClientsService.GetAll()
	return clients, 200, err
}
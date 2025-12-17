package links

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListLinks(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "link_id is required", http.StatusBadRequest)
		return
	}

	Id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid link_id", http.StatusBadRequest)
		return
	}

	links, err := h.service.ListLinks(r.Context(), Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}

package links

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/yihune21/link-shortner/internal/database"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}


func (h *Handler) CreateLink(w http.ResponseWriter, r *http.Request) {
	type parameter struct{
		Link string `json:"link"`
	}
    
	decode := json.NewDecoder(r.Body)
	params := parameter{}
	err := decode.Decode(&params)
	
	if err != nil {
		http.Error(w, "link is required", http.StatusBadRequest)
		return
	}

	link, err := h.service.CreateLink(r.Context(), database.CreateLinkParams{
		ID: uuid.New(),
		Link: params.Link,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(link)
}

func (h *Handler) GetLink(w http.ResponseWriter, r *http.Request) {
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

	links, err := h.service.GetLink(r.Context(), Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}

func (h *Handler) ListLinks(w http.ResponseWriter, r *http.Request) {

	links, err := h.service.ListLinks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(links)
}
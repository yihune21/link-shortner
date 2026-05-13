package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yihune21/link-shortner/internal/service"
)


type UserHandler struct {
    us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us: us}
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"message": message})
}

func (h *UserHandler)CreateUser(w http.ResponseWriter , r *http.Request)()  {
	var req struct {
		Name    string `json:"name"`
		Email  string `json:"email"`
		Password string `jsom:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

}

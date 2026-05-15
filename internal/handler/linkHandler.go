package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yihune21/link-shortner/internal/database"
	"github.com/yihune21/link-shortner/internal/service"
)


type LinkHandler struct {
    ls *service.LinkService
}

func NewLinkHandler(ls *service.LinkService) *LinkHandler {
	return &LinkHandler{ls: ls}
}

func (h *LinkHandler)CreateLink(w http.ResponseWriter , r *http.Request , user *database.User)()  {
	var req struct {
		OriginalLink    string `json:"original_link"`
	} 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}
	dbLink , err := h.ls.CreateLink(r.Context() , req.OriginalLink , user)
    if err != nil {
		WriteError(w , 401 , err.Error())
	}
    
	WriteJSON(w , 200 , dbLink)
}

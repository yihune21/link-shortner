package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/yihune21/link-shortner/internal/service"
)


type LinkHandler struct {
    ls *service.LinkService
}

func NewLinkHandler(ls *service.LinkService) *LinkHandler {
	return &LinkHandler{ls: ls}
}

func (h *LinkHandler)CreateLink(w http.ResponseWriter , r *http.Request)()  {
	var req struct {
		OriginalLink    string `json:"original_link"`
	} 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}
    
    idStr := chi.URLParam(r,"id")
	if idStr == "" {
		WriteError(w ,400 , "Link id is required.")
		return
	}
	id, err :=uuid.Parse(idStr)
    if err != nil {
		WriteError(w ,400 , "Couldn't parse link id.")
		return
	}

	dbLink , err := h.ls.CreateLink(r.Context() , req.OriginalLink ,id)
    if err != nil {
		WriteError(w , 401 , err.Error())
		return
	}
    
	WriteJSON(w , 200 , dbLink)
}
func (h *LinkHandler)ListLinks(w http.ResponseWriter , r *http.Request )  {
	links , err := h.ls.ListLink(r.Context())
    if err != nil {
		WriteError(w , 404 , err.Error())
		return
	}

	WriteJSON(w , 200 , links)
}
func (h *LinkHandler)GetLinkById(w http.ResponseWriter , r *http.Request )  {
	idStr := chi.URLParam(r,"id")
	if idStr == "" {
		WriteError(w ,400 , "Link id is required.")
		return
	}
	id, err :=uuid.Parse(idStr)
    if err != nil {
		WriteError(w ,400 , "Couldn't parse link id.")
		return
	}
	link , err := h.ls.GetLinkById(r.Context() , id)
	if err != nil{
		WriteError(w ,404 , "Link not found.")
		return
	}

	WriteJSON(w , 200 , link)
}
func (h *LinkHandler)GetLinksByShortLink(w http.ResponseWriter , r *http.Request) {
	shortId := chi.URLParam(r , "shortId")
	link , err := h.ls.GetLinksByShortLink(r.Context(),shortId)
	if err != nil {
		WriteError(w , 404 , "Short link not found")
		return
	}
	http.Redirect(w, r, link.OriginalLink, http.StatusFound)
}
func (h *LinkHandler)GetLinksByUserId(w http.ResponseWriter , r *http.Request)  {
    idStr := chi.URLParam(r,"id")
	if idStr == "" {
		WriteError(w ,400 , "Link id is required.")
		return
	}
	id, err :=uuid.Parse(idStr)
    if err != nil {
		WriteError(w ,400 , "Couldn't parse link id.")
		return
	}
	links,err:= h.ls.GetLinksByUserId(r.Context(),id)
	if err != nil{
		WriteError(w ,404 , "Link not found.")
		return
        }
	WriteJSON(w,200,links)
}
func (h *LinkHandler)DeleteLink(w http.ResponseWriter , r *http.Request)  {
	idStr :=chi.URLParam(r,"id")
	if idStr == "" {
		WriteError(w ,400 , "Link id is required.")
		return
	}
	id, err :=uuid.Parse(idStr)
    if err != nil {
		WriteError(w ,400 , "Link id is required.")
		return
	}
	err = h.ls.DeleteLink(r.Context(),id)
	if err != nil{
		WriteError(w,400,err.Error())
		return
	}
	WriteJSON(w ,200 ,nil)
}

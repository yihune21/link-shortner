package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/yihune21/link-shortner/internal/database"
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
	user , err :=h.us.CreateUser(r.Context() , req.Name , req.Email , req.Password)
	if err != nil {
		WriteError(w,400,err.Error())
		return
	}
	WriteJSON(w,200,user)

}
func (h *UserHandler)ListUser(w http.ResponseWriter , r *http.Request)  {
	users,err := h.us.ListUser(r.Context())
	if err != nil {
		WriteError(w,400,err.Error())
		return
	}
	WriteJSON(w,200,users)
}
func (h *UserHandler)GetUserById(w http.ResponseWriter , r *http.Request)  {
	idStr := chi.URLParam(r,"id")
	if idStr == "" {
		WriteError(w ,400 , "User id is required.")
		return
	}
	id, err :=uuid.Parse(idStr)
    if err != nil {
		WriteError(w ,400 , "Couldn't parse user id.")
		return
	}
	user , err := h.us.GetUserById(r.Context(),id)
	if err != nil {
		WriteError(w ,400 , "User id is required.")
		return
	}
	WriteJSON(w,200,user)
}
func (h *UserHandler)GetUserByEmail(w http.ResponseWriter , r *http.Request)  {
	var req struct {
		Email  string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}
	user , err :=h.us.GetUserByEmail(r.Context(),req.Email)
	if err !=nil {
	    WriteError(w,400,err.Error())
		return
	}

	WriteJSON(w,200,user)
}

func (h *UserHandler)UpdateUserName(w http.ResponseWriter , r *http.Request,user database.User)  {
	var req struct {
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

	dbUser , err :=h.us.UpdateUserName(r.Context(),req.Name , user)
	if err !=nil {
	    WriteError(w,400,err.Error())
		return
	}

	WriteJSON(w,200,dbUser)
}


func (h *UserHandler)DeleteUser(w http.ResponseWriter , r *http.Request)  {
	idStr := chi.URLParam(r,"id")
	if idStr == "" {
		WriteError(w ,400 , "User id is required.")
		return 
	}
	id, err :=uuid.Parse(idStr)
    if err != nil {
		WriteError(w ,400 , "Couldn't parse user id.")
		return
	}
    
	err = h.us.DeleteUser(r.Context(),id)

	if err != nil {
		WriteError(w,400,err.Error())
		return
	}

	WriteJSON(w,200 , nil)



}
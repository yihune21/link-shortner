package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
type CreateUserRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// CreateUser creates a new user
//
// @Summary Create a user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} map[string]interface{}
// @Router /v1/users [post]
func (h *UserHandler)CreateUser(w http.ResponseWriter , r *http.Request)  {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}
    
	err := IsValidPassword(req.Password)
    if err != nil {
		WriteError(w,302 , err.Error())
		return
	}
    
	err = IsValidEmail(req.Email)
    if err != nil {
		WriteError(w,302 , err.Error())
		return
	}

	hashedPass , err := HashPassword(req.Password) 
	if err != nil {
		WriteError(w, 302 , err.Error())
		return
	}
	user , err :=h.us.CreateUser(r.Context() , req.Name , req.Email , hashedPass)
	if err != nil {
		WriteError(w,400,err.Error())
		return
	}

	accessToken, refreshToken, err := h.us.LoginTokens(r.Context(), user)
	if err != nil {
		WriteError(w, 500, err.Error())
		return
	}
     
	WriteJSON(w, 200, map[string]interface{}{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
// ListUser gets all users
// @Summary List all users
// @Description Returns a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /v1/users [get]
func (h *UserHandler)ListUser(w http.ResponseWriter , r *http.Request)  {
	users,err := h.us.ListUser(r.Context())
	if err != nil {
		WriteError(w,400,err.Error())
		return
	}
	WriteJSON(w,200,users)
}
// GetUserById gets a user by ID
// @Summary Get user by ID
// @Description Returns a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /v1/users/{id} [get]
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
func (h *UserHandler)UpdateUserName(w http.ResponseWriter , r *http.Request)  {
	var req struct {
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

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

	dbUser , err :=h.us.UpdateUserName(r.Context(),req.Name , id)
	if err !=nil {
	    WriteError(w,400,err.Error())
		return
	}

	WriteJSON(w,200,dbUser)
}
func (h *UserHandler)UpdateUserPassword(w http.ResponseWriter , r *http.Request)  {
	var req struct {
		Password  string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

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

	dbUser , err :=h.us.UpdateUserPassword(r.Context(),req.Password , id)
	if err !=nil {
	    WriteError(w,400,err.Error())
		return
	}

	WriteJSON(w,200,dbUser)

}
// DeleteUser deletes a user
// @Summary Delete user
// @Description Deletes a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Router /v1/users/{id} [delete]
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
// Login authenticates a user
// @Summary User login
// @Description Authenticates user and returns tokens
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Router /v1/users/login [post]
func (h *UserHandler)Login(w http.ResponseWriter , r *http.Request)  {
	var req struct {
		Email  string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}
    
	user, err := h.us.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		WriteError(w, 400, "invalid email or password")
		return
	}

	if !VerifyPassword(req.Password, user.Password) {
		WriteError(w, 400, "invalid email or password")
		return
	}

	accessToken, refreshToken, err := h.us.LoginTokens(r.Context(), user)
	if err != nil {
		WriteError(w, 500, err.Error())
		return
	}

	WriteJSON(w, 200, map[string]interface{}{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
// Logout logs out a user
// @Summary User logout
// @Description Logs out the current user
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Security BearerAuth
// @Router /v1/logout/{id} [post]
func (h *UserHandler)Logout(w http.ResponseWriter , r *http.Request)  {
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
	err = h.us.Logout(r.Context() ,id)
	if err != nil {
		WriteError(w,400,err.Error())
		return
	}
	WriteJSON(w,200,nil)
}
func (h *UserHandler)SendOtp(w http.ResponseWriter , r *http.Request)  {
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
	dbUser ,err := h.us.GetUserById(r.Context(),id)
    user_email := dbUser.Email
	otp := generateSecureOTP(6)
    err  = h.us.CreateOtp(r.Context() , otp, id)
	if err != nil {
		WriteError(w , 400, err.Error())
		return
	}
	SendEmail(user_email , otp)
}
func (h *UserHandler)VerifyOtp(w http.ResponseWriter , r *http.Request)  {
	var req struct {
		Otp    string `json:"otp"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

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

	otp ,err := h.us.GetOtpByUserId(r.Context(),id)
    if err != nil {
		WriteError(w,400 , err.Error())
		return
	}
    if time.Now().After(otp.ExpAt) {
		WriteError(w, 400, "OTP has expired")
		return
	}

    isOtpMatched := VerifyOTP(otp.Otp,req.Otp)
    if !isOtpMatched {
		WriteError(w,400 , "wrong otp.")
		return
	}
	user , err := h.us.VerifyUser(r.Context() ,id)
	if err != nil {
		WriteError(w,400,err.Error())
		return
	}
	_=h.us.DeleteOtp(r.Context() , id)
	WriteJSON(w,200, user)
}
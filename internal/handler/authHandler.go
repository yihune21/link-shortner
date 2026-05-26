package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yihune21/link-shortner/internal/database"
	jsonwebtoken "github.com/yihune21/link-shortner/internal/jsonWebToken"
	"github.com/yihune21/link-shortner/internal/service"
)

type AuthHandler struct{
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{as:as}
}

func (a *AuthHandler) MiddlewareAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        accessToken, err := a.as.GetToken(r.Header)
        if err != nil {
            WriteError(w, 401, err.Error())
            return
        }

        if !jsonwebtoken.VerifyToken(accessToken) {
            WriteError(w, 401, "access token expired")
            return
        }

        userID, err := jsonwebtoken.ExtractUserIDFromToken(accessToken)
        if err != nil {
            WriteError(w, 400, err.Error())
            return
        }

        _, err = a.as.GetUserById(r.Context(), userID)
        if err != nil {
            WriteError(w, 404, "user not found")
            return
        }

        next.ServeHTTP(w, r)
    })
}

func (a *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	userID, err := jsonwebtoken.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	dbToken, err := a.as.GetRefreshToken(r.Context(), userID)
	if err != nil || dbToken.Token != req.RefreshToken {
		WriteError(w, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	user, err := a.as.GetUserById(r.Context(), userID)
	if err != nil {
		WriteError(w, http.StatusUnauthorized, "user not found")
		return
	}

	dbUser := database.User{
		ID:   user.ID,
	}

	newAccessToken := jsonwebtoken.GenerateAccessToken(dbUser)

	WriteJSON(w, http.StatusOK, map[string]string{
		"access_token": newAccessToken,
	})
}
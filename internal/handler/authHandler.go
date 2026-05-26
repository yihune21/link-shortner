package handler

import (
	"net/http"

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
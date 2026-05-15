package service

import (
	"errors"
	"net/http"
	"strings"

	"github.com/yihune21/link-shortner/internal/database"
)

type AuthService struct{
	q *database.Queries
}

func NewAuthService(q *database.Queries) *AuthService  {
	return &AuthService{q:q}
}

func (s *AuthService)GetToken(headers http.Header) (string , error)  {
	val := headers.Get("Authorization")

	if val == ""{
		return "",errors.New("no auth header found")
	}
	vals := strings.Split(val, " ")
	if len(vals) !=2 {
		return " ",errors.New("malformed auth header")
	}
	if vals[0] != "Bearer"{
		return "",errors.New("malformed first part auth of header found")
	}
	return vals[1],nil
}


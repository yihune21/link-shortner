package handler

import (
	"github.com/yihune21/link-shortner/internal/service"
)


type UserHandler struct {
    us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{us: us}
}

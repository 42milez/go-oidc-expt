package model

import (
	"github.com/42milez/go-oidc-server/app/typedef"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	ID   typedef.UserID
	Name string
}

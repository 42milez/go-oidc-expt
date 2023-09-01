package model

import (
	"github.com/42milez/go-oidc-server/app/typedef"
)

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserResponse struct {
	ID   typedef.UserID `json:"id" validate:"required"`
	Name string         `json:"name" validate:"required"`
}

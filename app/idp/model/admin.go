package model

import (
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

type AdminCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminResponse struct {
	ID   typedef.AdminID
	Name string
}

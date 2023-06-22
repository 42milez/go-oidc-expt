package model

import "github.com/42milez/go-oidc-server/app/idp/ent/alias"

type AdminResponse struct {
	ID   alias.AdminID
	Name string
}

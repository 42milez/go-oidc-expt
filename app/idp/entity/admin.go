package entity

import (
	"github.com/42milez/go-oidc-server/app/idp/ent/alias"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type Admin *ent.Admin
type Admins []*ent.Admin

type AdminResponse struct {
	ID   alias.AdminID
	Name string
}

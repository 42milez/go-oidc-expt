package entity

import "github.com/42milez/go-oidc-server/src/ent/ent"

type User *ent.User
type Users []*ent.User

type UserResponse struct {
	ID   int
	Name string
}

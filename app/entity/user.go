package entity

import "github.com/42milez/go-oidc-server/app/ent/ent"

func NewUser(entity *ent.User) *User {
	return &User{
		entity: entity,
	}
}

type User struct {
	entity *ent.User
}

func (u *User) Name() string {
	return u.entity.Name
}

func (u *User) Password() string {
	return u.entity.Password
}

package entity

import "github.com/42milez/go-oidc-server/pkg/typedef"

type UserInfo struct {
	ID   typedef.UserID `json:"id"`
	Name string         `json:"name"`
}

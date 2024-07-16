package entity

import "github.com/42milez/go-oidc-expt/pkg/typedef"

type UserInfo struct {
	ID   typedef.UserID `json:"id"`
	Name string         `json:"name"`
}

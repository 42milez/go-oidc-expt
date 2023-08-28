package repository

import "github.com/42milez/go-oidc-server/app/ent/typedef"

type IDGenerator interface {
	NextID() (typedef.UserID, error)
}

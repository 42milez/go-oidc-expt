package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

type UserCreator interface {
	Create(ctx context.Context, name string, pw typedef.PasswordHash) (*ent.User, error)
}

type UserSelector interface {
	SelectByName(ctx context.Context, name string) (*ent.User, error)
}

type TokenGenerator interface {
	GenerateAccessToken(name string) ([]byte, error)
}

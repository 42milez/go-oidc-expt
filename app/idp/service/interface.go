package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

type Identity interface {
	ent.Admin
}

type IdentityCreator[T Identity] interface {
	Create(ctx context.Context, name string, pw typedef.PasswordHash) (*T, error)
}

type IdentitySelector[T Identity] interface {
	SelectByName(ctx context.Context, name string) (*T, error)
}

type TokenGenerator interface {
	GenerateAccessToken(name string) ([]byte, error)
}

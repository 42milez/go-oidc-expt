package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type Identity interface {
	ent.Admin
}

type IdentityCreator[T Identity] interface {
	Create(ctx context.Context, name, pw string) (*T, error)
}

type IdentitySelector[T Identity] interface {
	SelectByName(ctx context.Context, name string) (*T, error)
}

type TokenGenerator interface {
	GenerateAccessToken(name string) ([]byte, error)
}

package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type AdminCreater interface {
	Create(ctx context.Context) error
}

type AdminReader interface {
	SelectAdmin(ctx context.Context, db *ent.Client) (*ent.Admin, error)
}

type AdminsReader interface {
	ReadAdmins(ctx context.Context, db *ent.Client) ([]*ent.Admin, error)
}

type Identity interface {
	ent.Admin
}

type IdentitySelector[T Identity] interface {
	SelectByName(ctx context.Context, name string) (*T, error)
}

type TokenGenerator interface {
	GenerateAccessToken(name string) ([]byte, error)
}

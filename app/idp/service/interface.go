package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/entity"
)

type AdminCreater interface {
	Create(ctx context.Context) error
}

type AdminReader interface {
	ReadAdmin(ctx context.Context, db *ent.Client) (entity.Admin, error)
}

type AdminsReader interface {
	ReadAdmins(ctx context.Context, db *ent.Client) (entity.Admins, error)
}

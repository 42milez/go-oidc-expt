package handler

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/entity"
)

type CheckHealthService interface {
	PingCache(ctx context.Context) error
	PingDB(ctx context.Context) error
}

type SignInService interface {
	SignIn(ctx context.Context, name, pw string) (string, error)
}

type ReadAdminService interface {
	ReadAdmin(ctx context.Context) (entity.Admin, error)
}

type ReadAdminsService interface {
	ReadAdmins(ctx context.Context) (entity.Admins, error)
}

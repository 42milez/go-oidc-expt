package handler

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/entity"
)

type CheckHealthService interface {
	PingCache(ctx context.Context) error
	PingDB(ctx context.Context) error
}

type ReadUserService interface {
	ReadUser(ctx context.Context) (entity.User, error)
}

type ReadUsersService interface {
	ReadUsers(ctx context.Context) (entity.Users, error)
}

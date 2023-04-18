package handler

import (
	"context"

	"github.com/42milez/go-oidc-server/src/entity"
)

type CheckHealthService interface {
	PingCache(ctx context.Context) error
	PingDB(ctx context.Context) error
}

type ReadUserService interface {
	Read(ctx context.Context) (entity.User, error)
}

type ReadUsersService interface {
	ReadBulk(ctx context.Context) (entity.Users, error)
}

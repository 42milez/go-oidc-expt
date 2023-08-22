package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type HealthChecker interface {
	PingCache(ctx context.Context) error
	PingDB(ctx context.Context) error
}

type TokenGenerator interface {
	GenerateAccessToken(name string) ([]byte, error)
}

type UserCreator interface {
	Create(ctx context.Context, name string, pw typedef.PasswordHash) (*ent.User, error)
}

type UserSelector interface {
	SelectByName(ctx context.Context, name string) (*ent.User, error)
}

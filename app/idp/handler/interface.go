package handler

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type HealthChecker interface {
	PingCache(ctx context.Context) error
	PingDB(ctx context.Context) error
}

type Authorizer interface {
	Authorize(ctx context.Context) error
}

type Authenticator interface {
	Authenticate(ctx context.Context, name, pw string) (string, error)
}

type UserCreator interface {
	CreateUser(ctx context.Context, name, pw string) (*ent.User, error)
}

type UserSelector interface {
	SelectUser(ctx context.Context) (*ent.User, error)
}

type SessionManager interface {
	SaveID(ctx context.Context)
	LoadID(ctx context.Context)
}

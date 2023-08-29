package handler

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/model"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/typedef"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type CacheStatusChecker interface {
	CheckCacheStatus(ctx context.Context) error
}

type DBStatusChecker interface {
	CheckDBStatus(ctx context.Context) error
}

type HealthChecker interface {
	CacheStatusChecker
	DBStatusChecker
}

type Authenticator interface {
	Authenticate(ctx context.Context, name, pw string) (typedef.UserID, error)
}

type Authorizer interface {
	Authorize(ctx context.Context, userID typedef.UserID, param *model.AuthorizeRequest) (string, error)
}

type SessionCreator interface {
	Create(ctx context.Context, sess *entity.UserSession) (string, error)
}

type SessionRestorer interface {
	Restore(r *http.Request, sessionID string) (*http.Request, error)
}

type UserCreator interface {
	CreateUser(ctx context.Context, name, pw string) (*ent.User, error)
}

type UserSelector interface {
	SelectUser(ctx context.Context) (*ent.User, error)
}

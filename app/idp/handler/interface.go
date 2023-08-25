package handler

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/entity"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

	"github.com/42milez/go-oidc-server/app/idp/model"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type HealthChecker interface {
	CheckCacheStatus(ctx context.Context) error
	CheckDBStatus(ctx context.Context) error
}

type Authorizer interface {
	Authorize(ctx context.Context, param *model.AuthorizeRequest) (string, error)
}

type Authenticator interface {
	Authenticate(ctx context.Context, name, pw string) (typedef.UserID, error)
}

type SessionCreator interface {
	Create(item *entity.UserSession) (string, error)
}

type SessionRestorer interface {
	Restore(r *http.Request) (*http.Request, error)
}

type UserCreator interface {
	CreateUser(ctx context.Context, name, pw string) (*ent.User, error)
}

type UserSelector interface {
	SelectUser(ctx context.Context) (*ent.User, error)
}

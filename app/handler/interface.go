package handler

import (
	"context"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/typedef"
	"github.com/42milez/go-oidc-server/app/model"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=$GOPACKAGE

type HealthChecker interface {
	CheckCacheStatus(ctx context.Context) error
	CheckDBStatus(ctx context.Context) error
}

type Authorizer interface {
	Authorize(ctx context.Context, userID typedef.UserID, param *model.AuthorizeRequest) (string, error)
}

type Authenticator interface {
	Authenticate(ctx context.Context, name, pw string) (typedef.UserID, error)
}

type SessionCreator interface {
	Create(item *UserSession) (string, error)
}

type SessionRestorer interface {
	Restore(r *http.Request) (*http.Request, error)
}

type TokenExtractor interface {
	ExtractToken(r *http.Request) (jwt.Token, error)
}

type UserCreator interface {
	CreateUser(ctx context.Context, name, pw string) (*ent.User, error)
}

type UserSelector interface {
	SelectUser(ctx context.Context) (*ent.User, error)
}

// TODO: Separate methods ( SessionManager )

type SessionManager interface {
	SaveUserID(ctx context.Context, key string, id typedef.UserID) error
	LoadUserID(ctx context.Context, key string) (typedef.UserID, error)
}

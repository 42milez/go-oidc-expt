package xutil

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/lestrrat-go/jwx/v2/jwt"
)

type SessionManager interface {
	SaveUserID(ctx context.Context, key string, id typedef.UserID) error
	LoadUserID(ctx context.Context, key string) (typedef.UserID, error)
}

type TokenExtractor interface {
	ExtractToken(r *http.Request) (jwt.Token, error)
}

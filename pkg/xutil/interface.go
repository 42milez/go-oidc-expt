package xutil

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type SessionManager interface {
	SaveID(ctx context.Context, key string, id typedef.UserID) error
	LoadID(ctx context.Context, key string) (typedef.UserID, error)
}

type TokenExtractor interface {
	ExtractToken(r *http.Request) (jwt.Token, error)
}

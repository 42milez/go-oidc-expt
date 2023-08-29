package session

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type TokenExtractor interface {
	ExtractToken(r *http.Request) (jwt.Token, error)
}

type Reader interface {
	Read(ctx context.Context, key string) (*entity.UserSession, error)
}

type Writer interface {
	Write(ctx context.Context, key string, sess *entity.UserSession) (bool, error)
}

type ReadWriter interface {
	Reader
	Writer
}

package session

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type TokenExtractor interface {
	Extract(r *http.Request) (jwt.Token, error)
}

type Creator interface {
	Create(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (bool, error)
}

type Reader interface {
	Read(ctx context.Context, sid typedef.SessionID) (*entity.Session, error)
}

type Updater interface {
	Update(ctx context.Context, sid typedef.SessionID, sess *entity.Session) (string, error)
}

type ReadUpdateWriter interface {
	Creator
	Reader
	Updater
}

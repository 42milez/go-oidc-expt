package session

import (
	"context"
	"github.com/42milez/go-oidc-server/app/idp/jwt"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/redis/go-redis/v9"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const (
	ErrFailedToExtractToken xerr.Err = "failed to extract token"
)

func NewUtil(redisClient *redis.Client, jwtUtil *jwt.Util) *Util {
	return &Util{
		Repo: &repository.Session{
			Cache: redisClient,
		},
		Token: jwtUtil,
	}
}

type IDKey struct{}

type Util struct {
	Repo  xutil.SessionManager
	Token xutil.TokenExtractor
}

func (p *Util) FillContext(r *http.Request) (*http.Request, error) {
	token, err := p.Token.ExtractToken(r)

	if err != nil {
		return nil, xerr.Wrap(ErrFailedToExtractToken, err)
	}

	id, err := p.Repo.LoadUserID(r.Context(), token.JwtID())

	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(r.Context(), IDKey{}, id)

	return r.Clone(ctx), nil
}

func GetUserID(ctx context.Context) (typedef.UserID, bool) {
	id, ok := ctx.Value(IDKey{}).(typedef.UserID)
	return id, ok
}

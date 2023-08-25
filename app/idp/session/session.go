package session

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/jwt"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

type Session struct {
	repo  xutil.SessionManager
	token xutil.TokenExtractor
}

type IDKey struct{}

func (p *Session) Create(item *entity.UserSession) (string, error) {
	ret, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	return ret.String(), nil
}

func (p *Session) Restore(r *http.Request) (*http.Request, error) {
	token, err := p.token.ExtractToken(r)

	if err != nil {
		return nil, xerr.FailedToExtractToken.Wrap(err)
	}

	id, err := p.repo.LoadUserID(r.Context(), token.JwtID())

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

func NewSession(redisClient *redis.Client, jwtUtil *jwt.Util) *Session {
	return &Session{
		repo: &repository.Session{
			Cache: redisClient,
		},
		token: jwtUtil,
	}
}

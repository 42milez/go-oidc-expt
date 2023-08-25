package handler

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/entity"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/42milez/go-oidc-server/pkg/xutil"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

type Session struct {
	repo  xutil.SessionManager
	token xutil.TokenExtractor
}

type UserIDKey struct{}

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

	ctx := context.WithValue(r.Context(), UserIDKey{}, id)

	return r.Clone(ctx), nil
}

func NewSession(redisClient *redis.Client, jwtUtil *auth.Util) *Session {
	return &Session{
		repo: &repository.Session{
			Cache: redisClient,
		},
		token: jwtUtil,
	}
}

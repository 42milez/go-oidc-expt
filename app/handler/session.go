package handler

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/app/auth"
	"github.com/42milez/go-oidc-server/app/repository"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/42milez/go-oidc-server/pkg/xerr"
)

type UserIDKey struct{}

type Session struct {
	repo  SessionManager
	token TokenExtractor
}

func (p *Session) Create(item *UserSession) (string, error) {
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

type UserSession struct {
	ID typedef.UserID
}

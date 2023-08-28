package handler

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xerr"
	"github.com/google/uuid"
)

type UserIDKey struct{}

type Session struct {
	Repo     SessionManager
	TokenExt TokenExtractor
}

func (p *Session) Create(item *UserSession) (string, error) {
	ret, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	return ret.String(), nil
}

func (p *Session) Restore(r *http.Request) (*http.Request, error) {
	token, err := p.TokenExt.ExtractToken(r)

	if err != nil {
		return nil, xerr.FailedToExtractToken.Wrap(err)
	}

	id, err := p.Repo.LoadUserID(r.Context(), token.JwtID())

	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(r.Context(), UserIDKey{}, id)

	return r.Clone(ctx), nil
}

type UserSession struct {
	ID typedef.UserID
}

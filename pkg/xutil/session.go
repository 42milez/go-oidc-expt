package xutil

import (
	"context"
	"net/http"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"
	"github.com/42milez/go-oidc-server/pkg/xerr"
)

const (
	ErrFailedToExtractToken xerr.Err = "failed to extract token"
)

type IDKey struct{}

type Session struct {
	Repo  SessionManager
	Token TokenExtractor
}

func (p *Session) GetUserID(ctx context.Context) (typedef.UserID, bool) {
	id, ok := ctx.Value(IDKey{}).(typedef.UserID)
	return id, ok
}

func (p *Session) FillContext(r *http.Request) (*http.Request, error) {
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

package service

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

func (p *Session) SetID(ctx context.Context, id typedef.UserID) context.Context {
	return context.WithValue(ctx, IDKey{}, id)
}

func (p *Session) GetID(ctx context.Context) (typedef.UserID, bool) {
	id, ok := ctx.Value(IDKey{}).(typedef.UserID)
	return id, ok
}

func (p *Session) FillContext(r *http.Request) (*http.Request, error) {
	token, err := p.Token.ExtractToken(r)

	if err != nil {
		return nil, xerr.Wrap(ErrFailedToExtractToken, err)
	}

	id, err := p.Repo.LoadID(r.Context(), token.JwtID())

	if err != nil {
		return nil, err
	}

	ctx := p.SetID(r.Context(), id)

	return r.Clone(ctx), nil
}

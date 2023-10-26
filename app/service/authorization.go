package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/42milez/go-oidc-server/app/httpstore"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/entity"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"github.com/42milez/go-oidc-server/app/config"
	"golang.org/x/exp/slices"
)

func NewAuthorize(repo Authorizer) *Authorize {
	return &Authorize{
		repo: repo,
		rCtx: &httpstore.ReadContext{},
	}
}

type Authorize struct {
	repo Authorizer
	rCtx ContextReader
}

func (a *Authorize) Authorize(ctx context.Context, clientID, redirectURI, state string) (string, error) {
	code, err := xrandom.MakeCryptoRandomString(config.AuthCodeLength)
	if err != nil {
		return "", err
	}

	uid, ok := a.rCtx.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		return "", xerr.SessionNotFound
	}

	if _, err = a.repo.CreateAuthCode(ctx, code, clientID, uid); err != nil {
		return "", err
	}

	ru, err := a.repo.ReadRedirectUriByClientID(ctx, clientID)
	if err != nil {
		return "", err
	}

	if !a.validateRedirectUri(ru, redirectURI) {
		return "", errors.New("invalid redirect uri")
	}

	return fmt.Sprintf("%s?code=%s&state=%s", redirectURI, code, state), nil
}

func (a *Authorize) validateRedirectUri(s []*entity.RedirectUri, v string) bool {
	return slices.ContainsFunc(s, func(uri *entity.RedirectUri) bool {
		if uri.URI() != v {
			return false
		}
		return true
	})
}

type AuthorizeParams struct {
	RedirectUri string
	State       string
}

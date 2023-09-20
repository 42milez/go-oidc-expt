package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"golang.org/x/exp/slices"
)

type Authorize struct {
	repo Authorizer
}

func (a *Authorize) Authorize(ctx context.Context, clientID, redirectURI, state string) (string, error) {
	code, err := xrandom.MakeCryptoRandomString(config.AuthCodeLength)

	if err != nil {
		return "", err
	}

	sess, ok := GetSession(ctx)

	if !ok {
		return "", xerr.SessionNotFound
	}

	if _, err = a.repo.CreateAuthCode(ctx, code, clientID, sess.UserID); err != nil {
		return "", err
	}

	ru, err := a.repo.ReadRedirectUriByClientID(ctx, clientID)

	if err != nil {
		return "", err
	}

	if !a.validateRedirectURI(ru, redirectURI) {
		return "", errors.New("invalid redirect uri")
	}

	return fmt.Sprintf("%s?code=%s&state=%s", redirectURI, code, state), nil
}

func (a *Authorize) validateRedirectURI(s []*ent.RedirectURI, v string) bool {
	return slices.ContainsFunc(s, func(uri *ent.RedirectURI) bool {
		if uri.URI != v {
			return false
		}
		return true
	})
}

type AuthorizeParams struct {
	RedirectUri string
	State       string
}

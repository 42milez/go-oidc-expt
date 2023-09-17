package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"golang.org/x/exp/slices"
)

type Authorize struct {
	repo Authorizer
}

func (a *Authorize) Authorize(ctx context.Context, userID typedef.UserID, param *AuthorizeParams) (string, error) {
	code, err := xrandom.MakeCryptoRandomString(config.AuthCodeLength)

	if err != nil {
		return "", err
	}

	if _, err = a.repo.CreateAuthCode(ctx, userID, code); err != nil {
		return "", err
	}

	ru, err := a.repo.ReadRedirectUriByUserID(ctx, userID)

	if err != nil {
		return "", err
	}

	if !a.validateRedirectURI(ru, param.RedirectUri) {
		return "", errors.New("invalid redirect uri")
	}

	return fmt.Sprintf("%s?code=%s&state=%s", param.RedirectUri, code, param.State), nil
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

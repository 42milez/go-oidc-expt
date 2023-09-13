package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/model"
	"github.com/42milez/go-oidc-server/app/repository"
	"golang.org/x/exp/slices"
)

type Authorize struct {
	Repo *repository.User
}

func (a *Authorize) Authorize(ctx context.Context, userID typedef.UserID, param *model.AuthorizeRequest) (string, error) {
	code, err := xrandom.MakeCryptoRandomString(config.AuthCodeLength)

	if err != nil {
		return "", err
	}

	if _, err = a.Repo.CreateAuthorizationCode(ctx, userID, code); err != nil {
		return "", err
	}

	ru, err := a.Repo.ReadRedirectUriByUserID(ctx, userID)

	if err != nil {
		return "", err
	}

	if !validateRedirectURI(ru, param.RedirectURI) {
		return "", errors.New("invalid redirect uri")
	}

	return fmt.Sprintf("%s?code=%s&state=%s", param.RedirectURI, code, param.State), nil
}

func validateRedirectURI(s []*ent.RedirectURI, v string) bool {
	return slices.ContainsFunc(s, func(uri *ent.RedirectURI) bool {
		if uri.URI != v {
			return false
		}
		return true
	})
}

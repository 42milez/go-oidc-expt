package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"golang.org/x/exp/slices"

	"github.com/42milez/go-oidc-server/app/idp/repository"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/model"
)

const authCodeLen = 20

type Authorize struct {
	Repo *repository.User
}

func (p *Authorize) Authorize(ctx context.Context, param *model.AuthorizeRequest) (string, error) {
	code, err := xutil.MakeCryptoRandomString(authCodeLen)

	if err != nil {
		return "", err
	}

	userID, ok := xutil.GetUserID(ctx)

	if !ok {
		return "", errors.New("user id not found")
	}

	if _, err = p.Repo.SaveAuthorizationCode(ctx, userID, code); err != nil {
		return "", err
	}

	ru, err := p.Repo.SelectRedirectURIByUserID(ctx, userID)

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

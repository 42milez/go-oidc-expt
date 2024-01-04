package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/httpstore"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"golang.org/x/exp/slices"
)

func NewAuthorize(opt *option.Option) *Authorize {
	return &Authorize{
		repo:    repository.NewRelyingParty(opt.DB),
		cache:   httpstore.NewCache(opt),
		context: &httpstore.Context{},
	}
}

type Authorize struct {
	repo    Authorizer
	cache   iface.OpenIdParamWriter
	context iface.ContextReader
}

func (a *Authorize) Authorize(ctx context.Context, clientID, redirectURI, state string) (*url.URL, string, error) {
	code, err := xrandom.GenerateCryptoRandomString(config.AuthCodeLength)
	if err != nil {
		return nil, "", err
	}

	sess, ok := a.context.Read(ctx, httpstore.SessionKey{}).(*httpstore.Session)
	if !ok {
		return nil, "", xerr.UnauthorizedRequest
	}

	if _, err = a.repo.CreateAuthCode(ctx, code, clientID, sess.UserID); err != nil {
		return nil, "", err
	}

	ru, err := a.repo.ReadRedirectUris(ctx, clientID)
	if err != nil {
		return nil, "", err
	}

	if !a.validateRedirectUri(ru, redirectURI) {
		return nil, "", xerr.InvalidRedirectURI
	}

	uri, err := url.Parse(fmt.Sprintf("%s?code=%s&state=%s", redirectURI, code, state))
	if err != nil {
		return nil, "", err
	}

	return uri, code, nil
}

func (a *Authorize) validateRedirectUri(s []*entity.RedirectUri, v string) bool {
	return slices.ContainsFunc(s, func(uri *entity.RedirectUri) bool {
		if uri.URI() != v {
			return false
		}
		return true
	})
}

func (a *Authorize) SaveRequestFingerprint(ctx context.Context, redirectURI, clientID, authCode string) error {
	sess, ok := a.context.Read(ctx, httpstore.SessionKey{}).(*httpstore.Session)
	if !ok {
		return xerr.UnauthorizedRequest
	}

	oidcParam := &typedef.OIDCParam{
		RedirectURI: redirectURI,
		UserId:      sess.UserID,
		AuthTime:    sess.AuthTime,
	}

	return a.cache.WriteOpenIdParam(ctx, oidcParam, clientID, authCode)
}

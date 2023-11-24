package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/42milez/go-oidc-server/app/idp/httpstore"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/entity"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
	"github.com/42milez/go-oidc-server/app/typedef"

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

	uid, ok := a.context.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		return nil, "", xerr.UserIdNotFoundInContext
	}

	if _, err = a.repo.CreateAuthCode(ctx, code, clientID, uid); err != nil {
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
	uid, ok := a.context.Read(ctx, typedef.UserIdKey{}).(typedef.UserID)
	if !ok {
		return xerr.UserIdNotFoundInContext
	}

	authParam := &typedef.OpenIdParam{
		RedirectURI: redirectURI,
		UserId:      uid,
	}

	return a.cache.WriteOpenIdParam(ctx, authParam, clientID, authCode)
}

package service

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/httpstore"

	"github.com/42milez/go-oidc-expt/cmd/config"
	"github.com/42milez/go-oidc-expt/cmd/entity"
	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/repository"
	"github.com/42milez/go-oidc-expt/pkg/xerr"
	"github.com/42milez/go-oidc-expt/pkg/xrandom"

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
	cache   iface.AuthorizationRequestFingerprintWriter
	context iface.ContextReader
}

func (a *Authorize) Authorize(ctx context.Context, clientID typedef.ClientID, redirectURI, state string) (*url.URL, string, error) {
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

	ru, err := a.repo.ReadRedirectURIs(ctx, clientID)
	if err != nil {
		return nil, "", err
	}

	if !a.validateRedirectURI(ru, redirectURI) {
		return nil, "", xerr.InvalidRedirectURI
	}

	uri, err := url.Parse(fmt.Sprintf("%s?code=%s&state=%s", redirectURI, code, state))
	if err != nil {
		return nil, "", err
	}

	return uri, code, nil
}

func (a *Authorize) validateRedirectURI(s []*entity.RedirectURI, v string) bool {
	return slices.ContainsFunc(s, func(uri *entity.RedirectURI) bool {
		// TODO: compare domains
		if !strings.HasPrefix(v, uri.URI()) {
			return false
		}
		return true
	})
}

func (a *Authorize) SaveAuthorizationRequestFingerprint(ctx context.Context, clientID typedef.ClientID, redirectURI, nonce, authCode string) error {
	sess, ok := a.context.Read(ctx, httpstore.SessionKey{}).(*httpstore.Session)
	if !ok {
		return xerr.FailedToReadSession
	}
	fp := &typedef.AuthorizationRequestFingerprint{
		AuthTime:    sess.AuthTime,
		Nonce:       nonce,
		RedirectURI: redirectURI,
		UserID:      sess.UserID,
	}
	return a.cache.WriteAuthorizationRequestFingerprint(ctx, clientID, authCode, fp)
}

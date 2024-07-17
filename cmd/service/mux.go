package service

import (
	"context"

	"github.com/42milez/go-oidc-expt/cmd/iface"
	"github.com/42milez/go-oidc-expt/cmd/option"
	"github.com/42milez/go-oidc-expt/cmd/repository"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/42milez/go-oidc-expt/pkg/typedef"
)

func NewOapiAuthenticate(opt *option.Option) *OapiAuthenticate {
	return &OapiAuthenticate{
		repo:  repository.NewRelyingParty(opt.DB),
		token: opt.Token,
	}
}

type OapiAuthenticate struct {
	repo  CredentialReader
	token iface.TokenParser
}

func (oa *OapiAuthenticate) ValidateCredential(ctx context.Context, clientID typedef.ClientID, clientSecret string) error {
	if _, err := oa.repo.ReadCredential(ctx, clientID, clientSecret); err != nil {
		return err
	}
	return nil
}

func (oa *OapiAuthenticate) ParseAccessToken(token string) (jwt.Token, error) {
	t, err := oa.token.Parse(token)
	if err != nil {
		return nil, err
	}
	return t, nil
}

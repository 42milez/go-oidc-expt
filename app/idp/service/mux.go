package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"
)

func NewOapiAuthenticate(repo CredentialReader) *OapiAuthenticate {
	return &OapiAuthenticate{
		repo: repo,
	}
}

type OapiAuthenticate struct {
	repo CredentialReader
}

func (oa *OapiAuthenticate) ValidateCredential(ctx context.Context, clientID typedef.ClientID, clientSecret string) error {
	if _, err := oa.repo.ReadCredential(ctx, clientID, clientSecret); err != nil {
		return err
	}
	return nil
}

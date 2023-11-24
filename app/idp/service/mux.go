package service

import (
	"context"
)

func NewOapiAuthenticate(repo CredentialReader) *OapiAuthenticate {
	return &OapiAuthenticate{
		repo: repo,
	}
}

type OapiAuthenticate struct {
	repo CredentialReader
}

func (oa *OapiAuthenticate) ValidateCredential(ctx context.Context, clientID, clientSecret string) error {
	if _, err := oa.repo.ReadCredential(ctx, clientID, clientSecret); err != nil {
		return err
	}
	return nil
}

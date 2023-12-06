package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/option"
	"github.com/42milez/go-oidc-server/app/idp/repository"
)

func NewConsent(opt *option.Option) *AcceptConsent {
	return &AcceptConsent{
		repo: repository.NewUser(opt.DB, opt.IdGen),
	}
}

type AcceptConsent struct {
	repo ConsentCreator
}

func (ac *AcceptConsent) AcceptConsent(ctx context.Context, userID typedef.UserID, clientID string) error {
	if _, err := ac.repo.CreateConsent(ctx, userID, clientID); err != nil {
		return err
	}
	return nil
}

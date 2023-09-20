package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/typedef"
)

type AcceptConsent struct {
	repo ConsentCreator
}

func (ac *AcceptConsent) AcceptConsent(ctx context.Context, userID typedef.UserID, clientID string) error {
	// NOT IMPLEMENTED
	return nil
}

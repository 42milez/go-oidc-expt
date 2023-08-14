package service

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/model"
)

type Authorize struct{}

func (p *Authorize) Authorize(ctx context.Context, query *model.Authorize) error {
	return nil
}

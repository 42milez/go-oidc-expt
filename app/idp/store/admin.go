package store

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/entity"
)

func (p *Repository) ReadAdmin(ctx context.Context, db *ent.Client) (entity.Admin, error) {
	return nil, errors.New("not implemented")
}

func (p *Repository) ReadAdmins(ctx context.Context, db *ent.Client) (entity.Admins, error) {
	return nil, errors.New("not implemented")
}

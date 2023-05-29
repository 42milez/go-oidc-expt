package store

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/src/ent/ent"
	"github.com/42milez/go-oidc-server/src/entity"
)

func (p *Repository) ReadUser(ctx context.Context, db *ent.Client) (entity.User, error) {
	return nil, errors.New("not implemented")
}

func (p *Repository) ReadUsers(ctx context.Context, db *ent.Client) (entity.Users, error) {
	return nil, errors.New("not implemented")
}

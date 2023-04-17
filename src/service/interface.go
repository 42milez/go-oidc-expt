package service

import (
	"context"

	"github.com/42milez/go-oidc-server/src/ent/ent"
	"github.com/42milez/go-oidc-server/src/entity"
)

type UserCreater interface {
	Create(ctx context.Context) error
}

type UserReader interface {
	Read(ctx context.Context, db *ent.Client) (entity.User, error)
	ReadBulk(ctx context.Context, db *ent.Client) ([]entity.Users, error)
}

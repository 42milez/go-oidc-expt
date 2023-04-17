package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/src/ent/ent"
	"github.com/42milez/go-oidc-server/src/entity"
)

type CreateUser struct {
	DB   *ent.UserClient
	Repo UserCreater
}

func (p *CreateUser) Create(ctx context.Context) error {
	return nil
}

type ReadUser struct {
	DB   *ent.UserClient
	Repo UserReader
}

func (p *ReadUser) Read(ctx context.Context, db *ent.Client) (entity.User, error) {
	return nil, errors.New("not implemented")
}

func (p *ReadUser) ReadBulk(ctx context.Context, db *ent.Client) (entity.Users, error) {
	return nil, errors.New("not implemented")
}

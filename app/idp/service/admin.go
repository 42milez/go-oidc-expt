package service

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type CreateAdmin struct {
	DB   *ent.Client
	Repo AdminCreater
}

func (p *CreateAdmin) Create(ctx context.Context) error {
	return nil
}

type ReadAdmin struct {
	DB   *ent.Client
	Repo AdminReader
}

func (p *ReadAdmin) ReadAdmin(ctx context.Context) (*ent.Admin, error) {
	return nil, errors.New("not implemented")
}

type ReadAdmins struct {
	DB   *ent.Client
	Repo AdminsReader
}

func (p *ReadAdmins) ReadAdmins(ctx context.Context) ([]*ent.Admin, error) {
	return nil, errors.New("not implemented")
}

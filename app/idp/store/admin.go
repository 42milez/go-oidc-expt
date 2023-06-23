package store

import (
	"context"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/admin"
)

type AdminRepository struct {
	Clock xutil.Clocker
	DB    *ent.Client
}

func (p *AdminRepository) Create(ctx context.Context, name, pw string) (*ent.Admin, error) {
	return nil, nil
}

func (p *AdminRepository) SelectByName(ctx context.Context, name string) (*ent.Admin, error) {
	return p.DB.Admin.Query().Where(admin.NameEQ(name)).First(ctx)
}

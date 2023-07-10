package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
	"github.com/42milez/go-oidc-server/app/idp/ent/ent/admin"
)

type Admin struct {
	Clock xutil.Clocker
	DB    *ent.Client
}

func (p *Admin) Create(ctx context.Context, name string, pw typedef.PasswordHash) (*ent.Admin, error) {
	return p.DB.Admin.Create().SetName(name).SetPasswordHash(pw).Save(ctx)
}

func (p *Admin) SelectByName(ctx context.Context, name string) (*ent.Admin, error) {
	return p.DB.Admin.Query().Where(admin.NameEQ(name)).First(ctx)
}

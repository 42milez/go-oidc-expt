package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent/user"

	"github.com/42milez/go-oidc-server/app/idp/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xutil"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent"
)

type User struct {
	Clock xutil.Clocker
	DB    *ent.Client
}

func (p *User) Create(ctx context.Context, name string, pw typedef.PasswordHash) (*ent.User, error) {
	return p.DB.User.Create().SetName(name).SetPasswordHash(pw).Save(ctx)
}

func (p *User) SelectByName(ctx context.Context, name string) (*ent.User, error) {
	return p.DB.User.Query().Where(user.NameEQ(name)).First(ctx)
}

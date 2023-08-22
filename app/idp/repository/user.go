package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/idp/ent/ent/redirecturi"

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

func (p *User) SelectRedirectURIByUserID(ctx context.Context, userID typedef.UserID) ([]*ent.RedirectURI, error) {
	return p.DB.RedirectURI.Query().Where(redirecturi.UserIDEQ(userID)).All(ctx)
}

func (p *User) SaveAuthorizationCode(ctx context.Context, userID typedef.UserID, code string) (*ent.AuthCode, error) {
	return p.DB.AuthCode.Create().SetUserID(userID).SetCode(code).Save(ctx)
}

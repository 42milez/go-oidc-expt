package repository

import (
	"context"

	ent2 "github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/ent/ent/user"
	"github.com/42milez/go-oidc-server/app/ent/typedef"

	"github.com/42milez/go-oidc-server/pkg/xutil"
)

type User struct {
	Clock xutil.Clocker
	DB    *ent2.Client
}

func (p *User) Create(ctx context.Context, name string, pw string) (*ent2.User, error) {
	return p.DB.User.Create().SetName(name).SetPassword(pw).Save(ctx)
}

func (p *User) SelectByName(ctx context.Context, name string) (*ent2.User, error) {
	return p.DB.User.Query().Where(user.NameEQ(name)).First(ctx)
}

func (p *User) SelectRedirectURIByUserID(ctx context.Context, userID typedef.UserID) ([]*ent2.RedirectURI, error) {
	return p.DB.RedirectURI.Query().Where(redirecturi.UserIDEQ(userID)).All(ctx)
}

func (p *User) SaveAuthorizationCode(ctx context.Context, userID typedef.UserID, code string) (*ent2.AuthCode, error) {
	return p.DB.AuthCode.Create().SetUserID(userID).SetCode(code).Save(ctx)
}

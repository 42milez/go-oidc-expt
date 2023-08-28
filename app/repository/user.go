package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/pkg/xtime"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-server/app/ent/ent/user"
	"github.com/42milez/go-oidc-server/app/ent/typedef"
)

type User struct {
	Clock xtime.Clocker
	DB    *ent.Client
	IDGen IDGenerator
}

func (p *User) Create(ctx context.Context, name string, pw string) (*ent.User, error) {
	return p.DB.User.Create().SetName(name).SetPassword(pw).Save(ctx)
}

func (p *User) SelectByName(ctx context.Context, name string) (*ent.User, error) {
	return p.DB.User.Query().Where(user.NameEQ(name)).First(ctx)
}

func (p *User) SelectRedirectUriByUserID(ctx context.Context, userID typedef.UserID) ([]*ent.RedirectURI, error) {
	return p.DB.RedirectURI.Query().Where(redirecturi.UserIDEQ(userID)).All(ctx)
}

func (p *User) SaveAuthorizationCode(ctx context.Context, userID typedef.UserID, code string) (*ent.AuthCode, error) {
	return p.DB.AuthCode.Create().SetUserID(userID).SetCode(code).Save(ctx)
}

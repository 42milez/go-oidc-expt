package repository

import (
	"context"
	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/datastore"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/redirecturi"
	_ "github.com/42milez/go-oidc-server/app/ent/ent/runtime"
	"github.com/42milez/go-oidc-server/app/ent/ent/user"
)

type User struct {
	db    *datastore.Database
	idGen IDGenerator
}

func (u *User) CreateAuthCode(ctx context.Context, id typedef.RelyingPartyID, code string) (*ent.AuthCode, error) {
	return u.db.Client.AuthCode.Create().SetRelyingPartyID(id).SetCode(code).Save(ctx)
}

func (u *User) CreateUser(ctx context.Context, name string, pw string) (*ent.User, error) {
	return u.db.Client.User.Create().SetName(name).SetPassword(pw).Save(ctx)
}

func (u *User) ReadUserByName(ctx context.Context, name string) (*ent.User, error) {
	return u.db.Client.User.Query().Where(user.NameEQ(name)).First(ctx)
}

func (u *User) ReadRedirectUriByUserID(ctx context.Context, id typedef.RelyingPartyID) ([]*ent.RedirectURI, error) {
	return u.db.Client.RedirectURI.Query().Where(redirecturi.RelyingPartyIDEQ(id)).All(ctx)
}

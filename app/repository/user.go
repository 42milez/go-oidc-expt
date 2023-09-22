package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/datastore"

	"github.com/42milez/go-oidc-server/app/ent/ent"
	_ "github.com/42milez/go-oidc-server/app/ent/ent/runtime"
	"github.com/42milez/go-oidc-server/app/ent/ent/user"
)

type User struct {
	db    *datastore.Database
	idGen IDGenerator
}

func (u *User) CreateUser(ctx context.Context, name string, pw string) (*ent.User, error) {
	return u.db.Client.User.Create().SetName(name).SetPassword(pw).Save(ctx)
}

func (u *User) CreateConsent(ctx context.Context, userID typedef.UserID, clientID string) (*ent.Consent, error) {
	tx, err := u.db.Client.Tx(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	targetUser, err := tx.User.Query().Where(user.ID(userID)).Only(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	consent, err := tx.Consent.Create().SetUser(targetUser).SetClientID(clientID).Save(ctx)

	if err != nil {
		return nil, rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}

	return consent, nil
}

func (u *User) ReadUserByName(ctx context.Context, name string) (*ent.User, error) {
	return u.db.Client.User.Query().Where(user.NameEQ(name)).First(ctx)
}

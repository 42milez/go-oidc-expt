package repository

import (
	"context"
	"errors"

	"github.com/42milez/go-oidc-server/app/pkg/typedef"

	"github.com/42milez/go-oidc-server/app/idp/entity"

	"github.com/42milez/go-oidc-server/app/ent/ent/consent"
	_ "github.com/42milez/go-oidc-server/app/ent/ent/runtime"
	"github.com/42milez/go-oidc-server/app/ent/ent/user"
	"github.com/42milez/go-oidc-server/app/idp/datastore"
	"github.com/42milez/go-oidc-server/app/idp/iface"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewUser(db *datastore.Database, idGen iface.IdGenerator) *User {
	return &User{
		db:    db,
		idGen: idGen,
	}
}

type User struct {
	db    *datastore.Database
	idGen iface.IdGenerator
}

func (u *User) CreateUser(ctx context.Context, name string, pw string) (*entity.User, error) {
	v, err := u.db.Client.User.Create().SetName(name).SetPassword(pw).Save(ctx)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(v), nil
}

func (u *User) CreateConsent(ctx context.Context, userID typedef.UserID, clientID string) (*entity.Consent, error) {
	tx, err := u.db.Client.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	targetUser, err := tx.User.Query().Where(user.ID(userID)).ForShare().Only(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	c, err := tx.Consent.Create().SetUser(targetUser).SetClientID(clientID).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}

	return entity.NewConsent(c), nil
}

func (u *User) ReadConsent(ctx context.Context, userID typedef.UserID, clientID string) (*entity.Consent, error) {
	c, err := u.db.Client.Consent.Query().Where(consent.UserID(userID), consent.ClientID(clientID)).Only(ctx)
	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.ConsentNotFound
		} else {
			return nil, err
		}
	}
	return entity.NewConsent(c), nil
}

func (u *User) ReadUser(ctx context.Context, name string) (*entity.User, error) {
	v, err := u.db.Client.User.Query().Where(user.NameEQ(name)).First(ctx)
	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.UserNotFound
		}
		return nil, err
	}
	return entity.NewUser(v), err
}

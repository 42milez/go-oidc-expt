package repository

import (
	"context"
	"errors"
	"time"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/authcode"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/pkg/xerr"
)

func NewAuthCode(db *datastore.Database) *AuthCode {
	return &AuthCode{
		db: db,
	}
}

type AuthCode struct {
	db *datastore.Database
}

func (a *AuthCode) MarkAuthCodeUsed(ctx context.Context, code, clientId string) (*ent.AuthCode, error) {
	tx, err := a.db.Client.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	ac, err := tx.AuthCode.Query().Where(authcode.Code(code)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientId))
	}).ForShare().Only(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	ac, err = ac.Update().SetUsedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}

	return ac, nil
}

func (a *AuthCode) ReadAuthCode(ctx context.Context, code, clientId string) (*ent.AuthCode, error) {
	ret, err := a.db.Client.AuthCode.Query().Where(authcode.Code(code)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientId))
	}).Only(ctx)

	if err != nil {
		if errors.Is(err, &ent.NotFoundError{}) {
			return nil, xerr.AuthCodeNotFound
		} else {
			return nil, err
		}
	}

	return ret, nil
}

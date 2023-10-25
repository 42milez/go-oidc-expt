package repository

import (
	"context"
	"errors"
	"time"

	"github.com/42milez/go-oidc-server/app/entity"

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

func (ac *AuthCode) MarkAuthCodeUsed(ctx context.Context, code, clientId string) (*entity.AuthCode, error) {
	tx, err := ac.db.Client.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	v, err := tx.AuthCode.Query().Where(authcode.Code(code)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientId))
	}).ForShare().Only(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	v, err = v.Update().SetUsedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}

	return entity.NewAuthCode(v), nil
}

func (ac *AuthCode) ReadAuthCode(ctx context.Context, code, clientId string) (*entity.AuthCode, error) {
	v, err := ac.db.Client.AuthCode.Query().Where(authcode.Code(code)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientId))
	}).Only(ctx)

	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.AuthCodeNotFound
		} else {
			return nil, err
		}
	}

	return entity.NewAuthCode(v), nil
}

package repository

import (
	"context"
	"errors"
	"time"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/pkg/ent/ent"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/authcode"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/relyingparty"

	"github.com/42milez/go-oidc-expt/cmd/datastore"
	"github.com/42milez/go-oidc-expt/cmd/entity"

	"github.com/42milez/go-oidc-expt/pkg/xerr"
)

func NewAuthCode(db *datastore.Database) *AuthCode {
	return &AuthCode{
		db: db,
	}
}

type AuthCode struct {
	db *datastore.Database
}

func (ac *AuthCode) RevokeAuthCode(ctx context.Context, code string, clientID typedef.ClientID) (*entity.AuthCode, error) {
	tx, err := ac.db.Client.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	v, err := tx.AuthCode.Query().Where(authcode.Code(code)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientID))
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

func (ac *AuthCode) ReadAuthCode(ctx context.Context, code string, clientID typedef.ClientID) (*entity.AuthCode, error) {
	v, err := ac.db.Client.AuthCode.Query().Where(authcode.Code(code)).WithRelyingParty(func(q *ent.RelyingPartyQuery) {
		q.Where(relyingparty.ClientID(clientID))
	}).Only(ctx)

	if err != nil {
		if errors.As(err, &errEntNotFoundError) {
			return nil, xerr.RecordNotFound
		} else {
			return nil, err
		}
	}

	return entity.NewAuthCode(v), nil
}

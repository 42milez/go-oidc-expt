package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
)

func NewRelyingParty(db *datastore.Database) *RelyingParty {
	return &RelyingParty{
		db: db,
	}
}

type RelyingParty struct {
	db *datastore.Database
}

func (rp *RelyingParty) CreateAuthCode(ctx context.Context, code string, clientID string, userID typedef.UserID) (*ent.AuthCode, error) {
	tx, err := rp.db.Client.Tx(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	owner, err := tx.RelyingParty.Query().Where(relyingparty.ClientID(clientID)).ForShare().Only(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	authCode, err := tx.AuthCode.Create().SetRelyingParty(owner).SetCode(code).SetUserID(userID).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return nil, rollback(tx, err)
	}

	return authCode, nil
}

func (rp *RelyingParty) ReadRedirectUriByClientID(ctx context.Context, clientID string) ([]*ent.RedirectURI, error) {
	return rp.db.Client.RelyingParty.
		Query().
		Where(relyingparty.ClientID(clientID)).
		QueryRedirectUris().
		All(ctx)
}

func (rp *RelyingParty) ReadCredential(ctx context.Context, clientID, clientSecret string) (*ent.RelyingParty, error) {
	return rp.db.Client.RelyingParty.Query().
		Where(relyingparty.ClientID(clientID), relyingparty.ClientSecret(clientSecret)).Only(ctx)
}

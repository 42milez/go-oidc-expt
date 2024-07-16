package repository

import (
	"context"

	"github.com/42milez/go-oidc-expt/pkg/ent/ent/relyingparty"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/cmd/entity"

	"github.com/42milez/go-oidc-expt/cmd/datastore"
)

func NewRelyingParty(db *datastore.Database) *RelyingParty {
	return &RelyingParty{
		db: db,
	}
}

type RelyingParty struct {
	db *datastore.Database
}

func (rp *RelyingParty) CreateAuthCode(ctx context.Context, code string, clientID typedef.ClientID, userID typedef.UserID) (*entity.AuthCode, error) {
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

	return entity.NewAuthCode(authCode), nil
}

func (rp *RelyingParty) ReadRedirectURIs(ctx context.Context, clientID typedef.ClientID) ([]*entity.RedirectURI, error) {
	redirectUris, err := rp.db.Client.RelyingParty.
		Query().
		Where(relyingparty.ClientID(clientID)).
		QueryRedirectUris().
		All(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]*entity.RedirectURI, len(redirectUris))

	for i, v := range redirectUris {
		ret[i] = entity.NewRedirectURI(v)
	}

	return ret, nil
}

func (rp *RelyingParty) ReadCredential(ctx context.Context, clientID typedef.ClientID, clientSecret string) (*entity.RelyingParty, error) {
	v, err := rp.db.Client.RelyingParty.Query().
		Where(relyingparty.ClientID(clientID), relyingparty.ClientSecret(clientSecret)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return entity.NewRelyingParty(v), err
}

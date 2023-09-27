package repository

import (
	"context"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-server/app/typedef"
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
	//return rp.db.Client.AuthCode.Create().SetRelyingPartyID(id).SetCode(code).Save(ctx)
	return nil, nil
}

func (rp *RelyingParty) ReadRedirectUriByClientID(ctx context.Context, clientID string) ([]*ent.RedirectURI, error) {
	return rp.db.Client.RelyingParty.
		Query().
		Where(relyingparty.ClientID(clientID)).
		QueryRedirectUris().
		All(ctx)
}

package main

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/typedef"
)

const nConsentMin = 1

func insertConsents(ctx context.Context, db *datastore.Database, users []*ent.User, relyingParties []*ent.RelyingParty, nConsent int) ([]*ent.Consent, error) {
	if db == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nConsent < nUserMin {
		return nil, fmt.Errorf("the number of consents must be greater than or equal to %d", nConsentMin)
	}

	nUser := len(users)
	nRelyingParty := len(relyingParties)

	params := make([]struct {
		UserID         typedef.UserID
		RelyingPartyID int
	}, nConsent*nUser)

	for i := range params {
		params[i].UserID = users[i%nUser].ID
		params[i].RelyingPartyID = relyingParties[i%nRelyingParty].ID
	}

	printSeeds(params)

	builders := make([]*ent.ConsentCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.Consent.Create().SetUserID(v.UserID).SetRelyingPartyID(v.RelyingPartyID)
	}

	return db.Client.Consent.CreateBulk(builders...).Save(ctx)
}

package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/pkg/ent/ent"

	"github.com/42milez/go-oidc-server/app/idp/datastore"
)

const nConsentMin = 1

type Consent struct {
	ClientID string
	User     *ent.User
}

func InsertConsents(ctx context.Context, db *datastore.Database, users []*ent.User, relyingParties []*ent.RelyingParty, nConsent int) ([]*ent.Consent, error) {
	if nConsent < nUserMin {
		return nil, fmt.Errorf("the number of consents must be greater than or equal to %d", nConsentMin)
	}

	nUser := len(users)
	nRelyingParty := len(relyingParties)

	params := make([]Consent, nConsent*nUser)

	for i := range params {
		params[i].ClientID = relyingParties[i%nRelyingParty].ClientID
		params[i].User = users[i%nUser]
	}

	builders := make([]*ent.ConsentCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.Consent.Create().SetClientID(v.ClientID).SetUser(v.User)
	}

	return db.Client.Consent.CreateBulk(builders...).Save(ctx)
}

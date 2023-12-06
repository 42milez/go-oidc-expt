package main

import (
	"context"
	"fmt"

	ent2 "github.com/42milez/go-oidc-server/app/pkg/ent/ent"

	"github.com/42milez/go-oidc-server/app/idp/datastore"
)

const nConsentMin = 1

type Consent struct {
	ClientID string
	User     *ent2.User
}

func InsertConsents(ctx context.Context, db *datastore.Database, users []*ent2.User, relyingParties []*ent2.RelyingParty, nConsent int) ([]*ent2.Consent, error) {
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

	builders := make([]*ent2.ConsentCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.Consent.Create().SetClientID(v.ClientID).SetUser(v.User)
	}

	return db.Client.Consent.CreateBulk(builders...).Save(ctx)
}

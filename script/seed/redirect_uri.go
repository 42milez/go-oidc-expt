package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
)

const nRedirectUriMin = 1

func insertRedirectURIs(ctx context.Context, db *datastore.Database, relyingParties []*ent.RelyingParty, nRedirectURI int) ([]*ent.RedirectURI, error) {
	if db == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nRedirectURI < nRedirectUriMin {
		return nil, fmt.Errorf("the number of redirect uris must be greater than or equal to %d", nRedirectUriMin)
	}

	nRelyingParty := len(relyingParties)

	params := make([]struct {
		URI            string
		RelyingPartyID typedef.RelyingPartyID
	}, nRedirectURI*nRelyingParty)

	for i := range params {
		params[i].URI = fmt.Sprintf("https://example.com/cb%d", i)
		params[i].RelyingPartyID = relyingParties[i%nRelyingParty].ID
	}

	printSeeds(params)

	builders := make([]*ent.RedirectURICreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RedirectURI.Create().SetURI(v.URI).SetRelyingPartyID(v.RelyingPartyID)
	}

	return db.Client.RedirectURI.CreateBulk(builders...).Save(ctx)
}

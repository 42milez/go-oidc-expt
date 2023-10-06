package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
)

const nRedirectUriMin = 1

type RedirectURI struct {
	URI          string
	RelyingParty *ent.RelyingParty
}

func InsertRedirectURIs(ctx context.Context, db *datastore.Database, relyingParties []*ent.RelyingParty, nRedirectURI int) ([]*ent.RedirectURI, error) {
	if nRedirectURI < nRedirectUriMin {
		return nil, fmt.Errorf("the number of redirect uris must be greater than or equal to %d", nRedirectUriMin)
	}

	nRelyingParty := len(relyingParties)

	params := make([]RedirectURI, nRedirectURI*nRelyingParty)

	for i := range params {
		params[i].URI = fmt.Sprintf("https://example.com/cb%d", i)
		params[i].RelyingParty = relyingParties[i%nRelyingParty]
	}

	builders := make([]*ent.RedirectURICreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RedirectURI.Create().SetURI(v.URI).SetRelyingParty(v.RelyingParty)
	}

	return db.Client.RedirectURI.CreateBulk(builders...).Save(ctx)
}

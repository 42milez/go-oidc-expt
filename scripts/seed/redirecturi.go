package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/idp/datastore"

	"github.com/42milez/go-oidc-server/app/ent/ent"
)

const nRedirectUriMin = 1

type RedirectUri struct {
	URI          string
	RelyingParty *ent.RelyingParty
}

func InsertRedirectUris(ctx context.Context, db *datastore.Database, relyingParties []*ent.RelyingParty, nRedirectUri int) ([]*ent.RedirectUri, error) {
	if nRedirectUri < nRedirectUriMin {
		return nil, fmt.Errorf("the number of redirect uris must be greater than or equal to %d", nRedirectUriMin)
	}

	nRelyingParty := len(relyingParties)

	params := make([]RedirectUri, nRedirectUri*nRelyingParty)

	for i := range params {
		params[i].URI = fmt.Sprintf("https://example.com/cb%d", i)
		params[i].RelyingParty = relyingParties[i%nRelyingParty]
	}

	builders := make([]*ent.RedirectUriCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RedirectUri.Create().SetURI(v.URI).SetRelyingParty(v.RelyingParty)
	}

	return db.Client.RedirectUri.CreateBulk(builders...).Save(ctx)
}

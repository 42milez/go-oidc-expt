package main

import (
	"context"
	"fmt"

	ent2 "github.com/42milez/go-oidc-server/app/pkg/ent/ent"

	"github.com/42milez/go-oidc-server/app/idp/datastore"
)

const nRedirectUriMin = 1

type RedirectUri struct {
	URI          string
	RelyingParty *ent2.RelyingParty
}

func InsertRedirectUris(ctx context.Context, db *datastore.Database, relyingParties []*ent2.RelyingParty, nRedirectUri int) ([]*ent2.RedirectUri, error) {
	if nRedirectUri < nRedirectUriMin {
		return nil, fmt.Errorf("the number of redirect uris must be greater than or equal to %d", nRedirectUriMin)
	}

	nRelyingParty := len(relyingParties)

	params := make([]RedirectUri, nRedirectUri*nRelyingParty)

	for i := range params {
		params[i].URI = fmt.Sprintf("https://example.com/cb%d", i)
		params[i].RelyingParty = relyingParties[i%nRelyingParty]
	}

	builders := make([]*ent2.RedirectUriCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RedirectUri.Create().SetURI(v.URI).SetRelyingParty(v.RelyingParty)
	}

	return db.Client.RedirectUri.CreateBulk(builders...).Save(ctx)
}

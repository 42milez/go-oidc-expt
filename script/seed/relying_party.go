package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
	"github.com/42milez/go-oidc-server/app/typedef"
)

const nRelyingPartyMin = 1

func insertRelyingParties(ctx context.Context, db *datastore.Database, nRelyingParty int) ([]*ent.RelyingParty, error) {
	if db == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nRelyingParty < nRelyingPartyMin {
		return nil, fmt.Errorf("the number of relying parties must be greater than or equal to %d", nRelyingPartyMin)
	}

	params := make([]struct {
		clientID     typedef.ClientID
		clientSecret typedef.ClientSecret
	}, nRelyingParty)

	for i := 0; i < nRelyingParty; i++ {
		v, err := xrandom.MakeCryptoRandomString(config.ClientIdLength)
		if err != nil {
			return nil, err
		}
		params[i].clientID = typedef.ClientID(v)
		v, err = xrandom.MakeCryptoRandomString(config.ClientSecretLength)
		if err != nil {
			return nil, err
		}
		params[i].clientSecret = typedef.ClientSecret(v)
	}

	printSeeds(params)

	builders := make([]*ent.RelyingPartyCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RelyingParty.Create().SetClientID(v.clientID).SetClientSecret(v.clientSecret)
	}

	return db.Client.RelyingParty.CreateBulk(builders...).Save(ctx)
}

package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
)

const nRelyingPartyMin = 1

type RelyingParty struct {
	ClientID     string
	ClientSecret string
}

func InsertRelyingParties(ctx context.Context, db *datastore.Database, nRelyingParty int) ([]*ent.RelyingParty, error) {
	if db == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nRelyingParty < nRelyingPartyMin {
		return nil, fmt.Errorf("the number of relying parties must be greater than or equal to %d", nRelyingPartyMin)
	}

	params := make([]RelyingParty, nRelyingParty)

	for i := 0; i < nRelyingParty; i++ {
		v, err := xrandom.MakeCryptoRandomString(config.ClientIdLength)
		if err != nil {
			return nil, err
		}
		params[i].ClientID = string(v)
		v, err = xrandom.MakeCryptoRandomString(config.ClientSecretLength)
		if err != nil {
			return nil, err
		}
		params[i].ClientSecret = string(v)
	}

	builders := make([]*ent.RelyingPartyCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RelyingParty.Create().SetClientID(v.ClientID).SetClientSecret(v.ClientSecret)
	}

	return db.Client.RelyingParty.CreateBulk(builders...).Save(ctx)
}

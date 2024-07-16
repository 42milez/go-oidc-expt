package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"github.com/42milez/go-oidc-expt/pkg/ent/ent"

	"github.com/42milez/go-oidc-expt/cmd/config"
	"github.com/42milez/go-oidc-expt/cmd/datastore"

	"github.com/42milez/go-oidc-expt/pkg/xrandom"
)

const nRelyingPartyMin = 1

type RelyingParty struct {
	ClientID     string
	ClientSecret string
}

func InsertRelyingParties(ctx context.Context, db *datastore.Database, nRelyingParty int) ([]*ent.RelyingParty, error) {
	if nRelyingParty < nRelyingPartyMin {
		return nil, fmt.Errorf("the number of relying parties must be greater than or equal to %d", nRelyingPartyMin)
	}

	params := make([]RelyingParty, nRelyingParty)

	for i := 0; i < nRelyingParty; i++ {
		v, err := xrandom.GenerateCryptoRandomString(config.ClientIDLength)
		if err != nil {
			return nil, err
		}
		params[i].ClientID = string(v)
		v, err = xrandom.GenerateCryptoRandomString(config.ClientSecretLength)
		if err != nil {
			return nil, err
		}
		params[i].ClientSecret = string(v)
	}

	builders := make([]*ent.RelyingPartyCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.RelyingParty.Create().SetClientID(typedef.ClientID(v.ClientID)).SetClientSecret(v.ClientSecret)
	}

	return db.Client.RelyingParty.CreateBulk(builders...).Save(ctx)
}

package main

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/app/typedef"
	"time"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
)

const nAuthCodeMin = 1

func insertAuthCodes(ctx context.Context, db *datastore.Database, relyingParties []*ent.RelyingParty, nAuthCode int) ([]*ent.AuthCode, error) {
	if db == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nAuthCode < nAuthCodeMin {
		return nil, fmt.Errorf("the number of auth codes must be greater than or equal to %d", nAuthCodeMin)
	}

	nRelyingParty := len(relyingParties)

	params := make([]struct {
		code           string
		usedAt         *time.Time
		relyingPartyId typedef.RelyingPartyID
	}, nAuthCode*nRelyingParty)

	for i := range params {
		code, err := xrandom.MakeCryptoRandomString(config.AuthCodeLength)
		if err != nil {
			return nil, err
		}
		params[i].code = code
		params[i].relyingPartyId = relyingParties[i%nRelyingParty].ID
	}

	printSeeds(params)

	builders := make([]*ent.AuthCodeCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.AuthCode.Create().SetCode(v.code).SetRelyingPartyID(v.relyingPartyId)
	}

	return db.Client.AuthCode.CreateBulk(builders...).Save(ctx)
}

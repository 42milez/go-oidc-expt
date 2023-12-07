package main

import (
	"context"
	"fmt"
	"time"

	ent2 "github.com/42milez/go-oidc-server/app/pkg/ent/ent"

	"github.com/42milez/go-oidc-server/app/idp/config"
	"github.com/42milez/go-oidc-server/app/idp/datastore"

	"github.com/42milez/go-oidc-server/app/pkg/xrandom"
)

const nAuthCodeMin = 1

type AuthCode struct {
	Code         string
	UsedAt       *time.Time
	RelyingParty *ent2.RelyingParty
	User         *ent2.User
}

func InsertAuthCodes(ctx context.Context, db *datastore.Database, relyingParties []*ent2.RelyingParty, users []*ent2.User, nAuthCode int) ([]*ent2.AuthCode, error) {
	if nAuthCode < nAuthCodeMin {
		return nil, fmt.Errorf("the number of auth codes must be greater than or equal to %d", nAuthCodeMin)
	}

	nRelyingParty := len(relyingParties)
	nUser := len(users)

	params := make([]AuthCode, nAuthCode*nRelyingParty)

	for i := range params {
		code, err := xrandom.GenerateCryptoRandomString(config.AuthCodeLength)
		if err != nil {
			return nil, err
		}
		params[i].Code = code
		params[i].RelyingParty = relyingParties[i%nRelyingParty]
		params[i].User = users[i%nUser]
	}

	builders := make([]*ent2.AuthCodeCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.AuthCode.Create().SetCode(v.Code).SetRelyingParty(v.RelyingParty).SetUserID(v.User.ID)
	}

	return db.Client.AuthCode.CreateBulk(builders...).Save(ctx)
}

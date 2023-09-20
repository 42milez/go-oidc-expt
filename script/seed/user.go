package main

import (
	"context"
	"fmt"

	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
)

const nUserMin = 1

func insertUsers(ctx context.Context, db *datastore.Database, nUser int) ([]*ent.User, error) {
	if db == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nUser < nUserMin {
		return nil, fmt.Errorf("the number of users must be greater than or equal to %d", nUserMin)
	}

	params := make([]struct {
		name     string
		password string
	}, nUser)

	for i := 0; i < nUser; i++ {
		params[i].name = fmt.Sprintf("username%d", i)
	}

	for i, v := range params {
		pwHash, err := xargon2.HashPassword(v.name)
		if err != nil {
			return nil, err
		}
		params[i].password = pwHash
	}

	printSeeds(params)

	builders := make([]*ent.UserCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.User.Create().SetName(v.name).SetPassword(v.password)
	}

	return db.Client.User.CreateBulk(builders...).Save(ctx)
}

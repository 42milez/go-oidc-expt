package main

import (
	"context"
	"fmt"

	ent2 "github.com/42milez/go-oidc-server/app/pkg/ent/ent"

	"github.com/42milez/go-oidc-server/app/idp/datastore"
	"github.com/42milez/go-oidc-server/app/idp/security"
)

const nUserMin = 1

type User struct {
	Name     string
	Password string
}

func InsertUsers(ctx context.Context, db *datastore.Database, nUser int) ([]*ent2.User, error) {
	if nUser < nUserMin {
		return nil, fmt.Errorf("the number of users must be greater than or equal to %d", nUserMin)
	}

	params := make([]User, nUser)

	for i := 0; i < nUser; i++ {
		params[i].Name = fmt.Sprintf("username%d", i)
	}

	for i, v := range params {
		pwHash, err := security.HashPassword(v.Name)
		if err != nil {
			return nil, err
		}
		params[i].Password = pwHash
	}

	builders := make([]*ent2.UserCreate, len(params))

	for i, v := range params {
		builders[i] = db.Client.User.Create().SetName(v.Name).SetPassword(v.Password)
	}

	return db.Client.User.CreateBulk(builders...).Save(ctx)
}

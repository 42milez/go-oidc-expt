package repository

import (
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/pkg/xtime"
)

func NewCheckHealth(db *datastore.Database, cache *datastore.Cache) *CheckHealth {
	return &CheckHealth{
		db:    db,
		cache: cache,
	}
}

func NewUser(db *datastore.Database, idGen IDGenerator) *User {
	return &User{
		db:    db,
		clock: xtime.RealClocker{},
		idGen: idGen,
	}
}

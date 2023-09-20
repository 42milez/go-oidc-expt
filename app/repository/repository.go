package repository

import (
	"github.com/42milez/go-oidc-server/app/datastore"
)

func NewCheckHealth(db *datastore.Database, cache *datastore.Cache) *CheckHealth {
	return &CheckHealth{
		db:    db,
		cache: cache,
	}
}

func NewRelyingParty(db *datastore.Database) *RelyingParty {
	return &RelyingParty{
		db: db,
	}
}

func NewSession(cache *datastore.Cache) *Session {
	return &Session{
		cache: cache,
	}
}

func NewUser(db *datastore.Database, idGen IDGenerator) *User {
	return &User{
		db:    db,
		idGen: idGen,
	}
}

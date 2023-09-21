package repository

import (
	"fmt"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
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

func rollback(tx *ent.Tx, err error) error {
	if retErr := tx.Rollback(); retErr != nil {
		return fmt.Errorf("%w: %v", err, retErr)
	}
	return err
}

package main

import (
	"context"
	"fmt"
	"github.com/42milez/go-oidc-server/app/datastore"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	"log"
	"reflect"

	"github.com/42milez/go-oidc-server/app/config"
	_ "github.com/42milez/go-oidc-server/app/ent/ent/runtime"
)

func printSeeds(data any) {
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("%+v\n", v.Index(i).Interface())
	}
}

func run(ctx context.Context, db *datastore.Database) error {
	nRelyingParty := 10
	nAuthCodeByClient := 3
	nRedirectUriByRelyingParty := 3

	nUser := 5
	nConsentByUser := 3

	var relyingParties []*ent.RelyingParty
	var users []*ent.User
	var err error

	if relyingParties, err = insertRelyingParties(ctx, db, nRelyingParty); err != nil {
		return err
	}

	if _, err = insertAuthCodes(ctx, db, relyingParties, nAuthCodeByClient); err != nil {
		return err
	}

	if _, err = insertRedirectURIs(ctx, db, relyingParties, nRedirectUriByRelyingParty); err != nil {
		return err
	}

	if users, err = insertUsers(ctx, db, nUser); err != nil {
		return err
	}

	if _, err = insertConsents(ctx, db, users, relyingParties, nConsentByUser); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()

	var cfg *config.Config
	var err error

	if cfg, err = config.New(); err != nil {
		log.Fatal(err)
	}

	cfg.DBPort = 13306
	db, err := datastore.NewDatabase(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err = run(ctx, db); err != nil {
		log.Fatal(err)
	}
}

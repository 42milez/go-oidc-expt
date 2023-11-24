package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/42milez/go-oidc-server/app/idp/config"
	datastore2 "github.com/42milez/go-oidc-server/app/idp/datastore"

	"github.com/42milez/go-oidc-server/app/ent/ent"

	_ "github.com/42milez/go-oidc-server/app/ent/ent/runtime"
)

func printSeeds(data any) {
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("%+v\n", v.Index(i).Interface())
	}
}

func run(ctx context.Context, db *datastore2.Database) error {
	var err error

	//  For Swagger
	// --------------------------------------------------

	if err = insertFixedData(ctx, db); err != nil {
		return err
	}

	//  Owner Edges
	// --------------------------------------------------

	nUser := 10
	nConsentByUser := 3

	nRelyingParty := nUser * nConsentByUser
	nAuthCodeByRelyingParty := 3
	nRedirectUriByRelyingParty := 3

	var relyingParties []*ent.RelyingParty
	var users []*ent.User

	if relyingParties, err = InsertRelyingParties(ctx, db, nRelyingParty); err != nil {
		return err
	}

	printSeeds(relyingParties)

	if users, err = InsertUsers(ctx, db, nUser); err != nil {
		return err
	}

	printSeeds(users)

	//  Other Edges
	// --------------------------------------------------

	var authCodes []*ent.AuthCode
	var redirectURIs []*ent.RedirectUri
	var consents []*ent.Consent

	if authCodes, err = InsertAuthCodes(ctx, db, relyingParties, users, nAuthCodeByRelyingParty); err != nil {
		return err
	}

	printSeeds(authCodes)

	if redirectURIs, err = InsertRedirectUris(ctx, db, relyingParties, nRedirectUriByRelyingParty); err != nil {
		return err
	}

	printSeeds(redirectURIs)

	if consents, err = InsertConsents(ctx, db, users, relyingParties, nConsentByUser); err != nil {
		return err
	}

	printSeeds(consents)

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
	db, err := datastore2.NewMySQL(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err = run(ctx, db); err != nil {
		log.Fatal(err)
	}
}

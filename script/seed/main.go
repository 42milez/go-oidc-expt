package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/42milez/go-oidc-server/app/pkg/xargon2"
	"github.com/42milez/go-oidc-server/app/pkg/xid"
	"github.com/42milez/go-oidc-server/app/pkg/xrandom"

	"github.com/42milez/go-oidc-server/app/typedef"

	"github.com/42milez/go-oidc-server/app/config"
	"github.com/42milez/go-oidc-server/app/ent/ent"
	_ "github.com/42milez/go-oidc-server/app/ent/ent/runtime"
	"github.com/42milez/go-oidc-server/app/repository"
)

const nUserMin = 1
const nAuthCodeMin = 1
const nRedirectUriMin = 1

func printSeeds(data any) {
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		fmt.Printf("%+v\n", v.Index(i).Interface())
	}
}

func insertUsers(ctx context.Context, client *ent.Client, nUser int) ([]*ent.User, error) {
	if client == nil {
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
		builders[i] = client.User.Create().SetName(v.name).SetPassword(v.password)
	}

	return client.User.CreateBulk(builders...).Save(ctx)
}

func insertAuthCodes(ctx context.Context, client *ent.Client, users []*ent.User, nCodeByUser int) ([]*ent.AuthCode, error) {
	if client == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nCodeByUser < nAuthCodeMin {
		return nil, fmt.Errorf("the number of auth codes must be greater than or equal to %d", nAuthCodeMin)
	}

	type param struct {
		code     string
		expireAt time.Time
		userID   typedef.UserID
	}

	nUser := len(users)
	params := make([]param, nUser*nCodeByUser)
	expireAt := time.Now().Add(config.AuthCodeLifetime)

	for i := range params {
		code, err := xrandom.MakeCryptoRandomString(config.AuthCodeLength)
		if err != nil {
			return nil, err
		}
		params[i].code = code
		params[i].expireAt = expireAt
		params[i].userID = users[i%nUser].ID
	}

	printSeeds(params)

	builders := make([]*ent.AuthCodeCreate, len(params))

	for i, v := range params {
		builders[i] = client.AuthCode.Create().SetCode(v.code).SetExpireAt(v.expireAt).SetUserID(v.userID)
	}

	return client.AuthCode.CreateBulk(builders...).Save(ctx)
}

func insertRedirectURIs(ctx context.Context, client *ent.Client, users []*ent.User, nUriByUser int) ([]*ent.RedirectURI, error) {
	if client == nil {
		return nil, fmt.Errorf("database client required")
	}

	if nUriByUser < nRedirectUriMin {
		return nil, fmt.Errorf("the number of auth codes must be greater than or equal to %d", nAuthCodeMin)
	}

	type param struct {
		uri    string
		userID typedef.UserID
	}

	nUser := len(users)
	params := make([]param, nUser*nUriByUser)

	for i := range params {
		params[i].uri = fmt.Sprintf("http://example.com/cb%d", i)
		params[i].userID = users[i%nUser].ID
	}

	printSeeds(params)

	builders := make([]*ent.RedirectURICreate, len(params))

	for i, v := range params {
		builders[i] = client.RedirectURI.Create().SetURI(v.uri).SetUserID(v.userID)
	}

	return client.RedirectURI.CreateBulk(builders...).Save(ctx)
}

func run(ctx context.Context, client *ent.Client) error {
	nUser := 10
	nAuthCodeByUser := 3
	nRedirectUriByUser := 3

	users, err := insertUsers(ctx, client, nUser)

	if err != nil {
		return err
	}

	_, err = insertAuthCodes(ctx, client, users, nAuthCodeByUser)

	if err != nil {
		return err
	}

	_, err = insertRedirectURIs(ctx, client, users, nRedirectUriByUser)

	if err != nil {
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

	if err = xid.Init(); err != nil {
		log.Fatal(err)
	}

	cfg.DBPort = 13306
	dbClient, err := repository.NewDBClient(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}

	entClient := repository.NewEntClient(dbClient)

	if err = run(ctx, entClient); err != nil {
		log.Fatal(err)
	}
}
